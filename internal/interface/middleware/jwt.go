package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // Use the same key as in your auth logic

type contextKey string

const userIDKey contextKey = "user_id"

// JWTAuthMiddleware checks for a valid JWT token in the Authorization header.
func JWTAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)

			w.Write([]byte(`{"error": "Missing or invalid Authorization header"}`))
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Invalid or expired token"}`))
			return
		}

		// Extract claims and add user_id to context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userID, exists := claims["user_id"]; exists {
				ctx := context.WithValue(r.Context(), userIDKey, userID)
				r = r.WithContext(ctx)
			}
		}

		next(w, r)
	}
}

// GetUserIDFromContext extracts user ID from request context
func GetUserIDFromContext(r *http.Request) (uint, error) {
	userIDValue := r.Context().Value(userIDKey)
	if userIDValue == nil {
		return 0, http.ErrNoCookie
	}

	// Handle different types that might come from JWT claims
	switch v := userIDValue.(type) {
	case float64:
		return uint(v), nil
	case int:
		return uint(v), nil
	case string:
		if id, err := strconv.ParseUint(v, 10, 32); err == nil {
			return uint(id), nil
		}
	}
	return 0, http.ErrNoCookie
}
