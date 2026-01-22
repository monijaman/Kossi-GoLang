package feedback

import (
	"encoding/json"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/usecase/feedback"
	"net/http"
	"strconv"
	"strings"
)

// Request/Response structures
type CreateFeedbackRequest struct {
	Content string `json:"content"`
}

type UpdateFeedbackRequest struct {
	Content *string `json:"content,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

type CreateTranslationRequest struct {
	FeedbackID        uint   `json:"feedback_id"`
	Locale            string `json:"locale"`
	TranslatedContent string `json:"translated_content"`
}

type FeedbackResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Content   string `json:"content"`
	Status    int    `json:"status"`
	CreatedBy *uint  `json:"created_by"`
	UpdatedBy *uint  `json:"updated_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
		UserID:    f.UserID,
		Content:   f.Content,
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

	_, err = feedback.CreateFeedback(r.Context(), h.repo, userID, uint(productID), req.Content)
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

	feedbacks, err := feedback.GetAllFeedback(r.Context(), h.repo, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var feedbackResponses []FeedbackResponse
	for _, f := range feedbacks {
		feedbackResponses = append(feedbackResponses, convertFeedbackToResponse(f))
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

	f, err := feedback.UpdateFeedback(r.Context(), h.repo, uint(id), req.Content, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	path := strings.TrimPrefix(r.URL.Path, "/feedback/")
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
		feedbackResponses = append(feedbackResponses, convertFeedbackToResponse(f))
	}

	response := FeedbackListResponse{
		Data: feedbackResponses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
