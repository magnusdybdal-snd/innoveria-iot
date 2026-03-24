package services

import (
	"context"
	"log/slog"

	"innoveria-iot/auth-service/internal/domain"
)

// FactoryAreaServiceImpl implements factory area use-cases for the auth service.
type FactoryAreaServiceImpl struct {
	factoryAreaRepo domain.FactoryAreaRepo
}

// NewFactoryAreaService creates a new FactoryAreaServiceImpl instance.
func NewFactoryAreaService(factoryAreaRepo domain.FactoryAreaRepo) *FactoryAreaServiceImpl {
	return &FactoryAreaServiceImpl{factoryAreaRepo: factoryAreaRepo}
}

// RegisterFactoryArea creates a new factory area.
func (s *FactoryAreaServiceImpl) RegisterFactoryArea(ctx context.Context, payload domain.FactoryArea) (domain.FactoryArea, error) {
	area, err := s.factoryAreaRepo.Create(ctx, payload)
	if err != nil {
		return domain.FactoryArea{}, err
	}

	slog.Info("successfully registered factory area", "id", area.ID)
	return area, nil
}

// GetAllFactoryAreas retrieves all factory areas.
func (s *FactoryAreaServiceImpl) GetAllFactoryAreas(ctx context.Context) ([]domain.FactoryArea, error) {
	areas, err := s.factoryAreaRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("successfully found all factory areas")
	return areas, nil
}

// GetOneFactoryArea retrieves a single factory area by its ID.
func (s *FactoryAreaServiceImpl) GetOneFactoryArea(ctx context.Context, areaID string) (domain.FactoryArea, error) {
	area, err := s.factoryAreaRepo.FindByID(ctx, areaID)
	if err != nil {
		return domain.FactoryArea{}, err
	}

	slog.Info("successfully found factory area", "id", area.ID)
	return area, nil
}

// DeleteFactoryArea deletes a factory area by ID.
func (s *FactoryAreaServiceImpl) DeleteFactoryArea(ctx context.Context, areaID string) error {
	if err := s.factoryAreaRepo.Delete(ctx, areaID); err != nil {
		return err
	}

	slog.Info("successfully deleted factory area", "id", areaID)
	return nil
}
