package domain

import (
	"context"
	"time"
)

type Gateway struct {
	Id         string
	Name       string
	Status     int // 0=online, 1=never_seen, 2=offline
	LastSeenAt time.Time
}

type GatewayService interface {
	Create()
	GetAll(ctx context.Context) ([]Gateway, error)
}
