package admin

import (
	"encoding/json"
	"kossti/internal/domain/repository"
	"kossti/internal/interface/middleware"
	"net/http"
)

// RegisterAdminRoutes registers admin-only routes
func RegisterAdminRoutes(mux *http.ServeMux, userRepo repository.UserRepository, productRepo repository.ProductRepository) {
	// Example: Admin-only endpoint to delete a product
	// This endpoint requires JWT authentication AND admin role
	mux.HandleFunc("/admin/products/", middleware.JWTAuthMiddleware(
		middleware.RequireRole(userRepo, "admin")(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			
			if r.Method == http.MethodDelete {
				// Extract product ID from path
				// Implementation would go here
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Product deleted successfully (admin only)",
				})
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Only DELETE method is allowed",
				})
			}
		}),
	))

	// Example: Get admin dashboard stats (admin only)
	mux.HandleFunc("/admin/stats", middleware.JWTAuthMiddleware(
		middleware.RequireRole(userRepo, "admin")(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Only GET method is allowed",
				})
				return
			}

			// Get user type from context (set by middleware)
			userType, _ := middleware.GetUserTypeFromContext(r)
			
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Admin dashboard stats",
				"your_role": userType,
				"stats": map[string]int{
					"total_products": 100,
					"total_users": 50,
				},
			})
		}),
	))

	// Example: Reviewer can access this, but not guests
	mux.HandleFunc("/admin/reviews/pending", middleware.JWTAuthMiddleware(
		middleware.RequireRole(userRepo, "admin", "reviewer")(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Only GET method is allowed",
				})
				return
			}

			userType, _ := middleware.GetUserTypeFromContext(r)
			
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Pending reviews (admin and reviewer only)",
				"your_role": userType,
				"reviews": []string{},
			})
		}),
	))
}
