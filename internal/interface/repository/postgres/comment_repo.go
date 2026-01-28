package postgres

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// PostgresCommentRepo implements the CommentRepository interface
type PostgresCommentRepo struct {
	db *gorm.DB
}

// NewPostgresCommentRepo creates a new instance of PostgresCommentRepo
func NewPostgresCommentRepo(db *gorm.DB) *PostgresCommentRepo {
	return &PostgresCommentRepo{db: db}
}

// Create saves a new comment to the database
func (r *PostgresCommentRepo) Create(comment *models.CommentModel) error {
	return r.db.Create(comment).Error
}

// GetByID retrieves a comment by its ID
func (r *PostgresCommentRepo) GetByID(id uint) (*models.CommentModel, error) {
	var comment models.CommentModel
	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetByProductID retrieves all comments for a specific product
func (r *PostgresCommentRepo) GetByProductID(productID uint) ([]*models.CommentModel, error) {
	var comments []*models.CommentModel
	if err := r.db.Where("product_id = ?", productID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// GetByProductAndUsername retrieves a comment by product ID and username
func (r *PostgresCommentRepo) GetByProductAndUsername(productID uint, username string) (*models.CommentModel, error) {
	var comment models.CommentModel
	if err := r.db.Where("product_id = ? AND username = ?", productID, username).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// Update updates an existing comment
func (r *PostgresCommentRepo) Update(comment *models.CommentModel) error {
	return r.db.Save(comment).Error
}

// Delete deletes a comment by ID
func (r *PostgresCommentRepo) Delete(id uint) error {
	return r.db.Delete(&models.CommentModel{}, id).Error
}
