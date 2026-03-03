package domain

import (
	"context"
	"time"
)

type CompanyConfig struct {
	CompanyID               string
	ChirpstackTenantID      string
	ChirpstackApplicationID string
	CreatedAt               time.Time
}

type CompanyConfigRepository interface {
	Create(ctx context.Context, config CompanyConfig) (CompanyConfig, error)
	FindByCompanyID(ctx context.Context, companyID string) (CompanyConfig, error)
}
