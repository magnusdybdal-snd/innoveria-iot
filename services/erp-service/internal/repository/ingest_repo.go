// Package repository provides persistence implementations for ERP ingest writes.
package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	pgClassDataException       = "22"
	pgClassIntegrityConstraint = "23"
)

const (
	orderColsPerRow  = 12
	upsertOrderQuery = `
	INSERT INTO erp."order" (
		company_id,
		id,
		order_number,
		part_id,
		part_description,
		planned_start_date,
		planned_finish_date,
		actual_start_date,
		actual_finish_date,
		status,
		priority,
		received_at
	)
	VALUES %s
	ON CONFLICT (company_id, id) DO UPDATE SET
		order_number = EXCLUDED.order_number,
		part_id = EXCLUDED.part_id,
		part_description = EXCLUDED.part_description,
		planned_start_date = EXCLUDED.planned_start_date,
		planned_finish_date = EXCLUDED.planned_finish_date,
		actual_start_date = EXCLUDED.actual_start_date,
		actual_finish_date = EXCLUDED.actual_finish_date,
		status = EXCLUDED.status,
		priority = EXCLUDED.priority,
		received_at = EXCLUDED.received_at
	`
)

// IngestRepoImpl is the implementaion of ingest persistant storage write
// Which is responsible for writing erp data from a erp agent service
type IngestRepoImpl struct {
	db *dbutil.DB
}

// NewIngestRepo initializes a new Ingest repository.
func NewIngestRepo(db *dbutil.DB) *IngestRepoImpl {
	return &IngestRepoImpl{
		db: db,
	}
}

// CreateOrder upserts a batch of orders.
// Postgres allows at most 65535 bind parameters per statement.
// With 12 columns per row, this statement supports up to 5461 rows per batch.
func (r *IngestRepoImpl) CreateOrder(ctx context.Context, payload []domain.Order) error {
	if len(payload) == 0 {
		return nil
	}

	args := make([]any, 0, len(payload)*orderColsPerRow)

	var values strings.Builder
	for i, order := range payload {
		if i > 0 {
			values.WriteString(",")
		}

		start := i*orderColsPerRow + 1
		fmt.Fprintf(&values, "($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			start,
			start+1,
			start+2,
			start+3,
			start+4,
			start+5,
			start+6,
			start+7,
			start+8,
			start+9,
			start+10,
			start+11,
		)

		args = append(args,
			order.CompanyID,
			order.ID,
			order.OrderNumber,
			order.PartID,
			order.PartDescription,
			order.PlannedStartDate,
			order.PlannedFinishDate,
			order.ActualStartDate,
			order.ActualFinishDate,
			string(order.Status),
			order.Priority,
			order.ReceivedAt,
		)
	}

	query := fmt.Sprintf(upsertOrderQuery, values.String())

	_, err := r.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("upsert order batch: %w", errors.Join(mapPgError(err), err))
	}

	return nil
}

// CreateOrderOperation stores a batch of order operations.
func (r *IngestRepoImpl) CreateOrderOperation(ctx context.Context) error {
	return nil
}

// CreateOrderReport stores a batch of order reports.
func (r *IngestRepoImpl) CreateOrderReport(ctx context.Context) error {
	return nil
}

// CreateProductionResource stores a batch of production resources.
func (r *IngestRepoImpl) CreateProductionResource(ctx context.Context) error {
	return nil
}

func mapPgError(err error) error {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return domain.ErrDatabase
	}

	switch pgErr.Code {
	case pgerrcode.UniqueViolation, pgerrcode.ExclusionViolation:
		return domain.ErrConflict
	case pgerrcode.NotNullViolation,
		pgerrcode.ForeignKeyViolation,
		pgerrcode.CheckViolation,
		pgerrcode.InvalidTextRepresentation,
		pgerrcode.InvalidDatetimeFormat,
		pgerrcode.NumericValueOutOfRange:
		return domain.ErrInvalidInput
	default:
		if strings.HasPrefix(pgErr.Code, pgClassDataException) {
			return domain.ErrInvalidInput
		}
		if strings.HasPrefix(pgErr.Code, pgClassIntegrityConstraint) {
			return domain.ErrConflict
		}
		return domain.ErrDatabase
	}
}
