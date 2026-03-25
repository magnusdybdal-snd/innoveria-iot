package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/domain"
)

// PayloadSchemaServiceImpl implements domain.PayloadSchemaService.
type PayloadSchemaServiceImpl struct {
	repo domain.PayloadSchemaRepository
}

// NewPayloadSchemaService creates a new PayloadSchemaServiceImpl with the given repository.
func NewPayloadSchemaService(repo domain.PayloadSchemaRepository) *PayloadSchemaServiceImpl {
	return &PayloadSchemaServiceImpl{repo: repo}
}

// DiscoverKeys creates draft payload schema rows for the given profile and payload keys.
// Called when an admin submits discovered keys from a sample payload via the UI.
func (s *PayloadSchemaServiceImpl) DiscoverKeys(ctx context.Context, chirpstackProfileID string, payloadKeys []string) error {
	if err := s.repo.UpsertDrafts(ctx, chirpstackProfileID, payloadKeys); err != nil {
		return fmt.Errorf("discover keys: %w", err)
	}
	slog.Info("discovered payload schema keys", "profile_id", chirpstackProfileID, "count", len(payloadKeys))
	return nil
}

// SaveLabels labels payload keys with their canonical measurement types for a profile.
// Returns domain.ErrMissingMeasurementType if any schema has a nil MeasurementType.
// Returns domain.ErrInvalidMeasurementType if any slug does not exist in the vocabulary.
func (s *PayloadSchemaServiceImpl) SaveLabels(ctx context.Context, schemas []domain.PayloadSchema) error {
	for _, schema := range schemas {
		if schema.MeasurementType == nil {
			return fmt.Errorf("save labels: nil measurement_type for payload_key %q (programming erorr)", schema.PayloadKey)
		}
	}

	if err := s.repo.SaveLabels(ctx, schemas); err != nil {
		return fmt.Errorf("save labels: %w", err)
	}

	slog.Info("saved payload schema labels", "count", len(schemas))
	return nil
}

// GetByProfile returns all payload schema rows for a profile, including drafts.
func (s *PayloadSchemaServiceImpl) GetByProfile(ctx context.Context, chirpstackProfileID string) ([]domain.PayloadSchema, error) {
	schemas, err := s.repo.FindByProfileID(ctx, chirpstackProfileID)
	if err != nil {
		return nil, fmt.Errorf("get payload schema by profile: %w", err)
	}
	return schemas, nil
}

// GetDraftProfiles returns ChirpStack profile IDs that have unlabeled rows.
// Used to populate the admin dashboard badge.
func (s *PayloadSchemaServiceImpl) GetDraftProfiles(ctx context.Context) ([]string, error) {
	profiles, err := s.repo.FindDraftProfiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("get draft profiles: %w", err)
	}
	return profiles, nil
}
