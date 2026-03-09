package mappers

import "innoveria-iot/device-service/internal/chirpstackrest/dto"

// MapCreateChirpstackTenant maps a companyID to a CreateChirpstackTenant request.
func MapCreateChirpstackTenant(companyID string) dto.CreateChirpstackTenant {
	return dto.CreateChirpstackTenant{
		Tenant: dto.TenantPayload{Name: companyID},
	}
}

// MapCreateChirpstackApplication maps a companyID and tenantID to a CreateChirpstackApplication request.
func MapCreateChirpstackApplication(companyID string, tenantID string) dto.CreateChirpstackApplication {
	return dto.CreateChirpstackApplication{
		ApplicationPayload: dto.ApplicationPayload{
			Name:     companyID,
			TenantID: tenantID,
		},
	}
}
