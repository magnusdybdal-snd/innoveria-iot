package domain

import (
	"context"
	"time"
)

type SensorMeasurement struct {
	DeviceEUI string         `json:"device_eui"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload"`
	CompanyID string         `json:"company_id"`
}

type MeasurementRepository interface {
	Insert(ctx context.Context, measurement SensorMeasurement, tenantID string) error
	FindLatest(ctx context.Context, deviceEUI string) (SensorMeasurement, error)
	FindByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
}

type MeasurementService interface {
	Create(ctx context.Context, measurement SensorMeasurement, tenantID string) error
	GetLatest(ctx context.Context, deviceEUI string) (SensorMeasurement, error)
	GetByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
}
