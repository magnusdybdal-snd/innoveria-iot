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

		companyID, err := authCompanyIDFromHeader(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		result, err := svc.GetAll(ctx, companyID)
		if err != nil {
			status, message, cause := MapDomainError(err)
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

		companyID, err := authCompanyIDFromHeader(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		idStr := r.PathValue("id")
		idInt, err := strconv.ParseInt(idStr, 10, 64) // converts to int64
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid production resource id")
			return
		}

		result, err := svc.GetOne(ctx, idInt, companyID)
		if err != nil {
			status, message, cause := MapDomainError(err)
			json.HandleError(w, status, cause, message)
			return
		}

		response := handlerdto.MapProductionResourceDomainToDTOSingle(result)
		if err := json.Encode(w, http.StatusOK, response); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}
