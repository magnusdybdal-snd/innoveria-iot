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
		SELECT area_id, factory_id, name, description, created_at, updated_at
		FROM auth.factory_area
		ORDER BY created_at ASC
	`
	findFactoryAreaByIDQuery = `
		SELECT area_id, factory_id, name, description, created_at, updated_at
		FROM auth.factory_area
		WHERE area_id = $1
	`
	deleteFactoryAreaByIDQuery = `
		DELETE FROM auth.factory_area
		WHERE area_id = $1
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
	var description any
	if area.Description != nil {
		description = *area.Description
	}

	var outDescription sql.NullString
	err := r.db.Pool.QueryRow(ctx, createFactoryAreaQuery,
		area.FactoryID,
		area.Name,
		description,
	).Scan(
		&out.ID,
		&out.FactoryID,
		&out.Name,
		&outDescription,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			return domain.FactoryArea{}, fmt.Errorf("create factory area : %w", domain.ErrFactoryNotFound)
		}
		return domain.FactoryArea{}, fmt.Errorf("create factory area: %w", err)
	}

	if outDescription.Valid {
		out.Description = &outDescription.String
	}

	return out, nil
}

// FindAll retrieves all factory areas.
func (r *FactoryAreaRepoImpl) FindAll(ctx context.Context) ([]domain.FactoryArea, error) {
	rows, err := r.db.Pool.Query(ctx, findAllFactoryAreaQuery)
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

// FindByID retrieves a factory area by id.
func (r *FactoryAreaRepoImpl) FindByID(ctx context.Context, areaID string) (domain.FactoryArea, error) {
	var out domain.FactoryArea
	var description sql.NullString

	err := r.db.Pool.QueryRow(ctx, findFactoryAreaByIDQuery, areaID).Scan(
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

// DeleteByID deletes a factory area by id.
func (r *FactoryAreaRepoImpl) DeleteByID(ctx context.Context, areaID string) error {
	result, err := r.db.Pool.Exec(ctx, deleteFactoryAreaByIDQuery, areaID)
	if err != nil {
		return fmt.Errorf("delete factory area by id: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("delete factory area by id: %w", domain.ErrFactoryAreaNotFound)
	}

	return nil
}
