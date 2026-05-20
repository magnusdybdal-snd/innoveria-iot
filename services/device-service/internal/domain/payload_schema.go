package domain

import "context"

// PayloadSchema maps a raw payload key to a canonical measurement type for a given
// ChirpStack device profile. Used for fixed-schema sensors where all sensors on the
// same profile share the same payload structure.
type PayloadSchema struct {
	ID                  string
	ChirpstackProfileID string
	PayloadKey          string
	MeasurementType     string
	Unit                *string
}

// PayloadSchemaRepository handles persistence of payload schema rows in the database.
type PayloadSchemaRepository interface {
	// SaveLabels upserts measurement_type and unit for the given profile and payload keys.
	SaveLabels(ctx context.Context, schemas []PayloadSchema) error
	// FindByProfileID returns all rows for a given ChirpStack profile ID.
	FindByProfileID(ctx context.Context, chirpstackProfileID string) ([]PayloadSchema, error)
}

// PayloadSchemaService defines the business logic for managing payload schemas.
type PayloadSchemaService interface {
	// SaveLabels labels payload keys with their canonical measurement types.
	SaveLabels(ctx context.Context, schemas []PayloadSchema) error
	// GetByProfile returns all payload schema rows for a profile.
	GetByProfile(ctx context.Context, chirpstackProfileID string) ([]PayloadSchema, error)
}
