package specification

import (
	"kossti/internal/domain/repository"
	"net/http"
	"strings"
)

// RegisterSpecificationRoutes registers specification-related endpoints to the mux.
func RegisterSpecificationRoutes(mux *http.ServeMux, specRepo repository.SpecificationRepository, keyRepo repository.SpecificationKeyRepository) {

	// Public specification endpoints

	// POST /api/specifications - Create specification
	mux.HandleFunc("/api/specifications", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		CreateSpecificationHandler(w, r, specRepo, keyRepo)
	})

	// POST /api/specifications/bulk - Bulk upsert specifications
	mux.HandleFunc("/api/specifications/bulk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		BulkUpsertSpecificationHandler(w, r, specRepo, keyRepo)
	})

	// GET /api/get-specifications/{id} - Get specifications by product ID
	mux.HandleFunc("/api/get-specifications/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetSpecificationsByProductHandler(w, r, specRepo)
	})

	// GET/PUT /api/specifications/{id} - Get or update specification by ID
	mux.HandleFunc("/api/specifications/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetSpecificationByIDHandler(w, r, specRepo)
		} else if r.Method == http.MethodPut {
			UpdateSpecificationHandler(w, r, specRepo, keyRepo)
		} else if r.Method == http.MethodDelete {
			DeleteSpecificationHandler(w, r, specRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET, PUT, and DELETE methods are allowed"}`))
		}
	})

	// GET /api/specificationsearch - Search specifications
	mux.HandleFunc("/api/specificationsearch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		SearchSpecificationsHandler(w, r, specRepo)
	})

	// POST /api/spec_translation - Create specification translation
	mux.HandleFunc("/api/spec_translation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		CreateSpecificationTranslationHandler(w, r, specRepo)
	})

	// GET /api/spec_translation/{id} - Get specification translations
	mux.HandleFunc("/api/spec_translation/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetSpecificationTranslationHandler(w, r, specRepo)
	})

	// GET /api/get-public-spec/{id} - Get public specification
	mux.HandleFunc("/api/get-public-spec/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetPublicSpecificationHandler(w, r, specRepo)
	})

	// Specification Key endpoints (these would typically be authenticated)

	// POST /api/speckey - Create or update specification key
	mux.HandleFunc("/api/speckey", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateOrUpdateSpecificationKeyHandler(w, r, keyRepo)
		} else if r.Method == http.MethodGet {
			GetAllSpecificationKeysHandler(w, r, keyRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
		}
	})

	// GET /api/speckey/{id} - Get specification key by ID
	mux.HandleFunc("/api/speckey/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/speckey/")
		if path == "" || path == "/" {
			// This is the base /api/speckey/ endpoint, redirect to /api/speckey
			http.Redirect(w, r, "/api/speckey", http.StatusMovedPermanently)
			return
		}

		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetSpecificationKeyByIDHandler(w, r, keyRepo)
	})

	// POST /api/specremove/{id} - Delete specification key
	mux.HandleFunc("/api/specremove/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		DeleteSpecificationKeyHandler(w, r, keyRepo)
	})

	// Specification Key Translation endpoints

	// POST /api/speckey-translation - Create specification key translation
	// GET /api/speckey-translation - Get all specification key translations
	mux.HandleFunc("/api/speckey-translation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateSpecificationKeyTranslationHandler(w, r, keyRepo)
		} else if r.Method == http.MethodGet {
			GetAllSpecificationKeyTranslationsHandler(w, r, keyRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
		}
	})
}
