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
func mapChirpstackGateway(from chirpstackrest.ChirpstackGateway) domain.Gateway {
	return domain.Gateway{
		Id:         from.GatewayEUI, // TODO: Change this to internal database id
		GatewayEUI: from.GatewayEUI,
		Name:       from.Name,
		Status:     mapStatus(from.State, from.LastSeenAt),
		LastSeenAt: from.LastSeenAt.Format(time.RFC1123),
	}
}

func mapCreateChirpstackGateway(from domain.Gateway, companyId string) chirpstackrest.CreateChirpstackGatewayRequest {
	return chirpstackrest.CreateChirpstackGatewayRequest{
		CreateGatewayPayload: chirpstackrest.CreateGatewayPayload{
			GatewayEUI: from.GatewayEUI,
			Name:       from.Name,
			TenantID:   companyId,
		},
	}
}

// Mapping for status
// Do number instead. And it will display as offline if last seen is bigger then 5 min
func mapStatus(status string, lastSeen time.Time) domain.Status {
	switch status {
	case "ONLINE":
		return domain.StatusOnline
	case "NEVER_SEEN":
		return domain.StatusNeverSeen
	case "OFFLINE":
		return domain.StatusOffline
	default:
		// Fallback in case there is no status
		if time.Since(lastSeen) < 5*time.Minute {
			return domain.StatusOnline
		}
		return domain.StatusOffline
	}
}

/*
Sensor mapping
*/
func mapChirpstackSensor(from chirpstackrest.ChirpstackSensor) domain.Sensor {
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
