package dto

import (
	"time"

	"innoveria-iot/device-service/internal/domain"
)

// MapGatewayDomainToDTO maps a slice of domain Gateways to a GatewayListResponse.
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
		ID:            from.Id,
		CompanyID:     from.CompanyId,
		GatewayEUI:    from.GatewayEUI,
		Name:          from.Name,
		Description:   from.Description,
		Status:        int(from.Status),
		State:         string(from.State),
		FactoryAreaID: from.FactoryAreaID,
		LastSeenAt:    from.LastSeenAt,
		CreatedAt:     from.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     from.UpdatedAt.Format(time.RFC3339),
	}
}

// MapGatewayDTOToDomain maps a CreateGatewayRequest to a domain Gateway, setting defaults for State and Status.
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
		Description:         from.Description,
		FactoryAreaID:       from.FactoryAreaID,
		ChirpstackProfileID: from.ChirpstackProfileID,
		ProductionResource:  from.ProductionResource,
	}
}

// MapCreateSensorDTOToDomain maps a CreateSensorRequest to a domain Sensor.
func MapCreateSensorDTOToDomain(from CreateSensorRequest) domain.Sensor {
	return domain.Sensor{
		CompanyID:           from.CompanyID,
		Name:                from.Name,
		Description:         from.Description,
		DeviceEUI:           from.DeviceEUI,
		ChirpstackProfileID: from.ChirpstackProfileID,
		FactoryAreaID:       from.FactoryAreaID,
		ProductionResource:  from.ProductionResource,
		State:               domain.DeviceStateActive, // Default state ACTIVE when created
	}
}

// MapSensorDomainToDTO maps a slice of domain Sensors to a SensorListResponse.
func MapSensorDomainToDTO(from []domain.Sensor) SensorListResponse {
	tot := len(from)
	sensors := make([]SensorResponse, tot)

	for i, s := range from {
		sensors[i] = mapSensor(s)
	}

	return SensorListResponse{
		TotalCount: tot,
		Sensors:    sensors,
	}
}

func mapSensor(from domain.Sensor) SensorResponse {
	return SensorResponse{
		ID:                  from.Id,
		CompanyID:           from.CompanyID,
		Name:                from.Name,
		Description:         from.Description,
		DeviceEUI:           from.DeviceEUI,
		State:               string(from.State),
		FactoryAreaID:       from.FactoryAreaID,
		ProductionResource:  from.ProductionResource,
		ChirpstackProfileID: from.ChirpstackProfileID,
		Status:              int(from.Status),
		LastSeenAt:          from.LastSeenAt,
		CreatedAt:           from.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           from.UpdatedAt.Format(time.RFC3339),
	}
}

// MapSensorProfileDomainToDTO maps a slice of domain SensorProfiles to a SensorProfileListResponse.
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

// MapSensorGroupToDomain maps a CreateSensorGroup request to a domain SensorGroup.
func MapSensorGroupToDomain(from CreateSensorGroup) domain.SensorGroup {
	return domain.SensorGroup{
		Id:        "", // converted in chirpstack
		Name:      from.Name,
		CompanyId: from.CompanyId, // TODO: Change this to tennant id
		Location:  "",             // TODO: handle this somewhere
	}
}

// MapSensorGroupToDTO maps a slice of domain SensorGroups to a SensorGroupListResponse.
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
