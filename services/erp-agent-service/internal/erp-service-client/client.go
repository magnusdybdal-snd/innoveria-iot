// Package erpserviceclient provides a minimal ERP ingest client.
package erpserviceclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"innoveria-iot/erp-agent-service/internal/domain"
	"innoveria-iot/pkg/httpclient"
)

// Client posts payloads to ERP service ingest endpoints.
type Client struct {
	baseURL    string
	httpClient *httpclient.Client
	apiKey     string
}

// Endpoint is an ERP service API path.
type Endpoint = domain.ERPEndpoint

// ERP service ingest endpoints.
const (
	// Index is the API root prefix.
	Index Endpoint = "/api/v1"
	// Svc is the ERP service base route.
	Svc Endpoint = Index + "/erp"
	// Ingest is the ingest route used by agent services.
	Ingest Endpoint = Svc + "/ingest"
	// Order is the ingest endpoint for orders
	Order Endpoint = Ingest + "/orders"
	// OrderOperations is the ingest endpoint for manufacturing order operations.
	OrderOperations Endpoint = Ingest + "/order-operations"
	// OrderReportings is the ingest endpoint for operation reportings.
	OrderReportings Endpoint = Ingest + "/order-reportings"
	// Workcenters is the ingest endpoint for work center master data.
	Workcenters Endpoint = Ingest + "/workcenters"
)

// New builds an ERP ingest client for the given base URL.
func New(baseURL string, apiKey string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpclient.New(),
		apiKey:     apiKey,
	}
}

func (c *Client) authHeader() string {
	return "Bearer " + c.apiKey
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
		map[string]string{
			"Authorization": c.authHeader(),
		},
	)
	if err != nil {
		var httpErr *httpclient.HTTPError
		if errors.As(err, &httpErr) && httpErr.StatusCode == http.StatusUnauthorized {
			return fmt.Errorf("%w: %s", domain.ErrERPUnauthorized, httpErr.Error())
		}
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
