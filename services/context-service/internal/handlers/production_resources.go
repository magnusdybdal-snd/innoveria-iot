package handlers

import (
	"net/http"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers/dto"
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

		// TODO: replace with AUTH — use r.Header.Get("X-Auth-Company-Id") once
		// the auth middleware is propagated to this service.
		companyID := hardcodedCompanyID

		resources, err := svc.GetProductionResources(ctx, companyID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch production resources")
			return
		}

		if err := json.Encode(w, http.StatusOK, dto.MapProductionResourcesDomainToDTO(resources)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}
