package dto

import (
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// ProductionResourceResponse represents a work center in the API response.
type ProductionResourceResponse struct {
	ID          int64   `json:"id"`
	Number      string  `json:"number"`
	Description *string `json:"description"`
	Type        string  `json:"type"`
}

// OrderOperationResponse represents a single manufacturing operation in the API response.
type OrderOperationResponse struct {
	ID                       int64                      `json:"id"`
	ProductionResource       ProductionResourceResponse `json:"production_resource"`
	PlannedStartDate         time.Time                  `json:"planned_start_date"`
	PlannedFinishDate        time.Time                  `json:"planned_finish_date"`
	ActualStartDate          *time.Time                 `json:"actual_start_date"`
	ActualFinishDate         *time.Time                 `json:"actual_finish_date"`
	Status                   string                     `json:"status"`
	ProductionResourceStatus string                     `json:"production_resource_status"`
}

// OrderResponse is the API response shape for a single ERP order.
type OrderResponse struct {
	ID                int64                    `json:"id"`
	OrderNumber       string                   `json:"order_number"`
	PartDescription   string                   `json:"part_description"`
	PlannedStartDate  time.Time                `json:"planned_start_date"`
	PlannedFinishDate time.Time                `json:"planned_finish_date"`
	ActualStartDate   *time.Time               `json:"actual_start_date"`
	ActualFinishDate  *time.Time               `json:"actual_finish_date"`
	Status            string                   `json:"status"`
	Priority          int                      `json:"priority"`
	Operations        []OrderOperationResponse `json:"operations"`
}

// MapOrderDomainToDTO converts a domain ERPOrder to its API response shape.
func MapOrderDomainToDTO(o domain.ERPOrder) OrderResponse {
	ops := make([]OrderOperationResponse, len(o.Operations))
	for i, op := range o.Operations {
		ops[i] = OrderOperationResponse{
			ID: op.ID,
			ProductionResource: ProductionResourceResponse{
				ID:          op.ProductionResource.ID,
				Number:      op.ProductionResource.Number,
				Description: op.ProductionResource.Description,
				Type:        op.ProductionResource.Type,
			},
			PlannedStartDate:         op.PlannedStartDate,
			PlannedFinishDate:        op.PlannedFinishDate,
			ActualStartDate:          op.ActualStartDate,
			ActualFinishDate:         op.ActualFinishDate,
			Status:                   op.Status,
			ProductionResourceStatus: op.ProductionResourceStatus,
		}
	}
	return OrderResponse{
		ID:                o.ID,
		OrderNumber:       o.OrderNumber,
		PartDescription:   o.PartDescription,
		PlannedStartDate:  o.PlannedStartDate,
		PlannedFinishDate: o.PlannedFinishDate,
		ActualStartDate:   o.ActualStartDate,
		ActualFinishDate:  o.ActualFinishDate,
		Status:            o.Status,
		Priority:          o.Priority,
		Operations:        ops,
	}
}

// MapOrdersDomainToDTO converts a slice of domain ERPOrder to its API response shape.
func MapOrdersDomainToDTO(orders []domain.ERPOrder) []OrderResponse {
	result := make([]OrderResponse, len(orders))
	for i, o := range orders {
		result[i] = MapOrderDomainToDTO(o)
	}
	return result
}
