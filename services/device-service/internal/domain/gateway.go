package domain

import (
	"context"
)

type Gateway struct {
	Id         string
	CompanyId  string
	GatewayEUI string
	Name       string
	Status     Status // 0=online, 1=never_seen, 2=offline
	LastSeenAt string
}

type GatewayService interface {
	Create(ctx context.Context, payload Gateway) error
	Update(ctx context.Context, gatewayId string, payload Gateway) error
	GetAll(ctx context.Context) ([]Gateway, error)
}
