package clients

import (
	"context"
	"fmt"
	"net/http"

	"innoveria-iot/onboarding-service/internal/clients/dto"
	"innoveria-iot/pkg/httpclient"
)

// CollectionClient is an HTTP client for the collection service.
type CollectionClient struct {
	baseURL string
	client  *httpclient.Client
}

// NewCollectionClient creates a new CollectionClient targeting the given base URL.
func NewCollectionClient(baseURL string) *CollectionClient {
	return &CollectionClient{
		baseURL: baseURL,
		client:  httpclient.New(),
	}
}

// CreateCompanyConfig calls the collection service to store the companyID-tenantID mapping.
func (c *CollectionClient) CreateCompanyConfig(ctx context.Context, companyID string, tenantID string) error {
	resp, err := httpclient.DoRaw(
		c.client,
		ctx,
		c.baseURL+"/api/v1/collection/company-config",
		http.MethodPost,
		dto.CreateTenantMappingRequest{CompanyID: companyID, TenantID: tenantID},
		nil,
	)
	if err != nil {
		return fmt.Errorf("collection client create company config: %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}
