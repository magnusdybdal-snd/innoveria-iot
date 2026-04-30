package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers/dto"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/json"
)

// GetOrderByID returns a single ERP order with full detail (operations, production resources, reports).
// @Summary		Get Order By ID
// @Tags		orders
// @Produce		json
// @Param		id	path	int	true	"ERP Order ID"
// @Success		200	{object}	dto.OrderResponse
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router		/orders/{id} [get]
func GetOrderByID(svc domain.ContextService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		idStr := r.PathValue("id")
		orderID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid order ID, expected an integer")
			return
		}

		order, err := svc.GetOrderByID(ctx, auth.CompanyID, auth.UserID, string(auth.Role), orderID)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "order not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch order")
			return
		}

		if err := json.Encode(w, http.StatusOK, dto.MapOrderDomainToDTO(*order)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}

// GetOrderContext aggregates ERP, sensor, and measurement data for a single order.
// @Summary		Get Order Context
// @Tags		orders
// @Produce		json
// @Param		id	path	int	true	"ERP Order ID"
// @Success		200	{object}	dto.OrderContextResponse
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router		/orders/{id}/context [get]
func GetOrderContext(svc domain.ContextService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		idStr := r.PathValue("id")
		orderID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid order ID, expected an integer")
			return
		}

		orderCtx, err := svc.GetOrderContext(ctx, auth.CompanyID, auth.UserID, string(auth.Role), orderID)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "order not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch order context")
			return
		}

		if err := json.Encode(w, http.StatusOK, dto.MapOrderContextDomainToDTO(orderCtx)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}

// GetOrders returns all ERP orders enriched with their full detail (reportings, workcenter).
// @Summary		Get Orders
// @Tags		orders
// @Produce		json
// @Success		200	{array}		dto.OrderSummaryResponse
// @Failure		401
// @Failure		500
// @Router		/orders [get]
func GetOrders(svc domain.ContextService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		orders, err := svc.GetOrders(ctx, auth.CompanyID, auth.UserID, string(auth.Role))
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch orders")
			return
		}

		if err := json.Encode(w, http.StatusOK, dto.MapOrdersSummaryToDTO(orders)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}
