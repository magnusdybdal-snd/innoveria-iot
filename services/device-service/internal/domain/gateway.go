package domain

import (
	"context"
)

type Gateway struct {
	Id         string
	GatewayEUI string
	Name       string
	Status     Status // 0=online, 1=never_seen, 2=offline
	LastSeenAt string
}

type GatewayService interface {
	Create(ctx context.Context, payload Gateway, companyId string) error
	GetAll(ctx context.Context) ([]Gateway, error)
}
