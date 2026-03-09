package domain

import (
	"context"
	"time"
)

// TenantMapping holds the mapping between a Chirpstack tenant ID and an internal company ID.
type TenantMapping struct {
	ChirpstackTenantID string
	CompanyID          string
	CreatedAt          time.Time
}

// TenantMappingRepository handles persistence of tenant mappings in the database.
type TenantMappingRepository interface {
	Create(ctx context.Context, mapping TenantMapping) (TenantMapping, error)
	Delete(ctx context.Context, companyID string) error
}

// TenantMappingService defines the business logic for managing tenant mappings.
type TenantMappingService interface {
	CreateTenantMapping(ctx context.Context, companyID string, tenantID string) error
	DeleteTenantMapping(ctx context.Context, companyID string) error
}
