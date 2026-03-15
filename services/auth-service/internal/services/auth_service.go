// Package services contains auth-service business logic and orchestrates
// authentication workflows between repositories and security components.
package services

import (
	"context"
	"log/slog"

	"innoveria-iot/auth-service/internal/domain"
)

// AuthServiceImpl implements authentication use cases for the auth service.
type AuthServiceImpl struct {
	companyRepo     domain.CompanyRepo
	factoryRepo     domain.FactoryRepo
	factoryAreaRepo domain.FactoryAreaRepo
}

// NewAuthServiceImpl creates a new AuthServiceImpl instance.
func NewAuthServiceImpl(
	companyRepo domain.CompanyRepo,
	factoryRepo domain.FactoryRepo,
	factoryAreaRepo domain.FactoryAreaRepo,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		companyRepo:     companyRepo,
		factoryRepo:     factoryRepo,
		factoryAreaRepo: factoryAreaRepo,
	}
}

// RegisterCompany generates a new company and starts onboarding on device-service
// Its only possible for platform admin to register a new company
// So a permission check will happend which is only available for the platform admin
func (s *AuthServiceImpl) RegisterCompany(ctx context.Context, payload domain.Company) (domain.Company, error) {
	// 1. insert in database
	company, err := s.companyRepo.Create(ctx, payload)
	if err != nil {
		return domain.Company{}, err
	}

	slog.Info("succsessfully registered company", "id", company.ID)
	return company, nil
}

// GetOneCompany retrieves a single company by its ID
func (s *AuthServiceImpl) GetOneCompany(ctx context.Context, companyID string) (domain.Company, error) {
	company, err := s.companyRepo.FindByID(ctx, companyID)
	if err != nil {
		return domain.Company{}, err
	}

	slog.Info("successfully found company", "id", company.ID)
	return company, nil
}

// GetAllCompanies retrieves all companies.
func (s *AuthServiceImpl) GetAllCompanies(ctx context.Context) ([]domain.Company, error) {
	companies, err := s.companyRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("successfully found all companies")
	return companies, nil
}

// DeleteCompany deletes a company by ID.
func (s *AuthServiceImpl) DeleteCompany(ctx context.Context, companyID string) error {
	if err := s.companyRepo.DeleteByID(ctx, companyID); err != nil {
		return err
	}

	slog.Info("successfully deleted company", "id", companyID)
	return nil
}

// RegisterFactory creates a new factory.
func (s *AuthServiceImpl) RegisterFactory(ctx context.Context, payload domain.Factory) (domain.Factory, error) {
	factory, err := s.factoryRepo.Create(ctx, payload)
	if err != nil {
		return domain.Factory{}, err
	}

	slog.Info("successfully registered factory", "id", factory.ID)
	return factory, nil
}

// GetOneFactory retrieves a single factory by its ID.
func (s *AuthServiceImpl) GetOneFactory(ctx context.Context, factoryID string) (domain.Factory, error) {
	factory, err := s.factoryRepo.FindByID(ctx, factoryID)
	if err != nil {
		return domain.Factory{}, err
	}

	slog.Info("successfully found factory", "id", factory.ID)
	return factory, nil
}

// GetAllFactories retrieves all factories.
func (s *AuthServiceImpl) GetAllFactories(ctx context.Context) ([]domain.Factory, error) {
	factories, err := s.factoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("successfully found all factories")
	return factories, nil
}

// DeleteFactory deletes a factory by ID.
func (s *AuthServiceImpl) DeleteFactory(ctx context.Context, factoryID string) error {
	if err := s.factoryRepo.DeleteByID(ctx, factoryID); err != nil {
		return err
	}

	slog.Info("successfully deleted factory", "id", factoryID)
	return nil
}

// RegisterFactoryArea creates a new factory area
func (s *AuthServiceImpl) RegisterFactoryArea(ctx context.Context, payload domain.FactoryArea) (domain.FactoryArea, error) {
	area, err := s.factoryAreaRepo.Create(ctx, payload)
	if err != nil {
		return domain.FactoryArea{}, err
	}
	return area, nil
}

// GetAllFactoryAreas retrieves all factory areas.
func (s *AuthServiceImpl) GetAllFactoryAreas(ctx context.Context) ([]domain.FactoryArea, error) {
	areas, err := s.factoryAreaRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return areas, nil
}
