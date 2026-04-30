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
func (s *CompanyServiceImpl) RegisterCompany(ctx context.Context, payload domain.Company) (domain.Company, error) {
	company, err := s.companyRepo.Create(ctx, payload)
	if err != nil {
		return domain.Company{}, err
	}

	slog.Info("successfully registered company", "id", company.ID)
	return company, nil
}

func (s *CompanyServiceImpl) generateERPAgentRuntimeToken(companyID string) (string, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(s.erpAgentTokenTTL)

	claims := jwt.MapClaims{
		"sub":        companyID,
		"company_id": companyID,
		"token_use":  "erp_agent_runtime",
		"iss":        s.jwtIssuer,
		"aud":        "erp-agent-service",
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

// IssueERPAgentToken creates a new ERP agent runtime token for an existing company.
func (s *CompanyServiceImpl) IssueERPAgentToken(ctx context.Context, companyID string) (string, error) {
	if _, err := s.companyRepo.FindByID(ctx, companyID); err != nil {
		return "", err
	}

	token, err := s.generateERPAgentRuntimeToken(companyID)
	if err != nil {
		slog.Error("issue erp agent token failed", "error", err)
		return "", err
	}

	slog.Info("successfully issued erp agent token", "company_id", companyID)
	return token, nil
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
