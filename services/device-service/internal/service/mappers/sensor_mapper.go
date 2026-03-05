package mappers

import (
	"time"

	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/domain"
)

// MergeSensor merges Chirpstack runtime data with db metadata and returns domain Sensor
func MergeSensor(cs dto.ChirpstackSensor, db domain.Sensor) domain.Sensor {
	return domain.Sensor{
		// From database
		Id:                  db.Id,
		CompanyID:           db.CompanyID,
		DeviceEUI:           db.DeviceEUI,
		Name:                db.Name,
		Description:         db.Description,
		State:               db.State,
		FactoryAreaID:       db.FactoryAreaID,
		ProductionResource:  db.ProductionResource,
		ChirpstackProfileID: db.ChirpstackProfileID,
		CreatedAt:           db.CreatedAt,
		UpdatedAt:           db.UpdatedAt,
		// Runetime from Chirpstack
		Status:     mapStatusSensor(cs.LastSeenAt),
		LastSeenAt: cs.LastSeenAt.Format(time.RFC1123),
	}
}

// TODO: Find a better way to handle sensor status
func mapStatusSensor(lastSeen time.Time) domain.Status {
	if time.Since(lastSeen) < 5*time.Minute {
		return domain.StatusOnline
	}
	return domain.StatusOffline
}
