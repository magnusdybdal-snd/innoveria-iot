package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// PostCompanyConfig creates a Chirpstack tenant and application for a company and stores the mapping.
//
// @Summary		Create company config
// @Tags		company-config
// @Accept		json
// @Produce		json
// @Param		body	body		dto.CreateCompanyConfigRequest	true	"Company config payload"
// @Success		201		{object}	dto.CreateCompanyConfigResponse
// @Failure		400
// @Failure		500
// @Router		/company-config [post]
func PostCompanyConfig(svc domain.CompanyConfigService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateCompanyConfigRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if payload.CompanyID == "" || payload.Name == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "company_id and name are required")
			return
		}

		tenantID, err := svc.CreateCompanyConfig(ctx, payload.CompanyID, payload.Name)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		if err := json.Encode(w, http.StatusCreated, dto.CreateCompanyConfigResponse{TenantID: tenantID}); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// DeleteCompanyConfig removes a company's Chirpstack config and database record.
//
// @Summary		Delete company config
// @Tags		company-config
// @Param		companyID	path	string	true	"Company ID"
// @Success		204
// @Failure		400
// @Failure		500
// @Router		/company-config/{companyID} [delete]
func DeleteCompanyConfig(svc domain.CompanyConfigService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companyID := r.PathValue("companyID")
		if companyID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing companyID"), "bad request")
			return
		}

		if err := svc.DeleteCompanyConfig(ctx, companyID); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
