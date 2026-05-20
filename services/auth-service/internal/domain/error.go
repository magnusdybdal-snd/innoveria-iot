package domain

import "errors"

var (
	// ErrCompanyNotFound triggers when there now company is found
	ErrCompanyNotFound = errors.New("company not found")
	// ErrFactoryNotFound trigger when there is no factory in db
	ErrFactoryNotFound = errors.New("factory not found")
	// ErrFactoryAreaNotFound trigger when there is no factory area in db
	ErrFactoryAreaNotFound = errors.New("factory area not found")
	// ErrUnauthorized trigger when there is no valid token
	ErrUnauthorized = errors.New("unauthorized")
	// ErrUserNotFound trigger when there is no user in db
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExists triggers when a user with the same email already exists
	ErrUserAlreadyExists = errors.New("user already exists")
)
