// Package domain defines the core types and interfaces for the collection service.
package domain

import (
	"context"
	"time"
)

// SensorMeasurement represents a single decoded payload reading from a LoRaWAN sensor,
// tagged with the originating device and the internal company it belongs to.
type SensorMeasurement struct {
	DeviceEUI string         `json:"device_eui"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload"`
	CompanyID string         `json:"company_id"`
}

// MeasurementRepository is the persistence interface for sensor measurements.
type MeasurementRepository interface {
	Insert(ctx context.Context, measurement SensorMeasurement, tenantID string) error
	FindLatest(ctx context.Context, deviceEUI string) (SensorMeasurement, error)
	FindByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
	FindPayloadKeys(ctx context.Context, deviceEUI string) ([]string, error)
}

// MeasurementService is the business logic interface for sensor measurements.
type MeasurementService interface {
	Create(ctx context.Context, measurement SensorMeasurement, tenantID string) error
	GetLatest(ctx context.Context, deviceEUI string) (SensorMeasurement, error)
	GetByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
	GetPayloadKeys(ctx context.Context, deviceEUI string) ([]string, error)
}
