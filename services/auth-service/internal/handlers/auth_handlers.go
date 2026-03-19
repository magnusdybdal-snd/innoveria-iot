package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

func setRefreshCookie(w http.ResponseWriter, token string, ttl time.Duration) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/api/v1/auth/refresh",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(ttl.Seconds()),
	})
}

// PostLogin handles user login requests.
//
// @Summary Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "Login payload"
// @Success 201 {object} dto.LoginResult
// @Failure 400
// @Failure 500
// @Router /login [post]
func PostLogin(svc domain.AuthService, refreshTTL time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := json.Decode[dto.LoginRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.Email = strings.TrimSpace(payload.Email)
		payload.Password = strings.TrimSpace(payload.Password)

		loginResult, refreshToken, err := svc.Login(ctx, payload.Email, payload.Password)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrUnauthorized):
				json.HandleError(w, http.StatusUnauthorized, err, "invalid credentials")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		setRefreshCookie(w, refreshToken, refreshTTL)

		resp := dto.ToLoginResult(loginResult)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// PostRefresh refreshes an access token using the refresh_token cookie.
//
// @Summary Refresh access token
// @Description Validates the refresh_token cookie, rotates it, and returns a new access token.
// @Tags auth
// @Produce json
// @Success 200 {object} dto.LoginResult
// @Failure 401
// @Failure 500
// @Router /refresh [post]
func PostRefresh(svc domain.AuthService, refreshTTL time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c, err := r.Cookie("refresh_token")
		if err != nil || c.Value == "" {
			json.HandleError(w, http.StatusUnauthorized, domain.ErrUnauthorized, "missing refresh token")
			return
		}

		loginResult, newRawRefreshToken, err := svc.Refresh(ctx, c.Value)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "invalid refresh token")
			return
		}
		setRefreshCookie(w, newRawRefreshToken, refreshTTL)

		resp := dto.ToLoginResult(loginResult)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}
