package domain

import (
	"context"
	"time"
)

type Gateway struct {
	Id   string
	Name string

	LastSeenAt time.Time
}

type GatewayService interface {
	Create()
	GetAll(ctx context.Context) ([]Gateway, error)
}
