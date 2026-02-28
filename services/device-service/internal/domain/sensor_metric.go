package domain

import "context"

type SensorMetric struct {
	SensorID        string
	MeasurementType string
	Unit            *string
}

type SensorMetricRepository interface {
	Upsert(ctx context.Context, metric SensorMetric) error
	FindBySensorID(ctx context.Context, sensorID string) ([]SensorMetric, error)
}
