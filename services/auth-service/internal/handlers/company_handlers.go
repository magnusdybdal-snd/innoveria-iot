package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// PostCompany handles company creation requests.
//
// @Summary Register a new company
// @Tags companies
// @Accept json
// @Produce json
// @Param body body dto.CreateNewCompany true "Company payload"
// @Success 201 {object} dto.CompanyResponse
// @Failure 400
// @Failure 500
// @Router /companies [post]
func PostCompany(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateNewCompany](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		// Dto to comain, only name and address
		companyDomain := dto.MapCreateCompanyToDomain(payload)
		if companyDomain.Name == "" || companyDomain.Address == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "name and address are required")
			return
		}

		companyResp, err := svc.RegisterCompany(ctx, companyDomain)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		// returned response as domain
		resp := dto.MapCompanyFromDomain(companyResp)
		if err := json.Encode(w, http.StatusCreated, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetAllCompanies handles requests to fetch all companies.
//
// @Summary Get all companies
// @Tags companies
// @Produce json
// @Success 200 {object} dto.CompanyListResponse
// @Failure 500
// @Router /companies [get]
func GetAllCompanies(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companies, err := svc.GetAllCompanies(ctx)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapCompanyListFromDomain(companies)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// DeleteCompany handles requests to delete a company by ID.
//
// @Summary Delete company
// @Tags companies
// @Param companyID path string true "Company ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /companies/{companyID} [delete]
func DeleteCompany(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		companyID := r.PathValue("companyID")
		if companyID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing companyID path parameter"), "companyID is required")
			return
		}

		if err := svc.DeleteCompany(ctx, companyID); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
