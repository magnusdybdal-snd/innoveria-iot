package dto

import (
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// MeasurementResponse represents a single sensor measurement in the API response.
type MeasurementResponse struct {
	DeviceEUI string         `json:"device_eui"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload"`
}

// SensorMetricResponse represents a payload-key → measurement type mapping in the API response.
type SensorMetricResponse struct {
	PayloadKey      string  `json:"payload_key"`
	MeasurementType string  `json:"measurement_type"`
	Unit            *string `json:"unit"`
}

// SensorContextResponse represents a sensor enriched with its metrics and measurements.
// TotalPowerWh is non-null only for electricity sensors with a configured voltage.
type SensorContextResponse struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	DeviceEUI    string                 `json:"device_eui"`
	Metrics      []SensorMetricResponse `json:"metrics"`
	Measurements []MeasurementResponse  `json:"measurements"`
	TotalPowerWh *float64               `json:"total_power_wh"`
}

// OperationContextResponse represents a manufacturing operation enriched with sensor data.
type OperationContextResponse struct {
	Operation OrderOperationResponse  `json:"operation"`
	Sensors   []SensorContextResponse `json:"sensors"`
	// Degraded is true when sensor data could not be loaded due to a technical error.
	// A false value with an empty sensors array means no sensors are mapped (expected state).
	Degraded bool `json:"degraded"`
}

// OrderContextResponse is the top-level API response for GET /orders/{id}/context.
type OrderContextResponse struct {
	Order      OrderResponse              `json:"order"`
	Operations []OperationContextResponse `json:"operations"`
}

// mapOperation converts a domain ERPOrderOperation to its API response shape.
func mapOperation(op domain.ERPOrderOperation) OrderOperationResponse {
	reports := make([]OrderReportResponse, len(op.Reports))
	for i, r := range op.Reports {
		reports[i] = OrderReportResponse{
			ID:                 r.ID,
			Quantity:           r.Quantity,
			RestQuantity:       r.RestQuantity,
			Type:               r.Type,
			ReportingTimestamp: r.ReportingTimestamp,
			ActualReportedDate: r.ActualReportedDate,
		}
	}
	return OrderOperationResponse{
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
		Reports:                  reports,
	}
}

// MapOrderContextDomainToDTO converts a domain OrderContext to its API response shape.
func MapOrderContextDomainToDTO(oc *domain.OrderContext) OrderContextResponse {
	ops := make([]OperationContextResponse, len(oc.Operations))
	for i, opCtx := range oc.Operations {
		sensors := make([]SensorContextResponse, len(opCtx.Sensors))
		for j, sc := range opCtx.Sensors {
			metrics := make([]SensorMetricResponse, len(sc.Metrics))
			for k, m := range sc.Metrics {
				metrics[k] = SensorMetricResponse{
					PayloadKey:      m.PayloadKey,
					MeasurementType: m.MeasurementType,
					Unit:            m.Unit,
				}
			}

			measurements := make([]MeasurementResponse, len(sc.Measurements))
			for k, mr := range sc.Measurements {
				measurements[k] = MeasurementResponse{
					DeviceEUI: mr.DeviceEUI,
					Timestamp: mr.Timestamp,
					Payload:   mr.Payload,
				}
			}

			sensors[j] = SensorContextResponse{
				ID:           sc.Sensor.ID,
				Name:         sc.Sensor.Name,
				DeviceEUI:    sc.Sensor.DeviceEUI,
				Metrics:      metrics,
				Measurements: measurements,
				TotalPowerWh: sc.PowerConsumptionWh,
			}
		}

		ops[i] = OperationContextResponse{
			Operation: mapOperation(opCtx.Operation),
			Sensors:   sensors,
			Degraded:  opCtx.Degraded,
		}
	}

	return OrderContextResponse{
		Order:      MapOrderDomainToDTO(oc.Order),
		Operations: ops,
	}
}
