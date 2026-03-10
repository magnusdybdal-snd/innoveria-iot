package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

// CompanyConfigServiceImpl implements domain.CompanyConfigService.
type CompanyConfigServiceImpl struct {
	cc   *chirpstackrest.Client
	repo domain.CompanyConfigRepository
}

// NewCompanyConfigService creates a new CompanyConfigServiceImpl.
func NewCompanyConfigService(cc *chirpstackrest.Client, repo domain.CompanyConfigRepository) *CompanyConfigServiceImpl {
	return &CompanyConfigServiceImpl{
		cc:   cc,
		repo: repo,
	}
}

// CreateCompanyConfig runs the company config creation SAGA:
//  1. Create a Chirpstack tenant → tenantID
//  2. Create a Chirpstack application under that tenant → applicationID
//  3. Store {companyID, tenantID, applicationID} in the database
//
// Returns the tenantID so the onboarding service can pass it to the collection service.
// Compensating transactions are run on failure to keep Chirpstack and the database in sync.
func (s *CompanyConfigServiceImpl) CreateCompanyConfig(ctx context.Context, companyID string, name string) (string, error) {
	// Step 1: create Chirpstack tenant
	tenantID, err := s.cc.CreateTenant(ctx, mappers.MapCreateChirpstackTenant(name))
	if err != nil {
		return "", fmt.Errorf("create company config: create chirpstack tenant: %w", err)
	}

	// Step 2: create Chirpstack application under the new tenant
	applicationID, err := s.cc.CreateApplication(ctx, mappers.MapCreateChirpstackApplication(name, tenantID))
	if err != nil {
		// Compensate: delete the tenant we just created.
		// Use WithoutCancel so compensation runs even if the request context is already cancelled.
		compCtx := context.WithoutCancel(ctx)
		if compErr := s.cc.DeleteTenant(compCtx, tenantID); compErr != nil {
			slog.Error("saga compensation failed: could not delete chirpstack tenant after application creation failure",
				"tenantID", tenantID, "error", compErr)
		}
		return "", fmt.Errorf("create company config: create chirpstack application: %w", err)
	}

	// Step 3: store mapping in database
	_, err = s.repo.Create(ctx, domain.CompanyConfig{
		CompanyID:               companyID,
		ChirpstackTenantID:      tenantID,
		ChirpstackApplicationID: applicationID,
	})
	if err != nil {
		// Compensate: delete application and tenant from Chirpstack.
		// Use WithoutCancel so compensation runs even if the request context is already cancelled.
		compCtx := context.WithoutCancel(ctx)
		if compErr := s.cc.DeleteApplication(compCtx, applicationID); compErr != nil {
			slog.Error("saga compensation failed: could not delete chirpstack application after db insert failure",
				"applicationID", applicationID, "error", compErr)
		}
		if compErr := s.cc.DeleteTenant(compCtx, tenantID); compErr != nil {
			slog.Error("saga compensation failed: could not delete chirpstack tenant after db insert failure",
				"tenantID", tenantID, "error", compErr)
		}
		return "", fmt.Errorf("create company config: store in database: %w", err)
	}

	slog.Info("successfully created company config", "companyID", companyID)
	return tenantID, nil
}

// DeleteCompanyConfig removes a company's Chirpstack tenant and application and their database record.
// Called as a compensating transaction by the onboarding service if a later SAGA step fails.
// The database record is deleted first as it is the source of truth — Chirpstack cleanup is best-effort.
func (s *CompanyConfigServiceImpl) DeleteCompanyConfig(ctx context.Context, companyID string) error {
	// Fetch config to get the Chirpstack IDs before deleting
	cfg, err := s.repo.FindByCompanyID(ctx, companyID)
	if err != nil {
		return fmt.Errorf("delete company config: find config: %w", err)
	}

	// Delete from DB first — source of truth
	if err := s.repo.Delete(ctx, companyID); err != nil {
		return fmt.Errorf("delete company config: delete from database: %w", err)
	}

	// Best-effort Chirpstack cleanup — application must be deleted before tenant.
	// Use WithoutCancel so cleanup runs even if the request context is already cancelled.
	compCtx := context.WithoutCancel(ctx)
	chirpstackCleanupOK := true
	if err := s.cc.DeleteApplication(compCtx, cfg.ChirpstackApplicationID); err != nil {
		chirpstackCleanupOK = false
		slog.Error("failed to delete chirpstack application after db delete",
			"applicationID", cfg.ChirpstackApplicationID, "error", err)
	}
	if err := s.cc.DeleteTenant(compCtx, cfg.ChirpstackTenantID); err != nil {
		chirpstackCleanupOK = false
		slog.Error("failed to delete chirpstack tenant after db delete",
			"tenantID", cfg.ChirpstackTenantID, "error", err)
	}

	if chirpstackCleanupOK {
		slog.Info("successfully deleted company config", "companyID", companyID)
	} else {
		slog.Warn("deleted company config from database; chirpstack cleanup partially failed — manual remediation required",
			"companyID", companyID)
	}
	return nil
}
