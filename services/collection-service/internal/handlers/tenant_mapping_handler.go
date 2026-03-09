package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/collection-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// PostTenantMapping stores a companyID-tenantID mapping.
//
// @Summary		Create tenant mapping
// @Tags		company-config
// @Accept		json
// @Param		body	body	dto.CreateTenantMappingRequest	true	"Tenant mapping payload"
// @Success		201
// @Failure		400
// @Failure		500
// @Router		/company-config [post]
func PostTenantMapping(svc domain.TenantMappingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateTenantMappingRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if payload.CompanyID == "" || payload.TenantID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "company_id and tenant_id are required")
			return
		}

		if err := svc.CreateTenantMapping(ctx, payload.CompanyID, payload.TenantID); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
