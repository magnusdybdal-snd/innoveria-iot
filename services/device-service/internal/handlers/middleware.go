package handlers

import (
	"fmt"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/json"
	"net/http"
)

// AdminGuard wraps a handler and returns 401 if the request has no valid auth context, or 403 if the caller is not an admin.
func AdminGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}
		if !auth.IsAdmin() {
			json.HandleError(w, http.StatusForbidden, fmt.Errorf("forbidden"), "forbidden")
			return
		}
		next(w, r)
	}
}
