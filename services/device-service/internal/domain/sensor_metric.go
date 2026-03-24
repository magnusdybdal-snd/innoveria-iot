package domain

import "context"

// SensorMetric describes a single measurement type on a configurable sensor (e.g. UC300).
// Each row maps a raw payload key to a canonical measurement type for a specific sensor.
type SensorMetric struct {
	ID              string
	SensorID        string
	PayloadKey      string  // actual key in raw JSON payload, e.g. "bus1"
	MeasurementType string  // canonical slug, e.g. "nitrogen_ppm"
	Unit            *string // optional, e.g. "ppm"
}

// SensorMetricRepository handles persistence of sensor metrics in the database.
type SensorMetricRepository interface {
	// Upsert inserts a metric or updates measurement_type and unit if the combination
	// (sensor_id, payload_key) already exists. This handles reconfiguration cleanly.
	Upsert(ctx context.Context, metric SensorMetric) error
	FindBySensorID(ctx context.Context, sensorID string) ([]SensorMetric, error)
}
