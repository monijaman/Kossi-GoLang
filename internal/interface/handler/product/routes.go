package product

import (
	"kossti/internal/domain/repository"
	"net/http"
	"strings"
)

// RegisterProductRoutes registers product-related endpoints to the mux.
func RegisterProductRoutes(mux *http.ServeMux, productRepo repository.ProductRepository, imageRepo repository.ImageRepository, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) {
	// GET /products - List products with search, pagination, and filtering
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Check if Laravel-style parameters are present (locale, page, sortby)
			if r.URL.Query().Get("locale") != "" || r.URL.Query().Get("page") != "" || r.URL.Query().Get("sortby") != "" {
				GetFilteredProductsHandler(w, r, productRepo, categoryRepo, brandRepo)
			} else {
				ListProductsHandler(w, r, productRepo, categoryRepo, brandRepo)
			}
		} else if r.Method == http.MethodPost {
			CreateProductHandler(w, r, productRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
		}
	})

	// GET /products/{id} - Get product by ID, with support for multiple sub-paths
	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/products/")
		pathParts := strings.Split(path, "/")

		if len(pathParts) == 1 && pathParts[0] != "" {
			// Single ID or slug
			if r.Method == http.MethodGet {
				GetProductByIDHandler(w, r, productRepo, categoryRepo, brandRepo)
			} else if r.Method == http.MethodPatch {
				UpdateProductHandler(w, r, productRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only GET and PATCH methods are allowed"}`))
			}
		} else if len(pathParts) == 2 && pathParts[1] == "increment-views" {
			// POST /products/{id}/increment-views
			if r.Method == http.MethodPost {
				IncrementProductViewsHandler(w, r, productRepo, categoryRepo, brandRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			}
		} else if len(pathParts) == 2 && pathParts[1] == "image" {
			// POST /products/{product}/image
			if r.Method == http.MethodPost {
				AddProductImageAltHandler(w, r, imageRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			}
		} else if len(pathParts) == 2 && pathParts[1] == "images" {
			// GET /products/{product}/images
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

	// GET /products-by-slug/{slug} - Get product by slug (alternative endpoint)
	mux.HandleFunc("/products-by-slug/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetProductBySlugHandler(w, r, productRepo, categoryRepo, brandRepo)
	})

	// GET /popular-products - Get popular products
	mux.HandleFunc("/popular-products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetPopularProductsHandler(w, r, productRepo, categoryRepo, brandRepo)
	})

	// Image upload endpoints
	// POST /addproductimage/{productId}
	mux.HandleFunc("/addproductimage/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		AddProductImageHandler(w, r, imageRepo)
	})

	// GET /get-product-image/{productId}
	mux.HandleFunc("/get-product-image/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetProductImageHandler(w, r, imageRepo)
	})

	// POST /product-trans/{id} - Product translation
	mux.HandleFunc("/product-trans/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		CreateProductTranslationHandler(w, r, productRepo)
	})

	// GET /public-reviews/{id} - Get public reviews for a product (Laravel compatibility)
	mux.HandleFunc("/public-reviews/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetPublicReviewsHandler(w, r, productRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
		}
	})
}
