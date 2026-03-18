package handlers

import (
	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
	"net/http"
	"strings"
)

// PostLogin handles user login requests.
//
// @Summary Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "Login payload"
// @Success 201 {object} dto.LoginResult
// @Failure 400
// @Failure 500
// @Router /login [post]
func PostLogin(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := json.Decode[dto.LoginRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.Email = strings.TrimSpace(payload.Email)
		payload.Password = strings.TrimSpace(payload.Password)

		result, err := svc.Login(ctx, payload.Email, payload.Password)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.ToLoginResult(result)
		if err := json.Encode(w, http.StatusCreated, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}

	}
}
