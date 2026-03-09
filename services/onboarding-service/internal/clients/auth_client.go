// Package clients provides HTTP clients for downstream services.
package clients

import (
	"context"
	"fmt"
	"net/http"

	"innoveria-iot/onboarding-service/internal/clients/dto"
	"innoveria-iot/onboarding-service/internal/domain"
	"innoveria-iot/pkg/httpclient"
)

// AuthClient is an HTTP client for the auth service.
type AuthClient struct {
	baseURL string
	client  *httpclient.Client
}

// NewAuthClient creates a new AuthClient targeting the given base URL.
func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		baseURL: baseURL,
		client:  httpclient.New(),
	}
}

// CreateCompany calls the auth service to create a new company and returns the assigned companyID.
func (c *AuthClient) CreateCompany(ctx context.Context, company domain.Company) (string, error) {
	resp, err := httpclient.DoRequest[dto.CreateCompanyResponse](
		c.client,
		ctx,
		c.baseURL+"/api/v1/auth/companies",
		http.MethodPost,
		dto.CreateCompanyRequest{Name: company.Name, Address: company.Address},
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("auth client create company: %w", err)
	}

	return resp.CompanyID, nil
}

// DeleteCompany calls the auth service to delete a company by ID (compensating transaction).
// TODO: not yet implemented in auth-service.
func (c *AuthClient) DeleteCompany(_ context.Context, _ string) error {
	return fmt.Errorf("not implemented")
}
