package domain

import "errors"

// ErrNotFound is returned when a requested resource does not exist in the database.
var ErrNotFound = errors.New("not found")

// ErrAlreadyExists is returned when attempting to create a resource that already exists.
var ErrAlreadyExists = errors.New("already exists")
