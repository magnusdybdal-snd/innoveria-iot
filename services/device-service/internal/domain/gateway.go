package domain

import (
	"context"
)

type Gateway struct {
	Id         string
	DeviceEUI  string
	Name       string
	Status     int // 0=online, 1=never_seen, 2=offline
	LastSeenAt string
}

type GatewayService interface {
	Create()
	GetAll(ctx context.Context) ([]Gateway, error)
}
