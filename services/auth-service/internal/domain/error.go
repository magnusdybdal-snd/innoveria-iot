package domain

import "errors"

var (
	// ErrCompanyNotFound triggers when there is a sql violation for company id in factory
	ErrCompanyNotFound = errors.New("company not found")
	// ErrFactoryNotFound trigger when there is no factory in db
	ErrFactoryNotFound = errors.New("factory not found")
)
