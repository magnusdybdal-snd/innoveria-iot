package mappers

import "innoveria-iot/device-service/internal/chirpstackrest/dto"

// MapCreateChirpstackTenant maps a company name to a CreateChirpstackTenant request.
func MapCreateChirpstackTenant(name string) dto.CreateChirpstackTenant {
	return dto.CreateChirpstackTenant{
		Tenant: dto.TenantPayload{Name: name},
	}
}

// MapCreateChirpstackApplication maps a company name and tenantID to a CreateChirpstackApplication request.
func MapCreateChirpstackApplication(name string, tenantID string) dto.CreateChirpstackApplication {
	return dto.CreateChirpstackApplication{
		ApplicationPayload: dto.ApplicationPayload{
			Name:     name,
			TenantID: tenantID,
		},
	}
}
