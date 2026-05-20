// Package domain defines auth-service core interfaces and domain contracts.
package domain

import (
	"context"
	"net/netip"
	"time"
)

// LoginResult contains token metadata returned by a successful login.
type LoginResult struct {
	AccessToken string
	TokenType   string
	ExpiresIn   time.Duration
}

// AuthService defines authentication and session operations exposed by
// the auth domain service layer.
//
// Login, Refresh, Logout and Me.
type AuthService interface {
	// TODO: Add Register
	Login(ctx context.Context, email, password, deviceInfo string, ip *netip.Addr) (LoginResult, string, error)
	Refresh(ctx context.Context, rawRefreshToken string) (LoginResult, string, error)
	Logout(ctx context.Context, rawRefreshToken string) error
	Me(ctx context.Context, token string) (User, error)
}
