package mappers

import (
	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/domain"
)

// MapCreateChirpstackApplication turns domain into chirpstack models
func MapCreateChirpstackApplication(from domain.SensorGroup) dto.CreateChirpstackApplication {
	return dto.CreateChirpstackApplication{
		ApplicationPayload: dto.ApplicationPayload{
			Name:     from.Name,
			TenantID: from.CompanyId, // TODO handle mapping here
		},
	}
}

// MapChirpstackSensorGroupDtoToDomain TODO(@vinjar): add proper documentation.
func MapChirpstackSensorGroupDtoToDomain(from dto.ChirpstackApplication) domain.SensorGroup {
	return domain.SensorGroup{
		Id:        from.ID,
		Name:      from.Name,
		CompanyId: "", // TODO: Fill from databse
		Location:  "", // Handle elsewhere
	}
}
