// Package mappers provides functions for converting client DTOs to domain types.
package mappers

import (
	"innoveria-iot/collection-service/internal/clients/dto"
	"innoveria-iot/collection-service/internal/domain"
)

// ToSensorMetric converts a device-service sensor metric DTO to its domain representation.
func ToSensorMetric(r dto.SensorMetricResponse) domain.SensorMetric {
	return domain.SensorMetric{
		PayloadKey:      r.PayloadKey,
		MeasurementType: r.MeasurementType,
		Unit:            r.Unit,
	}
}
