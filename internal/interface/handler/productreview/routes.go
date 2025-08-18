package productreview

import (
	"kossti/internal/domain/repository"
	"net/http"
)

// RegisterProductReviewRoutes registers all product review related routes
func RegisterProductReviewRoutes(mux *http.ServeMux, reviewRepo repository.ProductReviewRepository) {
	// Protected routes - require authentication
	mux.HandleFunc("POST /reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		CreateReviewHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("GET /reviews", func(w http.ResponseWriter, r *http.Request) {
		GetAllReviewsHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /review/translation", func(w http.ResponseWriter, r *http.Request) {
		CreateReviewTranslationHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /product/{id}/review/{reviewid}", func(w http.ResponseWriter, r *http.Request) {
		UpdateReviewHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /imageremove/{id}", func(w http.ResponseWriter, r *http.Request) {
		RemoveImageHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /productimages", func(w http.ResponseWriter, r *http.Request) {
		UploadImagesHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /default-image/{id}", func(w http.ResponseWriter, r *http.Request) {
		MakeDefaultImageHandler(w, r, reviewRepo)
	})

	// Public routes - no authentication required
	mux.HandleFunc("GET /products/{productId}/reviews", func(w http.ResponseWriter, r *http.Request) {
		GetProductReviewsHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("GET /reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetReviewHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("GET /productimages/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetProductImagesHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("GET /public-reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetPublicReviewsHandler(w, r, reviewRepo)
	})
}
