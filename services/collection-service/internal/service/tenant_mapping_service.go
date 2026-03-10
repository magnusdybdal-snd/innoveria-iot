package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/collection-service/internal/domain"
)

// TenantMappingServiceImpl implements domain.TenantMappingService.
type TenantMappingServiceImpl struct {
	repo domain.TenantMappingRepository
}

// NewTenantMappingService creates a new TenantMappingServiceImpl.
func NewTenantMappingService(repo domain.TenantMappingRepository) *TenantMappingServiceImpl {
	return &TenantMappingServiceImpl{repo: repo}
}

// CreateTenantMapping stores a companyID-tenantID mapping in the database.
func (s *TenantMappingServiceImpl) CreateTenantMapping(ctx context.Context, companyID string, tenantID string) error {
	_, err := s.repo.Create(ctx, domain.TenantMapping{
		CompanyID:          companyID,
		ChirpstackTenantID: tenantID,
	})
	if err != nil {
		return fmt.Errorf("create tenant mapping: %w", err)
	}

	slog.Info("successfully created tenant mapping", "companyID", companyID)
	return nil
}

// DeleteTenantMapping removes a companyID-tenantID mapping from the database.
func (s *TenantMappingServiceImpl) DeleteTenantMapping(ctx context.Context, companyID string) error {
	if err := s.repo.Delete(ctx, companyID); err != nil {
		return fmt.Errorf("delete tenant mapping: %w", err)
	}

	slog.Info("successfully deleted tenant mapping", "companyID", companyID)
	return nil
}
