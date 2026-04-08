// Package erpserviceclient provides a minimal ERP ingest client.
package erpserviceclient

import (
	"context"
	"net/http"

	"innoveria-iot/pkg/httpclient"
)

// Client posts payloads to ERP service ingest endpoints.
type Client struct {
	baseURL    string
	httpClient *httpclient.Client
	// apiKey string // TODO: add this when ready
}

// Endpoint is an ERP service API path.
type Endpoint string

// ERP service ingest endpoints.
const (
	// Index is the API root prefix.
	Index Endpoint = "/api/v1"
	// Svc is the ERP service base route.
	Svc Endpoint = Index + "/erp" // TODO: change this?
	// Ingest is the ingest route used by agent services.
	Ingest Endpoint = Svc + "/ingest"
	// Orders is the ingest endpoint for manufacturing order operations.
	OrderOperations Endpoint = Ingest + "/orders-operations"
	// OrderReportings is the ingest endpoint for operation reportings.
	OrderReportings Endpoint = Ingest + "/order-reportings"
	// Workcenters is the ingest endpoint for work center master data.
	Workcenters Endpoint = Ingest + "/workcenter"
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

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}
