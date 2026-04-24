package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"

	"github.com/google/uuid"
)

// GetERPAgentCredentialByCompanyID returns ERP agent credential metadata for one company.
func GetERPAgentCredentialByCompanyID(svc domain.ERPAgentCredentialService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companyID := r.PathValue("company_id")
		if companyID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing company_id path parameter"), "company_id is required")
			return
		}

		if _, err := uuid.Parse(companyID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid company id (uuid)")
			return
		}

		credential, err := svc.GetERPAgentCredentialByCompanyID(ctx, companyID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrERPAgentCredentialNotFound):
				json.HandleError(w, http.StatusNotFound, err, "erp agent credential not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.MapERPAgentCredentialFromDomain(credential)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetERPAgentCredentialByKeyID returns ERP agent credential metadata for one key id.
func GetERPAgentCredentialByKeyID(svc domain.ERPAgentCredentialService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		keyID := r.PathValue("key_id")
		if keyID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing key_id path parameter"), "key_id is required")
			return
		}

		if _, err := uuid.Parse(keyID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid key id (uuid)")
			return
		}

		credential, err := svc.GetERPAgentCredentialByKeyID(ctx, keyID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrERPAgentCredentialNotFound):
				json.HandleError(w, http.StatusNotFound, err, "erp agent credential not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.MapERPAgentCredentialFromDomain(credential)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}
