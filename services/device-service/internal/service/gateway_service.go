package service

import (
	"context"
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
)

type GatewayServiceImpl struct {
	cc *chirpstackrest.Client
}

func NewGatewayService(cc *chirpstackrest.Client) *GatewayServiceImpl {
	return &GatewayServiceImpl{
		cc: cc,
	}
}

// TODO: Handle connection of new gateway to chirpstack
func (g *GatewayServiceImpl) Create(ctx context.Context, payload domain.Gateway, companyId string) error {
	// Convert to chirpstack models
	gatewayReq := mapCreateChirpstackGateway(payload, companyId)

	// call chirpstack
	err := g.cc.CreateGateway(ctx, gatewayReq)
	if err != nil {
		return err
	}
	// Store in db
	return nil
}

func (g *GatewayServiceImpl) GetAll(ctx context.Context) ([]domain.Gateway, error) {
	// 1. Get gateway from database
	limit := 1 // TODO: Get the actual meta data from device db

	// 2. Get status from chirpstack
	// TODO: get only the status
	resp, err := g.cc.GetAllGateways(ctx, limit) // limit is needed for chirpstack
	if err != nil {
		return nil, err
	}

	var result []domain.Gateway

	for _, gw := range resp.Result {
		result = append(result, mapChirpstackGateway(gw))
	}

	// 3. Merge status and gateway data

	return result, nil
}
