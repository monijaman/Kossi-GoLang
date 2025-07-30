package product

import (
	"encoding/json"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"net/http"
	"strconv"
)

// Handler for GET /api/product/{id} (public)
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Path[len("/api/product/"):] // crude extraction
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	// TODO: Use repo.GetByID for real DB fetch
	// product, err := repo.GetByID(r.Context(), uint(productID))
	// if err != nil || product == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
	// 	return
	// }

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":   productID,
		"name": "Sample Product",
	})
}

// Handler for POST /api/products (public, for demo; add auth as needed)
func CreateProductHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Only POST method is allowed"})
		return
	}

	var req entities.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	product, err := repo.Create(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
