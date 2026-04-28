package handlers

import (
	"net/http"
	"strconv"

	"innoveria-iot/erp-service/internal/domain"
	handlerdto "innoveria-iot/erp-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// GetOrderSummaryHandler TODO: add swagger
//
// It handles GET requests for listing available orders for a company and
// returns a lightweight response containing order IDs and order numbers.
// This endpoint is intended for selection/list screens before fetching full
// order details.
func GetOrderSummaryHandler(svc domain.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companyID := "a0000000-0000-0000-0000-000000000001" // TODO: remove hardcoded companyID

		result, err := svc.GetOrderSummary(ctx, companyID)
		if err != nil {
			status, message, cause := MapDomainError(err)
			json.HandleError(w, status, cause, message)
			return
		}

		response := handlerdto.MapOrderSummaryDomainToDTO(result)

		if err := json.Encode(w, http.StatusOK, response); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetOneOrderHandler TODO: add swagger
//
// It handles GET requests for one order by path parameter `id` and returns
// an aggregated order view, including the order, its operations with reports,
// and referenced production resources.
func GetOneOrderHandler(svc domain.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companyID := "a0000000-0000-0000-0000-000000000001" // TODO: remove hardcoded companyID

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

		response := handlerdto.MapOrderAggregateDomainToDTO(result)

		if err := json.Encode(w, http.StatusOK, response); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}
