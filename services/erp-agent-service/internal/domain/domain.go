package domain

import (
	"context"
	erpserviceclient "innoveria-iot/erp-agent-service/internal/erp-service-client"
	"net/url"
)

type Runner interface {
	RunCycle(ctx context.Context) error
}

type MonitorHandler interface {
	Query(ctx context.Context, path string, opts url.Values, out any) error
}

type ERPIngestClient interface {
	Ingest(ctx context.Context, path erpserviceclient.Endpoint) error
}
