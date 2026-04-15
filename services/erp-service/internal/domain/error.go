package domain

import "errors"

var (
	// ErrDatabase is returned when an unexpected database error occurs.
	ErrDatabase = errors.New("database error")
	// ErrConflict is returned when a write conflicts with existing data.
	ErrConflict = errors.New("conflict")
	// ErrInvalidInput is returned when data cannot be persisted due to invalid values.
	ErrInvalidInput = errors.New("invalid input")
)
