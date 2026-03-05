package dto

import (
	"innoveria-iot/device-service/internal/domain"
)

// MapGatewayDomainToDTO TODO(@Magnus Dybdal): add proper documentation.
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

// MapGatewayDTOToDomain TODO(@Magnus Dybdal): add proper documentation.
func MapGatewayDTOToDomain(from CreateGatewayRequest) domain.Gateway {
	return domain.Gateway{
		Id:         "", // converted later in db
		CompanyId:  from.CompanyId,
		GatewayEUI: from.GatewayEUI,
		Name:       from.Name,
		State:      domain.DeviceStateActive,
		Status:     domain.StatusNeverSeen,
		LastSeenAt: "", // converted later after chirpstack
	}
}

// MapUpdateSensorDTOToDomain maps an UpdateSensorRequest to a domain Sensor.
func MapUpdateSensorDTOToDomain(from UpdateSensorRequest) domain.Sensor {
	return domain.Sensor{
		Name:                from.Name,
		Description:         &from.Description,
		FactoryAreaID:       &from.FactoryAreaID,
		ChirpstackProfileID: from.ChirpstackProfileId,
	}
}

// MapSensorDomainToDTO TODO(@Magnus Dybdal): add proper documentation.
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
		Status:     int(from.Status),
		LastSeenAt: from.LastSeenAt,
	}
}

// MapSensorProfileDomainToDTO TODO(@Magnus Dybdal): add proper documentation.
func MapSensorProfileDomainToDTO(from []domain.SensorProfile) SensorProfileListResponse {
	tot := len(from)
	sensorProfiles := make([]SensorProfileResponse, tot)

	for i, sp := range from {
		sensorProfiles[i] = mapSensorProfiles(sp)
	}
	return SensorProfileListResponse{
		TotalCount:     tot,
		SensorProfiles: sensorProfiles,
	}
}

func mapSensorProfiles(from domain.SensorProfile) SensorProfileResponse {
	return SensorProfileResponse{
		Id:         from.Id,
		Name:       from.Name,
		Region:     from.Region,
		MACVersion: from.MACVersion,
		VendorId:   from.VendorId,
		VendorName: from.VendorName,
	}
}

// MapSensorGroupToDomain TODO(@Magnus Dybdal): add proper documentation.
func MapSensorGroupToDomain(from CreateSensorGroup) domain.SensorGroup {
	return domain.SensorGroup{
		Id:        "", // converted in chirpstack
		Name:      from.Name,
		CompanyId: from.CompanyId, // TODO: Change this to tennant id
		Location:  "",             // TODO: handle this somewhere
	}
}

// MapSensorGroupToDTO TODO(@Magnus Dybdal): add proper documentation.
func MapSensorGroupToDTO(from []domain.SensorGroup) SensorGroupListResponse {
	tot := len(from)

	sensorGroups := make([]SensorGroupResponse, tot)

	for i, sg := range from {
		sensorGroups[i] = mapSensorGroup(sg)
	}
	return SensorGroupListResponse{
		TotalCount:   tot,
		SensorGroups: sensorGroups,
	}
}

func mapSensorGroup(from domain.SensorGroup) SensorGroupResponse {
	return SensorGroupResponse{
		Id:        from.Id,
		Name:      from.Name,
		CompanyId: from.CompanyId,
		Location:  from.Location,
	}
}
