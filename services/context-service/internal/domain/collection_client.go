package domain

import (
	"context"
	"time"
)

// MeasurementReading is a single sensor measurement returned from the collection service.
type MeasurementReading struct {
	DeviceEUI string         
	Timestamp time.Time      
	Payload   map[string]any 
	CompanyID string         
}

// CollectionClient defines the operations context-service needs from the collection service.
type CollectionClient interface {
	GetMeasurements(ctx context.Context, deviceEUI string, from, to time.Time) ([]MeasurementReading, error)
}
