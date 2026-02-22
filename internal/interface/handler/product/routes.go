package product

import (
	"encoding/json"
	"kossti/internal/domain/repository"
	handlerproductreview "kossti/internal/interface/handler/productreview"
	"net/http"
	"os"
	"strings"
)

// RegisterProductRoutes registers product-related endpoints to the mux.
func RegisterProductRoutes(mux *http.ServeMux, productRepo repository.ProductRepository, imageRepo repository.ImageRepository, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository, reviewRepo repository.ProductReviewRepository) {
	// GET /proxy-image?path=... - Proxy for backend-authenticated S3 image access
	mux.HandleFunc("/proxy-image", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodOptions {
			ProxyImageHandler(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
		}
	})

	// GET /debug/env - Diagnostic endpoint to check environment variables (remove in production!)
	mux.HandleFunc("/debug/env", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		response := map[string]interface{}{
			"GO_ENV": os.Getenv("GO_ENV"),
			"SERVER_URL": os.Getenv("SERVER_URL"),
			"S3_BUCKET": os.Getenv("S3_BUCKET"),
			"AWS_REGION": os.Getenv("AWS_REGION"),
			"AWS_ACCESS_KEY_ID": obfuscateKey(os.Getenv("AWS_ACCESS_KEY_ID")),
			"AWS_SECRET_ACCESS_KEY": obfuscateKey(os.Getenv("AWS_SECRET_ACCESS_KEY")),
		}
		json.NewEncoder(w).Encode(response)
	})

	// GET /products - List products with search, pagination, and filtering
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Check if   parameters are present (locale, page, sortby)
			if r.URL.Query().Get("locale") != "" || r.URL.Query().Get("page") != "" || r.URL.Query().Get("sortby") != "" {
				GetFilteredProductsHandler(w, r, productRepo, categoryRepo, brandRepo, imageRepo)
			} else {
				ListProductsHandler(w, r, productRepo, categoryRepo, brandRepo, imageRepo)
			}
		} else if r.Method == http.MethodPost {
			CreateProductHandler(w, r, productRepo, imageRepo)
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
				GetProductByIDHandler(w, r, productRepo, categoryRepo, brandRepo, imageRepo)
			} else if r.Method == http.MethodPatch {
				UpdateProductHandler(w, r, productRepo, imageRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only GET and PATCH methods are allowed"}`))
			}
		} else if len(pathParts) == 2 && pathParts[1] == "increment-views" {
			// POST /products/{id}/increment-views
			if r.Method == http.MethodPost {
				IncrementProductViewsHandler(w, r, productRepo, categoryRepo, brandRepo, imageRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			}
		} else if len(pathParts) == 2 && (pathParts[1] == "translation" || pathParts[1] == "translations") {
			// GET and POST /products/{id}/translation or /products/{id}/translations
			if r.Method == http.MethodGet {
				GetProductTranslationsHandler(w, r, productRepo)
			} else if r.Method == http.MethodPost {
				CreateProductTranslationHandler(w, r, productRepo)
			} else if r.Method == http.MethodDelete {
				DeleteProductTranslationHandler(w, r, productRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only GET, POST, and DELETE methods are allowed"}`))
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
		} else if len(pathParts) == 2 && pathParts[1] == "reviews" {
			// GET /products/{product}/reviews -> forward to productreview handler
			if r.Method == http.MethodGet {
				handlerproductreview.GetProductReviewsHandler(w, r, reviewRepo)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
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

		// Check for specific sub-paths manually since we are using prefix matching
		if strings.HasSuffix(r.URL.Path, "/similar") {
			GetSimilarProductsHandler(w, r, productRepo, categoryRepo, brandRepo, imageRepo)
			return
		}

		GetProductBySlugHandler(w, r, productRepo, categoryRepo, brandRepo, imageRepo)
	})

	// GET /popular-products - Get popular products
	mux.HandleFunc("/popular-products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetPopularProductsHandler(w, r, productRepo, categoryRepo, brandRepo, imageRepo)
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

	// Debug route
	mux.HandleFunc("/debug-trans", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"debug": "translation route is working"}`))
	})

	// Debug route for testing JSON decoding
	mux.HandleFunc("/debug-json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var request struct {
			Locale                string  `json:"locale"`
			TranslatedName        string  `json:"translated_name"`
			TranslatedDescription *string `json:"translated_description"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		response := map[string]interface{}{
			"received":   request,
			"locale_len": len(request.Locale),
			"name_len":   len(request.TranslatedName),
			"raw_bytes":  []byte(request.TranslatedName),
		}
		json.NewEncoder(w).Encode(response)
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

	// GET /market-products - Get new products from market (CSV)
	mux.HandleFunc("/market-products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetMarketProductsHandler(w, r)
	})
}

// obfuscateKey masks sensitive keys for debugging (shows first 4 and last 4 chars)
func obfuscateKey(key string) string {
	if key == "" {
		return "(empty)"
	}
	if len(key) <= 8 {
		return "***"
	}
	return key[:4] + "..." + key[len(key)-4:]
}
