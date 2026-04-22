package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"innoveria-iot/erp-service/internal/domain"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// WrapMappedDBError maps database errors to domain errors.
func WrapMappedDBError(op string, err error) error {
	mapped := mapPgError(err)

	if errors.Is(mapped, context.Canceled) || errors.Is(mapped, context.DeadlineExceeded) {
		return fmt.Errorf("%s: %w", op, mapped)
	}

	return fmt.Errorf("%s: %w", op, mapped)
}

func mapPgError(err error) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrNotFound
	}

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
