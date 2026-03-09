// Package clients provides HTTP clients for downstream services.
package clients

import (
	"context"
	"fmt"
	"net/http"

	"innoveria-iot/onboarding-service/internal/domain"
)

// AuthClient is an HTTP client for the auth service.
type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewAuthClient creates a new AuthClient targeting the given base URL.
func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// CreateCompany calls the auth service to create a new company and returns the assigned companyID.
func (c *AuthClient) CreateCompany(_ context.Context, _ domain.Company) (string, error) {
	// TODO: implement
	return "", fmt.Errorf("not implemented")
}

// DeleteCompany calls the auth service to delete a company by ID (compensating transaction).
func (c *AuthClient) DeleteCompany(_ context.Context, _ string) error {
	// TODO: implement
	return fmt.Errorf("not implemented")
}
