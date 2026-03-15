package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	createFactoryAreaQuery = `
		INSERT INTO auth.factory_area (factory_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING area_id, factory_id, name, description, created_at, updated_at
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
	return nil, nil
}

// FindByID retrieves a factory area by id.
func (r *FactoryAreaRepoImpl) FindByID(ctx context.Context, areaID string) (domain.FactoryArea, error) {
	return domain.FactoryArea{}, nil
}

// DeleteByID deletes a factory area by id.
func (r *FactoryAreaRepoImpl) DeleteByID(ctx context.Context, areaID string) error {
	return nil
}
