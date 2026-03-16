package domain

import "context"

// SensorMetric describes a single measurement type on a sensor.
// A sensor can have multiple metrics, e.g. temperature and humidity.
type SensorMetric struct {
	SensorID        string
	MeasurementType string  // i.e. "temperature", "humidity", "electrical current"
	Unit            *string // can be nil, not all measurements has a unit.
}

// SensorMetricRepository handles persistence of sensor metrics in the database.
type SensorMetricRepository interface {
	// Upsert inserts a metric or updates its unit if the combination (sensor_id, measurement_type)
	// already exists. This handles reconfiguration cleanly
	Upsert(ctx context.Context, metric SensorMetric) error
	FindBySensorID(ctx context.Context, sensorID string) ([]SensorMetric, error)
}
