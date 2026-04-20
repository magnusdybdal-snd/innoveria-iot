package handlers

import (
	"net/http"

	"innoveria-iot/erp-service/internal/domain"
	handlerdto "innoveria-iot/erp-service/internal/handlers/dto"

	"innoveria-iot/pkg/json"
)

// GetAllProductionResources TODO: add swagger
func GetAllProductionResources(svc domain.ProductionResourceSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		companyID := "a0000000-0000-0000-0000-000000000001" // TODO: use real company id

		result, err := svc.GetAll(ctx, companyID)
		if err != nil {
			status, message, cause := MapIngestDomainError(err)
			json.HandleError(w, status, cause, message)
			return
		}

		response := handlerdto.MapProductionResourceDomainToDTO(result)

		if err := json.Encode(w, http.StatusOK, response); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// GetOneProductionResource returns a handler for retrieving a single production resource.
func GetOneProductionResource(svc domain.ProductionResourceSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
