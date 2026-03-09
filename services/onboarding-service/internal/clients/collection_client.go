package clients

import (
	"context"
	"fmt"

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
func (c *CollectionClient) CreateCompanyConfig(_ context.Context, _ string, _ string) error {
	// TODO: implement
	return fmt.Errorf("not implemented")
}
