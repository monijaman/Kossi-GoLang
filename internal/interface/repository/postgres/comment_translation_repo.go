package postgres

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// PostgresCommentTranslationRepo implements the CommentTranslationRepository interface
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
