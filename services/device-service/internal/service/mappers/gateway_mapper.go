// Package mappers provides functions for converting between domain, Chirpstack, and database models.
package mappers

import (
	"time"

	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/domain"
)

/*
	Gateway mapping
*/

// MergeGateway merges Chirpstack runtime data with database domain data into a domain Gateway.
func MergeGateway(cs dto.ChirpstackGateway, db domain.Gateway) domain.Gateway {
	return domain.Gateway{
		Id:            db.Id,
		CompanyId:     db.CompanyId,
		GatewayEUI:    db.GatewayEUI,
		Name:          db.Name,
		Description:   db.Description,
		State:         db.State,
		FactoryAreaID: db.FactoryAreaID,
		CreatedAt:     db.CreatedAt,
		UpdatedAt:     db.UpdatedAt,
		Status:        mapStatus(cs.State, cs.LastSeenAt),
		LastSeenAt:    cs.LastSeenAt.Format(time.RFC3339),
	}
}

// MapCreateChirpstackGateway maps a domain Gateway and a Chirpstack tenant ID to a CreateChirpstackGatewayRequest.
func MapCreateChirpstackGateway(from domain.Gateway, chirpstackTenantID string) dto.CreateChirpstackGatewayRequest {
	return dto.CreateChirpstackGatewayRequest{
		CreateGatewayPayload: dto.CreateGatewayPayload{
			GatewayEUI: from.GatewayEUI,
			Name:       from.Name,
			TenantID:   chirpstackTenantID,
		},
	}
}

// mapStatus converts a Chirpstack status string to a domain Status.
// Falls back to online/offline based on last seen time if the status string is unrecognised.
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
