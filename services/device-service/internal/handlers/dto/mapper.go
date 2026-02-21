package dto

import (
	"innoveria-iot/device-service/internal/domain"
	"time"
)

func MapDomainToDTO(g []domain.Gateway) GatewayListResponse {
	tot := len(g)
	devices := make([]GatewayResponse, tot)

	for i, d := range g {
		devices[i] = mapGateway(d)
	}

	return GatewayListResponse{
		TotalCount: tot,
		Gateways:   devices,
	}
}

func mapGateway(g domain.Gateway) GatewayResponse {
	return GatewayResponse{
		ID:         g.Id,
		Name:       g.Name,
		Status:     g.Status,
		LastSeenAt: g.LastSeenAt.Format(time.RFC822Z),
	}
}
