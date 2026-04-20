package handlers

import (
	"context"
	"errors"
	"net/http"

	"innoveria-iot/erp-service/internal/domain"
)

// MapIngestDomainError maps domain ingestion errors to an HTTP status code,
// public message, and wrapped error for handler responses.
func MapIngestDomainError(err error) (int, string, error) {
	if errors.Is(err, domain.ErrInvalidInput) {
		return http.StatusBadRequest, "bad request", err
	}

	if errors.Is(err, domain.ErrConflict) {
		return http.StatusConflict, "conflict", err
	}

	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return http.StatusServiceUnavailable, "request canceled", err
	}

	return http.StatusInternalServerError, "internal server error", err
}
