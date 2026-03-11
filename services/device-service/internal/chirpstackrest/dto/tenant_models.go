package dto

// CreateChirpstackTenant is the top-level body for POST /api/tenants
type CreateChirpstackTenant struct {
	Tenant TenantPayload `json:"tenant"`
}

// TenantPayload contains the fields required by Chirpstack to create a new tenant
type TenantPayload struct {
	Name string `json:"name"`
}

// ChirpstackTenantCreateResponse is the response from Chirpstack when creating a tenant.
type ChirpstackTenantCreateResponse struct {
	ID string `json:"id"`
}
