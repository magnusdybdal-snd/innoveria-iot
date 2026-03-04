// Package domain TODO(@Magnus Dybdal): add proper documentation.
package domain

import (
	"context"
	"time"
)

// SensorMeasurement TODO(@Magnus Dybdal): add proper documentation.
type SensorMeasurement struct {
	DeviceEUI string         `json:"device_eui"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload"`
	CompanyID string         `json:"company_id"`
}

// MeasurementRepository TODO(@Magnus Dybdal): add proper documentation.
type MeasurementRepository interface {
	Insert(ctx context.Context, measurement SensorMeasurement, tenantID string) error
	FindLatest(ctx context.Context, deviceEUI string) (SensorMeasurement, error)
	FindByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
}

// MeasurementService TODO(@Magnus Dybdal): add proper documentation.
type MeasurementService interface {
	Create(ctx context.Context, measurement SensorMeasurement, tenantID string) error
	GetLatest(ctx context.Context, deviceEUI string) (SensorMeasurement, error)
	GetByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
}
