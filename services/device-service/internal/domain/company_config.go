// Package domain defines the core business entities, interfaces, and types for the device service.
package domain

import (
	"context"
	"time"
)

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
}
