// Package dto contains request and response models for downstream service clients.
package dto

// AuthCreateCompanyRequest is the request body for POST /auth/companies.
type AuthCreateCompanyRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// AuthCreateCompanyResponse is the response body from POST /auth/companies.
type AuthCreateCompanyResponse struct {
	CompanyID string `json:"id"`
}
