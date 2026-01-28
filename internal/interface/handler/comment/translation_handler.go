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

	"gorm.io/gorm"
)

// CommentTranslationRepository defines the interface for comment translation operations
type CommentTranslationRepository interface {
	Create(translation *models.CommentTranslationModel) error
	GetByCommentID(commentID uint) ([]*models.CommentTranslationModel, error)
	GetByCommentIDAndLocale(commentID uint, locale string) (*models.CommentTranslationModel, error)
	Update(translation *models.CommentTranslationModel) error
	Delete(id uint) error
}

// PostgresCommentTranslationRepo implements CommentTranslationRepository for PostgreSQL
type PostgresCommentTranslationRepo struct {
	db *gorm.DB
}

// NewPostgresCommentTranslationRepo creates a new instance of PostgresCommentTranslationRepo
func NewPostgresCommentTranslationRepo(db *gorm.DB) *PostgresCommentTranslationRepo {
	return &PostgresCommentTranslationRepo{db: db}
}

// Create saves a new comment translation to the database
func (r *PostgresCommentTranslationRepo) Create(translation *models.CommentTranslationModel) error {
	return r.db.Create(translation).Error
}

// GetByCommentID retrieves all translations for a specific comment
func (r *PostgresCommentTranslationRepo) GetByCommentID(commentID uint) ([]*models.CommentTranslationModel, error) {
	var translations []*models.CommentTranslationModel
	if err := r.db.Where("comment_id = ?", commentID).Find(&translations).Error; err != nil {
		return nil, err
	}
	return translations, nil
}

// GetByCommentIDAndLocale retrieves a specific translation for a comment and locale
func (r *PostgresCommentTranslationRepo) GetByCommentIDAndLocale(commentID uint, locale string) (*models.CommentTranslationModel, error) {
	var translation models.CommentTranslationModel
	if err := r.db.Where("comment_id = ? AND locale = ?", commentID, locale).First(&translation).Error; err != nil {
		return nil, err
	}
	return &translation, nil
}

// Update updates an existing comment translation
func (r *PostgresCommentTranslationRepo) Update(translation *models.CommentTranslationModel) error {
	return r.db.Save(translation).Error
}

// Delete deletes a comment translation by ID
func (r *PostgresCommentTranslationRepo) Delete(id uint) error {
	return r.db.Delete(&models.CommentTranslationModel{}, id).Error
}

// CreateCommentTranslationRequest defines the structure for creating a comment translation
type CreateCommentTranslationRequest struct {
	CommentID uint   `json:"comment_id"`
	Locale    string `json:"locale"`
	Comment   string `json:"comment"`
}

// CommentTranslationResponse defines the structure for comment translation response
type CommentTranslationResponse struct {
	ID        uint      `json:"id"`
	CommentID uint      `json:"comment_id"`
	Locale    string    `json:"locale"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateCommentTranslationHandler handles POST /comment-translations
func CreateCommentTranslationHandler(w http.ResponseWriter, r *http.Request, repo CommentTranslationRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req CreateCommentTranslationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("ERROR: Failed to decode comment translation request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	log.Printf("DEBUG: Creating comment translation - CommentID: %d, Locale: %s", req.CommentID, req.Locale)

	// Validate required fields
	if req.CommentID == 0 || req.Locale == "" || req.Comment == "" {
		log.Printf("ERROR: Comment translation validation failed - CommentID: %d, Locale: %s", req.CommentID, req.Locale)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "comment_id, locale, and comment are required"})
		return
	}

	// Create comment translation model
	translation := &models.CommentTranslationModel{
		CommentID: req.CommentID,
		Locale:    req.Locale,
		Comment:   req.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to database
	if err := repo.Create(translation); err != nil {
		log.Printf("ERROR: Failed to create comment translation in database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to create comment translation: %v", err)})
		return
	}

	log.Printf("SUCCESS: Comment translation created with ID: %d for comment: %d", translation.ID, req.CommentID)

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": CommentTranslationResponse{
			ID:        translation.ID,
			CommentID: translation.CommentID,
			Locale:    translation.Locale,
			Comment:   translation.Comment,
			CreatedAt: translation.CreatedAt,
			UpdatedAt: translation.UpdatedAt,
		},
		"message": "Comment translation created successfully",
	})
}

// GetCommentTranslationsHandler handles GET /comment-translations/{id}
func GetCommentTranslationsHandler(w http.ResponseWriter, r *http.Request, repo CommentTranslationRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract comment ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/comment-translations/")
	idStr := strings.Trim(path, "/")

	commentID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid comment ID"})
		return
	}

	translations, err := repo.GetByCommentID(uint(commentID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch translations"})
		return
	}

	// Convert models to response
	response := make([]CommentTranslationResponse, 0)
	for _, translation := range translations {
		response = append(response, CommentTranslationResponse{
			ID:        translation.ID,
			CommentID: translation.CommentID,
			Locale:    translation.Locale,
			Comment:   translation.Comment,
			CreatedAt: translation.CreatedAt,
			UpdatedAt: translation.UpdatedAt,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
	})
}

// UpsertCommentTranslationHandler handles POST/PUT /comment-translations/upsert - Create or update a comment translation
func UpsertCommentTranslationHandler(w http.ResponseWriter, r *http.Request, repo CommentTranslationRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req CreateCommentTranslationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("ERROR: Failed to decode comment translation request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	log.Printf("DEBUG: Upserting comment translation - CommentID: %d, Locale: %s", req.CommentID, req.Locale)

	// Validate required fields
	if req.CommentID == 0 || req.Locale == "" || req.Comment == "" {
		log.Printf("ERROR: Comment translation validation failed - CommentID: %d, Locale: %s", req.CommentID, req.Locale)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "comment_id, locale, and comment are required"})
		return
	}

	// Check if translation already exists
	existingTranslation, err := repo.GetByCommentIDAndLocale(req.CommentID, req.Locale)
	var translation *models.CommentTranslationModel
	isUpdate := false

	if err == nil && existingTranslation != nil {
		// Translation exists, update it
		isUpdate = true
		existingTranslation.Comment = req.Comment
		existingTranslation.UpdatedAt = time.Now()
		translation = existingTranslation

		if err := repo.Update(translation); err != nil {
			log.Printf("ERROR: Failed to update comment translation in database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to update comment translation: %v", err)})
			return
		}
		log.Printf("SUCCESS: Comment translation updated with ID: %d for comment: %d", translation.ID, req.CommentID)
	} else {
		// Translation doesn't exist, create it
		translation = &models.CommentTranslationModel{
			CommentID: req.CommentID,
			Locale:    req.Locale,
			Comment:   req.Comment,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := repo.Create(translation); err != nil {
			log.Printf("ERROR: Failed to create comment translation in database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to create comment translation: %v", err)})
			return
		}
		log.Printf("SUCCESS: Comment translation created with ID: %d for comment: %d", translation.ID, req.CommentID)
	}

	// Return success response
	statusCode := http.StatusCreated
	if isUpdate {
		statusCode = http.StatusOK
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": CommentTranslationResponse{
			ID:        translation.ID,
			CommentID: translation.CommentID,
			Locale:    translation.Locale,
			Comment:   translation.Comment,
			CreatedAt: translation.CreatedAt,
			UpdatedAt: translation.UpdatedAt,
		},
		"message": map[string]string{
			"true":  "Comment translation updated successfully",
			"false": "Comment translation created successfully",
		}[fmt.Sprintf("%v", isUpdate)],
	})
}
