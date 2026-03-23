package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/domain"
)

// MeasurementTypeServiceImpl implements domain.MeasurementTypeService.
type MeasurementTypeServiceImpl struct {
	repo domain.MeasurementTypeRepository
}

// NewMeasurementTypeService creates a new MeasurementTypeServiceImpl with the given repository.
func NewMeasurementTypeService(repo domain.MeasurementTypeRepository) *MeasurementTypeServiceImpl {
	return &MeasurementTypeServiceImpl{repo: repo}
}

// Create adds a new measurement type to the vocabulary.
func (s *MeasurementTypeServiceImpl) Create(ctx context.Context, m domain.MeasurementType) error {
	if err := s.repo.Create(ctx, m); err != nil {
		return err
	}
	slog.Info("created measurement type", "slug", m.Slug)
	return nil
}

// ListActive returns all non-deprecated measurement types. Used to populate UI dropdowns.
func (s *MeasurementTypeServiceImpl) ListActive(ctx context.Context) ([]domain.MeasurementType, error) {
	all, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("list active measurement types: %w", err)
	}

	active := []domain.MeasurementType{}
	for _, m := range all {
		if !m.Deprecated {
			active = append(active, m)
		}
	}
	return active, nil
}

// ListAll returns all measurement types including deprecated. Used for the admin management page.
func (s *MeasurementTypeServiceImpl) ListAll(ctx context.Context) ([]domain.MeasurementType, error) {
	all, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("list all measurement types: %w", err)
	}
	return all, nil
}

// Deprecate marks a measurement type as retired. It will no longer appear in dropdowns
// but existing references to the slug remain valid.
func (s *MeasurementTypeServiceImpl) Deprecate(ctx context.Context, slug string) error {
	if err := s.repo.Deprecate(ctx, slug); err != nil {
		return fmt.Errorf("deprecate measurement type: %w", err)
	}
	slog.Info("deprecated measurement type", "slug", slug)
	return nil
}
