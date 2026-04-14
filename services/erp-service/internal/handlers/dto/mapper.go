package dto

import (
	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/monitor/dto"
)

func MapMonitorOrderOperationToDomain(from []dto.ManufacturingOrderOperation) []domain.OrderOperation {
	return []domain.OrderOperation{}
}

func MapMonitorOrderToDomain(from []dto.ManufacturingOrder) []domain.Order {
	return []domain.Order{}
}

func MapMonitorOrderReportToDomain(from []dto.ManufacturingOrderOperationReporting) []domain.OrderReport {
	return []domain.OrderReport{}
}

func MapMonitorWorkcenterToDomain(from []dto.WorkCenter) []domain.ProductionResource {
	return []domain.ProductionResource{}
}
