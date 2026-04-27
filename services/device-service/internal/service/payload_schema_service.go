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

// SaveLabels labels payload keys with their canonical measurement types for a profile.
// Returns domain.ErrInvalidMeasurementType if any slug does not exist in the vocabulary.
func (s *PayloadSchemaServiceImpl) SaveLabels(ctx context.Context, schemas []domain.PayloadSchema) error {
	if err := s.repo.SaveLabels(ctx, schemas); err != nil {
		return fmt.Errorf("save labels: %w", err)
	}

	slog.Info("saved payload schema labels", "count", len(schemas))
	return nil
}

// GetByProfile returns all labeled payload schema rows for a profile.
func (s *PayloadSchemaServiceImpl) GetByProfile(ctx context.Context, chirpstackProfileID string) ([]domain.PayloadSchema, error) {
	schemas, err := s.repo.FindByProfileID(ctx, chirpstackProfileID)
	if err != nil {
		return nil, fmt.Errorf("get payload schema by profile: %w", err)
	}
	return schemas, nil
}
