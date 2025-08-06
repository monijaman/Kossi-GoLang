package product

import (
	"kossti/internal/domain/repository"
	"net/http"
)

// RegisterProductRoutes registers product-related endpoints to the mux.
func RegisterProductRoutes(mux *http.ServeMux, productRepo repository.ProductRepository) {
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateProductHandler(w, r, productRepo)
		} else if r.Method == http.MethodGet {
			ListProductsHandler(w, r, productRepo)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET and POST methods are allowed"}`))
		}
	})

	mux.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetProductByIDHandler(w, r, productRepo)
	})
}
