// Package domain defines auth-service core interfaces and domain contracts.
package domain

// AuthRepository defines persistence operations needed by the auth domain.
type AuthRepository interface {
}

// AuthService defines authentication and session operations exposed by
// the auth domain service layer.
type AuthService interface {
	// TODO: add these methods
	// Login
	// Refresh
	// Logout
	// Me
}
