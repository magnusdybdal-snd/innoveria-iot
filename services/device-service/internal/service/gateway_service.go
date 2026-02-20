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

	// 2. Get status from chirpstack

	// 3. Merge status and gateway data

	return nil, nil
}
