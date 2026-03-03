package domain

import "context"

// SensorMetric describes a single measurement on a sensor.
// A sensor can combine multiple sensor metrics like temp and humidity
type SensorMetric struct {
	SensorID        string
	MeasurementType string  // i.e. "temperature", "humidity", "electrical current"
	Unit            *string // can be nil, not all measurements has a unit.
}

// SensorMetricRepository TODO(@Magnus Dybdal): add proper documentation.
type SensorMetricRepository interface {
	// Upsert inserts a metric or updates its unit if the combination (sensor_id, measurement_type)
	// already exists. This handles reconfiguration cleanly
	Upsert(ctx context.Context, metric SensorMetric) error
	FindBySensorID(ctx context.Context, sensorID string) ([]SensorMetric, error)
}
