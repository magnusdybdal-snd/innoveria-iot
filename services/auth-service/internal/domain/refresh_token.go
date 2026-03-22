package domain

import (
	"context"
	"net/netip"
	"time"
)

// RefreshToken is the domain model for Refresh Token
type RefreshToken struct {
	ID         string
	UserID     string
	TokenHash  string
	ExpiresAt  time.Time
	RevokedAt  time.Time
	CreatedAt  time.Time
	DeviceInfo string
	IPAddress  *netip.Addr
}

// RefreshTokenRepo is the interface contract for refresh token repository
type RefreshTokenRepo interface {
	Create(ctx context.Context, token RefreshToken) error
	UpdateRefreshToken(ctx context.Context, token RefreshToken) error
	UpsertForLogin(ctx context.Context, token RefreshToken) error
	FindActiveByHash(ctx context.Context, oldHash string) (RefreshToken, error)
}
