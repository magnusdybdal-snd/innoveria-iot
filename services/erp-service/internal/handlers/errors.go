package handlers

import (
	"context"
	"errors"
	"net/http"

	"innoveria-iot/erp-service/internal/domain"
)

// MapDomainError maps domain errors to an HTTP status code,
// public message, and wrapped error for handler responses.
func MapDomainError(err error) (int, string, error) {
	if errors.Is(err, domain.ErrInvalidInput) {
		return http.StatusBadRequest, "bad request", err
	}

	if errors.Is(err, domain.ErrNotFound) {
		return http.StatusNotFound, "resource not found", err
	}

	if errors.Is(err, domain.ErrConflict) {
		return http.StatusConflict, "conflict", err
	}

	if errors.Is(err, context.Canceled) {
		return http.StatusServiceUnavailable, "request canceled", err
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return http.StatusGatewayTimeout, "request timeout", err
	}

	return http.StatusInternalServerError, "internal server error", err
}
