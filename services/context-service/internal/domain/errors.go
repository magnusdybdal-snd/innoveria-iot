package domain

import "errors"

// ErrNotFound is returned when a requested resource does not exist.
var ErrNotFound = errors.New("not found")

// ErrDatabase is returned when an unexpected database error occurs.
var ErrDatabase = errors.New("database error")

// ErrConflict is returned when a resource already exists and cannot be duplicated.
var ErrConflict = errors.New("conflict")
