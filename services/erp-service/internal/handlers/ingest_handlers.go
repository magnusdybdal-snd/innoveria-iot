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
func PostIngestOrderReportings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r
		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}

// PostIngestWorkCenters handles batch ingest of work center master data pushed
// from erp-agent-service.
func PostIngestWorkCenters() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r
		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}
