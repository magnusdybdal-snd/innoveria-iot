package mappers

import (
	"time"

	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/domain"
)

/*
	Gateway mapping
*/
// MergeGateway merges Chirpstack runtime data with DB domain data into a domain gateway
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
