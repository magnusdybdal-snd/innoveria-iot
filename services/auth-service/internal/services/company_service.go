package services

import (
	"context"
	"log/slog"

	"innoveria-iot/auth-service/internal/domain"
)

// CompanyServiceImpl implements company use-cases for the auth service.
type CompanyServiceImpl struct {
	companyRepo           domain.CompanyRepo
	erpAgentCredentialSvc domain.ERPAgentCredentialService
}

// NewCompanyService creates a new CompanyServiceImpl instance.
func NewCompanyService(companyRepo domain.CompanyRepo, erpAgentCredentialSvc domain.ERPAgentCredentialService) *CompanyServiceImpl {
	return &CompanyServiceImpl{companyRepo: companyRepo, erpAgentCredentialSvc: erpAgentCredentialSvc}
}

// RegisterCompany creates a new company.
func (s *CompanyServiceImpl) RegisterCompany(ctx context.Context, payload domain.Company) (domain.RegisterCompanyResult, error) {
	company, err := s.companyRepo.Create(ctx, payload)
	if err != nil {
		return domain.RegisterCompanyResult{}, err
	}

	credential, err := s.erpAgentCredentialSvc.CreateERPAgentCredential(ctx, domain.ERPAgentCredential{CompanyID: company.ID})
	if err != nil {
		if deleteErr := s.companyRepo.DeleteByID(ctx, company.ID); deleteErr != nil {
			slog.Error("failed to compensate company after erp credential creation error")
		}
		return domain.RegisterCompanyResult{}, err
	}

	slog.Info("successfully registered company", "id", company.ID)
	return domain.RegisterCompanyResult{
		Company:            company,
		ERPAgentCredential: credential,
	}, nil
}

// GetOneCompany retrieves a single company by its ID.
func (s *CompanyServiceImpl) GetOneCompany(ctx context.Context, companyID string) (domain.Company, error) {
	company, err := s.companyRepo.FindByID(ctx, companyID)
	if err != nil {
		return domain.Company{}, err
	}

	slog.Info("successfully found company", "id", company.ID)
	return company, nil
}

// GetAllCompanies retrieves all companies.
func (s *CompanyServiceImpl) GetAllCompanies(ctx context.Context) ([]domain.Company, error) {
	companies, err := s.companyRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("successfully found all companies")
	return companies, nil
}

// DeleteCompany deletes a company by ID.
func (s *CompanyServiceImpl) DeleteCompany(ctx context.Context, companyID string) error {
	if err := s.companyRepo.DeleteByID(ctx, companyID); err != nil {
		return err
	}

	slog.Info("successfully deleted company", "id", companyID)
	return nil
}
