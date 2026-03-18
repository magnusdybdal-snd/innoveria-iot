package dto

import (
	"time"

	"innoveria-iot/auth-service/internal/domain"
)

// LoginResult is the JSON response returned from a successful login.
type LoginResult struct {
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	ExpiresIn   time.Duration `json:"expires_in"`
}

// ToLoginResult maps a domain login result to a login response DTO.
func ToLoginResult(from domain.LoginResult) LoginResult {
	return LoginResult{
		AccessToken: from.AccessToken,
		TokenType:   from.TokenType,
		ExpiresIn:   time.Duration(from.ExpiresIn.Minutes()),
	}
}

// LoginRequest is the JSON payload accepted by the login endpoint.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// MeResponse is the JSON response returned by the me endpoint.
type MeResponse struct {
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
