package brand

import (
	"encoding/json"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// BrandResponse represents the response format for brands
type BrandResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// BrandTranslationResponse represents the response format for brand translations
type BrandTranslationResponse struct {
	ID             uint   `json:"id"`
	BrandID        uint   `json:"brand_id"`
	Locale         string `json:"locale"`
	TranslatedName string `json:"translated_name"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// convertBrandToResponse converts entity to response format
func convertBrandToResponse(brand *entities.Brand) BrandResponse {
	return BrandResponse{
		ID:        brand.ID,
		Name:      brand.Name,
		Slug:      brand.Slug,
		Status:    brand.Status,
		CreatedAt: brand.CreatedAt.Format(time.RFC3339),
		UpdatedAt: brand.UpdatedAt.Format(time.RFC3339),
	}
}

// convertBrandTranslationToResponse converts entity to response format
func convertBrandTranslationToResponse(translation *entities.BrandTranslation) BrandTranslationResponse {
	return BrandTranslationResponse{
		ID:             translation.ID,
		BrandID:        translation.BrandID,
		Locale:         translation.Locale,
		TranslatedName: translation.TranslatedName,
		CreatedAt:      translation.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      translation.UpdatedAt.Format(time.RFC3339),
	}
}

// CreateBrandHandler handles POST /brands
func CreateBrandHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		Name string `json:"name"`
		Slug string `json:"slug,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "name is required"})
		return
	}

	// Generate slug if not provided
	if request.Slug == "" {
		request.Slug = strings.ToLower(strings.ReplaceAll(request.Name, " ", "-"))
	}

	brand := &entities.Brand{
		Name:      request.Name,
		Slug:      request.Slug,
		Status:    1, // default to active
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	savedBrand, err := brandRepo.Create(r.Context(), brand)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create brand"})
		return
	}

	response := convertBrandToResponse(savedBrand)
	json.NewEncoder(w).Encode(response)
}

func GetBrandsHandlero(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	search := r.URL.Query().Get("search")

	// Set defaults
	limit := 50
	offset := 0

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

	var brands []*entities.Brand
	var err error

	if search != "" {
		brands, err = brandRepo.Search(r.Context(), search, limit, offset)
	} else {
		brands, err = brandRepo.GetAll(r.Context(), limit, offset)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get brands"})
		return
	}

	responses := make([]BrandResponse, len(brands))
	for i, brand := range brands {
		responses[i] = convertBrandToResponse(brand)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"brands": responses,
		"count":  len(responses),
		"limit":  limit,
		"offset": offset,
	})
}

// GetBrandsHandler handles GET /brands
func GetBrandsHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	// func GetWideCategoriesHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	perPageStr := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")
	paginateStr := r.URL.Query().Get("paginate")
	locale := r.URL.Query().Get("locale")
	categoryIDStr := r.URL.Query().Get("category_id")
	statusStr := r.URL.Query().Get("status")
	pageStr := r.URL.Query().Get("page")

	// Set defaults
	perPage := 10
	if perPageStr != "" {
		if p, err := strconv.Atoi(perPageStr); err == nil && p > 0 {
			perPage = p
		}
	}

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if locale == "" {
		locale = "en" // default locale
	}

	paginate := paginateStr == "true"

	// Calculate offset for pagination
	offset := 0
	if paginate && page > 1 {
		offset = (page - 1) * perPage
	}

	var brands []*entities.Brand
	var err error

	// Use search if provided, otherwise get all categories
	if search != "" {
		brands, err = brandRepo.Search(r.Context(), search, perPage, offset)
	} else {
		// For GetAll, we need to handle pagination ourselves if paginate is true
		if paginate {
			brands, err = brandRepo.GetAll(r.Context(), perPage, offset)
		} else {
			brands, err = brandRepo.GetWideBrands(r.Context(), 0) // 0 means no limit
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get wide categories"})
		return
	}

	// Apply additional filters if needed
	filteredBrands := brands

	// Filter by category ID if provided
	if categoryIDStr != "" {
		if categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32); err == nil && categoryID > 0 {
			var filtered []*entities.Brand
			for _, brand := range filteredBrands {
				if brand.ID == uint(categoryID) {
					filtered = append(filtered, brand)
				}
			}
			filteredBrands = filtered
		}
	}

	// Filter by status if provided
	if statusStr != "" {
		if status, err := strconv.Atoi(statusStr); err == nil {
			var filtered []*entities.Brand
			for _, brand := range filteredBrands {
				if brand.Status == status {
					filtered = append(filtered, brand)
				}
			}
			filteredBrands = filtered
		}
	}

	responses := make([]BrandResponse, len(filteredBrands))
	for i, brand := range filteredBrands {
		responses[i] = convertBrandToResponse(brand)
	}

	var total int
	if paginate {
		// Get total count without limit/offset for pagination
		if search != "" {
			// Count all matching search results
			allBrands, err := brandRepo.Search(r.Context(), search, 0, 0)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get wide categories"})
				return
			}
			total = len(allBrands)
		} else {
			// Count all brands (no limit/offset)
			allBrands, err := brandRepo.GetAll(r.Context(), 0, 0)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get wide categories"})
				return
			}
			total = len(allBrands)
		}
	} else {
		total = len(responses)
	}

	if paginate {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"brands":  responses,
			"total":   total,
			"success": true,
		})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"brands": responses,
			"count":  len(responses),
		})
	}

}

// GetBrandByIDHandler handles GET /brands/{id}
func GetBrandByIDHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract brand ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/brands/")
	brandIDStr := strings.Trim(path, "/")

	brandID, err := strconv.ParseUint(brandIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid brand ID format",
			"received": brandIDStr,
		})
		return
	}

	brand, err := brandRepo.GetByID(r.Context(), uint(brandID))
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Brand not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get brand"})
		}
		return
	}

	response := convertBrandToResponse(brand)
	json.NewEncoder(w).Encode(response)
}

// UpdateBrandHandler handles PUT /brands/{id}
func UpdateBrandHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract brand ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/brands/")
	brandIDStr := strings.Trim(path, "/")

	brandID, err := strconv.ParseUint(brandIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid brand ID format",
			"received": brandIDStr,
		})
		return
	}

	var request struct {
		Name string `json:"name"`
		Slug string `json:"slug,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "name is required"})
		return
	}

	// Generate slug if not provided
	if request.Slug == "" {
		request.Slug = strings.ToLower(strings.ReplaceAll(request.Name, " ", "-"))
	}

	brand := &entities.Brand{
		Name:      request.Name,
		Slug:      request.Slug,
		UpdatedAt: time.Now(),
	}

	savedBrand, err := brandRepo.Update(r.Context(), uint(brandID), brand)
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Brand not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update brand"})
		}
		return
	}

	response := convertBrandToResponse(savedBrand)
	json.NewEncoder(w).Encode(response)
}

// DeleteBrandHandler handles DELETE /brands/{id}
func DeleteBrandHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract brand ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/brands/")
	brandIDStr := strings.Trim(path, "/")

	brandID, err := strconv.ParseUint(brandIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid brand ID format",
			"received": brandIDStr,
		})
		return
	}

	err = brandRepo.Delete(r.Context(), uint(brandID))
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Brand not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete brand"})
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Brand deleted successfully"})
}

// GetWideBrandsHandler handles GET /wide-brands
func GetWideBrandsHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	limit := 0 // 0 means no limit

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	brands, err := brandRepo.GetWideBrands(r.Context(), limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get wide brands"})
		return
	}

	responses := make([]BrandResponse, len(brands))
	for i, brand := range brands {
		responses[i] = convertBrandToResponse(brand)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"brands": responses,
		"count":  len(responses),
	})
}

// GetPublicBrandsHandler handles GET /public-brands
func GetPublicBrandsHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	// Set defaults
	limit := 50
	offset := 0

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

	brands, err := brandRepo.GetPublicBrands(r.Context(), limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get public brands"})
		return
	}

	responses := make([]BrandResponse, len(brands))
	for i, brand := range brands {
		responses[i] = convertBrandToResponse(brand)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"brands": responses,
		"count":  len(responses),
		"limit":  limit,
		"offset": offset,
	})
}

// CreateBrandTranslationHandler handles POST /brand-translation
func CreateBrandTranslationHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		BrandID        uint   `json:"brand_id"`
		Locale         string `json:"locale"`
		TranslatedName string `json:"translated_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.BrandID == 0 || request.Locale == "" || request.TranslatedName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "brand_id, locale, and translated_name are required"})
		return
	}

	translation := &entities.BrandTranslation{
		BrandID:        request.BrandID,
		Locale:         request.Locale,
		TranslatedName: request.TranslatedName,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	savedTranslation, err := brandRepo.CreateTranslation(r.Context(), translation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create brand translation"})
		return
	}

	response := convertBrandTranslationToResponse(savedTranslation)
	json.NewEncoder(w).Encode(response)
}

// GetBrandTranslationHandler handles GET /brand-translation/{id}
func GetBrandTranslationHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract brand ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/brand-translation/")
	brandIDStr := strings.Trim(path, "/")

	brandID, err := strconv.ParseUint(brandIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid brand ID format",
			"received": brandIDStr,
		})
		return
	}

	// Check for locale query parameter
	locale := r.URL.Query().Get("locale")

	if locale != "" {
		// Get specific translation by locale
		translation, err := brandRepo.GetTranslationByLocale(r.Context(), uint(brandID), locale)
		if err != nil {
			if err.Error() == "translation not found" {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "Translation not found"})
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get translation"})
			}
			return
		}

		response := convertBrandTranslationToResponse(translation)
		json.NewEncoder(w).Encode(response)
	} else {
		// Get all translations for the brand
		translations, err := brandRepo.GetTranslations(r.Context(), uint(brandID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get translations"})
			return
		}

		responses := make([]BrandTranslationResponse, len(translations))
		for i, translation := range translations {
			responses[i] = convertBrandTranslationToResponse(translation)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"brand_id":     brandID,
			"translations": responses,
			"count":        len(responses),
		})
	}
}

// UpdateBrandStatusHandler handles POST /brand-status/{id}
func UpdateBrandStatusHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract brand ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/brand-status/")
	brandIDStr := strings.Trim(path, "/")

	brandID, err := strconv.ParseUint(brandIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid brand ID format",
			"received": brandIDStr,
		})
		return
	}

	var request struct {
		Status int `json:"status"`
	}

	// Decode the JSON payload
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid JSON payload", "code": 400})
		return
	}

	err = brandRepo.UpdateStatus(r.Context(), uint(brandID), request.Status)
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Brand not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update brand status"})
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Brand status updated successfully",
		"brand_id": brandID,
		"status":   request.Status,
	})
}
