package dto

import (
	"strconv"
	"time"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/monitor/dto"
)

func MapMonitorOrderOperationToDomain(from []dto.ManufacturingOrderOperation) []domain.OrderOperation {
	to := make([]domain.OrderOperation, 0, len(from))
	receivedAt := time.Now().UTC()

	for _, item := range from {
		to = append(to, domain.OrderOperation{
			ID:                       item.ID,
			ProductionResourceID:     item.WorkCenterId,
			OrderId:                  item.ManufacturingOrderId,
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

func MapMonitorOrderToDomain(from []dto.ManufacturingOrder) []domain.Order {
	to := make([]domain.Order, 0, len(from))
	receivedAt := time.Now().UTC()

	for _, item := range from {
		to = append(to, domain.Order{
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

func MapMonitorOrderReportToDomain(from []dto.ManufacturingOrderOperationReporting) []domain.OrderReport {
	to := make([]domain.OrderReport, 0, len(from))
	receivedAt := time.Now().UTC()

	for _, item := range from {
		to = append(to, domain.OrderReport{
			ID:                   item.ID,
			OrderOperationID:     item.OperationId,
			ProductionResourceID: strconv.FormatInt(item.WorkCenterId, 10),
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

func MapMonitorWorkcenterToDomain(from []dto.WorkCenter) []domain.ProductionResource {
	to := make([]domain.ProductionResource, 0, len(from))
	receivedAt := time.Now().UTC()

	for _, item := range from {
		description := item.Description
		to = append(to, domain.ProductionResource{
			ID:          item.ID,
			Number:      item.Number,
			Description: &description,
			Type:        mapWorkCenterType(item.Type),
			ReceivedAt:  receivedAt,
		})
	}

	return to
}

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
		return domain.OrderStatusNotInitialized
	}
}

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
		return domain.OperationStatusNone
	}
}

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
		return domain.OrderReportTypeRegular
	}
}

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
		return domain.WorkCenterTypeMachine
	}
}
