package handlers

import (
	"errors"
	"net"
	"net/http"
	"net/netip"
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

func parseClientIP(r *http.Request) *netip.Addr {
	if xff := strings.TrimSpace(r.Header.Get("X-Client-IP")); xff != "" {
		first := strings.TrimSpace(strings.Split(xff, ",")[0])
		if ip, err := netip.ParseAddr(first); err == nil {
			return &ip
		}
	}

	if xrip := strings.TrimSpace(r.Header.Get("X-Real-IP")); xrip != "" {
		if ip, err := netip.ParseAddr(xrip); err == nil {
			return &ip
		}
	}

	// Fallback: gateway managed x-forwarded-for
	if xff := strings.TrimSpace(r.Header.Get("X-Forwarded-For")); xff != "" {
		first := strings.TrimSpace(strings.Split(xff, ",")[0])
		if ip, err := netip.ParseAddr(first); err == nil {
			return &ip
		}
	}

	// last fallback
	host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err == nil {
		if ip, err := netip.ParseAddr(host); err == nil {
			return &ip
		}
	}
	return nil
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

		ipAddress := parseClientIP(r)
		deviceInfo := strings.TrimSpace(r.UserAgent())

		loginResult, refreshToken, err := svc.Login(
			ctx,
			payload.Email,
			payload.Password,
			deviceInfo,
			ipAddress,
		)
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

// GetMe returns the authenticated user's profile.
//
// @Summary Get current user profile
// @Description Returns profile details for the authenticated user.
// @Tags auth
// @Produce json
// @Success 200 {object} dto.MeResponse
// @Failure 401
// @Failure 500
// @Router /me [get]
func GetMe(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := r.Header.Get("X-Auth-User-Id")
		if userID == "" {
			json.HandleError(w, http.StatusUnauthorized, errors.New("missing auth user id"), "unauthorized")
			return
		}

		user, err := svc.Me(ctx, userID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrUserNotFound):
				json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.ToMeResponse(user)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}
