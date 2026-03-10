package repository

import (
	"context"
	"fmt"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

const (
	createFactoryQuery = `
		INSERT INTO auth.factory (company_id, name, address)
		VALUES ($1, $2, $3)
		RETURNING factory_id, company_id, name, COALESCE(address, ''), created_at::text, updated_at::text
	`
	findFactoryByIDQuery = `
		SELECT factory_id, company_id, name, COALESCE(address, ''), created_at::text, updated_at::text
		FROM auth.factory
		WHERE factory_id = $1
	`
	findAllFactoriesQuery = `
		SELECT factory_id, company_id, name, COALESCE(address, ''), created_at::text, updated_at::text
		FROM auth.factory
		ORDER BY created_at ASC
	`
)

// FactoryRepoImpl is the  implementation of
// auth domain factory persistence operations.
type FactoryRepoImpl struct {
	db *dbutil.DB
}

// NewFactoryRepo initializes a new factory repository.
func NewFactoryRepo(db *dbutil.DB) *FactoryRepoImpl {
	return &FactoryRepoImpl{db: db}
}

// Create inserts a new factory and generates a factory id.
func (r *FactoryRepoImpl) Create(ctx context.Context, factory domain.Factory) (domain.Factory, error) {
	var out domain.Factory
	err := r.db.Pool.QueryRow(ctx, createFactoryQuery,
		factory.CompanyID,
		factory.Name,
		factory.Address,
	).Scan(
		&out.Id,
		&out.CompanyID,
		&out.Name,
		&out.Address,
		&out.Created_at,
		&out.Updated_at,
	)
	if err != nil {
		return domain.Factory{}, fmt.Errorf("create factory: %w", err)
	}

	return out, nil
}

// FindAll retrieves all factories.
func (r *FactoryRepoImpl) FindAll(ctx context.Context) ([]domain.Factory, error) {
	rows, err := r.db.Pool.Query(ctx, findAllFactoriesQuery)
	if err != nil {
		return nil, fmt.Errorf("find all factories: %w", err)
	}
	defer rows.Close()

	var out []domain.Factory

	for rows.Next() {
		var factory domain.Factory
		if err := rows.Scan(
			&factory.Id,
			&factory.CompanyID,
			&factory.Name,
			&factory.Address,
			&factory.Created_at,
			&factory.Updated_at,
		); err != nil {
			return nil, fmt.Errorf("scan factory: %w", err)
		}
		out = append(out, factory)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil
}

// FindByID retrieves a factory by id.
func (r *FactoryRepoImpl) FindByID(ctx context.Context, factoryID string) (domain.Factory, error) {
	var out domain.Factory
	err := r.db.Pool.QueryRow(ctx, findFactoryByIDQuery, factoryID).Scan(
		&out.Id,
		&out.CompanyID,
		&out.Name,
		&out.Address,
		&out.Created_at,
		&out.Updated_at,
	)
	if err != nil {
		return out, fmt.Errorf("find factory by id: %w", err)
	}

	return out, nil
}
