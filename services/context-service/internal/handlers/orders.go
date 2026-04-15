package handlers

import (
	"net/http"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// hardcodedCompanyID is a temporary placeholder used while auth middleware
// propagation is not yet wired up end-to-end.
//
// TODO: replace with AUTH — read company ID from r.Header.Get("X-Auth-Company-Id") once the gateway injects trusted headers into this service.
const hardcodedCompanyID = "00000000-0000-0000-0000-000000000001"

// GetOrders returns all ERP orders enriched with their full detail (reportings, workcenter).
// @Summary		Get Orders
// @Tags		orders
// @Produce		json
// @Success		200	{array}		dto.OrderResponse
// @Failure		500
// @Router		/orders [get]
func GetOrders(svc domain.ContextService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// TODO: replace with AUTH — use r.Header.Get("X-Auth-Company-Id") once
		// the auth middleware is propagated to this service.
		companyID := hardcodedCompanyID

		orders, err := svc.GetOrders(ctx, companyID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch orders")
			return
		}

		if err := json.Encode(w, http.StatusOK, dto.MapOrdersDomainToDTO(orders)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}
