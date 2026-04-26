package server

import (
	"errors"
	"net/http"

	"innoveria-iot/pkg/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

// RequireAgentAuth validates ERP agent bearer JWTs and injects a trusted
// company header for downstream handlers.
func RequireAgentAuth(jwtSecret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			return []byte(jwtSecret), nil
		}, jwt.WithIssuer("auth-service"))
		if err != nil || !token.Valid {
			json.HandleError(w, http.StatusUnauthorized, errors.New("invalid token"), "unauthorized")
			return
		}

		companyID, _ := claims["company_id"].(string)
		if companyID == "" {
			json.HandleError(w, http.StatusUnauthorized, errors.New("invalid company"), "unauthorized")
			return
		}

		// Dont trust any inbound header from agents
		r.Header.Del("X-Erp-Company-Id")

		// inject trusted header
		r.Header.Add("X-Erp-Company-Id", companyID)

		next.ServeHTTP(w, r)
	})
}
