// Package services contains auth-service business logic.
package services

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
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
		if errors.Is(err, domain.ErrUserNotFound) {
			return domain.LoginResult{}, domain.ErrUnauthorized
		}

		return domain.LoginResult{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return domain.LoginResult{}, domain.ErrUnauthorized
	}

	accessToken, expiresIn, err := s.generateAccessToken(user)
	if err != nil {
		return domain.LoginResult{}, err
	}

	rawRefreshToken, err := s.generateRefreshToken()
	if err != nil {
		return domain.LoginResult{}, fmt.Errorf("generating refresh token: %w", err)
	}

	refreshToken := s.hasRefreshTokenHMAC(rawRefreshToken)

	if err := s.refreshTokenRepo.Create(ctx, domain.RefreshToken{
		UserID:    user.ID,
		TokenHash: refreshToken,
		ExpiresAt: time.Now().UTC().Add(s.refreshTTL),
	}); err != nil {
		return domain.LoginResult{}, err
	}

	// Update last login

	slog.Info("successfully authenticate user", "id", user.ID)
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

// generateAccessToken generates a short lived jwt token used by the client
func (s *AuthServiceImpl) generateAccessToken(user domain.User) (string, string, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(s.accessTTL)

	claims := jwt.MapClaims{
		"sub":        user.ID,          // subject of the jwt token
		"company_id": user.CompanyID,   // company the user belongs to
		"role":       user.Role,        // user role
		"iss":        s.jwtIssuer,      // issuer of the jwt (auth-service)
		"exp":        expiresAt.Unix(), // expiration time, when jwt expires
		"iat":        now.Unix(),       // issued at time. Time at which jwt token was created
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", "", fmt.Errorf("sign access token: %w", err)
	}

	expiresIn := strconv.FormatInt(int64(s.accessTTL.Seconds()), 10)

	return signed, expiresIn, nil
}

// generateRefreshToken generates a long term jwt token stored in the db
// is applied hashing after this function
func (s *AuthServiceImpl) generateRefreshToken() (string, error) {
	b := make([]byte, 32) // 256 bit random gen
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

// hasRefreshTokenHMAC creates a deterministic lookup value for DB storage.
// only the server with secret can reproduce/check this
// important with good pepper incase db leak
func (s *AuthServiceImpl) hasRefreshTokenHMAC(token string) string {
	mac := hmac.New(sha256.New, s.jwtSecret) // TODO: Change this with refresh_token_pepper
	mac.Write([]byte(token))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
