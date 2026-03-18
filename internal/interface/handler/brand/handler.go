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
	Total     int    `json:"total,omitempty"`
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
func convertBrandToResponse(brand *entities.Brand, total ...int) BrandResponse {
	resp := BrandResponse{
		ID:        brand.ID,
		Name:      brand.Name,
		Slug:      brand.Slug,
		Status:    brand.Status,
		CreatedAt: brand.CreatedAt.Format(time.RFC3339),
		UpdatedAt: brand.UpdatedAt.Format(time.RFC3339),
	}
	if len(total) > 0 {
		resp.Total = total[0]
	}
	return resp
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
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid JSON payload",
		})
		return
	}

	if request.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "name is required",
		})
		return
	}

	// Check if a brand with the same name already exists
	slug := strings.ToLower(strings.ReplaceAll(request.Name, " ", "-"))
	existingBrand, err := brandRepo.GetBySlug(r.Context(), slug)
	if err == nil && existingBrand != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Brand with this name already exists",
		})
		return
	}

	brand := &entities.Brand{
		Name:      request.Name,
		Slug:      slug,
		Status:    1, // default to active
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	savedBrand, err := brandRepo.Create(r.Context(), brand)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to create brand",
		})
		return
	}

	response := convertBrandToResponse(savedBrand)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Brand created successfully",
	})
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
	// Cache for 10 minutes - brands with counts don't change frequently
	w.Header().Set("Cache-Control", "public, max-age=600")

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

	// Fetch product counts for all brands (OPTIMIZED: single batch query)
	brandIDs := make([]uint, len(filteredBrands))
	for i, brand := range filteredBrands {
		brandIDs[i] = brand.ID
	}

	countMap := make(map[uint]int)
	if len(brandIDs) > 0 {
		countMap, _ = brandRepo.GetProductCountsByBrandIDs(r.Context(), brandIDs)
	}

	for i, brand := range filteredBrands {
		responses[i] = convertBrandToResponse(brand, countMap[brand.ID])
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  false,
			"error":    "Invalid brand ID format",
			"received": brandIDStr,
		})
		return
	}

	brand, err := brandRepo.GetByID(r.Context(), uint(brandID))
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Brand not found",
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Failed to get brand",
			})
		}
		return
	}

	response := convertBrandToResponse(brand)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Brand retrieved successfully",
	})
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  false,
			"error":    "Invalid brand ID format",
			"received": brandIDStr,
		})
		return
	}

	var request struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid JSON payload",
		})
		return
	}

	if request.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "name is required",
		})
		return
	}

	// Generate slug if not provided

	// Get existing brand to preserve status (and handle not found / errors)
	existingBrand, err := brandRepo.GetByID(r.Context(), uint(brandID))
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Brand not found",
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Failed to get brand",
			})
		}
		return
	}

	brand := &entities.Brand{
		Name:      request.Name,
		Slug:      existingBrand.Slug,
		Status:    existingBrand.Status,
		UpdatedAt: time.Now(),
	}

	savedBrand, err := brandRepo.Update(r.Context(), uint(brandID), brand)
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Brand not found",
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Failed to update brand",
			})
		}
		return
	}

	response := convertBrandToResponse(savedBrand)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Brand updated successfully",
	})
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  false,
			"error":    "Invalid brand ID format",
			"received": brandIDStr,
		})
		return
	}

	err = brandRepo.Delete(r.Context(), uint(brandID))
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Brand not found",
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Failed to delete brand",
			})
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Brand deleted successfully",
	})
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to get wide brands",
		})
		return
	}

	responses := make([]BrandResponse, len(brands))
	for i, brand := range brands {
		responses[i] = convertBrandToResponse(brand)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    responses,
		"message": "Wide brands retrieved successfully",
		"count":   len(responses),
	})
}

// GetPublicBrandsHandler handles GET /public-brands
func GetPublicBrandsHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")
	// Cache for 10 minutes - brands with counts don't change frequently
	w.Header().Set("Cache-Control", "public, max-age=600")

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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to get public brands",
		})
		return
	}

	// Fetch product counts for all brands (OPTIMIZED: single batch query)
	brandIDs := make([]uint, len(brands))
	for i, brand := range brands {
		brandIDs[i] = brand.ID
	}

	countMap := make(map[uint]int)
	if len(brandIDs) > 0 {
		countMap, _ = brandRepo.GetProductCountsByBrandIDs(r.Context(), brandIDs)
	}

	responses := make([]BrandResponse, len(brands))
	for i, brand := range brands {
		responses[i] = convertBrandToResponse(brand, countMap[brand.ID])
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    responses,
		"message": "Public brands retrieved successfully",
		"count":   len(responses),
		"limit":   limit,
		"offset":  offset,
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid JSON payload",
		})
		return
	}

	if request.BrandID == 0 || request.Locale == "" || request.TranslatedName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "brand_id, locale, and translated_name are required",
		})
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to create brand translation",
		})
		return
	}

	response := convertBrandTranslationToResponse(savedTranslation)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Brand translation created successfully",
	})
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  false,
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
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"error":   "Translation not found",
				})
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"error":   "Failed to get translation",
				})
			}
			return
		}

		response := convertBrandTranslationToResponse(translation)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"data":    response,
			"message": "Brand translation retrieved successfully",
		})
	} else {
		// Get all translations for the brand
		translations, err := brandRepo.GetTranslations(r.Context(), uint(brandID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Failed to get translations",
			})
			return
		}

		responses := make([]BrandTranslationResponse, len(translations))
		for i, translation := range translations {
			responses[i] = convertBrandTranslationToResponse(translation)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  true,
			"data":     responses,
			"message":  "Brand translations retrieved successfully",
			"brand_id": brandID,
			"count":    len(responses),
		})
	}
}

// UpdateBrandTranslationHandler handles PUT /brand-translation/{id}
func UpdateBrandTranslationHandler(w http.ResponseWriter, r *http.Request, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract brand translation ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/brand-translation/")
	brandTranslationIDStr := strings.Trim(path, "/")

	brandTranslationID, err := strconv.ParseUint(brandTranslationIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  false,
			"error":    "Invalid brand translation ID format",
			"received": brandTranslationIDStr,
		})
		return
	}

	var request struct {
		Locale         string `json:"locale"`
		TranslatedName string `json:"translated_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid JSON payload",
		})
		return
	}

	if request.Locale == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "locale is required",
		})
		return
	}

	if request.TranslatedName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "translated_name is required",
		})
		return
	}

	// Create brand translation object
	brandTranslation := &entities.BrandTranslation{
		BrandID:        uint(brandTranslationID),
		Locale:         request.Locale,
		TranslatedName: request.TranslatedName,
		UpdatedAt:      time.Now(),
	}

	savedBrandTranslation, err := brandRepo.UpdateTranslation(r.Context(), uint(brandTranslationID), brandTranslation)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Brand translation not found",
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Failed to update brand translation",
			})
		}
		return
	}

	response := convertBrandTranslationToResponse(savedBrandTranslation)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Brand translation updated successfully",
	})
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  false,
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid JSON payload",
			"code":    400,
		})
		return
	}

	err = brandRepo.UpdateStatus(r.Context(), uint(brandID), request.Status)
	if err != nil {
		if err.Error() == "brand not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Brand not found",
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Failed to update brand status",
			})
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Brand status updated successfully",
		"brand_id": brandID,
		"status":   request.Status,
	})
}
