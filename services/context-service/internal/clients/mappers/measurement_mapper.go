// Package mappers provides functions for converting client DTOs to domain types.
package mappers

import (
	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/domain"
)

// ToMeasurementReading converts a collection service DTO to a domain MeasurementReading.
func ToMeasurementReading(measurement dto.MeasurementResponse) domain.MeasurementReading {
	return domain.MeasurementReading{
		DeviceEUI: measurement.DeviceEUI,
		Timestamp: measurement.Timestamp,
		Payload:   measurement.Payload,
		CompanyID: measurement.CompanyID,
	}
}
