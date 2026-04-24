package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	createFactoryAreaQuery = `
		INSERT INTO auth.factory_area (factory_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING area_id, factory_id, name, description, created_at, updated_at
	`
	findAllFactoryAreaQuery = `
		SELECT fa.area_id, fa.factory_id, fa.name, fa.description, fa.created_at, fa.updated_at
		FROM auth.factory_area fa
		JOIN auth.factory f ON fa.factory_id = f.factory_id
		WHERE f.company_id = $1 AND fa.factory_id = $2
		ORDER BY fa.created_at ASC
	`
	findFactoryAreaByIDQuery = `
		SELECT fa.area_id, fa.factory_id, fa.name, fa.description, fa.created_at, fa.updated_at
		FROM auth.factory_area fa
		JOIN auth.factory f ON fa.factory_id = f.factory_id
		WHERE fa.area_id = $1 AND f.company_id = $2
	`
	deleteFactoryAreaByIDQuery = `
		DELETE FROM auth.factory_area
		USING auth.factory
		WHERE auth.factory_area.factory_id = auth.factory.factory_id
		  AND auth.factory_area.area_id = $1
		  AND auth.factory.company_id = $2
	`
)

// FactoryAreaRepoImpl is the implementation of
// auth domain factory area persistence operations.
type FactoryAreaRepoImpl struct {
	db *dbutil.DB
}

// NewFactoryAreaRepo initializes a repository for factory area persistence operations.
func NewFactoryAreaRepo(db *dbutil.DB) *FactoryAreaRepoImpl {
	return &FactoryAreaRepoImpl{db: db}
}

// Create inserts a new factory area.
func (r *FactoryAreaRepoImpl) Create(ctx context.Context, area domain.FactoryArea) (domain.FactoryArea, error) {
	var out domain.FactoryArea

	err := r.db.Pool.QueryRow(ctx, createFactoryAreaQuery,
		area.FactoryID,
		area.Name,
		area.Description,
	).Scan(
		&out.ID,
		&out.FactoryID,
		&out.Name,
		&out.Description,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			return domain.FactoryArea{}, fmt.Errorf("create factory area: %w", domain.ErrFactoryNotFound)
		}
		return domain.FactoryArea{}, fmt.Errorf("create factory area: %w", err)
	}

	return out, nil
}

// FindAll retrieves all factory areas belonging to a factory, scoped to the given company.
func (r *FactoryAreaRepoImpl) FindAll(ctx context.Context, companyID string, factoryID string) ([]domain.FactoryArea, error) {
	rows, err := r.db.Pool.Query(ctx, findAllFactoryAreaQuery, companyID, factoryID)
	if err != nil {
		return nil, fmt.Errorf("find all factory areas: %w", err)
	}
	defer rows.Close()

	var out []domain.FactoryArea

	for rows.Next() {
		var area domain.FactoryArea
		var desc sql.NullString
		if err := rows.Scan(
			&area.ID,
			&area.FactoryID,
			&area.Name,
			&desc,
			&area.CreatedAt,
			&area.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan factory area: %w", err)
		}

		if desc.Valid {
			area.Description = &desc.String
		}
		out = append(out, area)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate factory area: %w", err)
	}
	return out, nil
}

// FindByID retrieves a factory area by id, scoped to the given company via JOIN.
func (r *FactoryAreaRepoImpl) FindByID(ctx context.Context, companyID string, areaID string) (domain.FactoryArea, error) {
	var out domain.FactoryArea
	var description sql.NullString

	err := r.db.Pool.QueryRow(ctx, findFactoryAreaByIDQuery, areaID, companyID).Scan(
		&out.ID,
		&out.FactoryID,
		&out.Name,
		&description,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.FactoryArea{}, fmt.Errorf("find factory area by id: %w", domain.ErrFactoryAreaNotFound)
		}

		return domain.FactoryArea{}, fmt.Errorf("find factory area by id: %w", err)
	}

	if description.Valid {
		out.Description = &description.String
	}

	return out, nil
}

// Delete deletes a factory area by id, scoped to the given company via JOIN.
func (r *FactoryAreaRepoImpl) Delete(ctx context.Context, companyID string, areaID string) error {
	result, err := r.db.Pool.Exec(ctx, deleteFactoryAreaByIDQuery, areaID, companyID)
	if err != nil {
		return fmt.Errorf("delete factory area by id: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("delete factory area by id: %w", domain.ErrFactoryAreaNotFound)
	}

	return nil
}
