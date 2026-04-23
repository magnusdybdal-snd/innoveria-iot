package handlers

import (
	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/json"
	"net/http"
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
			return
		}
	}
}
