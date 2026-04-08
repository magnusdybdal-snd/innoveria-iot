// Package domain defines service interfaces used across layers.
package domain

import (
	"context"
	"net/url"

	erpserviceclient "innoveria-iot/erp-agent-service/internal/erp-service-client"
	"innoveria-iot/erp-agent-service/internal/monitor"
)

// Runner executes one polling and ingest cycle.
type Runner interface {
	RunCycle(ctx context.Context) error
}

// MonitorHandler reads data from Monitor ERP endpoints.
type MonitorHandler interface {
	Query(ctx context.Context, path monitor.MonitorERPEndpoint, opts url.Values, out any) error
}

// ERPIngestClient sends payloads to ERP service ingest endpoints.
type ERPIngestClient interface {
	Post(ctx context.Context, path erpserviceclient.Endpoint, body any) error
}
