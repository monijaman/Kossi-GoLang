package productreview

import (
	"encoding/json"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/usecase/productreview"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Request/Response structures
type CreateReviewRequest struct {
	Rating            int    `json:"rating"`
	Reviews           string `json:"reviews"`
	AdditionalDetails string `json:"additional_details,omitempty"`
	Priority          int    `json:"priority,omitempty"`
}

type UpdateReviewRequest struct {
	Rating            int    `json:"rating"`
	Reviews           string `json:"reviews"`
	AdditionalDetails string `json:"additional_details,omitempty"`
	Priority          int    `json:"priority,omitempty"`
}

type ReviewTranslationRequest struct {
	ProductID         uint   `json:"product_id"`
	Locale            string `json:"locale"`
	Rating            int    `json:"rating"`
	Review            string `json:"review"`
	AdditionalDetails string `json:"additional_details,omitempty"`
}

type UploadImageRequest struct {
	Files     []string `json:"files"`
	ProductID uint     `json:"product_id"`
}

type ReviewResponse struct {
	ID                uint   `json:"id"`
	ProductID         uint   `json:"product_id"`
	UserID            uint   `json:"user_id"`
	Rating            int    `json:"rating"`
	Reviews           string `json:"reviews"`
	AdditionalDetails string `json:"additional_details"`
	Priority          int    `json:"priority"`
	Status            bool   `json:"status"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type TranslationResponse struct {
	ID               uint   `json:"id"`
	ProductReviewID  uint   `json:"product_review_id"`
	Locale           string `json:"locale"`
	TranslatedReview string `json:"translated_review"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type ImageResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	ProductID    uint   `json:"product_id"`
	DefaultPhoto bool   `json:"defaultphoto"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// convertReviewToResponse converts entity to response format
func convertReviewToResponse(review *entities.ProductReview) ReviewResponse {
	reviewText := ""
	if review.Review != nil {
		reviewText = *review.Review
	}

	return ReviewResponse{
		ID:        review.ID,
		ProductID: review.ProductID,
		UserID:    review.UserID,
		Rating:    review.Rating,
		Reviews:   reviewText,
		CreatedAt: review.CreatedAt.Format(time.RFC3339),
		UpdatedAt: review.UpdatedAt.Format(time.RFC3339),
	}
}

// convertTranslationToResponse converts translation entity to response format
func convertTranslationToResponse(translation *entities.ProductReviewTranslation) TranslationResponse {
	return TranslationResponse{
		ID:               translation.ID,
		ProductReviewID:  translation.ProductReviewID,
		Locale:           translation.Locale,
		TranslatedReview: translation.TranslatedReview,
		CreatedAt:        translation.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        translation.UpdatedAt.Format(time.RFC3339),
	}
}

// CreateReviewHandler handles POST /reviews/{id}
func CreateReviewHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/reviews/")
	productIDStr := strings.Trim(path, "/")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	var req CreateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Rating < 1 || req.Rating > 5 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Rating must be between 1 and 5"})
		return
	}

	if req.Reviews == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Review content is required"})
		return
	}

	// TODO: Extract user ID from JWT token
	userID := uint(1) // Placeholder

	review, err := productreview.CreateReview(r.Context(), reviewRepo, userID, uint(productID), req.Rating, req.Reviews, req.AdditionalDetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Review added successfully",
		"review":  convertReviewToResponse(review),
	})
}

// GetAllReviewsHandler handles GET /reviews
func GetAllReviewsHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}

	sortOrder := r.URL.Query().Get("sortOrder")
	if sortOrder == "" {
		sortOrder = "desc"
	}

	searchTerm := r.URL.Query().Get("searchterm")

	var reviews []*entities.ProductReview
	var total int
	var err error

	if searchTerm != "" {
		reviews, total, err = productreview.SearchReviews(r.Context(), reviewRepo, searchTerm, page, limit, sortOrder)
	} else {
		reviews, total, err = productreview.GetAllReviews(r.Context(), reviewRepo, page, limit, sortOrder)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	reviewResponses := make([]ReviewResponse, len(reviews))
	for i, review := range reviews {
		reviewResponses[i] = convertReviewToResponse(review)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"reviews":      reviewResponses,
		"totalReviews": total,
		"currentPage":  page,
		"perPage":      limit,
	})
}

// CreateReviewTranslationHandler handles POST /review/translation
func CreateReviewTranslationHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req ReviewTranslationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Locale == "" || req.Review == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Locale and review are required"})
		return
	}

	// TODO: Get actual review ID from product and user
	reviewID := uint(1) // Placeholder - would need to find existing review

	// Check if translation exists and update or create
	existingTranslation, err := reviewRepo.GetTranslation(r.Context(), reviewID, req.Locale)
	if err == nil && existingTranslation != nil {
		// Update existing translation
		translation, err := productreview.UpdateTranslation(r.Context(), reviewRepo, reviewID, req.Locale, req.Review)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":     "Review updated successfully",
			"translation": convertTranslationToResponse(translation),
		})
	} else {
		// Create new translation
		translation, err := productreview.CreateTranslation(r.Context(), reviewRepo, reviewID, req.Locale, req.Review)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":     "Translation created successfully",
			"translation": convertTranslationToResponse(translation),
		})
	}
}

// UpdateReviewHandler handles POST /product/{id}/review/{reviewid}
func UpdateReviewHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract review ID from URL path - pattern: /api/product/{id}/review/{reviewid}
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 6 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid URL format"})
		return
	}

	reviewID, err := strconv.ParseUint(parts[len(parts)-1], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid review ID"})
		return
	}

	var req UpdateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// TODO: Extract user ID from JWT token
	userID := uint(1) // Placeholder

	review, err := productreview.UpdateReview(r.Context(), reviewRepo, uint(reviewID), userID, req.Rating, req.Reviews, req.AdditionalDetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Review updated successfully",
		"review":  convertReviewToResponse(review),
	})
}

// GetProductReviewsHandler handles GET /products/{productId}/reviews
func GetProductReviewsHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path - pattern: /api/products/{productId}/reviews
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 5 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid URL format"})
		return
	}

	productID, err := strconv.ParseUint(parts[len(parts)-2], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	reviews, err := productreview.GetReviewsByProduct(r.Context(), reviewRepo, uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	reviewResponses := make([]ReviewResponse, len(reviews))
	for i, review := range reviews {
		reviewResponses[i] = convertReviewToResponse(review)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"product_id": productID,
		"reviews":    reviewResponses,
	})
}

// GetReviewHandler handles GET /reviews/{id}
func GetReviewHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract review ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/reviews/")
	reviewIDStr := strings.Trim(path, "/")

	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid review ID"})
		return
	}

	review, err := reviewRepo.GetByID(r.Context(), uint(reviewID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Review not found"})
		return
	}

	locale := r.URL.Query().Get("locale")
	response := map[string]interface{}{
		"review": convertReviewToResponse(review),
	}

	// If locale is specified, get translation
	if locale != "" {
		translation, err := reviewRepo.GetTranslation(r.Context(), uint(reviewID), locale)
		if err == nil {
			response["translation"] = convertTranslationToResponse(translation)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// UploadImagesHandler handles POST /productimages
func UploadImagesHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Implement image upload logic
	// For now, return placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Images uploaded successfully",
		"images":  []ImageResponse{},
	})
}

// GetProductImagesHandler handles GET /productimages/{id}
func GetProductImagesHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/productimages/")
	productIDStr := strings.Trim(path, "/")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	// TODO: Implement get product images logic
	// For now, return placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"product_id": productID,
		"images":     []ImageResponse{},
	})
}

// MakeDefaultImageHandler handles POST /default-image/{id}
func MakeDefaultImageHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract image ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/default-image/")
	imageIDStr := strings.Trim(path, "/")

	imageID, err := strconv.ParseUint(imageIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid image ID"})
		return
	}

	// TODO: Implement set default image logic
	err = reviewRepo.SetDefaultImage(r.Context(), uint(imageID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Image set as default successfully",
		"image_id": imageID,
	})
}

// RemoveImageHandler handles POST /imageremove/{id}
func RemoveImageHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract image ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/imageremove/")
	imageIDStr := strings.Trim(path, "/")

	imageID, err := strconv.ParseUint(imageIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid image ID"})
		return
	}

	// TODO: Implement remove image logic
	// For now, return placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Image removed successfully",
		"image_id": imageID,
	})
}

// GetPublicReviewsHandler handles GET /public-reviews/{id}
func GetPublicReviewsHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/public-reviews/")
	productIDStr := strings.Trim(path, "/")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	locale := r.URL.Query().Get("locale")

	review, err := reviewRepo.GetPublicReviewsByProduct(r.Context(), uint(productID), locale)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "No reviews found"})
		return
	}

	response := map[string]interface{}{
		"product_id": productID,
	}

	if review != nil {
		response["reviews"] = convertReviewToResponse(review)
	} else {
		response["reviews"] = map[string]interface{}{
			"reviews": nil,
			"rating":  nil,
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
