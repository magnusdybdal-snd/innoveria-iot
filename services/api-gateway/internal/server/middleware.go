// Package server TODO(@vinjar): add proper documentation.
package server

import (
	"errors"
	"net/http"
	"strings"

	"innoveria-iot/api-gateway/internal/config"
	"innoveria-iot/pkg/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

// corsMiddleware adds CORS headers for configured frontend origins.
func corsMiddleware(cfg *config.Config, next http.Handler) http.Handler {
	allowedOrigins := parseAllowedOrigins(cfg.CorsAllowedOrigins)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")
		w.Header().Add("Vary", "Access-Control-Request-Headers")

		if origin != "" && allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func parseAllowedOrigins(raw string) map[string]bool {
	allowed := make(map[string]bool)
	for _, origin := range strings.Split(raw, ",") {
		origin = strings.TrimSpace(origin)
		if origin == "" {
			continue
		}
		allowed[origin] = true
	}
	return allowed
}

// authMiddleware validates bearer JWTs and forwards trusted auth headers.
func authMiddleware(cfg *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if its a public path
		if isPublicPath(r) {
			next.ServeHTTP(w, r)
			return
		}

		tokenStr, err := request.AuthorizationHeaderExtractor.ExtractToken(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, errors.New("missing bearer token"), "unauthorized")
			return
		}
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("invalid signing method")
			}
			return []byte(cfg.JWTSecret), nil
		}, jwt.WithIssuer(cfg.JWTIssuer))
		if err != nil || !token.Valid {
			json.HandleError(w, http.StatusUnauthorized, errors.New("invalid token"), "unauthorized")
			return
		}

		userID, _ := claims["sub"].(string)
		companyID, _ := claims["company_id"].(string)
		role, _ := claims["role"].(string)
		if userID == "" {
			json.HandleError(w, http.StatusUnauthorized, errors.New("invalid subject"), "unauthorized")
			return
		}
		if companyID == "" {
			json.HandleError(w, http.StatusUnauthorized, errors.New("invalid company"), "unauthorized")
			return
		}
		if role == "" {
			json.HandleError(w, http.StatusUnauthorized, errors.New("invalid role"), "unauthorized")
			return
		}

		// Never trust inbound X-Auth-* from clients
		r.Header.Del("X-Auth-User-Id")
		r.Header.Del("X-Auth-Company-Id")
		r.Header.Del("X-Auth-Role")
		// Inject trusted identity for upstream services
		r.Header.Set("X-Auth-User-Id", userID)
		r.Header.Set("X-Auth-Company-Id", companyID)
		r.Header.Set("X-Auth-Role", role)

		next.ServeHTTP(w, r)
	})
}

// isPublicPath checks whether the request should bypass JWT auth.
func isPublicPath(r *http.Request) bool {
	p := r.URL.Path

	if r.Method == http.MethodOptions {
		return true
	}

	if p == "/swagger" || strings.HasPrefix(p, "/swagger/") {
		return true
	}

	if r.Method == http.MethodPost && p == AUTHENTICATION_ROUTE+"/login" {
		return true
	}

	if r.Method == http.MethodPost && p == AUTHENTICATION_ROUTE+"/refresh" {
		return true
	}

	if r.Method == http.MethodPost && p == AUTHENTICATION_ROUTE+"/logout" {
		return true
	}

	return false
}
