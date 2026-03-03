package service

import (
	"context"
	"fmt"
	"log"

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

	log.Printf("successfully created gateway: %s", gateway.Id)
	return nil
}

// TODO: UPDATE IN DATABASE
func (g *GatewayServiceImpl) Update(ctx context.Context, gatewayId string, payload domain.Gateway) error {
	// Verify that the gateway exists in db
	gateway, err := g.gatewayRepo.FindByID(ctx, gatewayId)
	if err != nil {
		return fmt.Errorf("update gateway: gateway %s not found in database: %w", gatewayId, err)
	}
	// check database for chirpstack tenant ID
	companycfg, err := g.companycfgRepo.FindByCompanyID(ctx, gateway.CompanyId)
	if err != nil {
		return fmt.Errorf("update gateway: finding company tenant ID: %w", err)
	}

	// Chirpstack put request, Chirpstack dont need gatewayID, just gatewayEUI
	gatewayReq := mappers.MapCreateChirpstackGateway(payload, companycfg.ChirpstackTenantID)
	err = g.cc.RenameGateway(ctx, gatewayReq)
	if err != nil {
		return fmt.Errorf("update gateway: update in chirpstack: %w", err)
	}

	log.Printf("successfully updated gateway: %s", gateway.Id)
	return nil
}

// Returns all gateway meta data with gateway status
func (g *GatewayServiceImpl) GetAll(ctx context.Context) ([]domain.Gateway, error) {
	// 1. Get gateway from database
	limit := 1 // TODO: Get the actual gateway total count from database

	// 2. Get status from chirpstack
	// TODO: get only the status
	resp, err := g.cc.GetAllGateways(ctx, limit) // limit is needed for chirpstack
	if err != nil {
		return nil, err
	}

	var result []domain.Gateway

	for _, gw := range resp.Result {
		result = append(result, mappers.MapChirpstackGateway(gw))
	}

	// 3. Merge status and gateway data

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

	log.Printf("successfully deleted gateway %s.", gatewayID)
	return nil
}
