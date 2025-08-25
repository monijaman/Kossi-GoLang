package specification

import (
	"kossti/internal/domain/repository"
	"net/http"
	"strings"
)

// RegisterSpecificationRoutes registers specification-related endpoints to the mux.
func RegisterSpecificationRoutes(mux *http.ServeMux, specRepo repository.SpecificationRepository, keyRepo repository.SpecificationKeyRepository) {

	// Public specification endpoints

	// POST /specifications - Create specification
	mux.HandleFunc("/specifications", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		CreateSpecificationHandler(w, r, specRepo, keyRepo)
	})

	// POST /specifications/bulk - Bulk upsert specifications
	mux.HandleFunc("/specifications/bulk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		BulkUpsertSpecificationHandler(w, r, specRepo, keyRepo)
	})

	// GET /get-specifications/{id} - Get specifications by product ID
	mux.HandleFunc("/get-specifications/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetSpecificationsByProductHandler(w, r, specRepo)
	})

	// GET/PUT /specifications/{id} - Get or update specification by ID
	mux.HandleFunc("/specifications/", func(w http.ResponseWriter, r *http.Request) {
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

	// GET /specificationsearch - Search specifications
	mux.HandleFunc("/specificationsearch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		SearchSpecificationsHandler(w, r, specRepo)
	})

	// POST /spec_translation - Create specification translation
	mux.HandleFunc("/spec_translation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		CreateSpecificationTranslationHandler(w, r, specRepo)
	})

	// GET /spec_translation/{id} - Get specification translations
	mux.HandleFunc("/spec_translation/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetSpecificationTranslationHandler(w, r, specRepo)
	})

	// GET /get-public-spec/{id} - Get public specification
	mux.HandleFunc("/get-public-spec/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetPublicSpecificationHandler(w, r, specRepo, keyRepo)
	})

	// Specification Key endpoints (these would typically be authenticated)

	// POST /speckey - Create or update specification key
	mux.HandleFunc("/speckey", func(w http.ResponseWriter, r *http.Request) {
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

	// GET /speckey/{id} - Get specification key by ID
	mux.HandleFunc("/speckey/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/speckey/")
		if path == "" || path == "/" {
			// This is the base /speckey/ endpoint, redirect to /speckey with query parameters preserved
			redirectURL := "/speckey"
			if r.URL.RawQuery != "" {
				redirectURL += "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
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

	// POST /specremove/{id} - Delete specification key
	mux.HandleFunc("/specremove/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only POST method is allowed"}`))
			return
		}
		DeleteSpecificationKeyHandler(w, r, keyRepo)
	})

	// Specification Key Translation endpoints

	// POST /speckey-translation - Create specification key translation
	// GET /speckey-translation - Get all specification key translations
	mux.HandleFunc("/speckey-translation", func(w http.ResponseWriter, r *http.Request) {
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
