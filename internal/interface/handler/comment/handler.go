package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"kossti/internal/infrastructure/database/models"
)

// CommentRepository defines the interface for comment operations
type CommentRepository interface {
	Create(comment *models.CommentModel) error
	GetByID(id uint) (*models.CommentModel, error)
	GetByProductID(productID uint) ([]*models.CommentModel, error)
	GetByProductAndUsername(productID uint, username string) (*models.CommentModel, error)
	Update(comment *models.CommentModel) error
	Delete(id uint) error
}

// CreateCommentRequest defines the structure for creating a comment
type CreateCommentRequest struct {
	ProductID uint   `json:"product_id"`
	Username  string `json:"username"`
	Location  string `json:"location"`
	Comment   string `json:"comment"`
	Src       string `json:"src"`
}

// CommentResponse defines the structure for comment response
type CommentResponse struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	Username  string    `json:"username"`
	Location  string    `json:"location"`
	Comment   string    `json:"comment"`
	Src       string    `json:"src"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateCommentHandler handles POST /comments
func CreateCommentHandler(w http.ResponseWriter, r *http.Request, repo CommentRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("ERROR: Failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Log the received request
	log.Printf("DEBUG: Creating comment - ProductID: %d, Username: %s, Comment: %s", req.ProductID, req.Username, req.Comment)

	// Validate required fields
	if req.ProductID == 0 || req.Username == "" || req.Comment == "" {
		log.Printf("ERROR: Validation failed - ProductID: %d, Username: %s, Comment empty: %v", req.ProductID, req.Username, req.Comment == "")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "product_id, username, and comment are required"})
		return
	}

	// Create comment model
	comment := &models.CommentModel{
		ProductID: int(req.ProductID),
		Username:  req.Username,
		Location:  req.Location,
		Comment:   req.Comment,
		Src:       req.Src,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to database
	if err := repo.Create(comment); err != nil {
		log.Printf("ERROR: Failed to create comment in database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to create comment: %v", err)})
		return
	}

	log.Printf("SUCCESS: Comment created with ID: %d for product: %d", comment.ID, req.ProductID)

	// Return success response with comment ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": CommentResponse{
			ID:        comment.ID,
			ProductID: req.ProductID,
			Username:  comment.Username,
			Location:  comment.Location,
			Comment:   comment.Comment,
			Src:       comment.Src,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		},
		"message": "Comment created successfully",
	})
}

// GetCommentsByProductIDHandler handles GET /comments/product/{id}
func GetCommentsByProductIDHandler(w http.ResponseWriter, r *http.Request, repo CommentRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/comments/product/")
	productIDStr := strings.Trim(path, "/")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	comments, err := repo.GetByProductID(uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch comments"})
		return
	}

	// Convert models to response
	response := make([]CommentResponse, 0)
	for _, comment := range comments {
		response = append(response, CommentResponse{
			ID:        comment.ID,
			ProductID: uint(comment.ProductID),
			Username:  comment.Username,
			Location:  comment.Location,
			Comment:   comment.Comment,
			Src:       comment.Src,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
	})
}

// GetCommentByIDHandler handles GET /comments/{id}
func GetCommentByIDHandler(w http.ResponseWriter, r *http.Request, repo CommentRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/comments/")
	idStr := strings.Trim(path, "/")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid comment ID"})
		return
	}

	comment, err := repo.GetByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Comment not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": CommentResponse{
			ID:        comment.ID,
			ProductID: uint(comment.ProductID),
			Username:  comment.Username,
			Location:  comment.Location,
			Comment:   comment.Comment,
			Src:       comment.Src,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		},
	})
}

// UpsertCommentHandler handles POST /comments/upsert - Create or update a comment
func UpsertCommentHandler(w http.ResponseWriter, r *http.Request, repo CommentRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("ERROR: Failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Log the received request
	log.Printf("DEBUG: Upserting comment - ProductID: %d, Username: %s, Comment: %s", req.ProductID, req.Username, req.Comment)

	// Validate required fields
	if req.ProductID == 0 || req.Username == "" || req.Comment == "" {
		log.Printf("ERROR: Validation failed - ProductID: %d, Username: %s, Comment empty: %v", req.ProductID, req.Username, req.Comment == "")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "product_id, username, and comment are required"})
		return
	}

	// Check if comment already exists for this product and username
	existingComment, err := repo.GetByProductAndUsername(req.ProductID, req.Username)
	var comment *models.CommentModel
	isUpdate := false

	if err == nil && existingComment != nil {
		// Comment exists, update it
		isUpdate = true
		existingComment.Comment = req.Comment
		existingComment.Location = req.Location
		existingComment.Src = req.Src
		existingComment.UpdatedAt = time.Now()
		comment = existingComment

		if err := repo.Update(comment); err != nil {
			log.Printf("ERROR: Failed to update comment in database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to update comment: %v", err)})
			return
		}
		log.Printf("SUCCESS: Comment updated with ID: %d for product: %d", comment.ID, req.ProductID)
	} else {
		// Comment doesn't exist, create it
		comment = &models.CommentModel{
			ProductID: int(req.ProductID),
			Username:  req.Username,
			Location:  req.Location,
			Comment:   req.Comment,
			Src:       req.Src,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := repo.Create(comment); err != nil {
			log.Printf("ERROR: Failed to create comment in database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to create comment: %v", err)})
			return
		}
		log.Printf("SUCCESS: Comment created with ID: %d for product: %d", comment.ID, req.ProductID)
	}

	// Return success response
	statusCode := http.StatusCreated
	if isUpdate {
		statusCode = http.StatusOK
	}

	message := "Comment created successfully"
	if isUpdate {
		message = "Comment updated successfully"
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": CommentResponse{
			ID:        comment.ID,
			ProductID: uint(comment.ProductID),
			Username:  comment.Username,
			Location:  comment.Location,
			Comment:   comment.Comment,
			Src:       comment.Src,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		},
		"message": message,
	})
}
