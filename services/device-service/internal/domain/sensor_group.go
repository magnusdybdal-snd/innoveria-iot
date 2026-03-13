package domain

import "context"

// SensorGroup is how LoRaWAN groups devices (sensors) — Chirpstack calls this an application.
// A SensorGroup is a prerequisite for adding sensors.
// Gateways only packet-forward data; the application is responsible for managing devices per project.
type SensorGroup struct {
	Id        string
	Name      string
	CompanyId string
	Location  string
}

// SensorGroupService defines the business logic operations for sensor groups.
type SensorGroupService interface {
	Create(ctx context.Context, payload SensorGroup) error
	Update(ctx context.Context, sensorGroupId string, payload SensorGroup) error
	GetAll(ctx context.Context, limit int) ([]SensorGroup, error)
	Delete(ctx context.Context, sensorGroupId string) error
}
