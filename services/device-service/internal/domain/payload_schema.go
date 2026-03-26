package domain

import "context"

// PayloadSchema maps a raw payload key to a canonical measurement type for a given
// ChirpStack device profile. Used for fixed-schema sensors where all sensors on the
// same profile share the same payload structure.
//
// A nil MeasurementType means the row is a draft — the key has been discovered from
// an incoming payload but not yet labeled by an admin.
type PayloadSchema struct {
	ID                  string
	ChirpstackProfileID string
	PayloadKey          string
	MeasurementType     *string // nil = draft, non-nil = labeled
	Unit                *string
}

// PayloadSchemaRepository handles persistence of payload schema rows in the database.
type PayloadSchemaRepository interface {
	// UpsertDrafts inserts payload keys as draft rows (measurement_type = NULL).
	// Skips keys that already exist — does not overwrite existing labels.
	UpsertDrafts(ctx context.Context, chirpstackProfileID string, payloadKeys []string) error
	// SaveLabels upserts measurement_type and unit for the given profile and payload keys.
	SaveLabels(ctx context.Context, schemas []PayloadSchema) error
	// FindByProfileID returns all rows for a given ChirpStack profile ID.
	FindByProfileID(ctx context.Context, chirpstackProfileID string) ([]PayloadSchema, error)
	// FindDraftProfiles returns distinct ChirpStack profile IDs that have at least one unlabeled row.
	FindDraftProfiles(ctx context.Context) ([]string, error)
}

// PayloadSchemaService defines the business logic for managing payload schemas.
type PayloadSchemaService interface {
	// DiscoverKeys creates draft rows for the given payload keys on a profile.
	DiscoverKeys(ctx context.Context, chirpstackProfileID string, payloadKeys []string) error
	// SaveLabels labels payload keys with their canonical measurement types.
	SaveLabels(ctx context.Context, schemas []PayloadSchema) error
	// GetByProfile returns all payload schema rows for a profile.
	GetByProfile(ctx context.Context, chirpstackProfileID string) ([]PayloadSchema, error)
	// GetDraftProfiles returns profile IDs that have unlabeled rows, for the admin badge.
	GetDraftProfiles(ctx context.Context) ([]string, error)
}
