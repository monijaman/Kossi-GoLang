package brand

import (
	"fmt"
	"kossti/internal/domain/repository"
	"net/http"
	"strings"
)

// RegisterBrandRoutes registers brand-related endpoints to the mux.
func RegisterBrandRoutes(mux *http.ServeMux, brandRepo repository.BrandRepository) {

	// Public brand endpoints

	// GET/POST /brands - Brand listing and creation
	mux.HandleFunc("/brands", func(w http.ResponseWriter, r *http.Request) {
		// Handle /brands (without trailing slash)
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

	// GET/PUT/DELETE /brands/{id} - Brand CRUD by ID
	mux.HandleFunc("/brands/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/brands/")

		fmt.Println("------------Handling request for path:", path)
		if path == "" {
			// Handle /brands/ - list all brands or create new brand
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

	// GET /wide-brands - Get wide brands (public access)
	mux.HandleFunc("/wide-brands", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetWideBrandsHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
		}
	})

	// GET /public-brands - Get public brands
	mux.HandleFunc("/public-brands", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetPublicBrandsHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
		}
	})

	// Brand translation endpoints

	// POST /brand-translation - Create brand translation
	mux.HandleFunc("/brand-translation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateBrandTranslationHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
		}
	})

	// GET /brand-translation/{id} - Get brand translations by brand ID
	mux.HandleFunc("/brand-translation/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/brand-translation/")

		if path == "" || path == "/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Brand ID required for status update"}`))
			return
		}

		// Handle /brand-translation/ - redirect to creation
		if r.Method == http.MethodPut {
			UpdateBrandTranslationHandler(w, r, brandRepo)
		} else if r.Method == http.MethodGet {
			GetBrandTranslationHandler(w, r, brandRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET, PUT methods are allowed"}`))
		}
	})

	// Brand status endpoints

	// POST /brand-status/{id} - Update brand status
	mux.HandleFunc("/brand-status/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/brand-status/")
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
