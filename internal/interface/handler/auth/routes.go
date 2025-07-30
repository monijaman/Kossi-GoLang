package auth

import (
	"encoding/json"
	"kossti/internal/domain/repository"
	"kossti/internal/interface/handler"
	"net/http"
)

// RegisterAuthRoutes registers authentication-related endpoints to the mux.
// Now accepts both userRepo and refreshTokenRepo for token saving.
func RegisterAuthRoutes(mux *http.ServeMux, userRepo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository) {
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
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
