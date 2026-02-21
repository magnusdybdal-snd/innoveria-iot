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

func (g *GatewayServiceImpl) Create() {

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

	for _, g := range resp.Result {
		result = append(result, mapGateway(g))
	}

	// 3. Merge status and gateway data

	return result, nil
}
