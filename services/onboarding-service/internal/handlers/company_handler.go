package handlers

import (
	"net/http"

	"innoveria-iot/onboarding-service/internal/domain"
	"innoveria-iot/onboarding-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// PostCompany onboards a new company by running the creation SAGA.
//
// @Summary		Onboard a new company
// @Tags		onboarding
// @Accept		json
// @Param		body	body	dto.CreateCompanyRequest	true	"Company payload"
// @Success		201
// @Failure		400
// @Failure		500
// @Router		/company [post]
func PostCompany(svc domain.OnboardingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateCompanyRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		company := dto.MapToDomain(payload)

		if err := svc.CreateCompany(ctx, company); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
