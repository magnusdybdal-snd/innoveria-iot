package domain

import (
	"context"
)

type Sensor struct {
	Id         string
	Name       string
	DeviceEUI  string
	GatewayEUI string
	Status     Status // 0=online, 1=never_seen, 2=offline
	// TODO: Add state (active or inactive)
	LastSeenAt string
}

type SensorService interface {
	GetAll(ctx context.Context) ([]Sensor, error)
}
