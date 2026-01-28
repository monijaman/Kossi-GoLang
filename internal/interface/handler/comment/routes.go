package handler

import (
	"net/http"
	"strings"

	pgRepo "kossti/internal/interface/repository/postgres"

	"gorm.io/gorm"
)

// RegisterCommentRoutes registers all comment-related routes
func RegisterCommentRoutes(mux *http.ServeMux, db *gorm.DB) {
	commentRepo := pgRepo.NewPostgresCommentRepo(db)
	translationRepo := pgRepo.NewPostgresCommentTranslationRepo(db)

	// POST /comments - Create a new comment
	mux.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateCommentHandler(w, r, commentRepo)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// POST /comments/upsert - Create or update a comment
	mux.HandleFunc("/comments/upsert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			UpsertCommentHandler(w, r, commentRepo)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// GET /comments/{id} or GET /comments/product/{id}
	mux.HandleFunc("/comments/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/comments/")

		// Check if it's a product ID query
		if strings.HasPrefix(path, "product/") {
			GetCommentsByProductIDHandler(w, r, commentRepo)
		} else {
			// Single comment by ID
			GetCommentByIDHandler(w, r, commentRepo)
		}
	})

	// POST /comment-translations - Create a new comment translation
	mux.HandleFunc("/comment-translations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateCommentTranslationHandler(w, r, translationRepo)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// POST/PUT /comment-translations/upsert - Create or update a comment translation
	mux.HandleFunc("/comment-translations/upsert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			UpsertCommentTranslationHandler(w, r, translationRepo)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// GET /comment-translations/{id} - Get translations for a comment
	mux.HandleFunc("/comment-translations/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		GetCommentTranslationsHandler(w, r, translationRepo)
	})
}
