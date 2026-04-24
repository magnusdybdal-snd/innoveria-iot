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

// GetAllFactoryAreas retrieves all factory areas for a factory, scoped to the given company.
func (s *FactoryAreaServiceImpl) GetAllFactoryAreas(ctx context.Context, companyID string, factoryID string) ([]domain.FactoryArea, error) {
	areas, err := s.factoryAreaRepo.FindAll(ctx, companyID, factoryID)
	if err != nil {
		return nil, err
	}

	slog.Debug("successfully found all factory areas")
	return areas, nil
}

// GetOneFactoryArea retrieves a single factory area by its ID, scoped to the given company.
func (s *FactoryAreaServiceImpl) GetOneFactoryArea(ctx context.Context, companyID string, areaID string) (domain.FactoryArea, error) {
	area, err := s.factoryAreaRepo.FindByID(ctx, companyID, areaID)
	if err != nil {
		return domain.FactoryArea{}, err
	}

	slog.Debug("successfully found factory area", "id", area.ID)
	return area, nil
}

// DeleteFactoryArea deletes a factory area by ID, scoped to the given company.
func (s *FactoryAreaServiceImpl) DeleteFactoryArea(ctx context.Context, companyID string, areaID string) error {
	if err := s.factoryAreaRepo.Delete(ctx, companyID, areaID); err != nil {
		return err
	}

	slog.Debug("successfully deleted factory area", "id", areaID)
	return nil
}
