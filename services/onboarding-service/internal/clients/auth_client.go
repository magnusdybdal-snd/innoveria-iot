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
	resp, err := httpclient.DoRequest[dto.AuthCreateCompanyResponse](
		c.client,
		ctx,
		fmt.Sprintf("%s/api/v1/auth/companies", c.baseURL),
		http.MethodPost,
		dto.AuthCreateCompanyRequest{Name: company.Name, Address: company.Address},
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("auth client create company: %w", err)
	}

	return resp.CompanyID, nil
}

// DeleteCompany calls the auth service to delete a company by ID (compensating transaction).
func (c *AuthClient) DeleteCompany(ctx context.Context, companyID string) error {
	_, err := httpclient.DoRaw(
		c.client,
		ctx,
		fmt.Sprintf("%s/api/v1/auth/companies/%s", c.baseURL, companyID),
		http.MethodDelete,
		nil,
		nil,
	)
	if err != nil {
		return fmt.Errorf("auth client delete company: %w", err)
	}

	return nil
}
