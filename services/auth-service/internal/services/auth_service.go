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
func (a *AuthServiceImpl) RegisterCompany(ctx context.Context, payload domain.Company) (domain.Company, error) {
	// 1. insert in database
	company, err := a.companyRepo.Create(ctx, payload)
	if err != nil {
		return domain.Company{}, err
	}

	slog.Info("succsessfully registered company", "id", company.Id)
	return company, nil
}

// GetOneCompany retrieves a single company by its ID
func (a *AuthServiceImpl) GetOneCompany(ctx context.Context, companyID string) (domain.Company, error) {
	company, err := a.companyRepo.FindByID(ctx, companyID)
	if err != nil {
		return domain.Company{}, err
	}

	slog.Info("successfully found company", "id", company.Id)
	return company, nil
}
