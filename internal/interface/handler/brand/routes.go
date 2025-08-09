package brand

import (
	"kossti/internal/domain/repository"
	"net/http"
	"strings"
)

// RegisterBrandRoutes registers brand-related endpoints to the mux.
func RegisterBrandRoutes(mux *http.ServeMux, brandRepo repository.BrandRepository) {

	// Public brand endpoints

	// GET /api/brands - List all brands
	mux.HandleFunc("/api/brands", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetBrandsHandler(w, r, brandRepo)
		} else if r.Method == http.MethodPost {
			CreateBrandHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
		}
	})

	// GET/PUT/DELETE /api/brands/{id} - Brand CRUD by ID
	mux.HandleFunc("/api/brands/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/brands/")
		if path == "" || path == "/" {
			// Handle /api/brands/ - redirect to /api/brands
			if r.Method == http.MethodGet {
				GetBrandsHandler(w, r, brandRepo)
			} else if r.Method == http.MethodPost {
				CreateBrandHandler(w, r, brandRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
			}
			return
		}

		// Handle specific brand ID operations
		switch r.Method {
		case http.MethodGet:
			GetBrandByIDHandler(w, r, brandRepo)
		case http.MethodPut:
			UpdateBrandHandler(w, r, brandRepo)
		case http.MethodDelete:
			DeleteBrandHandler(w, r, brandRepo)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET, PUT, and DELETE methods are allowed"}`))
		}
	})

	// GET /api/wide-brands - Get wide brands (public access)
	mux.HandleFunc("/api/wide-brands", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetWideBrandsHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
		}
	})

	// GET /api/public-brands - Get public brands
	mux.HandleFunc("/api/public-brands", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetPublicBrandsHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
		}
	})

	// Brand translation endpoints

	// POST /api/brand-translation - Create brand translation
	mux.HandleFunc("/api/brand-translation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateBrandTranslationHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
		}
	})

	// GET /api/brand-translation/{id} - Get brand translations by brand ID
	mux.HandleFunc("/api/brand-translation/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/brand-translation/")
		if path == "" || path == "/" {
			// Handle /api/brand-translation/ - redirect to creation
			if r.Method == http.MethodPost {
				CreateBrandTranslationHandler(w, r, brandRepo)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			}
			return
		}

		// Handle specific brand translation operations
		if r.Method == http.MethodGet {
			GetBrandTranslationHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
		}
	})

	// Brand status endpoints

	// POST /api/brand-status/{id} - Update brand status
	mux.HandleFunc("/api/brand-status/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/brand-status/")
		if path == "" || path == "/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Brand ID required for status update"}`))
			return
		}

		// Handle brand status update
		if r.Method == http.MethodPost {
			UpdateBrandStatusHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
		}
	})
}
