package product

import (
	"kossti/internal/domain/repository"
	"net/http"
	"strings"
)

// RegisterProductRoutes registers product-related endpoints to the mux.
func RegisterProductRoutes(mux *http.ServeMux, productRepo repository.ProductRepository, imageRepo repository.ImageRepository) {
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

	// GET /api/products/{id} - Get product by ID, with support for multiple sub-paths
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
		} else if len(pathParts) == 2 && pathParts[1] == "image" {
			// POST /api/products/{product}/image
			if r.Method == http.MethodPost {
				AddProductImageAltHandler(w, r, imageRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			}
		} else if len(pathParts) == 2 && pathParts[1] == "images" {
			// GET /api/products/{product}/images
			if r.Method == http.MethodGet {
				GetProductImagesHandler(w, r, imageRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only GET method is allowed"}`))
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

	// Image upload endpoints
	// POST /api/addproductimage/{productId}
	mux.HandleFunc("/api/addproductimage/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		AddProductImageHandler(w, r, imageRepo)
	})

	// GET /api/get-product-image/{productId}
	mux.HandleFunc("/api/get-product-image/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetProductImageHandler(w, r, imageRepo)
	})

	// POST /api/product-trans/{id} - Product translation
	mux.HandleFunc("/api/product-trans/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		CreateProductTranslationHandler(w, r, productRepo)
	})
}
