package service

import (
	"time"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
)

func mapGateway(from chirpstackrest.ChirpstackGateway) domain.Gateway {
	return domain.Gateway{
		Id:         from.GatewayID,
		Name:       from.Name,
		Status:     mapState(from.State, from.LastSeenAt),
		LastSeenAt: from.LastSeenAt,
	}
}

func mapState(state string, lastSeen time.Time) int {
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
