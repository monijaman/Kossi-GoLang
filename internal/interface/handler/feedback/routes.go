package feedback

import (
	"kossti/internal/domain/repository"
	"net/http"
)

// RegisterRoutes registers all feedback routes
func RegisterRoutes(mux *http.ServeMux, repo repository.FeedbackRepository) {
	handler := NewFeedbackHandler(repo)

	// Protected routes (TODO: Add JWT middleware)
	mux.HandleFunc("/feedback/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateFeedback(w, r)
		case http.MethodGet:
			// Determine if this is a single feedback or product feedback request
			// For simplicity, we'll handle both in the same handler for now
			handler.GetFeedbackByID(w, r)
		case http.MethodPut:
			handler.UpdateFeedback(w, r)
		case http.MethodDelete:
			handler.DeleteFeedback(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// V1 API routes (for compatibility)
	mux.HandleFunc("/api/v1/feedback/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateFeedback(w, r)
		case http.MethodGet:
			handler.GetFeedbackByID(w, r)
		case http.MethodPut:
			handler.UpdateFeedback(w, r)
		case http.MethodDelete:
			handler.DeleteFeedback(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Public routes
	mux.HandleFunc("/feedback", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetAllFeedback(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/feedback", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetAllFeedback(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
