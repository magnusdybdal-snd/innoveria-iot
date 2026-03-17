package domain

import "errors"

var (
	// ErrCompanyNotFound triggers when there now company is found
	ErrCompanyNotFound = errors.New("company not found")
	// ErrFactoryNotFound trigger when there is no factory in db
	ErrFactoryNotFound = errors.New("factory not found")
	// ErrFactoryAreaNotFound trigger when there is no factory area in db
	ErrFactoryAreaNotFound = errors.New("factory area not found")
)
