package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/json"
	"innoveria-iot/pkg/roles"

	"github.com/google/uuid"
)

// PostUser handles requests to create a new user. Admin only.
//
// @Summary		Create a user
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		body	body	dto.CreateUserRequest	true	"User payload"
// @Success		201		{object}	dto.UserResponse
// @Failure		400
// @Failure		403
// @Failure		409
// @Failure		500
// @Router		/users [post]
func PostUser(svc domain.UserService) http.HandlerFunc {
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

		payload, err := json.Decode[dto.CreateUserRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.CompanyID = strings.TrimSpace(payload.CompanyID)
		payload.Name = strings.TrimSpace(payload.Name)
		payload.Email = strings.TrimSpace(payload.Email)
		payload.Role = strings.TrimSpace(payload.Role)

		if payload.CompanyID == "" || payload.Name == "" || payload.Email == "" || payload.Password == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "company_id, name, email and password are required")
			return
		}

		if _, err := uuid.Parse(payload.CompanyID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid company_id (uuid)")
			return
		}

		if payload.Role == "" {
			payload.Role = string(roles.User)
		} else if payload.Role != string(roles.PlatformAdmin) && payload.Role != string(roles.User) {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("invalid role %q", payload.Role), "role must be PLATFORM_ADMIN or FACTORY_WORKER")
			return
		}

		user, err := svc.Create(ctx, dto.MapCreateUserDTOToDomain(payload))
		if err != nil {
			if errors.Is(err, domain.ErrUserAlreadyExists) {
				json.HandleError(w, http.StatusConflict, err, "user already exists")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		if err := json.Encode(w, http.StatusCreated, dto.MapUserDomainToDTO(user)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetUsers handles requests to list all users, optionally filtered by company. Admin only.
//
// @Summary		List users
// @Tags		users
// @Produce		json
// @Param		company_id	query	string	false	"Filter by company"
// @Success		200	{object}	dto.UserListResponse
// @Failure		403
// @Failure		500
// @Router		/users [get]
func GetUsers(svc domain.UserService) http.HandlerFunc {
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

		companyID := r.URL.Query().Get("company_id")
		if companyID != "" {
			if _, err := uuid.Parse(companyID); err != nil {
				json.HandleError(w, http.StatusBadRequest, err, "invalid company_id (uuid)")
				return
			}
		}

		users, err := svc.GetAll(ctx, companyID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		if err := json.Encode(w, http.StatusOK, dto.MapUserListDomainToDTO(users)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// PatchUser handles requests to update a user's editable fields. Admin only.
//
// @Summary		Update a user
// @Tags		users
// @Accept		json
// @Param		id		path	string					true	"User ID"
// @Param		body	body	dto.UpdateUserRequest	true	"Update payload"
// @Success		204
// @Failure		400
// @Failure		403
// @Failure		404
// @Failure		500
// @Router		/users/{id} [patch]
func PatchUser(svc domain.UserService) http.HandlerFunc {
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

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}
		if _, err := uuid.Parse(id); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid user id (uuid)")
			return
		}

		payload, err := json.Decode[dto.UpdateUserRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if payload.Name == nil && payload.Email == nil && payload.Password == nil {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no fields provided"), "at least one field is required")
			return
		}

		if err := svc.Update(ctx, id, dto.MapUpdateUserDTOToDomain(payload)); err != nil {
			if errors.Is(err, domain.ErrUserNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "user not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteUser handles requests to delete a user by ID. Admin only.
//
// @Summary		Delete a user
// @Tags		users
// @Param		id	path	string	true	"User ID"
// @Success		204
// @Failure		400
// @Failure		403
// @Failure		404
// @Failure		500
// @Router		/users/{id} [delete]
func DeleteUser(svc domain.UserService) http.HandlerFunc {
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

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}
		if _, err := uuid.Parse(id); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid user id (uuid)")
			return
		}

		if err := svc.Delete(ctx, id); err != nil {
			if errors.Is(err, domain.ErrUserNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "user not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
