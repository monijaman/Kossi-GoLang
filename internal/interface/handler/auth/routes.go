package auth

import (
	"encoding/json"
	"kossti/internal/domain/repository"
	"kossti/internal/interface/handler"
	"net/http"
)

// RegisterAuthRoutes registers authentication-related endpoints to the mux.
func RegisterAuthRoutes(mux *http.ServeMux, userRepo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository) {

	// Public auth endpoints (no auth required)

	// POST /register - User registration
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		RegisterHandler(w, r, userRepo)
	})

	// POST /login - User login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		LoginHandler(w, r, userRepo, refreshTokenRepo)
	})

	// V1 API endpoints

	// POST /v1/refresh-token - Refresh access token
	mux.HandleFunc("/v1/refresh-token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		RefreshTokenHandler(w, r, userRepo, refreshTokenRepo)
	})

	// POST /v1/registration - Alternative registration endpoint
	mux.HandleFunc("/v1/registration", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		RegisterHandler(w, r, userRepo)
	})

	// POST /v1/login - V1 login endpoint
	mux.HandleFunc("/v1/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		LoginHandler(w, r, userRepo, refreshTokenRepo)
	})

	// POST /v1/forgot-password - Password reset request
	mux.HandleFunc("/v1/forgot-password", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		ForgotPasswordHandler(w, r, userRepo)
	})

	// POST /v1/reset-password - Password reset
	mux.HandleFunc("/v1/reset-password", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		ResetPasswordHandler(w, r, userRepo)
	})

	// POST /v1/logout - User logout (auth required)
	mux.HandleFunc("/v1/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		// TODO: Add JWT middleware for authentication
		LogoutHandler(w, r, refreshTokenRepo)
	})

	// GET /v1/check-token - Check token validity (auth required)
	mux.HandleFunc("/v1/check-token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET method is allowed",
			})
			return
		}
		// TODO: Add JWT middleware for authentication
		CheckTokenHandler(w, r, userRepo)
	})

	// GET /v1/search_users - Search users (auth required)
	mux.HandleFunc("/v1/search_users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET method is allowed",
			})
			return
		}
		// TODO: Add JWT middleware for authentication
		SearchUsersHandler(w, r, userRepo)
	})

	// GET /v1/profile - Get user profile (auth required)
	mux.HandleFunc("/v1/profile", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// TODO: Add JWT middleware for authentication
			GetProfileHandler(w, r, userRepo)
		} else if r.Method == http.MethodPost {
			// TODO: Add JWT middleware for authentication
			UpdateProfileHandler(w, r, userRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET and POST methods are allowed",
			})
		}
	})

	// GET /v1/checkrole - Check user role (auth required)
	mux.HandleFunc("/v1/checkrole", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET method is allowed",
			})
			return
		}
		// TODO: Add JWT middleware and permission check for 'clients.create'
		GetProfileHandler(w, r, userRepo)
	})

	// Legacy endpoint compatibility - use old handler for backward compatibility
	mux.HandleFunc("/old-register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		handler.RegisterHandlerWithRepo(w, r, userRepo)
	})

	mux.HandleFunc("/old-login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}
		handler.LoginHandlerWithRepo(w, r, userRepo, refreshTokenRepo)
	})
}
