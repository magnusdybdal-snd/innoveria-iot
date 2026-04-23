package handlers

import (
	"net/http"
	"strconv"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/json"
)

// GetOrderSummaryHandler TODO: add swagger
func GetOrderSummaryHandler(svc domain.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companyID := "a0000000-0000-0000-0000-000000000001"

		result, err := svc.GetOrderSummary(ctx, companyID)
		if err != nil {
			status, message, cause := MapDomainError(err)
			json.HandleError(w, status, cause, message)
			return
		}

		if err := json.Encode(w, http.StatusOK, result); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetOneOrderHandler TODO: add swagger
func GetOneOrderHandler(svc domain.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companyID := "a0000000-0000-0000-0000-000000000001"

		idStr := r.PathValue("id")
		idInt, err := strconv.ParseInt(idStr, 10, 64) // converts to int64
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid order id")
			return
		}

		result, err := svc.GetOne(ctx, idInt, companyID)
		if err != nil {
			status, message, cause := MapDomainError(err)
			json.HandleError(w, status, cause, message)
			return
		}

		if err := json.Encode(w, http.StatusOK, result); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}
