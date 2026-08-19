package feedback

import (
	"encoding/json"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/interface/middleware"
	"kossti/internal/usecase/feedback"
	"net/http"
	"strconv"
	"strings"
)

// Request/Response structures
type CreateFeedbackRequest struct {
	Content   string  `json:"content"`
	Locale    string  `json:"locale,omitempty"`
	ContentEN string  `json:"content_en,omitempty"`
	ContentBN string  `json:"content_bn,omitempty"`
	Rating    string  `json:"rating,omitempty"`
	SourceURL *string `json:"source_url,omitempty"`
}

func (h *FeedbackHandler) CreateFeedbackTranslation(w http.ResponseWriter, r *http.Request) {
	var req CreateTranslationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.FeedbackID == 0 || req.Locale == "" || req.TranslatedContent == "" {
		http.Error(w, "invalid translation", http.StatusBadRequest)
		return
	}
	item, err := h.repo.GetByID(r.Context(), req.FeedbackID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if req.Locale == "bn" {
		item.ContentBN = req.TranslatedContent
	} else if req.Locale == "en" {
		item.ContentEN = req.TranslatedContent
		item.Content = req.TranslatedContent
	} else {
		http.Error(w, "locale must be en or bn", http.StatusBadRequest)
		return
	}
	updated, err := h.repo.Update(r.Context(), req.FeedbackID, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(convertFeedbackToResponse(updated))
}

// CreateProductFeedback handles POST /product-feedback/{product_id}.
func (h *FeedbackHandler) CreateProductFeedback(w http.ResponseWriter, r *http.Request) {
	productIDStr := strings.Trim(strings.TrimPrefix(r.URL.Path, "/product-feedback/"), "/")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	var req CreateFeedbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	content := strings.TrimSpace(strings.Join(strings.Fields(req.Content), " "))
	if len([]rune(content)) < 3 || len([]rune(content)) > 2000 {
		http.Error(w, "Comment must be between 3 and 2000 characters", http.StatusBadRequest)
		return
	}
	rating, ratingErr := strconv.Atoi(strings.TrimSpace(req.Rating))
	if ratingErr != nil || rating < 1 || rating > 5 {
		http.Error(w, "Rating must be between 1 and 5", http.StatusBadRequest)
		return
	}
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil || userID == 0 {
		http.Error(w, "You must be logged in to comment", http.StatusUnauthorized)
		return
	}
	created, err := feedback.CreateFeedbackWithDetails(r.Context(), h.repo, userID, uint(productID), content, req.Rating, req.SourceURL)
	if err != nil {
		if errors.Is(err, feedback.ErrDailyFeedbackLimitReached) {
			http.Error(w, err.Error(), http.StatusTooManyRequests)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if req.Locale == "bn" {
		created.ContentEN = ""
		created.ContentBN = content
		if _, err = h.repo.Update(r.Context(), created.ID, created); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		created.ContentEN = content
		if _, err = h.repo.Update(r.Context(), created.ID, created); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Comment submitted successfully", "feedback": convertFeedbackToResponse(created)})
}

type UpdateFeedbackRequest struct {
	Content   *string `json:"content,omitempty"`
	ContentEN *string `json:"content_en,omitempty"`
	ContentBN *string `json:"content_bn,omitempty"`
	Status    *int    `json:"status,omitempty"`
	Rating    *string `json:"rating,omitempty"`
	SourceURL *string `json:"source_url,omitempty"`
}

type CreateTranslationRequest struct {
	FeedbackID        uint   `json:"feedback_id"`
	Locale            string `json:"locale"`
	TranslatedContent string `json:"translated_content"`
}

type FeedbackResponse struct {
	ID           uint              `json:"id"`
	ProductID    uint              `json:"product_id"`
	Rating       string            `json:"rating"`
	SourceURL    *string           `json:"source_url,omitempty"`
	UserID       uint              `json:"user_id"`
	Content      string            `json:"content"`
	ContentEN    string            `json:"content_en,omitempty"`
	ContentBN    string            `json:"content_bn,omitempty"`
	Status       int               `json:"status"`
	CreatedBy    *uint             `json:"created_by"`
	UpdatedBy    *uint             `json:"updated_by"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
	Translation  string            `json:"translation,omitempty"`
	Translations map[string]string `json:"translations,omitempty"`
}

type TranslationResponse struct {
	ID                uint   `json:"id"`
	FeedbackID        uint   `json:"feedback_id"`
	Locale            string `json:"locale"`
	TranslatedContent string `json:"translated_content"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type FeedbackListResponse struct {
	Data    []FeedbackResponse `json:"data"`
	Message string             `json:"message,omitempty"`
}

// convertFeedbackToResponse converts entity to response format
func convertFeedbackToResponse(f *entities.Feedback) FeedbackResponse {
	return FeedbackResponse{
		ID:        f.ID,
		ProductID: f.ProductID,
		Rating:    f.Rating,
		SourceURL: f.SourceURL,
		UserID:    f.UserID,
		Content:   f.Content,
		ContentEN: f.ContentEN,
		ContentBN: f.ContentBN,
		Status:    f.Status,
		CreatedBy: f.CreatedBy,
		UpdatedBy: f.UpdatedBy,
		CreatedAt: f.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: f.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// convertTranslationToResponse converts translation entity to response format
func convertTranslationToResponse(t *entities.FeedbackTranslation) TranslationResponse {
	return TranslationResponse{
		ID:                t.ID,
		FeedbackID:        t.FeedbackID,
		Locale:            t.Locale,
		TranslatedContent: t.TranslatedContent,
		CreatedAt:         t.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:         t.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// FeedbackHandler handles HTTP requests for feedback operations
type FeedbackHandler struct {
	repo repository.FeedbackRepository
}

// NewFeedbackHandler creates a new feedback handler
func NewFeedbackHandler(repo repository.FeedbackRepository) *FeedbackHandler {
	return &FeedbackHandler{
		repo: repo,
	}
}

// CreateFeedback handles POST /api/v1/feedback/{product_id}
func (h *FeedbackHandler) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/feedback/")
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/v1/feedback/")
	}
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/feedback/")
	}

	productIDStr := strings.Split(path, "/")[0]
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var req CreateFeedbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Content == "" {
		http.Error(w, "content is required", http.StatusBadRequest)
		return
	}

	// TODO: Extract user ID from JWT token
	// For now, using a placeholder user ID
	userID := uint(1)

	_, err = feedback.CreateFeedbackWithDetails(r.Context(), h.repo, userID, uint(productID), req.Content, req.Rating, req.SourceURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Feedback submitted successfully.",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetAllFeedback handles GET /feedback
func (h *FeedbackHandler) GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10
	offset := 0
	const maxLimit = 100

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
			if limit > maxLimit {
				limit = maxLimit
			}
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	feedbacks, err := feedback.GetAllFeedback(r.Context(), h.repo, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var feedbackResponses []FeedbackResponse
	for _, f := range feedbacks {
		item := convertFeedbackToResponse(f)
		locale := r.URL.Query().Get("locale")
		if locale == "bn" {
			item.Content = f.ContentBN
		} else if locale == "en" {
			item.Content = f.ContentEN
		}
		if locale != "" && item.Content == "" {
			continue
		}
		feedbackResponses = append(feedbackResponses, item)
	}

	response := FeedbackListResponse{
		Data: feedbackResponses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetFeedbackByID handles GET /api/v1/feedback/{id}
func (h *FeedbackHandler) GetFeedbackByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/feedback/")
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/v1/feedback/")
	}
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/feedback/")
	}

	idStr := strings.Split(path, "/")[0]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid feedback ID", http.StatusBadRequest)
		return
	}

	f, err := feedback.GetFeedbackByID(r.Context(), h.repo, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := convertFeedbackToResponse(f)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateFeedback handles PUT /api/v1/feedback/{id}
func (h *FeedbackHandler) UpdateFeedback(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/feedback/")
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/v1/feedback/")
	}
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/feedback/")
	}

	idStr := strings.Split(path, "/")[0]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid feedback ID", http.StatusBadRequest)
		return
	}

	var req UpdateFeedbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	f, err := feedback.UpdateFeedback(r.Context(), h.repo, uint(id), req.Content, req.Status, req.Rating, req.SourceURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if req.ContentEN != nil || req.ContentBN != nil {
		if req.ContentEN != nil {
			f.ContentEN = *req.ContentEN
			f.Content = *req.ContentEN
		}
		if req.ContentBN != nil {
			f.ContentBN = *req.ContentBN
		}
		f, err = h.repo.Update(r.Context(), uint(id), f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	response := convertFeedbackToResponse(f)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteFeedback handles DELETE /api/v1/feedback/{id}
func (h *FeedbackHandler) DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/feedback/")
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/v1/feedback/")
	}
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/feedback/")
	}

	idStr := strings.Split(path, "/")[0]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid feedback ID", http.StatusBadRequest)
		return
	}

	err = feedback.DeleteFeedback(r.Context(), h.repo, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetProductFeedback handles GET /feedback/{productId} (for product-specific feedback)
func (h *FeedbackHandler) GetProductFeedback(w http.ResponseWriter, r *http.Request) {
	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/product-feedback/")
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/feedback/")
	}
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/v1/feedback/")
	}

	productIDStr := strings.Split(path, "/")[0]
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	feedbacks, err := feedback.GetProductFeedback(r.Context(), h.repo, uint(productID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var feedbackResponses []FeedbackResponse
	for _, f := range feedbacks {
		item := convertFeedbackToResponse(f)
		locale := r.URL.Query().Get("locale")
		if locale == "bn" {
			item.Content = f.ContentBN
		} else if locale == "en" {
			item.Content = f.ContentEN
		}
		if locale != "" && item.Content == "" {
			continue
		}
		feedbackResponses = append(feedbackResponses, item)
	}

	response := FeedbackListResponse{
		Data: feedbackResponses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
