package services

import (
	"context"
	"log/slog"

	"innoveria-iot/auth-service/internal/domain"
)

// FactoryServiceImpl implements factory use-cases for the auth service.
type FactoryServiceImpl struct {
	factoryRepo domain.FactoryRepo
}

// NewFactoryService creates a new FactoryServiceImpl instance.
func NewFactoryService(factoryRepo domain.FactoryRepo) *FactoryServiceImpl {
	return &FactoryServiceImpl{factoryRepo: factoryRepo}
}

// RegisterFactory creates a new factory.
func (s *FactoryServiceImpl) RegisterFactory(ctx context.Context, payload domain.Factory) (domain.Factory, error) {
	factory, err := s.factoryRepo.Create(ctx, payload)
	if err != nil {
		return domain.Factory{}, err
	}

	slog.Info("successfully registered factory", "id", factory.ID)
	return factory, nil
}

// GetOneFactory retrieves a single factory by its ID, scoped to the given company.
func (s *FactoryServiceImpl) GetOneFactory(ctx context.Context, companyID string, factoryID string) (domain.Factory, error) {
	factory, err := s.factoryRepo.FindByID(ctx, companyID, factoryID)
	if err != nil {
		return domain.Factory{}, err
	}

	slog.Debug("successfully found factory", "id", factory.ID)
	return factory, nil
}

// GetAllFactories retrieves all factories belonging to the given company.
func (s *FactoryServiceImpl) GetAllFactories(ctx context.Context, companyID string) ([]domain.Factory, error) {
	factories, err := s.factoryRepo.FindAll(ctx, companyID)
	if err != nil {
		return nil, err
	}

	slog.Debug("successfully found all factories")
	return factories, nil
}

// DeleteFactory deletes a factory by ID, scoped to the given company.
func (s *FactoryServiceImpl) DeleteFactory(ctx context.Context, companyID string, factoryID string) error {
	if err := s.factoryRepo.Delete(ctx, companyID, factoryID); err != nil {
		return err
	}

	slog.Debug("successfully deleted factory", "id", factoryID)
	return nil
}
