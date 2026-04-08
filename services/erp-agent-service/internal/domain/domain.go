package domain

import (
	"context"
	"net/url"

	erpserviceclient "innoveria-iot/erp-agent-service/internal/erp-service-client"
	"innoveria-iot/erp-agent-service/internal/monitor"
)

type Runner interface {
	RunCycle(ctx context.Context) error
}

type MonitorHandler interface {
	Query(ctx context.Context, path monitor.MonitorERPEndpoint, opts url.Values, out any) error
}

type ERPIngestClient interface {
	Post(ctx context.Context, path erpserviceclient.Endpoint, body any) error
}
