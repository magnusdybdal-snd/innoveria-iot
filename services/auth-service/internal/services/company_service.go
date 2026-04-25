package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"innoveria-iot/auth-service/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

// CompanyServiceImpl implements company use-cases for the auth service.
type CompanyServiceImpl struct {
	companyRepo      domain.CompanyRepo
	jwtSecret        []byte
	jwtIssuer        string
	erpAgentTokenTTL time.Duration
}

// NewCompanyService creates a new CompanyServiceImpl instance.
func NewCompanyService(companyRepo domain.CompanyRepo, jwtSecret, jwtIssuer string, erpAgentTokenTTL time.Duration) *CompanyServiceImpl {
	return &CompanyServiceImpl{
		companyRepo:      companyRepo,
		jwtSecret:        []byte(jwtSecret),
		jwtIssuer:        jwtIssuer,
		erpAgentTokenTTL: erpAgentTokenTTL,
	}
}

// RegisterCompany creates a new company.
func (s *CompanyServiceImpl) RegisterCompany(ctx context.Context, payload domain.Company) (domain.RegisterCompanyResult, error) {
	company, err := s.companyRepo.Create(ctx, payload)
	if err != nil {
		return domain.RegisterCompanyResult{}, err
	}

	erpAgentToken, err := s.generateERPAgentRuntimeToken(company.ID)
	if err != nil {
		return domain.RegisterCompanyResult{}, err
	}

	slog.Info("successfully registered company", "id", company.ID)
	return domain.RegisterCompanyResult{
		Company:       company,
		ERPAgentToken: erpAgentToken,
	}, nil
}

func (s *CompanyServiceImpl) generateERPAgentRuntimeToken(companyID string) (string, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(s.erpAgentTokenTTL)

	claims := jwt.MapClaims{
		"sub":        companyID,
		"company_id": companyID,
		"token_use":  "erp_agent_runtime",
		"iss":        s.jwtIssuer,
		"aud":        "erp-ingest",
		"iat":        now.Unix(),
		"exp":        expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("sign erp agent token: %w", err)
	}

	return signed, nil
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
