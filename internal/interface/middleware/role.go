package middleware

import (
	"context"
	"encoding/json"
	"kossti/internal/domain/repository"
	"net/http"
)

const userTypeKey contextKey = "user_type"

// RequireRole is a middleware that checks if the user has the required role
func RequireRole(userRepo repository.UserRepository, allowedRoles ...string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// First, get the user ID from the JWT middleware
			userID, err := GetUserIDFromContext(r)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
				return
			}

			// Get the user from the database
			user, err := userRepo.GetByID(r.Context(), userID)
			if err != nil || user == nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
				return
			}

			// Check if user's type is in the allowed roles
			allowed := false
			for _, role := range allowedRoles {
				if user.Type == role {
					allowed = true
					break
				}
			}

			if !allowed {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{
					"error":     "Access denied. Required role: " + allowedRoles[0],
					"your_role": user.Type,
				})
				return
			}

			// Add user type to context for downstream handlers
			ctx := context.WithValue(r.Context(), userTypeKey, user.Type)
			r = r.WithContext(ctx)

			next(w, r)
		}
	}
}

// GetUserTypeFromContext extracts user type from request context
func GetUserTypeFromContext(r *http.Request) (string, error) {
	userType := r.Context().Value(userTypeKey)
	if userType == nil {
		return "", http.ErrNoCookie
	}

	if typeStr, ok := userType.(string); ok {
		return typeStr, nil
	}

	return "", http.ErrNoCookie
}
