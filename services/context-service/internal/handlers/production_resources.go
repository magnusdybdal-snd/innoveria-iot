package handlers

import (
	"errors"
	"net/http"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers/dto"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/json"
)

// GetProductionResources returns all ERP production resources (work centers) for the authenticated company.
// @Summary      Get Production Resources
// @Tags         production-resources
// @Produce      json
// @Success      200  {array}   dto.ProductionResourceResponse
// @Failure      500
// @Router       /production-resources [get]
func GetProductionResources(svc domain.ContextService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}
		companyID := auth.CompanyID

		resources, err := svc.GetProductionResources(ctx, companyID)
		if err != nil {
			if errors.Is(err, domain.ErrUnauthorized) {
				json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch production resources")
			return
		}

		if err := json.Encode(w, http.StatusOK, dto.MapProductionResourcesDomainToDTO(resources)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}
