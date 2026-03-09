// Package repository implements the persistence layer for the device service.
package repository

import (
	"context"
	"fmt"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

const (
	createCompanyConfigQuery = `
		INSERT INTO device.company_config (company_id, chirpstack_tenant_id, chirpstack_application_id)
		VALUES ($1, $2, $3)
		RETURNING company_id, chirpstack_tenant_id, chirpstack_application_id, created_at
	`

	findCompanyConfigByCompanyIDQuery = `
		SELECT company_id, chirpstack_tenant_id, chirpstack_application_id, created_at
		FROM device.company_config
		WHERE company_id = $1
	`

	deleteCompanyConfigQuery = `
		DELETE FROM device.company_config
		WHERE company_id = $1
	`
)

// CompanyConfigRepository handles persistance of company configuration in the database.
// It stores Chirpstack tenant and application ID's that are created during onboarding for each company.
type CompanyConfigRepository struct {
	db *dbutil.DB
}

// NewCompanyConfigRepository creates a new CompanyConfigRepository with the given databse
func NewCompanyConfigRepository(db *dbutil.DB) *CompanyConfigRepository {
	return &CompanyConfigRepository{db: db}
}

// Create inserts a new company config into the databse and returns the newly inserted row.
// Called once during onboarding when the Chirpstack tenant and application ID are created.
// Returns the newly inserted company config or an error if failed.
func (r *CompanyConfigRepository) Create(ctx context.Context, config domain.CompanyConfig) (domain.CompanyConfig, error) {

	var out domain.CompanyConfig
	err := r.db.Pool.QueryRow(ctx, createCompanyConfigQuery,
		config.CompanyID,
		config.ChirpstackTenantID,
		config.ChirpstackApplicationID,
	).Scan(
		&out.CompanyID,
		&out.ChirpstackTenantID,
		&out.ChirpstackApplicationID,
		&out.CreatedAt,
	)
	if err != nil {
		return domain.CompanyConfig{}, fmt.Errorf("create company config: %w", err)
	}

	return out, nil
}

// FindByCompanyID retrieves the company config for the given company ID.
// Returns an error of no config exist for the given company ID.
func (r *CompanyConfigRepository) FindByCompanyID(ctx context.Context, companyID string) (domain.CompanyConfig, error) {

	var out domain.CompanyConfig
	err := r.db.Pool.QueryRow(ctx, findCompanyConfigByCompanyIDQuery, companyID).Scan(
		&out.CompanyID,
		&out.ChirpstackTenantID,
		&out.ChirpstackApplicationID,
		&out.CreatedAt,
	)
	if err != nil {
		return domain.CompanyConfig{}, fmt.Errorf("find company config by id: %w", err)
	}

	return out, nil
}

// Delete tries to delete a company config mapping from the database.
// Returns an error if deletion fails or no company config is found.
func (r *CompanyConfigRepository) Delete(ctx context.Context, companyID string) error {

	tag, err := r.db.Pool.Exec(ctx, deleteCompanyConfigQuery, companyID)
	if err != nil {
		return fmt.Errorf("delete company config %s: %w", companyID, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("company not found: %s", companyID)
	}

	return nil
}
