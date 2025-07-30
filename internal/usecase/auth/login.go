package auth

// This use case follows Clean Code Architecture principles:
// - Business logic (authentication) is separated from infrastructure (database, hashing).
// - The repository is injected (Dependency Injection), so this code is easy to test and not tied to a specific DB.
// - The Repository Pattern is used: repo.FindByEmail abstracts data access.
// - The Hash utility is injected, so password logic is also decoupled.
// This makes the code modular, testable, and easy to maintain.
import (
	"context"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/pkg/hash"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

func Login(
	ctx context.Context,
	userRepo repository.UserRepository,
	refreshTokenRepo repository.RefreshTokenRepository,
	email, password string,
) (string, string, error) {
	user, err := userRepo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return "", "", errors.New("user not found")
	}
	if !hash.CheckPasswordHash(password, user.Password) {
		return "", "", errors.New("invalid password")
	}

	// Create JWT claims
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiry
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days expiry
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshTokenObj.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Save refresh token to DB
	rt := &entities.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: refreshClaims["exp"].(int64),
	}
	if err := refreshTokenRepo.Save(ctx, rt); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
