package domain

import (
	"context"
	"time"
)

// MeasurementReading is a single sensor measurement returned from the collection service.
type MeasurementReading struct {
	DeviceEUI string         `json:"device_eui"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload"`
	CompanyID string         `json:"company_id"`
}

// CollectionClient defines the operations context-service needs from the collection service.
type CollectionClient interface {
	GetMeasurements(ctx context.Context, deviceEUI string, from, to time.Time) ([]MeasurementReading, error)
}
