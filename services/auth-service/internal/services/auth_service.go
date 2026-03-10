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
	companyRepo domain.CompanyRepo
}

// NewAuthServiceImpl creates a new AuthServiceImpl instance.
func NewAuthServiceImpl(companyRepo domain.CompanyRepo) *AuthServiceImpl {
	return &AuthServiceImpl{
		companyRepo: companyRepo,
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

	slog.Info("succsessfully registered company", "id", company.Id)
	return company, nil
}

// GetOneCompany retrieves a single company by its ID
func (s *AuthServiceImpl) GetOneCompany(ctx context.Context, companyID string) (domain.Company, error) {
	company, err := s.companyRepo.FindByID(ctx, companyID)
	if err != nil {
		return domain.Company{}, err
	}

	slog.Info("successfully found company", "id", company.Id)
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
