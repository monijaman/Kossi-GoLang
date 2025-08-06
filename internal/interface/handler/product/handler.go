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

// Handler for GET /api/products (list all products)
func ListProductsHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse optional query parameters for pagination
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10 // default
	offset := 0 // default

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	// TODO: Use repo.List for real DB fetch
	// products, err := repo.List(r.Context(), limit, offset)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	// 	return
	// }

	// For now, return sample data
	sampleProducts := []map[string]interface{}{
		{"id": 1, "name": "Sample Product 1", "price": 99.99},
		{"id": 2, "name": "Sample Product 2", "price": 149.99},
		{"id": 3, "name": "Sample Product 3", "price": 199.99},
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"products": sampleProducts,
		"count":    len(sampleProducts),
		"limit":    limit,
		"offset":   offset,
	})
}
