package handlers

import (
	"net/http"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/erp-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
	monitordto "innoveria-iot/pkg/monitor/dto"
)

func PostIngestOrder(svc domain.Ingest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := json.Decode[[]monitordto.ManufacturingOrder](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}
		result := dto.MapMonitorOrderToDomain(payload)
		if err := svc.CreateOrder(ctx, result); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// PostIngestOrderOperations handles batch ingest of manufacturing order operations
// pushed from erp-agent-service.
func PostIngestOrderOperations(svc domain.Ingest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := json.Decode[[]monitordto.ManufacturingOrderOperation](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}
		result := dto.MapMonitorOrderOperationToDomain(payload)
		if err := svc.CreateOrderOperation(ctx, result); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// PostIngestOrderReportings handles batch ingest of manufacturing order operation
// reporting events pushed from erp-agent-service.
func PostIngestOrderReports(svc domain.Ingest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := json.Decode[[]monitordto.ManufacturingOrderOperationReporting](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}
		result := dto.MapMonitorOrderReportToDomain(payload)
		if err := svc.CreateOrderReport(ctx, result); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// PostIngestWorkCenters handles batch ingest of work center master data pushed
// from erp-agent-service.
func PostIngestWorkCenters(svc domain.Ingest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := json.Decode[[]monitordto.WorkCenter](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}
		result := dto.MapMonitorWorkcenterToDomain(payload)
		if err := svc.CreateProductionResource(ctx, result); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}
