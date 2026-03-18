// Package domain defines auth-service core interfaces and domain contracts.
package domain

import "context"

// LoginResult contains token metadata returned by a successful login.
type LoginResult struct {
	AccessToken string
	TokenType   string
	ExpiresIn   string
}

// AuthService defines authentication and session operations exposed by
// the auth domain service layer.
//
// Login, Refresh, Logout and Me.
type AuthService interface {
	// TODO: Add Register
	Login(ctx context.Context, email, password string) (LoginResult, error)
	Me(ctx context.Context, token string)
}
