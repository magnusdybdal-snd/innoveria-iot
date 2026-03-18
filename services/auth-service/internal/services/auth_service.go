// Package services contains auth-service business logic.
package services

import (
	"context"
	"time"

	"innoveria-iot/auth-service/internal/domain"
)

// AuthServiceImpl implements authentication use cases for the auth service.
type AuthServiceImpl struct {
	userRepo  domain.UserRepo
	jwtSecret []byte // converted to byte in initializer
	jwtIssuer string
	accessTTL time.Duration
}

// NewAuthServiceImpl creates a new AuthServiceImpl instance.
func NewAuthServiceImpl(
	userRepo domain.UserRepo,
	jwtSecret string,
	jwtIssuer string,
	accessTTL time.Duration,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo:  userRepo,
		jwtSecret: []byte(jwtSecret),
		jwtIssuer: jwtIssuer,
		accessTTL: accessTTL,
	}
}

// Login authenticates a user.
func (s *AuthServiceImpl) Login(ctx context.Context, email, password string) (domain.LoginResult, error) {
	_, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return domain.LoginResult{}, err
	}
	return domain.LoginResult{}, nil
}

// Me returns the authenticated user profile.
func (s *AuthServiceImpl) Me(ctx context.Context, token string) {
	_, _ = ctx, token
}
