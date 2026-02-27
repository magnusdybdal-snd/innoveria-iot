package mappers

import (
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"time"
)

/*
Sensor mapping
*/
func MapChirpstackSensor(from chirpstackrest.ChirpstackSensor) domain.Sensor {
	return domain.Sensor{
		Id:         from.DeviceEUI, // TODO: Change this to internal database id
		Name:       from.Name,
		DeviceEUI:  from.DeviceEUI,
		GatewayEUI: "1234", // TODO: Handle in database
		Status:     mapStatusSensor(from.LastSeenAt),
		LastSeenAt: from.LastSeenAt.Format(time.RFC1123),
	}
}

// TODO: Find a better way to handle sensor status
func mapStatusSensor(lastSeen time.Time) domain.Status {
	if time.Since(lastSeen) < 5*time.Minute {
		return domain.StatusOnline
	}
	return domain.StatusOffline
}
