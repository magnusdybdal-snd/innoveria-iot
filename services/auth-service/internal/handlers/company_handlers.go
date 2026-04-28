package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"

	"innoveria-iot/pkg/authctx"

	"github.com/google/uuid"
)

// PostCompany handles company creation requests.
//
// @Summary Register a new company
// @Tags companies
// @Accept json
// @Produce json
// @Param body body dto.CreateNewCompany true "Company payload"
// @Success 201 {object} dto.RegisterCompanyResponse
// @Failure 400
// @Failure 500
// @Router /companies [post]
func PostCompany(svc domain.CompanyService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateNewCompany](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.Name = strings.TrimSpace(payload.Name)
		payload.Address = strings.TrimSpace(payload.Address)

		// Dto to domain, only name and address
		companyDomain := dto.MapCreateCompanyToDomain(payload)
		if companyDomain.Name == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "name is required")
			return
		}

		result, err := svc.RegisterCompany(ctx, companyDomain)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		// Maps from domain to DTO
		resp := dto.MapCompanyFromDomain(result)
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
func GetAllCompanies(svc domain.CompanyService) http.HandlerFunc {
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

// GetOneCompany handles requests to fetch one company by ID.
//
// @Summary Get one company
// @Tags companies
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} dto.CompanyResponse
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /companies/{id} [get]
func GetOneCompany(svc domain.CompanyService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companyID := r.PathValue("id")
		if companyID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		// Check for valid uuid
		if _, err := uuid.Parse(companyID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid company id (uuid)")
			return
		}

		company, err := svc.GetOneCompany(ctx, companyID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrCompanyNotFound):
				json.HandleError(w, http.StatusNotFound, err, "company not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.MapCompanyFromDomain(company)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// DeleteCompany handles requests to delete a company by ID.
//
// @Summary Delete company
// @Tags companies
// @Param id path string true "id"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /companies/{id} [delete]
func DeleteCompany(svc domain.CompanyService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		companyID := r.PathValue("id")
		if companyID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		// Check for valid uuid
		if _, err := uuid.Parse(companyID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid company id (uuid)")
			return
		}

		if err := svc.DeleteCompany(ctx, companyID); err != nil {
			switch {
			case errors.Is(err, domain.ErrCompanyNotFound):
				json.HandleError(w, http.StatusNotFound, err, "company not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// PostCompanyERPAgentToken handles requests to issue a new ERP agent token for a company.
//
// @Summary Issue ERP agent token
// @Tags companies
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} dto.ERPAgentTokenResponse
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /companies/{id}/erp-agent-token [post]
func PostCompanyERPAgentToken(svc domain.CompanyService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		if !auth.IsAdmin() {
			json.HandleError(w, http.StatusForbidden, fmt.Errorf("forbidden"), "forbidden")
			return
		}

		companyID := r.PathValue("id")
		if companyID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		if _, err := uuid.Parse(companyID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid company id (uuid)")
			return
		}

		token, err := svc.IssueERPAgentToken(ctx, companyID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrCompanyNotFound):
				json.HandleError(w, http.StatusNotFound, err, "company not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.MapERPAgentTokenResponse(token)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}
