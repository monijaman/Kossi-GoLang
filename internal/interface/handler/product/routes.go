package product

import (
	"kossti/internal/domain/repository"
	"net/http"
	"strings"
)

// RegisterProductRoutes registers product-related endpoints to the mux.
func RegisterProductRoutes(mux *http.ServeMux, productRepo repository.ProductRepository) {
	// GET /api/products - List products with search, pagination, and filtering
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListProductsHandler(w, r, productRepo)
		} else if r.Method == http.MethodPost {
			CreateProductHandler(w, r, productRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
		}
	})

	// GET /api/products/{id} - Get product by ID
	mux.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/products/")
		pathParts := strings.Split(path, "/")

		if len(pathParts) == 1 && pathParts[0] != "" {
			// Single ID or slug
			if r.Method == http.MethodGet {
				GetProductByIDHandler(w, r, productRepo)
			} else if r.Method == http.MethodPatch {
				UpdateProductHandler(w, r, productRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only GET and PATCH methods are allowed"}`))
			}
		} else if len(pathParts) == 2 && pathParts[1] == "increment-views" {
			// POST /api/products/{id}/increment-views
			if r.Method == http.MethodPost {
				IncrementProductViewsHandler(w, r, productRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Endpoint not found"}`))
		}
	})

	// GET /api/products-by-slug/{slug} - Get product by slug (alternative endpoint)
	mux.HandleFunc("/api/products-by-slug/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetProductBySlugHandler(w, r, productRepo)
	})

	// GET /api/popular-products - Get popular products
	mux.HandleFunc("/api/popular-products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetPopularProductsHandler(w, r, productRepo)
	})
}
