package category

import (
	"encoding/json"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CategoryResponse represents the response format for categories
type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CategoryTranslationResponse represents the response format for category translations
type CategoryTranslationResponse struct {
	ID             uint   `json:"id"`
	CategoryID     uint   `json:"category_id"`
	Locale         string `json:"locale"`
	TranslatedName string `json:"translated_name"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// BrandCategoryResponse represents the response format for brand-category relations
type BrandCategoryResponse struct {
	ID         uint `json:"id"`
	BrandID    uint `json:"brand_id"`
	CategoryID uint `json:"category_id"`
}

// convertCategoryToResponse converts entity to response format
func convertCategoryToResponse(category *entities.Category) CategoryResponse {
	return CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
		UpdatedAt: category.UpdatedAt.Format(time.RFC3339),
	}
}

// convertCategoryTranslationToResponse converts entity to response format
func convertCategoryTranslationToResponse(translation *entities.CategoryTranslation) CategoryTranslationResponse {
	return CategoryTranslationResponse{
		ID:             translation.ID,
		CategoryID:     translation.CategoryID,
		Locale:         translation.Locale,
		TranslatedName: translation.TranslatedName,
		CreatedAt:      translation.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      translation.UpdatedAt.Format(time.RFC3339),
	}
}

// convertBrandCategoryToResponse converts entity to response format
func convertBrandCategoryToResponse(relation *entities.BrandCategory) BrandCategoryResponse {
	return BrandCategoryResponse{
		ID:         relation.ID,
		BrandID:    relation.BrandID,
		CategoryID: relation.CategoryID,
	}
}

// CreateCategoryHandler handles POST /api/categories
func CreateCategoryHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
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

	category := &entities.Category{
		Name:      request.Name,
		Slug:      request.Slug,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	savedCategory, err := categoryRepo.Create(r.Context(), category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create category"})
		return
	}

	response := convertCategoryToResponse(savedCategory)
	json.NewEncoder(w).Encode(response)
}

// GetCategoriesHandler handles GET /api/categories
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
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

	var categories []*entities.Category
	var err error

	if search != "" {
		categories, err = categoryRepo.Search(r.Context(), search, limit, offset)
	} else {
		categories, err = categoryRepo.GetAll(r.Context(), limit, offset)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get categories"})
		return
	}

	responses := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = convertCategoryToResponse(category)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"categories": responses,
		"count":      len(responses),
		"limit":      limit,
		"offset":     offset,
	})
}

// GetCategoryByIDHandler handles GET /api/categories/{id}
func GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	categoryIDStr := strings.Trim(path, "/")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid category ID format",
			"received": categoryIDStr,
		})
		return
	}

	category, err := categoryRepo.GetByID(r.Context(), uint(categoryID))
	if err != nil {
		if err.Error() == "category not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get category"})
		}
		return
	}

	response := convertCategoryToResponse(category)
	json.NewEncoder(w).Encode(response)
}

// UpdateCategoryHandler handles PUT /api/categories/{id}
func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	categoryIDStr := strings.Trim(path, "/")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid category ID format",
			"received": categoryIDStr,
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

	category := &entities.Category{
		Name:      request.Name,
		Slug:      request.Slug,
		UpdatedAt: time.Now(),
	}

	savedCategory, err := categoryRepo.Update(r.Context(), uint(categoryID), category)
	if err != nil {
		if err.Error() == "category not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update category"})
		}
		return
	}

	response := convertCategoryToResponse(savedCategory)
	json.NewEncoder(w).Encode(response)
}

// DeleteCategoryHandler handles DELETE /api/categories/{id}
func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	categoryIDStr := strings.Trim(path, "/")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid category ID format",
			"received": categoryIDStr,
		})
		return
	}

	err = categoryRepo.Delete(r.Context(), uint(categoryID))
	if err != nil {
		if err.Error() == "category not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete category"})
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted successfully"})
}

// GetWideCategoriesHandler handles GET /api/wide-categories
func GetWideCategoriesHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	limit := 0 // 0 means no limit

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	categories, err := categoryRepo.GetWideCategories(r.Context(), limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get wide categories"})
		return
	}

	responses := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = convertCategoryToResponse(category)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"categories": responses,
		"count":      len(responses),
	})
}

// CreateCategoryTranslationHandler handles POST /api/category-translation
func CreateCategoryTranslationHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		CategoryID     uint   `json:"category_id"`
		Locale         string `json:"locale"`
		TranslatedName string `json:"translated_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.CategoryID == 0 || request.Locale == "" || request.TranslatedName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "category_id, locale, and translated_name are required"})
		return
	}

	translation := &entities.CategoryTranslation{
		CategoryID:     request.CategoryID,
		Locale:         request.Locale,
		TranslatedName: request.TranslatedName,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	savedTranslation, err := categoryRepo.CreateTranslation(r.Context(), translation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create category translation"})
		return
	}

	response := convertCategoryTranslationToResponse(savedTranslation)
	json.NewEncoder(w).Encode(response)
}

// GetCategoryTranslationHandler handles GET /api/category-translation/{id}
func GetCategoryTranslationHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/category-translation/")
	categoryIDStr := strings.Trim(path, "/")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid category ID format",
			"received": categoryIDStr,
		})
		return
	}

	// Check for locale query parameter
	locale := r.URL.Query().Get("locale")

	if locale != "" {
		// Get specific translation by locale
		translation, err := categoryRepo.GetTranslationByLocale(r.Context(), uint(categoryID), locale)
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

		response := convertCategoryTranslationToResponse(translation)
		json.NewEncoder(w).Encode(response)
	} else {
		// Get all translations for the category
		translations, err := categoryRepo.GetTranslations(r.Context(), uint(categoryID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get translations"})
			return
		}

		responses := make([]CategoryTranslationResponse, len(translations))
		for i, translation := range translations {
			responses[i] = convertCategoryTranslationToResponse(translation)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"category_id":  categoryID,
			"translations": responses,
			"count":        len(responses),
		})
	}
}

// CreateCategoryBrandRelationHandler handles POST /api/category-brands
func CreateCategoryBrandRelationHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		CategoryID uint `json:"category_id"`
		BrandID    uint `json:"brand_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.CategoryID == 0 || request.BrandID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "category_id and brand_id are required"})
		return
	}

	relation := &entities.BrandCategory{
		CategoryID: request.CategoryID,
		BrandID:    request.BrandID,
	}

	savedRelation, err := categoryRepo.CreateBrandRelation(r.Context(), relation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create brand-category relation"})
		return
	}

	response := convertBrandCategoryToResponse(savedRelation)
	json.NewEncoder(w).Encode(response)
}

// GetCategoryBrandRelationsHandler handles GET /api/category-brands
func GetCategoryBrandRelationsHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Check for category_id query parameter
	categoryIDStr := r.URL.Query().Get("category_id")

	if categoryIDStr != "" {
		// Get relations for specific category
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":    "Invalid category ID format",
				"received": categoryIDStr,
			})
			return
		}

		relations, err := categoryRepo.GetBrandRelations(r.Context(), uint(categoryID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get brand relations"})
			return
		}

		responses := make([]BrandCategoryResponse, len(relations))
		for i, relation := range relations {
			responses[i] = convertBrandCategoryToResponse(relation)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"category_id": categoryID,
			"relations":   responses,
			"count":       len(responses),
		})
	} else {
		// Get all category-brand relations
		relations, err := categoryRepo.GetCategoryBrandRelations(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get category-brand relations"})
			return
		}

		responses := make([]BrandCategoryResponse, len(relations))
		for i, relation := range relations {
			responses[i] = convertBrandCategoryToResponse(relation)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"relations": responses,
			"count":     len(responses),
		})
	}
}

// UpdateCategoryStatusHandler handles POST /api/category-status/{id}
func UpdateCategoryStatusHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/category-status/")
	categoryIDStr := strings.Trim(path, "/")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid category ID format",
			"received": categoryIDStr,
		})
		return
	}

	var request struct {
		Status bool `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	err = categoryRepo.UpdateStatus(r.Context(), uint(categoryID), request.Status)
	if err != nil {
		if err.Error() == "category not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update category status"})
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "Category status updated successfully",
		"category_id": categoryID,
		"status":      request.Status,
	})
}
