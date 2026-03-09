package clients

import (
	"context"
	"fmt"
	"net/http"
)

// CollectionClient is an HTTP client for the collection service.
type CollectionClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewCollectionClient creates a new CollectionClient targeting the given base URL.
func NewCollectionClient(baseURL string) *CollectionClient {
	return &CollectionClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// CreateCompanyConfig calls the collection service to store the companyID-tenantID mapping.
func (c *CollectionClient) CreateCompanyConfig(_ context.Context, _ string, _ string) error {
	// TODO: implement
	return fmt.Errorf("not implemented")
}
