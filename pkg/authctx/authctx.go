// Package authctx provides utilities for extracting authenticated user identity from HTTP request headers.
package authctx

import (
	"fmt"
	"net/http"
)

const (
	headerUserID    = "X-Auth-User-Id"
	headerCompanyID = "X-Auth-Company-Id"
	// TODO(RBAC): extract headerRole = "X-Auth-Role" here when implementing role-based access control.
	// The API gateway will need to inject this header from the JWT "role" claim alongside the existing headers.
)

// Auth holds the authenticated user's identity extracted from request headers.
// These headers are injected by the API gateway from the validated JWT — they can be trusted.
//
// TODO(RBAC): add Role string field here to carry the user's role (e.g. PLATFORM_ADMIN, USER).
type Auth struct {
	UserID    string
	CompanyID string
}

// FromRequest extracts the authenticated user's identity from the request headers.
// Returns an error if either header is missing, which should not happen for requests
// routed through the API gateway.
//
// TODO(RBAC): extract the X-Auth-Role header here and populate Auth.Role.
func FromRequest(r *http.Request) (Auth, error) {
	userID := r.Header.Get(headerUserID)
	if userID == "" {
		return Auth{}, fmt.Errorf("authctx: missing %s header", headerUserID)
	}

	companyID := r.Header.Get(headerCompanyID)
	if companyID == "" {
		return Auth{}, fmt.Errorf("authctx: missing %s header", headerCompanyID)
	}

	return Auth{
		UserID:    userID,
		CompanyID: companyID,
	}, nil
}
