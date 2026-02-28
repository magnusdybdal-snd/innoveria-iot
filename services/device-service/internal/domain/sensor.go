package domain

import (
	"context"
)

type Sensor struct {
	Id              string
	Name            string
	DeviceEUI       string // how chirpstack handles id
	GatewayEUI      string
	Status          Status // 0=online, 1=never_seen, 2=offline
	SensorProfileId string // LoRaWAN templates to describe device model
	// TODO: Add state (active or inactive)
	LastSeenAt string
}

type SensorService interface {
	GetAll(ctx context.Context) ([]Sensor, error)
	Create(ctx context.Context) error
}
