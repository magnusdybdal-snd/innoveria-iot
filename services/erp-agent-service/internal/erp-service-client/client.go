package erpserviceclient

import (
	"context"
	"net/http"

	"innoveria-iot/pkg/httpclient"
)

type Client struct {
	baseURL    string
	httpClient *httpclient.Client
	// apiKey string // TODO: add this when ready
}

// API endpoints for the erp-svc
type Endpoint string

const (
	Index           Endpoint = "/api/v1"
	Svc             Endpoint = Index + "/erp"  // TODO: change this?
	Ingest          Endpoint = Svc + "/ingest" // Endpoint for agent service
	Orders          Endpoint = Ingest + "/orders"
	OrderReportings Endpoint = Ingest + "/order-reportings"
	Workcenters     Endpoint = Ingest + "/workcenter"
)

func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpclient.New(),
	}
}

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
