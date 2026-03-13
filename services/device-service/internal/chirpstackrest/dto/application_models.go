// Package dto contains request and response models for the Chirpstack REST API.
package dto

// Chirpstack v4 model for applications
import "time"

// ChirpstackApplicationList is the paginated response from the Chirpstack GET /api/applications endpoint.
type ChirpstackApplicationList struct {
	Result     []ChirpstackApplication `json:"result"`
	TotalCount int                     `json:"totalCount"`
}

// ChirpstackApplication represents a single application as returned by the Chirpstack API.
type ChirpstackApplication struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateChirpstackApplication is the top-level body for POST /api/applications.
type CreateChirpstackApplication struct {
	ApplicationPayload `json:"application"`
}

// ApplicationPayload contains the fields required by Chirpstack to create a new application.
type ApplicationPayload struct {
	Name     string `json:"name"`
	TenantID string `json:"tenantId"`
}

// ChirpstackApplicationCreateResponse is the response from Chirpstack when creating a tenant.
type ChirpstackApplicationCreateResponse struct {
	ID string `json:"id"`
}
