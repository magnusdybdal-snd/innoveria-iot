package dto

import (
	"time"

	"innoveria-iot/auth-service/internal/domain"
)

// UserResponse is the response payload for a single user.
type UserResponse struct {
	ID           string  `json:"id"`
	CompanyID    string  `json:"company_id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Role         string  `json:"role"`
	LastLoggedIn *string `json:"last_logged_in"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

// UserListResponse wraps a slice of UserResponse with a total count.
type UserListResponse struct {
	TotalCount int            `json:"total_count"`
	Users      []UserResponse `json:"users"`
}

// CreateUserRequest is the request payload for creating a new user.
type CreateUserRequest struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

// UpdateUserRequest is the request payload for updating a user.
// All fields are optional — only provided fields are applied.
type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

// MapCreateUserDTOToDomain maps a CreateUserRequest to a domain User.
// Password is placed in PasswordHash; the service layer will hash it before storage.
func MapCreateUserDTOToDomain(from CreateUserRequest) domain.User {
	return domain.User{
		CompanyID:    from.CompanyID,
		Name:         from.Name,
		Email:        from.Email,
		PasswordHash: from.Password,
		Role:         domain.RoleType(from.Role),
	}
}

// MapUpdateUserDTOToDomain maps an UpdateUserRequest to a domain User.
// Only non-nil pointer fields are populated; the service layer merges with current values.
func MapUpdateUserDTOToDomain(from UpdateUserRequest) domain.User {
	var u domain.User
	if from.Name != nil {
		u.Name = *from.Name
	}
	if from.Email != nil {
		u.Email = *from.Email
	}
	if from.Password != nil {
		u.PasswordHash = *from.Password
	}
	return u
}

// MapUserDomainToDTO maps a domain User to a UserResponse DTO.
func MapUserDomainToDTO(from domain.User) UserResponse {
	var lastLoggedIn *string
	if from.LastLoggedIn != nil {
		formatted := from.LastLoggedIn.Format(time.RFC3339)
		lastLoggedIn = &formatted
	}

	return UserResponse{
		ID:           from.ID,
		CompanyID:    from.CompanyID,
		Name:         from.Name,
		Email:        from.Email,
		Role:         string(from.Role),
		LastLoggedIn: lastLoggedIn,
		CreatedAt:    from.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    from.UpdatedAt.Format(time.RFC3339),
	}
}

// MapUserListDomainToDTO maps a slice of domain Users to a UserListResponse DTO.
func MapUserListDomainToDTO(from []domain.User) UserListResponse {
	users := make([]UserResponse, 0, len(from))
	for _, u := range from {
		users = append(users, MapUserDomainToDTO(u))
	}
	return UserListResponse{
		TotalCount: len(users),
		Users:      users,
	}
}
