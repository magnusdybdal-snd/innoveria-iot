package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"innoveria-iot/auth-service/internal/domain"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

// UserServiceImpl implements user management use-cases for the auth service.
type UserServiceImpl struct {
	userRepo domain.UserRepo
}

// NewUserService creates a new UserServiceImpl instance.
func NewUserService(userRepo domain.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo}
}

// Create creates a new user, hashing the plaintext password before storage.
func (s *UserServiceImpl) Create(ctx context.Context, user domain.User) (domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: hash password: %w", err)
	}
	user.PasswordHash = string(hash)

	created, err := s.userRepo.Create(ctx, user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return domain.User{}, domain.ErrUserAlreadyExists
		}
		return domain.User{}, err
	}

	slog.Info("successfully created user", "id", created.ID)
	return created, nil
}

// GetAll retrieves all users. If companyID is non-empty, results are filtered to that company.
func (s *UserServiceImpl) GetAll(ctx context.Context, companyID string) ([]domain.User, error) {
	users, err := s.userRepo.FindAll(ctx, companyID)
	if err != nil {
		return nil, err
	}

	slog.Info("successfully fetched users", "count", len(users))
	return users, nil
}

// Update updates the editable fields of a user.
// Only non-empty fields in payload are applied; the rest retain their current values.
func (s *UserServiceImpl) Update(ctx context.Context, userID string, payload domain.User) error {
	current, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if payload.Name != "" {
		current.Name = payload.Name
	}
	if payload.Email != "" {
		current.Email = payload.Email
	}
	if payload.PasswordHash != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(payload.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("update user: hash password: %w", err)
		}
		current.PasswordHash = string(hash)
	}

	if err := s.userRepo.Update(ctx, userID, current); err != nil {
		return err
	}

	slog.Info("successfully updated user", "id", userID)
	return nil
}

// Delete removes a user by ID.
func (s *UserServiceImpl) Delete(ctx context.Context, userID string) error {
	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return err
	}

	slog.Info("successfully deleted user", "id", userID)
	return nil
}
