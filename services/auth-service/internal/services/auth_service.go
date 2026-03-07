// Package services contains auth-service business logic and orchestrates
// authentication workflows between repositories and security components.
package services

// AuthServiceImpl implements authentication use cases for the auth service.
type AuthServiceImpl struct {
}

// NewAuthServiceImpl creates a new AuthServiceImpl instance.
func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}
