package server

import (
	"errors"
	"net/http"

	"innoveria-iot/pkg/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

// SeededCompanyID is a dev only const used to insert correct erp data with the dev company
const seededCompanyID = "a0000000-0000-0000-0000-000000000001"

// RequireAgentAuth validates ERP agent bearer JWTs and injects a trusted
// company header for downstream handlers.
func RequireAgentAuth(jwtSecret string, goEnv string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := request.AuthorizationHeaderExtractor.ExtractToken(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, errors.New("missing bearer token"), "unauthorized")
			return
		}

		// dev mode
		if tokenStr == "dev" {
			if goEnv != "development" {
				json.HandleError(w, http.StatusUnauthorized, errors.New("invalid token"), "unauthorized")
				return
			}

			r.Header.Del("X-Erp-Company-Id")
			r.Header.Add("X-Erp-Company-Id", seededCompanyID)
			next.ServeHTTP(w, r)
			return
		}

		// Prod mode
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
		if companyID == "" && goEnv == "development" {
			companyID = seededCompanyID
		}
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
