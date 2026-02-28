package mappers

import (
	"time"

	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/domain"
)

/*
	Gateway mapping
*/
// Mapping for the chirpstack gateway domain to device service gateway domain
func MapChirpstackGateway(from dto.ChirpstackGateway) domain.Gateway {
	return domain.Gateway{
		Id:         from.GatewayEUI, // TODO: Change this to internal database id
		CompanyId:  from.TenantID,   // TODO: look up company mapping in DB
		GatewayEUI: from.GatewayEUI,
		Name:       from.Name,
		Status:     mapStatus(from.State, from.LastSeenAt),
		LastSeenAt: from.LastSeenAt.Format(time.RFC1123),
	}
}

// Mapping for gateway domain to chirpstack post and put requests
// Tennant id is chirpstacks internal understanding of companies
func MapCreateChirpstackGateway(from domain.Gateway, chirpstackTennantId string) dto.CreateChirpstackGatewayRequest {
	return dto.CreateChirpstackGatewayRequest{
		CreateGatewayPayload: dto.CreateGatewayPayload{
			GatewayEUI: from.GatewayEUI,
			Name:       from.Name,
			TenantID:   chirpstackTennantId,
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
