// Package domain defines the core business entities, interfaces, and types for the device service.
package domain

import (
	"context"
	"errors"
	"time"
)

// ErrNotFound is returned when a requested resource does not exist in the database.
var ErrNotFound = errors.New("not found")

// CompanyConfig holds the Chirpstack tenant and application IDs associated with a company, created during onboarding.
type CompanyConfig struct {
	CompanyID               string
	ChirpstackTenantID      string
	ChirpstackApplicationID string
	CreatedAt               time.Time
}

// CompanyConfigRepository handles persistence of company configuration in the database.
type CompanyConfigRepository interface {
	Create(ctx context.Context, config CompanyConfig) (CompanyConfig, error)
	FindByCompanyID(ctx context.Context, companyID string) (CompanyConfig, error)
	Delete(ctx context.Context, companyID string) error
}

// CompanyConfigService defines the business logic for managing company configurations.
type CompanyConfigService interface {
	CreateCompanyConfig(ctx context.Context, companyID string, name string) (tenantID string, err error)
	DeleteCompanyConfig(ctx context.Context, companyID string) error
}
