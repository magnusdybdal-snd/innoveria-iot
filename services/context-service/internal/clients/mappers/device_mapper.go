package mappers

import (
	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/domain"
)

// ToDeviceSensor converts a device-service sensor DTO to its domain representation.
func ToDeviceSensor(r dto.DeviceSensorResponse) domain.DeviceSensor {
	return domain.DeviceSensor{
		ID:                 r.ID,
		Name:               r.Name,
		DeviceEUI:          r.DeviceEUI,
		ProductionResource: r.ProductionResource,
	}
}

// ToSensorMetric converts a device-service sensor metric DTO to its domain representation.
func ToSensorMetric(r dto.SensorMetricResponse) domain.SensorMetric {
	return domain.SensorMetric{
		PayloadKey:      r.PayloadKey,
		MeasurementType: r.MeasurementType,
		Unit:            r.Unit,
	}
}
