// Package dto TODO(@vinjar): add proper documentation.
package dto

// Chirpstack v4 model for applications
import "time"

// ChirpstackApplicationList TODO(@vinjar): add proper documentation.
type ChirpstackApplicationList struct {
	Result     []ChirpstackApplication `json:"result"`
	TotalCount int                     `json:"totalCount"`
}

// ChirpstackApplication TODO(@vinjar): add proper documentation.
type ChirpstackApplication struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateChirpstackApplication TODO(@vinjar): add proper documentation.
type CreateChirpstackApplication struct {
	ApplicationPayload `json:"application"`
}

// ApplicationPayload TODO(@vinjar): add proper documentation.
type ApplicationPayload struct {
	Name        string            `json:"name"`
	TenantID    string            `json:"tenantId"`
	Description string            `json:"description,omitempty"` // Not using description
	Tags        map[string]string `json:"tags,omitempty"`        // Not using tags
}
