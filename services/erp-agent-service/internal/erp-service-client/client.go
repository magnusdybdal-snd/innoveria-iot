// Package erpserviceclient provides a minimal ERP ingest client.
package erpserviceclient

import (
	"context"
	"io"
	"net/http"

	"innoveria-iot/erp-agent-service/internal/domain"
	"innoveria-iot/pkg/httpclient"
)

// Client posts payloads to ERP service ingest endpoints.
type Client struct {
	baseURL    string
	httpClient *httpclient.Client
	// apiKey string // TODO: add this when ready
}

// Endpoint is an ERP service API path.
type Endpoint = domain.ERPEndpoint

// ERP service ingest endpoints.
const (
	// Index is the API root prefix.
	Index Endpoint = "/api/v1"
	// Svc is the ERP service base route.
	Svc Endpoint = Index + "/erp" // TODO: change this?
	// Ingest is the ingest route used by agent services.
	Ingest Endpoint = Svc + "/ingest"
	// OrderOperations is the ingest endpoint for manufacturing order operations.
	OrderOperations Endpoint = Ingest + "/order-operations"
	// OrderReportings is the ingest endpoint for operation reportings.
	OrderReportings Endpoint = Ingest + "/order-reportings"
	// Workcenters is the ingest endpoint for work center master data.
	Workcenters Endpoint = Ingest + "/workcenters"
)

// New builds an ERP ingest client for the given base URL.
func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpclient.New(),
	}
}

// Post sends a JSON payload to the selected ERP endpoint.
func (c *Client) Post(ctx context.Context, path Endpoint, body any) error {
	url := c.baseURL + string(path)
	resp, err := httpclient.DoRaw(
		c.httpClient,
		ctx,
		url,
		http.MethodPost,
		body,
		nil, // TODO: jwt token
	)
	if err != nil {
		return err
	}

	// Drain tcp connection
	_, copyErr := io.Copy(io.Discard, resp.Body)
	// Close response body
	closeErr := resp.Body.Close()

	// handle drain error
	if copyErr != nil {
		return copyErr
	}
	// handle close response body error
	if closeErr != nil {
		return closeErr
	}
	return nil
}
