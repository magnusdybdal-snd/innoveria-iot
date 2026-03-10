package repository

import (
	"context"
	"fmt"

	"innoveria-iot/collection-service/internal/db"
	"innoveria-iot/collection-service/internal/domain"
)

const (
	createTenantMappingQuery = `
		INSERT INTO collection.tenant_mapping (chirpstack_tenant_id, company_id)
		VALUES ($1, $2)
		RETURNING chirpstack_tenant_id, company_id, created_at
	`

	deleteTenantMappingQuery = `
		DELETE FROM collection.tenant_mapping
		WHERE company_id = $1
	`
)

// TenantMappingRepository handles persistence of tenant mappings in the database.
type TenantMappingRepository struct {
	db *db.DB
}

// NewTenantMappingRepository creates a new TenantMappingRepository with the given database.
func NewTenantMappingRepository(db *db.DB) *TenantMappingRepository {
	return &TenantMappingRepository{db: db}
}

// Create inserts a new tenant mapping into the database and returns the inserted row.
func (r *TenantMappingRepository) Create(ctx context.Context, mapping domain.TenantMapping) (domain.TenantMapping, error) {
	var out domain.TenantMapping
	err := r.db.Pool.QueryRow(ctx, createTenantMappingQuery,
		mapping.ChirpstackTenantID,
		mapping.CompanyID,
	).Scan(
		&out.ChirpstackTenantID,
		&out.CompanyID,
		&out.CreatedAt,
	)
	if err != nil {
		return domain.TenantMapping{}, fmt.Errorf("create tenant mapping: %w", err)
	}

	return out, nil
}

// Delete removes a tenant mapping by company ID.
func (r *TenantMappingRepository) Delete(ctx context.Context, companyID string) error {
	tag, err := r.db.Pool.Exec(ctx, deleteTenantMappingQuery, companyID)
	if err != nil {
		return fmt.Errorf("delete tenant mapping %s: %w", companyID, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("delete tenant mapping: %w", domain.ErrNotFound)
	}

	return nil
}
