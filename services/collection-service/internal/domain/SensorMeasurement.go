package domain

import (
	"context"
	"time"
)

type SensorMeasurement struct {
	DeviceEUI string         `json:"device_eui"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload"`
	CompanyId string         `json:"company_id"`
}

type SensorRepository interface {
	Insert(ctx context.Context, measurement SensorMeasurement) error
	FindLatest(ctx context.Context, deviceEUI string) (SensorMeasurement, error)
	FindByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
}

type SensorService interface {
	Create(ctx context.Context, measurement SensorMeasurement) error
	GetLatest(ctx context.Context, deviceEUI string) (SensorMeasurement, error)
	GetByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
}
