package productreview

import (
	"kossti/internal/domain/repository"
	"net/http"
)

// RegisterProductReviewRoutes registers all product review related routes
func RegisterProductReviewRoutes(mux *http.ServeMux, reviewRepo repository.ProductReviewRepository) {
	// Protected routes - require authentication
	mux.HandleFunc("POST /api/reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		CreateReviewHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("GET /api/reviews", func(w http.ResponseWriter, r *http.Request) {
		GetAllReviewsHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /api/review/translation", func(w http.ResponseWriter, r *http.Request) {
		CreateReviewTranslationHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /api/product/{id}/review/{reviewid}", func(w http.ResponseWriter, r *http.Request) {
		UpdateReviewHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /api/imageremove/{id}", func(w http.ResponseWriter, r *http.Request) {
		RemoveImageHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /api/productimages", func(w http.ResponseWriter, r *http.Request) {
		UploadImagesHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("POST /api/default-image/{id}", func(w http.ResponseWriter, r *http.Request) {
		MakeDefaultImageHandler(w, r, reviewRepo)
	})

	// Public routes - no authentication required
	mux.HandleFunc("GET /api/products/{productId}/reviews", func(w http.ResponseWriter, r *http.Request) {
		GetProductReviewsHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("GET /api/reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetReviewHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("GET /api/productimages/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetProductImagesHandler(w, r, reviewRepo)
	})

	mux.HandleFunc("GET /api/public-reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetPublicReviewsHandler(w, r, reviewRepo)
	})
}
