// Package dto contains transport-to-domain mapping helpers for ERP handlers.
package dto

import (
	"log/slog"
	"strconv"
	"time"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/monitor/dto"
)

// MapMonitorOrderOperationToDomain maps Monitor order operations into domain entities.
func MapMonitorOrderOperationToDomain(from []dto.ManufacturingOrderOperation) []domain.OrderOperation {
	to := make([]domain.OrderOperation, 0, len(from))
	receivedAt := time.Now().UTC()

	for _, item := range from {
		to = append(to, domain.OrderOperation{
			CompanyID:                "a0000000-0000-0000-0000-000000000001", // TODO: Change this. Its hardcoded to innoveria for now, which ofc is bad
			ID:                       item.ID,
			ProductionResourceID:     item.WorkCenterId,
			OrderID:                  item.ManufacturingOrderId,
			PlannedStartDate:         item.PlannedStartDate,
			PlannedFinishDate:        item.PlannedFinishDate,
			ActualStartDate:          item.ActualStartDate,
			ActualFinishDate:         item.ActualFinishDate,
			Status:                   mapOperationStatus(item.Status),
			ProductionResourceStatus: mapOperationStatus(item.WorkshopOperationStatus),
			ReceivedAt:               receivedAt,
		})
	}

	return to
}

// MapMonitorOrderToDomain maps Monitor manufacturing orders into domain entities.
func MapMonitorOrderToDomain(from []dto.ManufacturingOrder) []domain.Order {
	to := make([]domain.Order, 0, len(from))
	receivedAt := time.Now().UTC()

	for _, item := range from {
		to = append(to, domain.Order{
			CompanyID:         "a0000000-0000-0000-0000-000000000001", // TODO: change this
			ID:                item.ID,
			OrderNumber:       item.OrderNumber,
			PartID:            item.PartID,
			PartDescription:   item.PartDescription,
			PlannedStartDate:  item.PlannedStartDate,
			PlannedFinishDate: item.PlannedFinishDate,
			ActualStartDate:   item.ActualStartDate,
			ActualFinishDate:  item.ActualFinishDate,
			Status:            mapOrderStatus(item.Status),
			Priority:          item.Priority,
			ReceivedAt:        receivedAt,
		})
	}

	return to
}

// MapMonitorOrderReportToDomain maps Monitor operation reportings into domain entities.
func MapMonitorOrderReportToDomain(from []dto.ManufacturingOrderOperationReporting) []domain.OrderReport {
	to := make([]domain.OrderReport, 0, len(from))
	receivedAt := time.Now().UTC()

	for _, item := range from {
		to = append(to, domain.OrderReport{
			CompanyID:            "a0000000-0000-0000-0000-000000000001", // TODO: change this
			ID:                   item.ID,
			OrderOperationID:     item.OperationId,
			ProductionResourceID: item.WorkCenterId,
			Quantity:             item.Quantity,
			RestQuantity:         item.RestQuantity,
			Type:                 mapOrderReportType(item.Type),
			ReportingTimestamp:   item.ReportingTimestamp,
			ActualReportedDate:   item.ActualReportedDate,
			ReceivedAt:           receivedAt,
		})
	}

	return to
}

// MapMonitorWorkcenterToDomain maps Monitor work centers into production resources.
func MapMonitorWorkcenterToDomain(from []dto.WorkCenter) []domain.ProductionResource {
	to := make([]domain.ProductionResource, 0, len(from))
	receivedAt := time.Now().UTC()

	for _, item := range from {
		description := item.Description
		to = append(to, domain.ProductionResource{
			CompanyID:   "a0000000-0000-0000-0000-000000000001",
			ID:          item.ID,
			Number:      item.Number,
			Description: description,
			Type:        mapWorkCenterType(item.Type),
			ReceivedAt:  receivedAt,
		})
	}

	return to
}

// mapOrderStatus converts Monitor order status values to domain order statuses.
func mapOrderStatus(status int) domain.OrderStatus {
	switch status {
	case 0:
		return domain.OrderStatusNotInitialized
	case 1:
		return domain.OrderStatusRegistered
	case 2:
		return domain.OrderStatusPrinted
	case 3:
		return domain.OrderStatusStarted
	case 4:
		return domain.OrderStatusFinished
	case 5:
		return domain.OrderStatusPostCalculated
	case 6:
		return domain.OrderStatusDelivered
	case 7:
		return domain.OrderStatusHistorical
	default:
		slog.Warn("unknown monitor order status, preserving raw value", "status", status)
		return domain.OrderStatus(strconv.Itoa(status))
	}
}

// mapOperationStatus converts Monitor operation status values to domain operation statuses.
func mapOperationStatus(status int) domain.OperationStatus {
	switch status {
	case 1:
		return domain.OperationStatusStarted
	case 2:
		return domain.OperationStatusPartiallyShipped
	case 3:
		return domain.OperationStatusFullyShipped
	case 4:
		return domain.OperationStatusPartiallyReported
	case 5:
		return domain.OperationStatusFinished
	default:
		slog.Warn("unknown monitor operation status, preserving raw value", "status", status)
		return domain.OperationStatus(strconv.Itoa(status))
	}
}

// mapOrderReportType converts Monitor report type values to domain report types.
func mapOrderReportType(reportType int) domain.OrderReportType {
	switch reportType {
	case 1:
		return domain.OrderReportTypeSendToSubcontractor
	case 2:
		return domain.OrderReportTypeReceiveFromSubcontractor
	case 3:
		return domain.OrderReportTypeCancelRest
	case 4:
		return domain.OrderReportTypeMaterialOnly
	case 5:
		return domain.OrderReportTypeSubcontractorInvoicePrice
	case 6:
		return domain.OrderReportTypeRecordingTerminal
	case 7:
		return domain.OrderReportTypeAdjustRecording
	case 8:
		return domain.OrderReportTypeUndoRegular
	case 9:
		return domain.OrderReportTypeUndoRecordingTerminal
	case 10:
		return domain.OrderReportTypeUndoAdjustRecording
	case 11:
		return domain.OrderReportTypeSubcontractorPosteriorReport
	case 12:
		return domain.OrderReportTypeSubcontractorCostsManual
	case 13:
		return domain.OrderReportTypePickWorkCenter
	default:
		slog.Warn("unknown monitor order report type, preserving raw value", "type", reportType)
		return domain.OrderReportType(strconv.Itoa(reportType))
	}
}

// mapWorkCenterType converts Monitor work center type values to domain resource types.
func mapWorkCenterType(workCenterType int) domain.WorkCenterType {
	switch workCenterType {
	case 1:
		return domain.WorkCenterTypeMachine
	case 2:
		return domain.WorkCenterTypeManualWork
	case 3:
		return domain.WorkCenterTypeSubContract
	case 4:
		return domain.WorkCenterTypePool
	case 5:
		return domain.WorkCenterTypePick
	default:
		slog.Warn("unknown monitor work center type, preserving raw value", "type", workCenterType)
		return domain.WorkCenterType(strconv.Itoa(workCenterType))
	}
}
