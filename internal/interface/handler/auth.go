// auth.go - HTTP Handlers for Authentication
//
// This file is part of the Interface Adapters Layer in Clean Architecture.
// It contains HTTP handler functions for authentication-related endpoints (register, login).
// These handlers:
//   - Parse incoming HTTP requests (JSON bodies)
//   - Call the appropriate use case (business logic)
//   - Write HTTP responses (JSON)
//   - Are decoupled from infrastructure and database details
//
// Handlers here can be used directly or with dependency-injected repositories for testability and modularity.

package handler

import (
	"encoding/json"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/usecase/auth"
	"net/http"
) // This file is part of the Interface Adapters Layer in Clean Code Architecture.
// - These handlers adapt HTTP requests to use case calls.
// - They parse input, call the appropriate use case, and write the response.
// - Keeps delivery (HTTP) logic separate from business logic and infrastructure.
// - Makes it easy to swap out the web framework or test the handlers in isolation.

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandlerWithRepo(w http.ResponseWriter, r *http.Request, repo repository.UserRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	user := &entities.User{
		Name:     req.Username, // Using username as name for compatibility
		Email:    req.Email,
		Password: req.Password,
	}

	err := auth.Register(r.Context(), repo, user)
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

func LoginHandlerWithRepo(w http.ResponseWriter, r *http.Request, repo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	accessToken, refreshToken, err := auth.Login(r.Context(), repo, refreshTokenRepo, req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":       "Login successful",
		"email":         req.Email,
		"token":         accessToken,
		"refresh_token": refreshToken,
	})
}
