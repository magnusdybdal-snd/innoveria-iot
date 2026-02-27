package dto

import (
	"innoveria-iot/device-service/internal/domain"
)

// Domain to dto mapping. So the response has the json tag and shows the total count

/*
	Gateway mapping
*/

// Gateway Domain -> DTO
func MapGatewayDomainToDTO(from []domain.Gateway) GatewayListResponse {
	tot := len(from)
	gateways := make([]GatewayResponse, tot)

	for i, g := range from {
		gateways[i] = mapGateway(g)
	}

	return GatewayListResponse{
		TotalCount: tot,
		Gateways:   gateways,
	}
}

func mapGateway(from domain.Gateway) GatewayResponse {
	return GatewayResponse{
		ID:         from.Id,
		CompanyId:  from.CompanyId,
		GatewayEUI: from.GatewayEUI,
		Name:       from.Name,
		Status:     int(from.Status),
		LastSeenAt: from.LastSeenAt,
	}
}

// Post requests Gateway DTO -> domain
func MapGatewayDTOToDomain(from CreateGatewayRequest) domain.Gateway {
	return domain.Gateway{
		Id:         "", // converted later in db
		CompanyId:  from.CompanyId,
		GatewayEUI: from.GatewayEUI,
		Name:       from.Name,
		Status:     domain.StatusNeverSeen,
		LastSeenAt: "", // converted later after chirpstack
	}
}

/*
Sensors mapping
*/
func MapSensorDomainToDTO(from []domain.Sensor) SensorListResponse {
	tot := len(from)
	sensors := make([]SensorResponse, tot)

	for i, s := range from {
		sensors[i] = mapSensors(s)
	}

	return SensorListResponse{
		TotalCount: tot,
		Sensors:    sensors,
	}
}

func mapSensors(from domain.Sensor) SensorResponse {
	return SensorResponse{
		ID:         from.Id,
		Name:       from.Name,
		DeviceEUI:  from.DeviceEUI,
		GatewayEUI: from.GatewayEUI,
		Status:     int(from.Status),
		LastSeenAt: from.LastSeenAt,
	}
}
