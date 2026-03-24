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
	// UpsertBatch inserts or updates a batch of sensor metrics in a single query.
	// Conflicts on (sensor_id, payload_key) update measurement_type and unit.
	UpsertBatch(ctx context.Context, metrics []SensorMetric) error
	FindBySensorID(ctx context.Context, sensorID string) ([]SensorMetric, error)
}

// SensorMetricService defines the business logic for managing sensor metrics.
type SensorMetricService interface {
	// UpsertMetrics saves operator-defined metric labels for a configurable sensor.
	UpsertMetrics(ctx context.Context, metrics []SensorMetric) error
	// GetEffectiveMetrics resolves the payload key mappings for a sensor.
	// Checks per-sensor metrics first, falls back to the profile-level payload schema.
	// Returns an empty slice if the sensor is not yet configured.
	GetEffectiveMetrics(ctx context.Context, deviceEUI string) ([]SensorMetric, error)
}
