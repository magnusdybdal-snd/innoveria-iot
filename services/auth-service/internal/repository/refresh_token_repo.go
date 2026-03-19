package repository

import (
	"context"
	"fmt"
	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

// TODO: Handle ip and device info
const (
	createRefreshTokenQuery = `
		INSERT INTO auth.refresh_token (user_id, token_hash, expires_at)
		VALUES ($1, $2, $3)
	`
)

// RefreshTokenRepoImpl is the PostgreSQL implementation of refresh token persistence.
type RefreshTokenRepoImpl struct {
	db *dbutil.DB
}

// NewRefreshTokenRepo creates a new refresh token repository.
func NewRefreshTokenRepo(db *dbutil.DB) *RefreshTokenRepoImpl {
	return &RefreshTokenRepoImpl{db: db}
}

// Create stores a refresh token record.
func (r *RefreshTokenRepoImpl) Create(ctx context.Context, token domain.RefreshToken) error {
	_, err := r.db.Pool.Exec(ctx, createRefreshTokenQuery,
		token.UserID,
		token.TokenHash,
		token.ExpiresAt,
	)
	if err != nil {
		return fmt.Errorf("create refresh token: %w", err)
	}

	return nil
}

// UpdateRefreshToken stores a new fresh refresh token in the db
func (r *RefreshTokenRepoImpl) UpdateRefreshToken(ctx context.Context, token domain.RefreshToken) error {
	return nil
}

// FindActiveByHash will look for an active refresh token in db
func (r *RefreshTokenRepoImpl) FindActiveByHash(ctx context.Context, oldHash string) (domain.RefreshToken, error) {
	return domain.RefreshToken{}, nil
}
