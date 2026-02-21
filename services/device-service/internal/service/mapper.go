package service

import (
	"time"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
)

/*
	Gateway mapping
*/
// Mapping for the chirpstack gateway domain to device service gateway domain
func mapGateway(from chirpstackrest.ChirpstackGateway) domain.Gateway {
	return domain.Gateway{
		Id:         from.GatewayEUI, // TODO: Change this to internal database id
		DeviceEUI:  from.GatewayEUI,
		Name:       from.Name,
		Status:     mapStatus(from.State, from.LastSeenAt),
		LastSeenAt: from.LastSeenAt.Format(time.RFC1123),
	}
}

// Mapping for status
// Do number instead. And it will display as offline if last seen is bigger then 5 min
func mapStatus(state string, lastSeen time.Time) int {
	switch state {
	case "ONLINE":
		return 0
	case "NEVER_SEEN":
		return 1
	case "OFFLINE":
		return 2
	default:
		if time.Since(lastSeen) < 5*time.Minute {
			return 0
		}
		return 2
	}
}

/*
Sensor mapping
*/
func mapSensor(from chirpstackrest.ChirpstackSensor) domain.Sensor {
	return domain.Sensor{
		Id:         from.DeviceEUI, // TODO: Change this to internal database id
		Name:       from.Name,
		DeviceEUI:  from.DeviceEUI,
		GatewayEUI: "1234", // TODO: Handle in database
		Status:     mapStatusSensor(from.LastSeenAt),
		LastSeenAt: from.LastSeenAt.Format(time.RFC1123),
	}
}

func mapStatusSensor(lastSeen time.Time) int {
	if time.Since(lastSeen) < 5*time.Minute {
		return 0
	}
	return 2
}
