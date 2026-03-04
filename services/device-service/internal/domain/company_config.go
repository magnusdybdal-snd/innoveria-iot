// Package domain TODO(@Magnus Dybdal): add proper documentation.
package domain

import (
	"context"
	"time"
)

// CompanyConfig TODO(@Magnus Dybdal): add proper documentation.
type CompanyConfig struct {
	CompanyID               string
	ChirpstackTenantID      string
	ChirpstackApplicationID string
	CreatedAt               time.Time
}

// CompanyConfigRepository TODO(@Magnus Dybdal): add proper documentation.
type CompanyConfigRepository interface {
	Create(ctx context.Context, config CompanyConfig) (CompanyConfig, error)
	FindByCompanyID(ctx context.Context, companyID string) (CompanyConfig, error)
}
