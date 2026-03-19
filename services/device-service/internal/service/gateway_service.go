// Package service implements the business logic for the device service.
package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

// GatewayServiceImpl implements domain.GatewayService, coordinating between the database and Chirpstack.
type GatewayServiceImpl struct {
	cc             *chirpstackrest.Client
	gatewayRepo    domain.GatewayRepository
	companycfgRepo domain.CompanyConfigRepository
}

// NewGatewayService creates a new GatewayServiceImpl with the given Chirpstack client and repositories.
func NewGatewayService(cc *chirpstackrest.Client, gatewayRepo domain.GatewayRepository, companyCfgRepo domain.CompanyConfigRepository) *GatewayServiceImpl {
	return &GatewayServiceImpl{
		cc:             cc,
		gatewayRepo:    gatewayRepo,
		companycfgRepo: companyCfgRepo,
	}
}

// Create adds a new gateway to Chirpstack and the database. The gateway will always
// be added to Chirpstack first. If this fails we return early and don't make an entry
// in the database. If the database insert fails, the gateway is deleted from Chirpstack
// as a compensating transaction to keep both systems in sync.
func (g *GatewayServiceImpl) Create(ctx context.Context, payload domain.Gateway) error {
	// Find the company's chirpstack tenant ID.
	companycfg, err := g.companycfgRepo.FindByCompanyID(ctx, payload.CompanyId)
	if err != nil {
		return fmt.Errorf("create gateway: finding company tenant ID: %w", err)
	}

	// Sending post request to chirpstack
	gatewayReq := mappers.MapChirpstackGatewayRequest(payload, companycfg.ChirpstackTenantID)
	if err := g.cc.CreateGateway(ctx, gatewayReq); err != nil {
		return fmt.Errorf("create gateway: add to chirpstack: %w", err)
	}

	// On successful creation in Chirpstack, try to store in DB.
	gateway, err := g.gatewayRepo.Create(ctx, payload)
	if err != nil {
		// Compensate: remove from Chirpstack so systems stay in sync.
		if compErr := g.cc.DeleteGateway(ctx, payload.GatewayEUI); compErr != nil {
			slog.Error("saga compensation failed: could not delete gateway from chirpstack after db insert failure",
				"eui", payload.GatewayEUI, "error", compErr)
		}
		return fmt.Errorf("create gateway: add to database: %w", err)
	}

	slog.Info("successfully created gateway", "id", gateway.Id)
	return nil
}

// Update updates the gateway's metadata. If the name changed, Chirpstack is updated first
// and the database second. If the database update fails, the Chirpstack rename is reverted
// as a compensating transaction. If only non-name fields changed, only the database is
// updated since Chirpstack stores no other gateway fields.
func (g *GatewayServiceImpl) Update(ctx context.Context, gatewayId string, payload domain.Gateway) error {
	// Verify that the gateway exists in db
	gateway, err := g.gatewayRepo.FindByID(ctx, gatewayId)
	if err != nil {
		return fmt.Errorf("update gateway: gateway %s not found in database: %w", gatewayId, err)
	}

	oldGateway := gateway // save before merge

	// Update values if not nil / empty string
	if payload.Name != "" {
		gateway.Name = payload.Name
	}
	if payload.Description != nil {
		gateway.Description = payload.Description
	}
	if payload.FactoryID != "" {
		gateway.FactoryID = payload.FactoryID
	}
	if payload.FactoryAreaID != "" {
		gateway.FactoryAreaID = payload.FactoryAreaID
	}

	// Only update in Chirpstack (and apply saga) if the name has changed (Chirpstack only allows name updates)
	if gateway.Name != oldGateway.Name {

		// check database for chirpstack tenant ID
		companycfg, err := g.companycfgRepo.FindByCompanyID(ctx, gateway.CompanyId)
		if err != nil {
			return fmt.Errorf("update gateway: finding company tenant ID: %w", err)
		}

		// Chirpstack put request, Chirpstack dont need gatewayID, just gatewayEUI
		newReq := mappers.MapChirpstackGatewayRequest(gateway, companycfg.ChirpstackTenantID)
		if err := g.cc.RenameGateway(ctx, newReq); err != nil {
			return fmt.Errorf("update gateway: update in chirpstack: %w", err)
		}

		// On successful update in Chirpstack, try to update the database.
		if err := g.gatewayRepo.Update(ctx, gatewayId, gateway); err != nil {
			// Compensate: revert Chirpstack to the old name.
			oldReq := mappers.MapChirpstackGatewayRequest(oldGateway, companycfg.ChirpstackTenantID)
			if compErr := g.cc.RenameGateway(ctx, oldReq); compErr != nil {
				slog.Error("saga compensation failed: could not revert gateway name in chirpstack after db update failure",
					"id", gatewayId, "error", compErr)
			}
			return fmt.Errorf("update gateway: update in database: %w", err)
		}
	} else {
		// Name unchanged — Chirpstack stores no other gateway fields, so only update the DB.
		if err := g.gatewayRepo.Update(ctx, gatewayId, gateway); err != nil {
			return fmt.Errorf("update gateway: update in database: %w", err)
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
// Chirpstack is always tried first so we keep the database entry if Chirpstack fails.
// If the database delete fails, the gateway is re-created in Chirpstack as a compensating
// transaction to keep both systems in sync.
func (g *GatewayServiceImpl) Delete(ctx context.Context, gatewayID string) error {
	// Get the gateway EUI from database
	gateway, err := g.gatewayRepo.FindByID(ctx, gatewayID)
	if err != nil {
		return fmt.Errorf("delete gateway: gateway %s not found in database: %w", gatewayID, err)
	}

	// Fetch companycfg before deletion — needed for compensation if DB delete fails.
	companycfg, err := g.companycfgRepo.FindByCompanyID(ctx, gateway.CompanyId)
	if err != nil {
		return fmt.Errorf("delete gateway: finding company tenant ID: %w", err)
	}

	// Delete in chirpstack
	if err := g.cc.DeleteGateway(ctx, gateway.GatewayEUI); err != nil {
		return fmt.Errorf("delete gateway: delete in chirpstack: %w", err)
	}

	// Delete in database after successfully deleting in Chirpstack.
	if err := g.gatewayRepo.Delete(ctx, gatewayID); err != nil {
		// Compensate: re-create in Chirpstack so systems stay in sync.
		gatewayReq := mappers.MapChirpstackGatewayRequest(gateway, companycfg.ChirpstackTenantID)
		if compErr := g.cc.CreateGateway(ctx, gatewayReq); compErr != nil {
			slog.Error("saga compensation failed: could not re-create gateway in chirpstack after db delete failure",
				"eui", gateway.GatewayEUI, "error", compErr)
		}
		return fmt.Errorf("delete gateway: delete in database: %w", err)
	}

	slog.Info("successfully deleted gateway", "id", gatewayID)
	return nil
}
