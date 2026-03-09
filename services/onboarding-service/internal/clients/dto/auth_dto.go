// Package dto contains request and response models for downstream service clients.
package dto

// CreateCompanyRequest is the request body for POST /auth/companies.
type CreateCompanyRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// CreateCompanyResponse is the response body from POST /auth/companies.
type CreateCompanyResponse struct {
	CompanyID string `json:"company_id"`
}
