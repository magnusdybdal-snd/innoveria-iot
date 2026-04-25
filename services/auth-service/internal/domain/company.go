package domain

import (
	"context"
	"time"
)

// Company is the domain model for company
type Company struct {
	ID        string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// RegisterCompanyResult holds the created company and initial ERP agent credential.
type RegisterCompanyResult struct {
	Company       Company
	ERPAgentToken string
}

// CompanyRepo defines the company repository needed by the auth domain.
type CompanyRepo interface {
	Create(ctx context.Context, company Company) (Company, error)
	FindAll(ctx context.Context) ([]Company, error)
	FindByID(ctx context.Context, companyID string) (Company, error)
	DeleteByID(ctx context.Context, companyID string) error
}

// CompanyService defines company use-cases exposed by the service layer.
type CompanyService interface {
	RegisterCompany(ctx context.Context, payload Company) (RegisterCompanyResult, error)
	GetOneCompany(ctx context.Context, companyID string) (Company, error)
	GetAllCompanies(ctx context.Context) ([]Company, error)
	DeleteCompany(ctx context.Context, companyID string) error
}
