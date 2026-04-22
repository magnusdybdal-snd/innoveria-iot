package domain

import "context"

// DeviceClient defines the operations context-service needs from the device service.
type DeviceClient interface {
	// GetSensorsByProductionResourceID returns all sensors assigned to the given
	// production resource.
	//
	// TODO: verify — device-service stores ERPProductionResource.ID (int64) as a
	// string in the production_resource field. Double-check this mapping when the
	// real device-service integration is live.
	GetSensorsByProductionResourceID(ctx context.Context, productionResourceID string) ([]DeviceSensor, error)

	// GetSensorMetrics returns the payload-key → measurement type/unit mappings
	// for the given sensor EUI. Returns an empty slice if the sensor exists but
	// has no configured metrics.
	GetSensorMetrics(ctx context.Context, deviceEUI string) ([]SensorMetric, error)
}

// DeviceSensor is a sensor returned by the device service.
type DeviceSensor struct {
	ID                 string
	Name               string
	DeviceEUI          string
	ProductionResource *int64
}

// SensorMetric maps a sensor payload key to its measurement type and unit.
type SensorMetric struct {
	PayloadKey      string
	MeasurementType string
	Unit            *string
}
