// Package services contains auth-service business logic.
package services

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"innoveria-iot/auth-service/internal/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthServiceImpl implements authentication use cases for the auth service.
type AuthServiceImpl struct {
	userRepo         domain.UserRepo
	refreshTokenRepo domain.RefreshTokenRepo
	jwtSecret        []byte // converted to byte in initializer
	jwtIssuer        string
	accessTTL        time.Duration
	refreshTTL       time.Duration
}

// NewAuthServiceImpl creates a new AuthServiceImpl instance.
func NewAuthServiceImpl(
	userRepo domain.UserRepo,
	refreshTokenRepo domain.RefreshTokenRepo,
	jwtSecret string,
	jwtIssuer string,
	accessTTL time.Duration,
	refreshTokenTTL time.Duration,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        []byte(jwtSecret),
		jwtIssuer:        jwtIssuer,
		accessTTL:        accessTTL,
		refreshTTL:       refreshTokenTTL,
	}
}

// Login authenticates a user.
func (s *AuthServiceImpl) Login(ctx context.Context, email, password string) (domain.LoginResult, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return domain.LoginResult{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return domain.LoginResult{}, nil
	}

	accessToken, expiresIn, err := s.generateAccessToken(user)
	if err != nil {
		return domain.LoginResult{}, err
	}

	refreshToken := s.generateRefreshToken()

	if err := s.refreshTokenRepo.Create(ctx, domain.RefreshToken{
		UserID:    user.ID,
		TokenHash: refreshToken,
		ExpiresAt: time.Now().UTC().Add(s.refreshTTL),
	}); err != nil {
		return domain.LoginResult{}, err
	}

	// Update last login

	return domain.LoginResult{
		AccessToken: accessToken,
		TokenType:   "Bearer", // How the token is sendt over http
		ExpiresIn:   expiresIn,
	}, nil
}

// Me returns the authenticated user profile.
func (s *AuthServiceImpl) Me(ctx context.Context, token string) {
	_, _ = ctx, token
}

// generateAccessToken generates a short lived token used by the client
func (s *AuthServiceImpl) generateAccessToken(user domain.User) (string, string, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(s.accessTTL)

	claims := jwt.MapClaims{
		"sub":        user.ID,          // subject of the jwt token
		"company_id": user.CompanyID,   // company the user belongs to
		"role":       user.Role,        // user role
		"iss":        s.jwtIssuer,      // issuer of the jwt (auth-service)
		"exp":        expiresAt.Unix(), // expiration time, when jwt expires
		"iat":        now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", "", fmt.Errorf("sign access token: %w", err)
	}

	expiresIn := strconv.FormatInt(int64(s.accessTTL.Seconds()), 10)

	return signed, expiresIn, nil
}

// generateRefreshToken generates a long term token stored in the db
func (s *AuthServiceImpl) generateRefreshToken() string {
	return ""
}
