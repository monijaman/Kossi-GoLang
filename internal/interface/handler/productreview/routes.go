package productreview

import (
	"kossti/internal/domain/repository"
	"log"
	"net/http"
	"strings"
)

// RegisterProductReviewRoutes registers all product review related routes
func RegisterProductReviewRoutes(mux *http.ServeMux, reviewRepo repository.ProductReviewRepository, productRepo repository.ProductRepository, imageRepo repository.ImageRepository) {
	// Protected routes - require authentication

	// Create review -> POST /reviews/{id}
	mux.HandleFunc("/reviews/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateReviewHandler(w, r, reviewRepo, productRepo)
			return
		}

		if r.Method == http.MethodGet {
			// Could be GET /reviews/{id}
			GetReviewHandler(w, r, reviewRepo)
			return
		}

		http.NotFound(w, r)
	})

	// List all reviews -> GET /reviews
	mux.HandleFunc("/reviews", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetAllReviewsHandler(w, r, reviewRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Get reviews by product ID -> GET /product-reviews/{product_id}
	mux.HandleFunc("/product-reviews/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetReviewsByProductHandler(w, r, reviewRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Create translation -> POST /review/translation
	mux.HandleFunc("/review/translation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateReviewTranslationHandler(w, r, reviewRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Presign S3 PUT URL -> POST /s3/presign
	mux.HandleFunc("/s3/presign", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			PresignS3Handler(w, r)
			return
		}
		http.NotFound(w, r)
	})

	// Presign S3 GET URL -> POST /s3/presign-get
	mux.HandleFunc("/s3/presign-get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			PresignGetHandler(w, r)
			return
		}
		http.NotFound(w, r)
	})

	// Update review -> POST /product/{id}/review/{reviewid}
	mux.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && strings.Contains(r.URL.Path, "/review/") {
			UpdateReviewHandler(w, r, reviewRepo)
			return
		}

		// Image related endpoints under /product/ path are handled elsewhere
		http.NotFound(w, r)
	})

	// Remove image -> POST /imageremove/{id}
	mux.HandleFunc("/imageremove/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			RemoveImageHandler(w, r, reviewRepo, imageRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Upload images -> POST /productimages
	mux.HandleFunc("/productimages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			UploadImagesHandler(w, r, reviewRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Register S3 uploaded images -> POST /productimages/s3
	mux.HandleFunc("/productimages/s3", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[route] /productimages/s3 -> %s %s", r.Method, r.URL.Path)
		if r.Method == http.MethodPost {
			RegisterS3ImagesHandler(w, r, reviewRepo, imageRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Accept trailing-slash variant as well to avoid ServeMux prefix/routing surprises
	mux.HandleFunc("/productimages/s3/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[route] /productimages/s3/ -> %s %s", r.Method, r.URL.Path)
		if r.Method == http.MethodPost {
			RegisterS3ImagesHandler(w, r, reviewRepo, imageRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Make default image -> POST /default-image/{id}
	mux.HandleFunc("/default-image/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			MakeDefaultImageHandler(w, r, reviewRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Note: product review GET /products/{id}/reviews is routed from the product handler
	// to avoid registering the same "/products/" pattern twice on the default ServeMux.

	// Get product images -> GET /productimages/{id}
	mux.HandleFunc("/productimages/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetProductImagesHandler(w, r, reviewRepo, imageRepo)
			return
		}
		http.NotFound(w, r)
	})

	// Note: GET /public-reviews/{id} is handled by the product routes to avoid duplicate ServeMux registration
}
