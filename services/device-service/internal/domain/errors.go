package domain

import "errors"

// ErrNotFound is returned when a requested resource does not exist in the database.
var ErrNotFound = errors.New("not found")

// ErrAlreadyExists is returned when attempting to create a resource that already exists.
var ErrAlreadyExists = errors.New("already exists")

// ErrInvalidMeasurementType is returned when a measurement type slug does not exist in the vocabulary.
var ErrInvalidMeasurementType = errors.New("invalid measurement type")

// ErrMissingMeasurementType is returned when a payload schema is saved without a measurement type.
var ErrMissingMeasurementType = errors.New("measurement type is required")
