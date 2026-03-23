package repository

import (
	"context"
	"errors"
	"fmt"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	createMeasurementTypeQuery = `
		INSERT INTO device.measurement_type (slug, display_name, description, default_unit)
		VALUES ($1, $2, $3, $4)
	`

	findAllMeasurementTypesQuery = `
		SELECT slug, display_name, description, default_unit, deprecated
		FROM device.measurement_type
		ORDER BY slug ASC
	`

	deprecateMeasurementTypeQuery = `
		UPDATE device.measurement_type
		SET deprecated = true
		WHERE slug = $1
	`
)

// MeasurementTypeRepository handles persistence of measurement types in the database.
type MeasurementTypeRepository struct {
	db *dbutil.DB
}

// NewMeasurementTypeRepository creates a new MeasurementTypeRepository with the given database.
func NewMeasurementTypeRepository(db *dbutil.DB) *MeasurementTypeRepository {
	return &MeasurementTypeRepository{db: db}
}

// Create inserts a new measurement type into the database.
// Returns domain.ErrAlreadyExists if a type with the same slug already exists.
func (r *MeasurementTypeRepository) Create(ctx context.Context, m domain.MeasurementType) error {
	_, err := r.db.Pool.Exec(ctx, createMeasurementTypeQuery, m.Slug, m.DisplayName, m.Description, m.DefaultUnit)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return domain.ErrAlreadyExists
		}
		return fmt.Errorf("create measurement type: %w", err)
	}
	return nil
}

// FindAll retrieves all measurement types ordered by slug.
func (r *MeasurementTypeRepository) FindAll(ctx context.Context) ([]domain.MeasurementType, error) {
	rows, err := r.db.Pool.Query(ctx, findAllMeasurementTypesQuery)
	if err != nil {
		return nil, fmt.Errorf("find all measurement types: %w", err)
	}
	defer rows.Close()

	out := []domain.MeasurementType{}

	for rows.Next() {
		var m domain.MeasurementType
		err := rows.Scan(&m.Slug, &m.DisplayName, &m.Description, &m.DefaultUnit, &m.Deprecated)
		if err != nil {
			return nil, fmt.Errorf("scan measurement type: %w", err)
		}
		out = append(out, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil
}

// Deprecate marks a measurement type as deprecated by its slug.
// Returns domain.ErrNotFound if no measurement type with the given slug exists.
func (r *MeasurementTypeRepository) Deprecate(ctx context.Context, slug string) error {
	tag, err := r.db.Pool.Exec(ctx, deprecateMeasurementTypeQuery, slug)
	if err != nil {
		return fmt.Errorf("deprecate measurement type: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}
