package product

import (
	"kossti/internal/domain/repository"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterProductRoutes registers product-related endpoints to the mux.
func RegisterProductRoutes(r *mux.Router, productRepo repository.ProductRepository) {
	r.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		CreateProductHandler(w, r, productRepo)
	}).Methods("POST")

	r.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Only GET method is allowed"}`))
			return
		}
		GetProductByIDHandler(w, r, productRepo)
	})
}
