package user

import (
	"encoding/json"
	"kossti/internal/domain/repository"
	"kossti/internal/interface/handler"
	"kossti/internal/interface/middleware"
	"net/http"
)

// RegisterUserRoutes registers all user-related endpoints to the mux.
func RegisterUserRoutes(mux *http.ServeMux, userRepo repository.UserRepository) {
	mux.HandleFunc("/api/users", middleware.JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET method is allowed",
			})
			return
		}
		handler.ListUsersHandler(w, r, userRepo)
	}))

	mux.HandleFunc("/api/users/all", middleware.JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET method is allowed",
			})
			return
		}
		handler.GetAllUsersHandler(w, r, userRepo)
	}))

	mux.HandleFunc("/api/users/search", middleware.JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET method is allowed",
			})
			return
		}
		handler.SearchUsersHandler(w, r, userRepo)
	}))

	mux.HandleFunc("/api/users/count", middleware.JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET method is allowed",
			})
			return
		}
		handler.GetUserCountHandler(w, r, userRepo)
	}))

	mux.HandleFunc("/api/users/", middleware.JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only GET method is allowed",
			})
			return
		}
		handler.GetUserByIDHandler(w, r, userRepo)
	}))
}
