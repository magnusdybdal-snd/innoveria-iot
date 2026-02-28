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
