package handlers

import (
	"innoveria-iot/auth-service/internal/domain"
	"net/http"
)

// PostFactoryArea handles factory area creation requests.
//
// @Summary Register a new factory area
// @Tags factory-areas
// @Accept json
// @Produce json
// @Param body body dto.CreateNewFactoryArea true "Factory area payload"
// @Success 201 {object} dto.FactoryAreaResponse
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /factory-areas [post]
func PostFactoryArea(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
