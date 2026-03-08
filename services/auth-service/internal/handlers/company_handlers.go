package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// PostCompany handles company creation requests.
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
