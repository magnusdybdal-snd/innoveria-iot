package domain

import "context"

// DeviceClient defines the operations context-service needs from the device service.
type DeviceClient interface {
	// GetSensorsByProductionResourceID returns all sensors assigned to the given
	// production resource, scoped to the given company. companyID, userID and role
	// are forwarded as auth headers to device-service.
	GetSensorsByProductionResourceID(ctx context.Context, companyID, userID, role string, productionResourceID int64) ([]DeviceSensor, error)

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
	ElectricitySensor  bool
	Voltage            *int // 230 or 400 — only set when ElectricitySensor is true
}

// SensorMetric maps a sensor payload key to its measurement type and unit.
type SensorMetric struct {
	PayloadKey      string
	MeasurementType string
	Unit            *string
}
