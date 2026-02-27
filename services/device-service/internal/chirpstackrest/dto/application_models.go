package dto

// Chirpstack v4 model for applications
import "time"

type ChirpstackApplicationList struct {
	Result     []ChirpstackApplication `json:"result"`
	TotalCount int                     `json:"totalCount"`
}

type ChirpstackApplication struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Chirpstack application payload
type CreateChirpstackApplication struct {
	ApplicationPayload `json:"application"`
}

type ApplicationPayload struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	TenantID    string            `json:"tenantId"`
	Description string            `json:"description,omitempty"` // Not using description
	Tags        map[string]string `json:"tags,omitempty"`        // Not using tags
}
