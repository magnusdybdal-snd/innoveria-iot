package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

type GatewayServiceImpl struct {
	cc             *chirpstackrest.Client
	gatewayRepo    domain.GatewayRepository
	companycfgRepo domain.CompanyConfigRepository
}

func NewGatewayService(cc *chirpstackrest.Client, gatewayRepo domain.GatewayRepository, companyCfgRepo domain.CompanyConfigRepository) *GatewayServiceImpl {
	return &GatewayServiceImpl{
		cc:             cc,
		gatewayRepo:    gatewayRepo,
		companycfgRepo: companyCfgRepo,
	}
}

// Create adds a new gateway to Chirpstack and the database. The gateway will always
// be added to Chirpstack first. If this fails we return early and don't make an entry
// in the database.
func (g *GatewayServiceImpl) Create(ctx context.Context, payload domain.Gateway) error {
	// Find the company's chirpstack tenant ID.
	companycfg, err := g.companycfgRepo.FindByCompanyID(ctx, payload.CompanyId)
	if err != nil {
		return fmt.Errorf("create gateway: finding company tenant ID: %w", err)
	}

	// Sending post request to chirpstack
	gatewayReq := mappers.MapCreateChirpstackGateway(payload, companycfg.ChirpstackTenantID)
	// call chirpstack
	err = g.cc.CreateGateway(ctx, gatewayReq)
	if err != nil {
		return fmt.Errorf("create gateway: add to chirpstack: %w", err)
	}
	// On successful creation in chirpstack we store in DB
	gateway, err := g.gatewayRepo.Create(ctx, payload)
	if err != nil {
		return fmt.Errorf("create gateway: add to database: %w", err)
	}

	slog.Info("successfully created gateway", "id", gateway.Id)
	return nil
}

// Update updates the gateways metadata in our database and the name in Chirpstack
// DB is called before Chirpstack here because if Chirpstack fails we still want
// to update the fields in the database.
func (g *GatewayServiceImpl) Update(ctx context.Context, gatewayId string, payload domain.Gateway) error {
	// Verify that the gateway exists in db
	gateway, err := g.gatewayRepo.FindByID(ctx, gatewayId)
	if err != nil {
		return fmt.Errorf("update gateway: gateway %s not found in database: %w", gatewayId, err)
	}

	// Update the gateway in the database
	err = g.gatewayRepo.Update(ctx, gatewayId, payload)
	if err != nil {
		return fmt.Errorf("update gateway: update in database: %w", err)
	}

	// Only update in Chirpstack if the name has changed (Chirpstack only allows name updates)
	if payload.Name != gateway.Name {

		// check database for chirpstack tenant ID
		companycfg, err := g.companycfgRepo.FindByCompanyID(ctx, gateway.CompanyId)
		if err != nil {
			return fmt.Errorf("update gateway: finding company tenant ID: %w", err)
		}

		// Ensure the payload EUI field is populated before mapping to Chirpstack
		payload.GatewayEUI = gateway.GatewayEUI

		// Chirpstack put request, Chirpstack dont need gatewayID, just gatewayEUI
		gatewayReq := mappers.MapCreateChirpstackGateway(payload, companycfg.ChirpstackTenantID)
		err = g.cc.RenameGateway(ctx, gatewayReq)
		if err != nil {
			return fmt.Errorf("update gateway: update in chirpstack: %w", err)
		}
	}

	slog.Info("successfully updated gateway", "id", gateway.Id)
	return nil
}

// GetAll retrieves all gateways belonging to a companyID from the database and merges
// it with the status from Chirpstack (status and last seen).
func (g *GatewayServiceImpl) GetAll(ctx context.Context) ([]domain.Gateway, error) {

	// Fetch all gateways belonging to the company in DB
	gateways, err := g.gatewayRepo.FindAllByCompanyID(ctx, "a0000000-0000-0000-0000-000000000001") // TODO: REPLACE HARDCODED COMPANYID WITH PROPER AUTH
	if err != nil {
		return nil, fmt.Errorf("get all gateways: getting gateways from db: %w", err)
	}

	// Gateways to be returned
	var result []domain.Gateway

	// Loop over gateways, get their chirpstack status, merge and append response
	for _, gw := range gateways {
		status, err := g.cc.GetOneGateway(ctx, gw.GatewayEUI)
		if err != nil {
			// If no status from Chirpstack, append gateway without status / last seen
			slog.Warn("failed to fetch gateway status from chirpstack", "error", err)
			result = append(result, gw)
			continue
		}
		result = append(result, mappers.MergeGateway(status, gw))
	}

	return result, nil
}

// Delete deletes a gateway from both Chirpstack and from the database. Deletion in
// Chirpstack is always tried first so we keep database entry if we fail
// Any failure will return early to prevent desyncing chirpstack and the database.
func (g *GatewayServiceImpl) Delete(ctx context.Context, gatewayID string) error {
	// Get the gateway EUI from database
	gateway, err := g.gatewayRepo.FindByID(ctx, gatewayID)
	if err != nil {
		return fmt.Errorf("delete gateway: gateway %s not found in database: %w", gatewayID, err)
	}

	// Delete in chirpstack
	err = g.cc.DeleteGateway(ctx, gateway.GatewayEUI)
	if err != nil {
		return fmt.Errorf("delete gateway: delete in chirpstack: %w", err)
	}

	// Delete in databse after successfully deleting in Chirpstack
	err = g.gatewayRepo.Delete(ctx, gatewayID)
	if err != nil {
		return fmt.Errorf("delete gateway: delete in database: %w", err)
	}

	slog.Info("successfully deleted gateway", "id", gatewayID)
	return nil
}
