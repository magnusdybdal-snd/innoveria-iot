package service

import (
	"context"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

type GatewayServiceImpl struct {
	cc *chirpstackrest.Client
}

func NewGatewayService(cc *chirpstackrest.Client) *GatewayServiceImpl {
	return &GatewayServiceImpl{
		cc: cc,
	}
}

func (g *GatewayServiceImpl) Create(ctx context.Context, payload domain.Gateway) error {
	// Convert to chirpstack models
	// TODO: DB generates the gateway Id, which is not the same as gatewayEUI

	// Sending post request to chirpstack
	gatewayReq := mappers.MapCreateChirpstackGateway(payload, payload.CompanyId) // TODO: Change this to chirpstack tennant id in db
	// call chirpstack
	err := g.cc.CreateGateway(ctx, gatewayReq)
	if err != nil {
		return err
	}
	// Store in db
	return nil
}

// Handles the put request from handler
// remember gatewayId is not gatewayEUI
func (g *GatewayServiceImpl) Update(ctx context.Context, gatewayId string, payload domain.Gateway) error {
	// TODO: check gatewayId in db
	// check database for tennant id (chirpstack tennant id)

	// Chirpstack put request, Chirpstack dont need gatewayID, just gatewayEUI
	gatewayReq := mappers.MapCreateChirpstackGateway(payload, payload.CompanyId) // TODO: Change this to chirpstack tennant id in db
	err := g.cc.RenameGateway(ctx, gatewayReq)
	if err != nil {
		return err
	}

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

func (g *GatewayServiceImpl) Delete(ctx context.Context, gatewayID string) error {
	// delete from database

	// delete in chirpstack
	err := g.cc.DeleteGateway(ctx, gatewayID) // TODO: Change this to gateway eui
	if err != nil {
		return err
	}

	// No error return nil
	return nil
}
