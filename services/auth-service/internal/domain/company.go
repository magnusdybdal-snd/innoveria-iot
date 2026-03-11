package domain

import (
	"context"
	"time"
)

// Company is the domain model for company
type Company struct {
	Id        string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CompanyRepo defines the company repository needed by the auth domain.
type CompanyRepo interface {
	Create(ctx context.Context, company Company) (Company, error)
	FindAll(ctx context.Context) ([]Company, error)
	FindByID(ctx context.Context, companyID string) (Company, error)
	DeleteByID(ctx context.Context, companyID string) error
}
