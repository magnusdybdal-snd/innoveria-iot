// Package authctx provides utilities for extracting authenticated user identity from HTTP request headers.
package authctx

import (
	"fmt"
	"net/http"

	"innoveria-iot/pkg/roles"
)

const (
	headerUserID    = "X-Auth-User-Id"
	headerCompanyID = "X-Auth-Company-Id"
	headerRole      = "X-Auth-Role"
)

// Auth holds the authenticated user's identity extracted from request headers.
// These headers are injected by the API gateway from the validated JWT — they can be trusted.
type Auth struct {
	UserID    string
	CompanyID string
	Role      roles.RoleType
}

// IsAdmin returns true if the user has the PLATFORM_ADMIN role.
func (a Auth) IsAdmin() bool {
	return a.Role == roles.PlatformAdmin
}

// FromRequest extracts the authenticated user's identity from the request headers.
// Returns an error if any header is missing, which should not happen for requests
// routed through the API gateway.
func FromRequest(r *http.Request) (Auth, error) {
	userID := r.Header.Get(headerUserID)
	if userID == "" {
		return Auth{}, fmt.Errorf("authctx: missing %s header", headerUserID)
	}

	companyID := r.Header.Get(headerCompanyID)
	if companyID == "" {
		return Auth{}, fmt.Errorf("authctx: missing %s header", headerCompanyID)
	}
	role := r.Header.Get(headerRole)
	if role == "" {
		return Auth{}, fmt.Errorf("authctx: missing %s header", headerRole)
	}

	return Auth{
		UserID:    userID,
		CompanyID: companyID,
		Role:      roles.RoleType(role),
	}, nil
}
