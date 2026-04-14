package dto

import (
	"strings"
	"time"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/ptrutil"
)

// MapDraftProfilesToDTO maps a slice of draft profile IDs to a DraftProfilesResponse.
func MapDraftProfilesToDTO(from []string) DraftProfilesResponse {
	profileIDs := make([]string, len(from))
	copy(profileIDs, from)
	return DraftProfilesResponse{
		TotalCount: len(profileIDs),
		ProfileIDs: profileIDs,
	}
}

// MapPayloadSchemaDomainToDTO maps a slice of domain PayloadSchemas to a PayloadSchemaListResponse.
func MapPayloadSchemaDomainToDTO(from []domain.PayloadSchema) PayloadSchemaListResponse {
	schemas := make([]PayloadSchemaResponse, len(from))
	for i, s := range from {
		schemas[i] = PayloadSchemaResponse{
			ID:                  s.ID,
			ChirpstackProfileID: s.ChirpstackProfileID,
			PayloadKey:          s.PayloadKey,
			MeasurementType:     s.MeasurementType,
			Unit:                s.Unit,
		}
	}
	return PayloadSchemaListResponse{
		TotalCount: len(schemas),
		Schemas:    schemas,
	}
}

// MapSaveLabelsRequestToDomain maps a SavePayloadSchemaLabelsRequest to a slice of domain PayloadSchemas.
func MapSaveLabelsRequestToDomain(chirpstackProfileID string, from SavePayloadSchemaLabelsRequest) []domain.PayloadSchema {
	schemas := make([]domain.PayloadSchema, len(from.Labels))
	for i, l := range from.Labels {
		mt := l.MeasurementType
		schemas[i] = domain.PayloadSchema{
			ChirpstackProfileID: chirpstackProfileID,
			PayloadKey:          l.PayloadKey,
			MeasurementType:     &mt,
			Unit:                l.Unit,
		}
	}
	return schemas
}

// MapSensorMetricDomainToDTO maps a slice of domain SensorMetrics to a SensorMetricListResponse.
func MapSensorMetricDomainToDTO(from []domain.SensorMetric) SensorMetricListResponse {
	metrics := make([]SensorMetricResponse, len(from))
	for i, m := range from {
		metrics[i] = SensorMetricResponse{
			PayloadKey:      m.PayloadKey,
			MeasurementType: m.MeasurementType,
			Unit:            m.Unit,
		}
	}
	return SensorMetricListResponse{
		TotalCount: len(metrics),
		Metrics:    metrics,
	}
}

// MapUpsertMetricsRequestToDomain maps an UpsertSensorMetricsRequest to a slice of domain SensorMetrics.
func MapUpsertMetricsRequestToDomain(from UpsertSensorMetricsRequest) []domain.SensorMetric {
	metrics := make([]domain.SensorMetric, len(from.Metrics))
	for i, m := range from.Metrics {
		metrics[i] = domain.SensorMetric{
			PayloadKey:      m.PayloadKey,
			MeasurementType: m.MeasurementType,
			Unit:            m.Unit,
		}
	}
	return metrics
}

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
		FactoryID:     from.FactoryID,
		FactoryAreaID: from.FactoryAreaID,
		LastSeenAt:    from.LastSeenAt,
		CreatedAt:     from.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     from.UpdatedAt.Format(time.RFC3339),
	}
}

// MapGatewayDTOToDomain maps a CreateGatewayRequest to a domain Gateway, setting defaults for State and Status.
func MapGatewayDTOToDomain(from CreateGatewayRequest) domain.Gateway {
	return domain.Gateway{
		Id:            "", // converted later in db
		CompanyId:     from.CompanyId,
		GatewayEUI:    strings.ToLower(from.GatewayEUI),
		Name:          from.Name,
		Description:   from.Description,
		FactoryID:     from.FactoryID,
		FactoryAreaID: from.FactoryAreaID,
		State:         domain.DeviceStateActive,
		Status:        domain.StatusNeverSeen,
		LastSeenAt:    "", // converted later after chirpstack
	}
}

// MapUpdateSensorDTOToDomain maps an UpdateSensorRequest to a domain Sensor.
func MapUpdateSensorDTOToDomain(from UpdateSensorRequest) domain.Sensor {
	return domain.Sensor{
		Name:                ptrutil.Deref(from.Name),
		Description:         from.Description,
		ElectricitySensor:   from.ElectricitySensor,
		Voltage:             from.Voltage,
		FactoryID:           ptrutil.Deref(from.FactoryID),
		FactoryAreaID:       ptrutil.Deref(from.FactoryAreaID),
		ChirpstackProfileID: ptrutil.Deref(from.ChirpstackProfileID),
		ProductionResource:  from.ProductionResource,
	}
}

// MapUpdateGatewayDTOToDomain maps an UpdateGatewayRequest to a domain Gateway.
func MapUpdateGatewayDTOToDomain(from UpdateGatewayRequest) domain.Gateway {
	return domain.Gateway{
		Name:          ptrutil.Deref(from.Name),
		Description:   from.Description,
		FactoryID:     ptrutil.Deref(from.FactoryID),
		FactoryAreaID: ptrutil.Deref(from.FactoryAreaID),
	}
}

// MapCreateSensorDTOToDomain maps a CreateSensorRequest to a domain Sensor.
func MapCreateSensorDTOToDomain(from CreateSensorRequest) domain.Sensor {
	return domain.Sensor{
		CompanyID:           from.CompanyID,
		Name:                from.Name,
		Description:         from.Description,
		ElectricitySensor:   &from.ElectricitySensor,
		Voltage:             from.Voltage,
		DeviceEUI:           strings.ToLower(from.DeviceEUI),
		AppKey:              from.AppKey,
		ChirpstackProfileID: from.ChirpstackProfileID,
		FactoryID:           from.FactoryID,
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
		ElectricitySensor:   from.ElectricitySensor != nil && *from.ElectricitySensor,
		Voltage:             from.Voltage,
		DeviceEUI:           from.DeviceEUI,
		AppKey:              from.AppKey,
		State:               string(from.State),
		FactoryID:           from.FactoryID,
		FactoryAreaID:       from.FactoryAreaID,
		ProductionResource:  from.ProductionResource,
		ChirpstackProfileID: from.ChirpstackProfileID,
		Status:              int(from.Status),
		LastSeenAt:          from.LastSeenAt,
		CreatedAt:           from.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           from.UpdatedAt.Format(time.RFC3339),
	}
}

// MapMeasurementTypeDomainToDTO maps a slice of domain MeasurementTypes to a MeasurementTypeListResponse.
func MapMeasurementTypeDomainToDTO(from []domain.MeasurementType) MeasurementTypeListResponse {
	tot := len(from)
	types := make([]MeasurementTypeResponse, tot)

	for i, m := range from {
		types[i] = MeasurementTypeResponse{
			Slug:        m.Slug,
			DisplayName: m.DisplayName,
			Description: m.Description,
			DefaultUnit: m.DefaultUnit,
			Deprecated:  m.Deprecated,
		}
	}

	return MeasurementTypeListResponse{
		TotalCount:       tot,
		MeasurementTypes: types,
	}
}

// MapCreateMeasurementTypeDTOToDomain maps a CreateMeasurementTypeRequest to a domain MeasurementType.
func MapCreateMeasurementTypeDTOToDomain(from CreateMeasurementTypeRequest) domain.MeasurementType {
	return domain.MeasurementType{
		Slug:        from.Slug,
		DisplayName: from.DisplayName,
		Description: from.Description,
		DefaultUnit: from.DefaultUnit,
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
