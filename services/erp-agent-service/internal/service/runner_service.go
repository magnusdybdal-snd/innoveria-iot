// Package service orchestrates per-cycle Monitor to ERP sync.
package service

import (
	"context"
	"fmt"
	"net/url"

	"innoveria-iot/erp-agent-service/internal/domain"
	erpserviceclient "innoveria-iot/erp-agent-service/internal/erp-service-client"
	"innoveria-iot/erp-agent-service/internal/monitor"
	"innoveria-iot/erp-agent-service/internal/monitor/dto"
)

// RunnerServiceImpl coordinates endpoint fetches and ERP ingest posts.
type RunnerServiceImpl struct {
	monitorClient domain.MonitorHandler
	erpSvcClient  domain.ERPIngestClient
}

// New creates a runner service with monitor and ERP clients.
func New(monitor domain.MonitorHandler, erpSvc domain.ERPIngestClient) *RunnerServiceImpl {
	return &RunnerServiceImpl{
		monitorClient: monitor,
		erpSvcClient:  erpSvc,
	}
}

// RunCycle executes one sync cycle across configured endpoints.
func (r *RunnerServiceImpl) RunCycle(ctx context.Context) error {
	if err := syncRows[dto.ManufacturingOrderOperationReporting](
		ctx, r.monitorClient, r.erpSvcClient, monitor.OrderReportings, nil, erpserviceclient.OrderReportings,
	); err != nil {
		return err
	}

	if err := syncRows[dto.ManufacturingOrderOperation](
		ctx, r.monitorClient, r.erpSvcClient, monitor.OrderOperations, nil, erpserviceclient.OrderOperations,
	); err != nil {
		return err
	}

	if err := syncRows[dto.WorkCenter](
		ctx, r.monitorClient, r.erpSvcClient, monitor.Workcenters, nil, erpserviceclient.Workcenters,
	); err != nil {
		return err
	}

	return nil
}

func syncRows[T any](
	ctx context.Context,
	monitor domain.MonitorHandler,
	erpSvc domain.ERPIngestClient,
	monitorPath monitor.MonitorERPEndpoint,
	opts url.Values,
	erpPath erpserviceclient.Endpoint,
) error {
	rows := make([]T, 0)
	if err := monitor.Query(ctx, monitorPath, opts, &rows); err != nil {
		return fmt.Errorf("monitor query failed (%s): %w", monitorPath, err)
	}

	if len(rows) == 0 {
		return nil
	}

	if err := erpSvc.Post(ctx, erpPath, rows); err != nil {
		return fmt.Errorf("request to erp-service failed (%s): %w", erpPath, err)
	}
	return nil
}
