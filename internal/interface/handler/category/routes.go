package category

import (
	"kossti/internal/domain/repository"
	"net/http"
	"strings"
)

// RegisterCategoryRoutes registers category-related endpoints to the mux.
func RegisterCategoryRoutes(mux *http.ServeMux, categoryRepo repository.CategoryRepository) {

	// Public category endpoints

	// GET /api/categories - List all categories
	mux.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetCategoriesHandler(w, r, categoryRepo)
		} else if r.Method == http.MethodPost {
			CreateCategoryHandler(w, r, categoryRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
		}
	})

	// GET/PUT/DELETE /api/categories/{id} - Category CRUD by ID
	mux.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/categories/")
		if path == "" || path == "/" {
			// This is the base /api/categories/ endpoint, redirect to /api/categories
			http.Redirect(w, r, "/api/categories", http.StatusMovedPermanently)
			return
		}

		if r.Method == http.MethodGet {
			GetCategoryByIDHandler(w, r, categoryRepo)
		} else if r.Method == http.MethodPut {
			UpdateCategoryHandler(w, r, categoryRepo)
		} else if r.Method == http.MethodDelete {
			DeleteCategoryHandler(w, r, categoryRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET, PUT, and DELETE methods are allowed"}`))
		}
	})

	// GET /api/wide-categories - Get wide categories
	mux.HandleFunc("/api/wide-categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetWideCategoriesHandler(w, r, categoryRepo)
	})

	// POST /api/category-translation - Create category translation
	mux.HandleFunc("/api/category-translation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		CreateCategoryTranslationHandler(w, r, categoryRepo)
	})

	// GET /api/category-translation/{id} - Get category translations
	mux.HandleFunc("/api/category-translation/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetCategoryTranslationHandler(w, r, categoryRepo)
	})

	// POST/GET /api/category-brands - Category-Brand relations
	mux.HandleFunc("/api/category-brands", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateCategoryBrandRelationHandler(w, r, categoryRepo)
		} else if r.Method == http.MethodGet {
			GetCategoryBrandRelationsHandler(w, r, categoryRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
		}
	})

	// POST /api/category-status/{id} - Update category status
	mux.HandleFunc("/api/category-status/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		UpdateCategoryStatusHandler(w, r, categoryRepo)
	})
}
