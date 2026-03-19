package repository

import (
	"context"
	"errors"
	"fmt"
	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5"
)

// TODO: Handle ip and device info
const (
	createRefreshTokenQuery = `
		INSERT INTO auth.refresh_token (user_id, token_hash, expires_at)
		VALUES ($1, $2, $3)
	`
	updateRefreshTokenQuery = `
		UPDATE auth.refresh_token
		SET token_hash = $1, expires_at = $2, revoked_at = NULL
		WHERE token_id = $3
			AND revoked_at IS NULL
			AND expires_at > NOW()
	`
	findActiveRefreshTokenHashQuery = `
		SELECT token_id, user_id, token_hash, expires_at, created_at
		FROM auth.refresh_token
		WHERE token_hash = $1
			AND revoked_at is NULL
			AND expires_at > NOW()
		ORDER BY created_at DESC
		LIMIT 1
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
	resp, err := r.db.Pool.Exec(ctx, updateRefreshTokenQuery,
		token.TokenHash,
		token.ExpiresAt,
		token.ID,
	)
	if err != nil {
		return fmt.Errorf("update refresh token: %w", err)
	}

	if resp.RowsAffected() == 0 {
		return fmt.Errorf("update refresh token: %w", domain.ErrUnauthorized)
	}
	return nil
}

// FindActiveByHash will look for an active refresh token in db
func (r *RefreshTokenRepoImpl) FindActiveByHash(ctx context.Context, oldHash string) (domain.RefreshToken, error) {
	var out domain.RefreshToken
	err := r.db.Pool.QueryRow(ctx, findActiveRefreshTokenHashQuery,
		oldHash,
	).Scan(
		&out.ID,
		&out.UserID,
		&out.TokenHash,
		&out.ExpiresAt,
		&out.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.RefreshToken{}, domain.ErrUnauthorized
		}
		return domain.RefreshToken{}, fmt.Errorf("find active refresh token by hash: %w", err)
	}
	return out, nil
}
