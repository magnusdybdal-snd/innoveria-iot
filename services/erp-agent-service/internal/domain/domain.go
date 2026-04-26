// Package domain defines service interfaces used across layers.
package domain

import (
	"context"
	"errors"
	"net/url"
)

var (
	// ErrERPUnauthorized marks a non-retriable auth failure to ERP service.
	ErrERPUnauthorized = errors.New("erp service unauthorized")
)

// MonitorERPEndpoint is a relative Monitor ERP API endpoint path.
type MonitorERPEndpoint string

// ERPEndpoint is an ERP service API path.
type ERPEndpoint string

// Runner executes one polling and ingest cycle.
type Runner interface {
	RunCycle(ctx context.Context) error
}

// MonitorHandler reads data from Monitor ERP endpoints.
type MonitorHandler interface {
	Query(ctx context.Context, path MonitorERPEndpoint, opts url.Values, out any) error
}

// ERPIngestClient sends payloads to ERP service ingest endpoints.
type ERPIngestClient interface {
	Post(ctx context.Context, path ERPEndpoint, body any) error
}
