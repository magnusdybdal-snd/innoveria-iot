package domain

import "context"

// SensorMetric is the collection-service view of a device metric mapping.
type SensorMetric struct {
	PayloadKey      string
	MeasurementType string
	Unit            *string
}

// DeviceClient defines the operations this service needs from the device service.
type DeviceClient interface {
	GetSensorMetrics(ctx context.Context, deviceEUI string) ([]SensorMetric, error)
}
