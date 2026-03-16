// Package domain defines auth-service core interfaces and domain contracts.
package domain

import "context"

// AuthService defines authentication and session operations exposed by
// the auth domain service layer.
//
// Login, Refresh, Logout and Me.
type AuthService interface {
	// TODO: Add Register
	Login(ctx context.Context, email, password string)
	Me(ctx context.Context, token string)
}
