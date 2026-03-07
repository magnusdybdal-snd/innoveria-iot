// Package services contains auth-service business logic and orchestrates
// authentication workflows between repositories and security components.
package services

import "context"

// AuthServiceImpl implements authentication use cases for the auth service.
type AuthServiceImpl struct {
}

// NewAuthServiceImpl creates a new AuthServiceImpl instance.
func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

// RegisterCompany generates a new company and starts onboarding on device-service
func (a *AuthServiceImpl) RegisterCompany(ctx context.Context) error {
	return nil
}
