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
	"net/netip"
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

// Login authenticates a user. By creating an accesss token and a refresh token
func (s *AuthServiceImpl) Login(ctx context.Context, email, password, deviceInfo string, ip *netip.Addr) (domain.LoginResult, string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return domain.LoginResult{}, "", domain.ErrUnauthorized
		}

		return domain.LoginResult{}, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return domain.LoginResult{}, "", domain.ErrUnauthorized
	}

	accessToken, expiresIn, err := s.generateAccessToken(user)
	if err != nil {
		return domain.LoginResult{}, "", err
	}

	rawRefreshToken, err := s.generateRefreshToken()
	if err != nil {
		return domain.LoginResult{}, "", fmt.Errorf("generating refresh token: %w", err)
	}

	// Update last login in user db
	if err := s.userRepo.UpdateLastLoggedIn(ctx, user.ID); err != nil {
		return domain.LoginResult{}, "", err
	}

	refreshToken := s.hashRefreshTokenHMAC(rawRefreshToken)

	if err := s.refreshTokenRepo.UpsertForLogin(ctx, domain.RefreshToken{
		UserID:     user.ID,
		TokenHash:  refreshToken, // storing hmac version in db
		ExpiresAt:  time.Now().UTC().Add(s.refreshTTL),
		DeviceInfo: deviceInfo,
		IPAddress:  ip,
	}); err != nil {
		return domain.LoginResult{}, "", err
	}

	slog.Info("successfully authenticate user", "id", user.ID)
	return domain.LoginResult{
		AccessToken: accessToken,
		TokenType:   "Bearer", // How the token is sendt over http
		ExpiresIn:   expiresIn,
	}, rawRefreshToken, nil // returns the rawrefresh token, and the hmac version is in db
}

// Refresh is used to revoke the refresh token so the user has a fresh access token
func (s *AuthServiceImpl) Refresh(ctx context.Context, refreshToken string) (domain.LoginResult, string, error) {
	if refreshToken == "" {
		return domain.LoginResult{}, "", domain.ErrUnauthorized
	}
	tokenHash := s.hashRefreshTokenHMAC(refreshToken)

	// check db for active hash
	stored, err := s.refreshTokenRepo.FindActiveByHash(ctx, tokenHash)
	if err != nil {
		return domain.LoginResult{}, "", domain.ErrUnauthorized
	}

	// check to see if expired, else log out
	if time.Now().UTC().After(stored.ExpiresAt) {
		return domain.LoginResult{}, "", domain.ErrUnauthorized
	}

	user, err := s.userRepo.FindByID(ctx, stored.UserID)
	if err != nil {
		return domain.LoginResult{}, "", domain.ErrUnauthorized
	}

	// Generate a fresh access token
	accessToken, expiresIn, err := s.generateAccessToken(user)
	if err != nil {
		return domain.LoginResult{}, "", fmt.Errorf("generate access token: %w", err)
	}

	// Generate a fresh refresh token
	newRefreshToken, err := s.generateRefreshToken()
	if err != nil {
		return domain.LoginResult{}, "", fmt.Errorf("generate refresh token: %w", err)
	}

	// Updates the refresh token in db
	if err := s.refreshTokenRepo.UpdateRefreshToken(ctx, domain.RefreshToken{
		ID:        stored.ID,
		TokenHash: s.hashRefreshTokenHMAC(newRefreshToken), // hmac hash in db
		ExpiresAt: time.Now().UTC().Add(s.refreshTTL),
	}); err != nil {
		return domain.LoginResult{}, "", err
	}

	return domain.LoginResult{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
	}, newRefreshToken, nil
}

// Me returns the authenticated user profile.
func (s *AuthServiceImpl) Me(ctx context.Context, token string) (domain.User, error) {

	return domain.User{}, nil
}

// generateAccessToken generates a short lived jwt token used by the client
func (s *AuthServiceImpl) generateAccessToken(user domain.User) (string, time.Duration, error) {
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
		return "", time.Duration(0), fmt.Errorf("sign access token: %w", err)
	}

	expiresIn := s.accessTTL

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
func (s *AuthServiceImpl) hashRefreshTokenHMAC(token string) string {
	mac := hmac.New(sha256.New, s.jwtSecret) // TODO: Change this with refresh_token_pepper
	mac.Write([]byte(token))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
