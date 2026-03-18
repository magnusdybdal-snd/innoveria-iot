package dto

// LoginResult is the JSON response returned from a successful login.
type LoginResult struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}

// LoginRequest is the JSON payload accepted by the login endpoint.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"passwod"`
}

// MeResponse is the JSON response returned by the me endpoint.
type MeResponse struct {
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
