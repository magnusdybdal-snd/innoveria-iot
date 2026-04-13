package mappers

import (
	"time"

	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/ptrutil"
)

// MergeSensor merges Chirpstack runtime data with database metadata and returns a domain Sensor.
func MergeSensor(cs dto.ChirpstackSensor, db domain.Sensor) domain.Sensor {
	return domain.Sensor{
		// From database
		Id:                  db.Id,
		CompanyID:           db.CompanyID,
		DeviceEUI:           db.DeviceEUI,
		AppKey:              db.AppKey,
		Name:                db.Name,
		Description:         db.Description,
		ElectricitySensor:   db.ElectricitySensor,
		Voltage:             db.Voltage,
		State:               db.State,
		FactoryID:           db.FactoryID,
		FactoryAreaID:       db.FactoryAreaID,
		ProductionResource:  db.ProductionResource,
		ChirpstackProfileID: db.ChirpstackProfileID,
		CreatedAt:           db.CreatedAt,
		UpdatedAt:           db.UpdatedAt,
		// Runetime from Chirpstack
		Status:     mapStatusSensor(cs.LastSeenAt),
		LastSeenAt: cs.LastSeenAt.Format(time.RFC3339),
	}
}

// mapStatusSensor returns sensor status based on time since it last was seen in Chirpstack.
func mapStatusSensor(lastSeen time.Time) domain.Status {
	if lastSeen.IsZero() {
		return domain.StatusNeverSeen
	}
	// TODO: Sensors have different hearthbeats, should probably be stored alongside profileID.
	if time.Since(lastSeen) < 2*time.Hour {
		return domain.StatusOnline
	}
	return domain.StatusOffline
}

// MapChirpstackSensorRequest builds a Chirpstack registration/update request.
func MapChirpstackSensorRequest(sensor domain.Sensor, applicationID string) dto.ChirpstackSensorRequest {
	return dto.ChirpstackSensorRequest{
		SensorPayload: dto.SensorPayload{
			DeviceEUI:       sensor.DeviceEUI,
			Name:            sensor.Name,
			Description:     ptrutil.Deref(sensor.Description),
			ApplicationID:   applicationID,
			DeviceProfileID: sensor.ChirpstackProfileID,
			JoinEUI:         "0000000000000000", // Not an issue when we host Chirpstack privately.
		},
	}
}

// MapChirpstackSensorKeyRequest builds a Chirpstack device key request from a domain Sensor.
func MapChirpstackSensorKeyRequest(sensor domain.Sensor) dto.ChirpstackSensorKeyRequest {
	return dto.ChirpstackSensorKeyRequest{
		DeviceKeys: dto.SensorKeysPayload{
			DevEUI: sensor.DeviceEUI,
			NwkKey: sensor.AppKey,
		},
	}
}
