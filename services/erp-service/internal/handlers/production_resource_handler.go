package handlers

import (
	"net/http"
	"strconv"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/json"

	handlerdto "innoveria-iot/erp-service/internal/handlers/dto"
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
		ctx := r.Context()

		companyID := "a0000000-0000-0000-0000-000000000001" // TODO: use real company id

		idStr := r.PathValue("id")
		idInt, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid production resource id")
			return
		}

		result, err := svc.GetOne(ctx, idInt, companyID)
		if err != nil {
			status, message, cause := MapIngestDomainError(err)
			json.HandleError(w, status, cause, message)
			return
		}

		if err := json.Encode(w, http.StatusOK, result); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}
