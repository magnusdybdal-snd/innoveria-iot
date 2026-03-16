package repository

import (
	"context"
	"errors"
	"fmt"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn" // for error handling
)

const (
	createFactoryQuery = `
		INSERT INTO auth.factory (company_id, name, address)
		VALUES ($1, $2, $3)
		RETURNING factory_id, company_id, name, address, created_at, updated_at
	`
	findFactoryByIDQuery = `
		SELECT factory_id, company_id, name, address, created_at, updated_at
		FROM auth.factory
		WHERE factory_id = $1
	`
	findAllFactoriesQuery = `
		SELECT factory_id, company_id, name, address, created_at, updated_at
		FROM auth.factory
		ORDER BY created_at ASC
	`
	deleteFactoryByIDQuery = `
		DELETE FROM auth.factory
		WHERE factory_id = $1
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
		&out.ID,
		&out.CompanyID,
		&out.Name,
		&out.Address,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			return domain.Factory{}, fmt.Errorf("create factory: %w", domain.ErrCompanyNotFound)
		}
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
			&factory.ID,
			&factory.CompanyID,
			&factory.Name,
			&factory.Address,
			&factory.CreatedAt,
			&factory.UpdatedAt,
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
		&out.ID,
		&out.CompanyID,
		&out.Name,
		&out.Address,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// domain not found, code: 404
			return domain.Factory{}, fmt.Errorf("find factory by id: %w", domain.ErrFactoryNotFound)
		}
		// Internal server error, code: 500
		return domain.Factory{}, fmt.Errorf("find factory by id: %w", err)
	}

	return out, nil
}

// DeleteByID deletes a factory by id.
func (r *FactoryRepoImpl) DeleteByID(ctx context.Context, factoryID string) error {
	result, err := r.db.Pool.Exec(ctx, deleteFactoryByIDQuery, factoryID)

	// Internal server error, code: 500
	if err != nil {
		return fmt.Errorf("delete factory by id: %w", err)
	}

	// Domain not found, code: 404
	if result.RowsAffected() == 0 {
		return fmt.Errorf("delete factory by id: %w", domain.ErrFactoryNotFound)
	}

	return nil
}
