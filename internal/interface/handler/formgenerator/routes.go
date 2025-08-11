package formgenerator

import (
	"kossti/internal/domain/repository"
	"net/http"
)

// RegisterRoutes registers all form generator routes
func RegisterRoutes(mux *http.ServeMux, repo repository.FormGeneratorRepository) {
	handler := NewFormGeneratorHandler(repo)

	// Protected routes (TODO: Add JWT middleware)
	mux.HandleFunc("/formgenerator", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateFormGenerator(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/formgenerator/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetFormGeneratorByID(w, r)
		case http.MethodPut:
			handler.UpdateFormGenerator(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// V1 API routes (for compatibility)
	mux.HandleFunc("/v1/formgenerator", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateFormGenerator(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/formgenerator/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetFormGeneratorByID(w, r)
		case http.MethodPut:
			handler.UpdateFormGenerator(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Public routes
	mux.HandleFunc("/catgory-specs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetCategorySpec(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/catgory-specs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetCategorySpec(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
