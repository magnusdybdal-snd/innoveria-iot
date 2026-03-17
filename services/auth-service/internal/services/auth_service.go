// Package services contains auth-service business logic.
package services

import "context"

// AuthServiceImpl implements authentication use cases for the auth service.
type AuthServiceImpl struct{}

// NewAuthServiceImpl creates a new AuthServiceImpl instance.
func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

// Login authenticates a user.
func (s *AuthServiceImpl) Login(ctx context.Context, email, password string) {
	_, _, _ = ctx, email, password
}

// Me returns the authenticated user profile.
func (s *AuthServiceImpl) Me(ctx context.Context, token string) {
	_, _ = ctx, token
}
