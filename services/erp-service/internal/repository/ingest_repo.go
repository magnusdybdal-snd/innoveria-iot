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

// Postgres error prefixes
const (
	pgClassDataException       = "22"
	pgClassIntegrityConstraint = "23"
)

const (
	orderColsPerRow              = 12
	orderOperationColsPerRow     = 11
	orderReportColsPerRow        = 10
	productionResourceColsPerRow = 6
	maxBindParametersPerStmt     = 65535
	defaultMaxRowsPerChunk       = 1000

	upsertOrderQuery = `
	INSERT INTO erp_raw."order" (
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
		received_at = EXCLUDED.received_at,
		sync_status = 'pending',
		sync_error = NULL,
		next_retry_at = NULL,
		last_synced_at = NULL,
		last_seen_at = now()
	`
	upsertOrderOperationQuery = `
	INSERT INTO erp_raw.order_operation (
		company_id,
		id,
		production_resource_id,
		order_id,
		planned_start_date,
		planned_finish_date,
		actual_start_date,
		actual_finish_date,
		status,
		production_resource_status,
		received_at
	)
	VALUES %s
	ON CONFLICT (company_id, id) DO UPDATE SET
		production_resource_id = EXCLUDED.production_resource_id,
		order_id = EXCLUDED.order_id,
		planned_start_date = EXCLUDED.planned_start_date,
		planned_finish_date = EXCLUDED.planned_finish_date,
		actual_start_date = EXCLUDED.actual_start_date,
		actual_finish_date = EXCLUDED.actual_finish_date,
		status = EXCLUDED.status,
		production_resource_status = EXCLUDED.production_resource_status,
		received_at = EXCLUDED.received_at,
		sync_status = 'pending',
		sync_error = NULL,
		next_retry_at = NULL,
		last_synced_at = NULL,
		last_seen_at = now()
	`
	upsertOrderReportQuery = `
	INSERT INTO erp_raw.order_report (
		company_id,
		id,
		order_operation_id,
		production_resource_id,
		quantity,
		rest_quantity,
		type,
		reporting_timestamp,
		actual_reported_date,
		received_at
	)
	VALUES %s
	ON CONFLICT (company_id, id) DO UPDATE SET
		order_operation_id = EXCLUDED.order_operation_id,
		production_resource_id = EXCLUDED.production_resource_id,
		quantity = EXCLUDED.quantity,
		rest_quantity = EXCLUDED.rest_quantity,
		type = EXCLUDED.type,
		reporting_timestamp = EXCLUDED.reporting_timestamp,
		actual_reported_date = EXCLUDED.actual_reported_date,
		received_at = EXCLUDED.received_at,
		sync_status = 'pending',
		sync_error = NULL,
		next_retry_at = NULL,
		last_synced_at = NULL,
		last_seen_at = now()
	`
	upsertProductionResourceQuery = `
	INSERT INTO erp_raw.production_resource (
		company_id,
		id,
		number,
		description,
		type,
		received_at
	)
	VALUES %s
	ON CONFLICT (company_id, id) DO UPDATE SET
		number = EXCLUDED.number,
		description = EXCLUDED.description,
		type = EXCLUDED.type,
		received_at = EXCLUDED.received_at,
		sync_status = 'pending',
		sync_error = NULL,
		next_retry_at = NULL,
		last_synced_at = NULL,
		last_seen_at = now()
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

	// This looks like a big bottleneck at first glance
	// Although its still O(n) (in total work)
	// Outerloop does chunck by chunck
	// Inner loop builds up the SQL placeholders, and insert after the loop
	rowsPerChunk := chunkSizeForColumns(orderColsPerRow)
	for start := 0; start < len(payload); start += rowsPerChunk {
		end := min(start+rowsPerChunk, len(payload))
		chunk := payload[start:end]

		args := make([]any, 0, len(chunk)*orderColsPerRow)

		chunkNo := (start / rowsPerChunk) + 1
		totalChunks := (len(payload) + rowsPerChunk - 1) / rowsPerChunk
		rowFrom := start + 1
		rowTo := end

		var values strings.Builder
		for i, order := range chunk {
			if i > 0 {
				values.WriteString(",")
			}

			argStart := i*orderColsPerRow + 1
			fmt.Fprintf(&values, "($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
				argStart,
				argStart+1,
				argStart+2,
				argStart+3,
				argStart+4,
				argStart+5,
				argStart+6,
				argStart+7,
				argStart+8,
				argStart+9,
				argStart+10,
				argStart+11,
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
			return fmt.Errorf("upsert to erp_raw.\"order\" failed (chunk %d/%d, rows %d-%d of %d): %w",
				chunkNo, totalChunks, rowFrom, rowTo, len(payload),
				errors.Join(mapPgError(err), err))
		}
	}

	return nil
}

// CreateOrderOperation stores a batch of order operations in raw ingest storage.
func (r *IngestRepoImpl) CreateOrderOperation(ctx context.Context, payload []domain.OrderOperation) error {
	if len(payload) == 0 {
		return nil
	}

	rowsPerChunk := chunkSizeForColumns(orderOperationColsPerRow)
	for start := 0; start < len(payload); start += rowsPerChunk {
		end := min(start+rowsPerChunk, len(payload))
		chunk := payload[start:end]

		args := make([]any, 0, len(chunk)*orderOperationColsPerRow)

		chunkNo := (start / rowsPerChunk) + 1
		totalChunks := (len(payload) + rowsPerChunk - 1) / rowsPerChunk
		rowFrom := start + 1
		rowTo := end

		var values strings.Builder
		for i, operation := range chunk {
			if i > 0 {
				values.WriteString(",")
			}

			argStart := i*orderOperationColsPerRow + 1
			fmt.Fprintf(&values, "($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
				argStart,
				argStart+1,
				argStart+2,
				argStart+3,
				argStart+4,
				argStart+5,
				argStart+6,
				argStart+7,
				argStart+8,
				argStart+9,
				argStart+10,
			)

			args = append(args,
				operation.CompanyID,
				operation.ID,
				operation.ProductionResourceID,
				operation.OrderID,
				operation.PlannedStartDate,
				operation.PlannedFinishDate,
				operation.ActualStartDate,
				operation.ActualFinishDate,
				string(operation.Status),
				string(operation.ProductionResourceStatus),
				operation.ReceivedAt,
			)
		}

		query := fmt.Sprintf(upsertOrderOperationQuery, values.String())

		_, err := r.db.Pool.Exec(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("upsert to erp_raw.order_operation failed (chunk %d/%d, rows %d-%d of %d): %w",
				chunkNo, totalChunks, rowFrom, rowTo, len(payload),
				errors.Join(mapPgError(err), err))
		}
	}

	return nil
}

// CreateOrderReport stores a batch of order reports in raw ingest storage.
func (r *IngestRepoImpl) CreateOrderReport(ctx context.Context, payload []domain.OrderReport) error {
	if len(payload) == 0 {
		return nil
	}

	rowsPerChunk := chunkSizeForColumns(orderReportColsPerRow)
	for start := 0; start < len(payload); start += rowsPerChunk {
		end := min(start+rowsPerChunk, len(payload))
		chunk := payload[start:end]

		args := make([]any, 0, len(chunk)*orderReportColsPerRow)

		chunkNo := (start / rowsPerChunk) + 1
		totalChunks := (len(payload) + rowsPerChunk - 1) / rowsPerChunk
		rowFrom := start + 1
		rowTo := end

		var values strings.Builder
		for i, report := range chunk {
			if i > 0 {
				values.WriteString(",")
			}

			argStart := i*orderReportColsPerRow + 1
			fmt.Fprintf(&values, "($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
				argStart,
				argStart+1,
				argStart+2,
				argStart+3,
				argStart+4,
				argStart+5,
				argStart+6,
				argStart+7,
				argStart+8,
				argStart+9,
			)

			args = append(args,
				report.CompanyID,
				report.ID,
				report.OrderOperationID,
				report.ProductionResourceID,
				report.Quantity,
				report.RestQuantity,
				string(report.Type),
				report.ReportingTimestamp,
				report.ActualReportedDate,
				report.ReceivedAt,
			)
		}

		query := fmt.Sprintf(upsertOrderReportQuery, values.String())

		_, err := r.db.Pool.Exec(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("upsert to erp_raw.order_report failed (chunk %d/%d, rows %d-%d of %d): %w",
				chunkNo, totalChunks, rowFrom, rowTo, len(payload),
				errors.Join(mapPgError(err), err))
		}
	}

	return nil
}

// CreateProductionResource stores a batch of production resources in raw ingest storage.
func (r *IngestRepoImpl) CreateProductionResource(ctx context.Context, payload []domain.ProductionResource) error {
	if len(payload) == 0 {
		return nil
	}

	rowsPerChunk := chunkSizeForColumns(productionResourceColsPerRow)
	for start := 0; start < len(payload); start += rowsPerChunk {
		end := min(start+rowsPerChunk, len(payload))
		chunk := payload[start:end]

		args := make([]any, 0, len(chunk)*productionResourceColsPerRow)

		chunkNo := (start / rowsPerChunk) + 1
		totalChunks := (len(payload) + rowsPerChunk - 1) / rowsPerChunk
		rowFrom := start + 1
		rowTo := end

		var values strings.Builder
		for i, resource := range chunk {
			if i > 0 {
				values.WriteString(",")
			}

			argStart := i*productionResourceColsPerRow + 1
			fmt.Fprintf(&values, "($%d,$%d,$%d,$%d,$%d,$%d)",
				argStart,
				argStart+1,
				argStart+2,
				argStart+3,
				argStart+4,
				argStart+5,
			)

			args = append(args,
				resource.CompanyID,
				resource.ID,
				resource.Number,
				resource.Description,
				string(resource.Type),
				resource.ReceivedAt,
			)
		}

		query := fmt.Sprintf(upsertProductionResourceQuery, values.String())

		_, err := r.db.Pool.Exec(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("upsert to erp_raw.production_resource failed (chunk %d/%d, rows %d-%d of %d): %w",
				chunkNo, totalChunks, rowFrom, rowTo, len(payload),
				errors.Join(mapPgError(err), err))
		}
	}

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

func chunkSizeForColumns(columnsPerRow int) int {
	maxRowsByBindLimit := maxBindParametersPerStmt / columnsPerRow
	return min(maxRowsByBindLimit, defaultMaxRowsPerChunk)
}
