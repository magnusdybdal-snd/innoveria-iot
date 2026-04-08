// Package service
package service

import (
	"context"
	"fmt"

	"innoveria-iot/erp-agent-service/internal/domain"
	erpserviceclient "innoveria-iot/erp-agent-service/internal/erp-service-client"
	"innoveria-iot/erp-agent-service/internal/monitor/dto"
)

type RunnerServiceImpl struct {
	monitorClient domain.MonitorHandler
	erpSvcClient  domain.ERPIngestClient
}

func New(monitor domain.MonitorHandler, erpSvc domain.ERPIngestClient) *RunnerServiceImpl {
	return &RunnerServiceImpl{
		monitorClient: monitor,
		erpSvcClient:  erpSvc,
	}
}

func (r *RunnerServiceImpl) RunCycle(ctx context.Context) error {
	var rows []dto.ManufacturingOrderOperationReporting
	if err := r.monitorClient.Query(ctx, "test", nil, rows); err != nil {
		return fmt.Errorf("monitor query failed: %w", err)
	}

	if len(rows) == 0 {
		return nil
	}

	if err := r.erpSvcClient.Post(ctx, erpserviceclient.OrderReportings, rows); err != nil {
		return fmt.Errorf("post to erp-service failed: %w", err)
	}
	return nil
}
