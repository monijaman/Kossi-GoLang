package auth

import (
	"encoding/json"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/interface/middleware"
	"kossti/internal/usecase/auth"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

// Request/Response structures
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email"`
	Token           string `json:"token"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation"`
}

type UpdateProfileRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// convertUserToResponse converts entity to response format
func convertUserToResponse(user *entities.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

// RegisterHandler handles POST /register and POST /v1/registration
func RegisterHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "username, email, and password are required"})
		return
	}

	user := &entities.User{
		Name:      req.Username,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if req.Type != "" {
		t := strings.ToLower(strings.TrimSpace(req.Type))
		switch t {
		case "admin", "reviewer", "guest":
			user.Type = t
		default:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid user type; must be admin, reviewer, or guest"})
			return
		}
	}

	err := auth.Register(r.Context(), userRepo, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
		"email":   user.Email,
	})
}

// LoginHandler handles POST /login
func LoginHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "email and password are required"})
		return
	}

	accessToken, refreshToken, user, err := auth.Login(r.Context(), userRepo, refreshTokenRepo, req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":       "Login successful",
		"email":         user.Email,
		"name":          user.Name,
		"type":          user.Type,
		"token":         accessToken,
		"refresh_token": refreshToken,
	})
}

// RefreshTokenHandler handles POST /v1/refresh-token
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Get refresh token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing or invalid Authorization header"})
		return
	}
	refreshToken := strings.TrimPrefix(authHeader, "Bearer ")

	if refreshToken == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "refresh_token is required"})
		return
	}

	// Get refresh token from database
	refreshTokenEntity, err := refreshTokenRepo.GetByToken(r.Context(), refreshToken)
	if err != nil || refreshTokenEntity == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid refresh token"})
		return
	}

	// Check if refresh token is expired
	if refreshTokenEntity.ExpiresAt < time.Now().Unix() {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Refresh token expired"})
		return
	}

	// Get user
	user, err := userRepo.GetByID(r.Context(), refreshTokenEntity.UserID)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	// Generate new access token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to generate token"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":      "Token refreshed successfully",
		"access_token": accessToken,
	})
}

// ForgotPasswordHandler handles POST /v1/forgot-password
func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "email is required"})
		return
	}

	// TODO: Implement password reset logic
	// For now, return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password reset email sent",
		"email":   req.Email,
	})
}

// ResetPasswordHandler handles POST /v1/reset-password
func ResetPasswordHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Email == "" || req.Token == "" || req.Password == "" || req.PasswordConfirm == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "email, token, password, and password_confirmation are required"})
		return
	}

	if req.Password != req.PasswordConfirm {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "password and password_confirmation do not match"})
		return
	}

	// TODO: Implement password reset verification and update
	// For now, return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password reset successfully",
		"email":   req.Email,
	})
}

// LogoutHandler handles POST /v1/logout
func LogoutHandler(w http.ResponseWriter, r *http.Request, refreshTokenRepo repository.RefreshTokenRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Get user ID from context (set by JWT middleware)
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user context"})
		return
	}

	// Delete all refresh tokens for this user
	if err := refreshTokenRepo.DeleteByUserID(r.Context(), userID); err != nil {
		// Log error but don't fail the logout
		log.Printf("Warning: Failed to delete refresh tokens for user %d: %v", userID, err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}

// CheckTokenHandler handles GET /v1/check-token
func CheckTokenHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Extract and validate JWT token from Authorization header
	// TODO: Return user info if token is valid
	// For now, return placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Token is valid",
		"status":  "authenticated",
	})
}

// SearchUsersHandler handles GET /v1/search_users
func SearchUsersHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	query := r.URL.Query().Get("q")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	// Set defaults
	limit := 10
	offset := 0

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "q parameter is required for search"})
		return
	}

	// TODO: Implement user search in repository
	// For now, return empty results
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":  []UserResponse{},
		"count":  0,
		"limit":  limit,
		"offset": offset,
		"query":  query,
	})
}

// GetProfileHandler handles GET /v1/profile
func GetProfileHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Extract user ID from JWT token
	// TODO: Get user profile from repository
	// For now, return placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         1,
		"name":       "Test User",
		"email":      "test@example.com",
		"created_at": time.Now().Format(time.RFC3339),
		"updated_at": time.Now().Format(time.RFC3339),
	})
}

// UpdateProfileHandler handles POST /v1/profile
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// TODO: Extract user ID from JWT token
	// TODO: Update user profile in repository
	// TODO: Validate email uniqueness if email is being updated
	// For now, return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Profile updated successfully",
	})
}
