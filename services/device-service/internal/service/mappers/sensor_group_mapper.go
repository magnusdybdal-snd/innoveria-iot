package mappers

import (
	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/domain"
)

// MapCreateChirpstackApplication maps a domain SensorGroup to a CreateChirpstackApplication request.
func MapCreateChirpstackApplication(from domain.SensorGroup) dto.CreateChirpstackApplication {
	return dto.CreateChirpstackApplication{
		ApplicationPayload: dto.ApplicationPayload{
			Name:     from.Name,
			TenantID: from.CompanyId, // TODO handle mapping here
		},
	}
}

// MapChirpstackSensorGroupDtoToDomain maps a Chirpstack Application response to a domain SensorGroup.
func MapChirpstackSensorGroupDtoToDomain(from dto.ChirpstackApplication) domain.SensorGroup {
	return domain.SensorGroup{
		Id:        from.ID,
		Name:      from.Name,
		CompanyId: "", // TODO: Fill from databse
		Location:  "", // Handle elsewhere
	}
}
