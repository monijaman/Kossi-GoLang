package category

import (
	"encoding/json"
	"fmt"
	"io"
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
	Status    int    `json:"status"`
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
		Status:    category.Status,
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

// CreateCategoryHandler handles POST /categories
func CreateCategoryHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Read the body to debug
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to read request body"})
		return
	}

	// Log what we received for debugging
	fmt.Printf("Received body: %s\n", string(body))
	fmt.Printf("Content-Type: %s\n", r.Header.Get("Content-Type"))

	var request struct {
		Name   string `json:"name"`
		Status *int   `json:"status,omitempty"`
	}

	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid JSON payload", "code": 400})
		return
	}

	if request.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "name is required"})
		return
	}

	// Create a new category
	baseSlug := strings.ToLower(strings.ReplaceAll(request.Name, " ", "-"))

	// Set status: use provided value or default to 1 (active)
	status := 1
	if request.Status != nil {
		status = *request.Status
	}

	// Generate a unique slug by first trying the base slug, then adding suffix if needed
	slug := baseSlug

	// Try to create with the base slug first
	category := &entities.Category{
		Name:      request.Name,
		Slug:      slug,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	savedCategory, err := categoryRepo.Create(r.Context(), category)
	if err != nil {
		// If it's a duplicate slug error, try with a timestamp suffix
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") && strings.Contains(err.Error(), "slug") {
			// Generate unique slug with timestamp
			timestamp := time.Now().Unix()
			slug = fmt.Sprintf("%s-%d", baseSlug, timestamp)

			category.Slug = slug
			savedCategory, err = categoryRepo.Create(r.Context(), category)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create category"})
				fmt.Println("Error creating category with unique slug:", err)
				return
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create category"})
			fmt.Println("Error creating category:", err)
			return
		}
	}

	response := convertCategoryToResponse(savedCategory)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Category created successfully",
	})

}

// GetCategoriesHandler handles GET /categories
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

// GetCategoryByIDHandler handles GET /categories/{id}
func GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/categories/")
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
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Category retrieved successfully",
	})
}

// UpdateCategoryHandler handles PUT /categories/{id}
func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/categories/")
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
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid JSON payload", "code": 400})
		return
	}

	if request.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "name is required"})
		return
	}

	// Get the existing category to preserve the slug
	existingCategory, err := categoryRepo.GetByID(r.Context(), uint(categoryID))
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

	// Update category with new name but keep existing slug and status
	category := &entities.Category{
		Name:      request.Name,
		Slug:      existingCategory.Slug,   // Keep the existing slug
		Status:    existingCategory.Status, // Keep the existing status
		UpdatedAt: time.Now(),
	}

	// fmt.Println("============================", categoryID, "to new name:", request.Name)

	savedCategory, err := categoryRepo.Update(r.Context(), uint(categoryID), category)
	if err != nil {
		if err.Error() == "category not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error updating category:", err)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update category"})
		}
		return
	}

	response := convertCategoryToResponse(savedCategory)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Category updated successfully",
	})
}

// DeleteCategoryHandler handles DELETE /categories/{id}
func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/categories/")
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

// GetWideCategoriesHandler handles GET /wide-categories
// This endpoint supports comprehensive category retrieval with multiple query parameters:
//
// Query Parameters:
//   - per_page: Number of items per page for pagination (default: 10)
//   - search: Search term to filter categories by name/content
//   - paginate: "true"/"false" to enable/disable pagination response format
//   - locale: Language locale for internationalization (default: "en")
//   - category_id: Filter results by specific category ID
//   - page: Current page number for pagination (default: 1)
//   - status: Category status filter (0 = inactive, 1 = active)
//
// Response Formats:
//   - When paginate=true: {"data": {"categories": [...], "total": number}, "success": true}
//   - When paginate=false: {"categories": [...], "count": number}
//
// The function intelligently chooses the appropriate repository method:
//   - Uses Search() when search term is provided
//   - Uses GetAll() with pagination when paginate=true
//   - Falls back to GetWideCategories() for simple listing
func GetWideCategoriesHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	perPageStr := r.URL.Query().Get("per_page")
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

	var categories []*entities.Category
	var err error

	// Use search if provided, otherwise get all categories
	if search != "" {
		categories, err = categoryRepo.Search(r.Context(), search, perPage, offset)
	} else {
		// For GetAll, we need to handle pagination ourselves if paginate is true
		if paginate {
			categories, err = categoryRepo.GetAll(r.Context(), perPage, offset)
		} else {
			categories, err = categoryRepo.GetWideCategories(r.Context(), 0) // 0 means no limit
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get wide categories"})
		return
	}

	// Apply additional filters if needed
	filteredCategories := categories

	// Filter by category ID if provided
	if categoryIDStr != "" {
		if categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32); err == nil && categoryID > 0 {
			var filtered []*entities.Category
			for _, category := range filteredCategories {
				if category.ID == uint(categoryID) {
					filtered = append(filtered, category)
				}
			}
			filteredCategories = filtered
		}
	}

	// Filter by status if provided
	if statusStr != "" {
		if status, err := strconv.Atoi(statusStr); err == nil {
			var filtered []*entities.Category
			for _, category := range filteredCategories {
				if category.Status == status {
					filtered = append(filtered, category)
				}
			}
			filteredCategories = filtered
		}
	}

	responses := make([]CategoryResponse, len(filteredCategories))
	for i, category := range filteredCategories {
		responses[i] = convertCategoryToResponse(category)
	}

	var total int
	if paginate {
		// Get total count without limit/offset for pagination
		if search != "" {
			// Count all matching search results
			allCategories, err := categoryRepo.Search(r.Context(), search, 0, 0)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get wide categories"})
				return
			}
			total = len(allCategories)
		} else {
			// Count all categories (no limit/offset)
			allCategories, err := categoryRepo.GetAll(r.Context(), 0, 0)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get wide categories"})
				return
			}
			total = len(allCategories)
		}
	} else {
		total = len(responses)
	}

	if paginate {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"categories": responses,
			"total":      total,
			"success":    true,
		})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"categories": responses,
			"count":      len(responses),
		})
	}

}

// CreateCategoryTranslationHandler handles POST /category-translation
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

func UpdateCategoryTranslationHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract translation ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/category-translation/")
	translationIDStr := strings.Trim(path, "/")
	translationID, err := strconv.ParseUint(translationIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid translation ID format",
			"received": translationIDStr,
		})
		return
	}

	fmt.Println("Updating translation ID:", translationID)

	var request struct {
		Locale         string `json:"locale"`
		TranslatedName string `json:"translated_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.Locale == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "locale is required"})
		return
	}

	if request.TranslatedName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "translated_name is required"})
		return
	}

	// First, get the existing translation to preserve CategoryID and other fields
	existingTranslation, err := categoryRepo.GetTranslationByID(r.Context(), uint(translationID))
	if err != nil {
		if err.Error() == "translation not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Translation not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get existing translation"})
		}
		return
	}

	// Create translation object with the correct ID and preserved CategoryID
	translation := &entities.CategoryTranslation{
		ID:             uint(translationID),            // The translation record ID
		CategoryID:     existingTranslation.CategoryID, // Preserve the original category ID
		Locale:         request.Locale,
		TranslatedName: request.TranslatedName,
		CreatedAt:      existingTranslation.CreatedAt, // Preserve original creation time
		UpdatedAt:      time.Now(),
	}

	savedTranslation, err := categoryRepo.UpdateTranslation(r.Context(), translation)
	if err != nil {
		if err.Error() == "translation not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Translation not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update category translation"})
		}
		return
	}

	response := convertCategoryTranslationToResponse(savedTranslation)
	json.NewEncoder(w).Encode(response)
} // GetCategoryTranslationHandler handles GET /category-translation/{id}
func GetCategoryTranslationHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/category-translation/")
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
			"success":     true,
			"data":        responses,
			"message":     "Category translations retrieved successfully",
			"category_id": categoryID,
			"count":       len(responses),
		})
	}
}

// CreateCategoryBrandRelationHandler handles POST /category-brands
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

// GetCategoryBrandRelationsHandler handles GET /category-brands
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

// UpdateCategoryStatusHandler handles POST /category-status/{id}
func UpdateCategoryStatusHandler(w http.ResponseWriter, r *http.Request, categoryRepo repository.CategoryRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/category-status/")
	categoryIDStr := strings.Trim(path, "/")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid category ID format",
			"received": categoryIDStr,
		})
		return
	}

	// Define the request struct with an integer status field
	var request struct {
		Status int `json:"status"`
	}

	// Decode the JSON payload
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid JSON payload", "code": 400})
		return
	}

	// Update the status in the database
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

	// Respond with success
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "Category status updated successfully",
		"category_id": categoryID,
		"status":      request.Status,
	})
}
