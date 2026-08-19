// Package postgresql provides PostgreSQL-specific implementations of repository interfaces.
// This package is part of the infrastructure layer in Clean Architecture.
// It contains concrete implementations that interact with PostgreSQL database using GORM.
package postgresql

import (
	"context"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	models "kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// feedbackRepository implements the FeedbackRepository interface using PostgreSQL
type feedbackRepository struct {
	db *gorm.DB
}

// NewFeedbackRepository creates a new PostgreSQL feedback repository
func NewFeedbackRepository(db *gorm.DB) repository.FeedbackRepository {
	return &feedbackRepository{
		db: db,
	}
}

// Create creates a new feedback
func (r *feedbackRepository) Create(ctx context.Context, feedback *entities.Feedback) (*entities.Feedback, error) {
	model := &models.FeedbackModel{}
	model.FromEntity(feedback)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}

// GetByID retrieves a feedback by ID
func (r *feedbackRepository) GetByID(ctx context.Context, id uint) (*entities.Feedback, error) {
	var model models.FeedbackModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// GetAll retrieves all feedback with pagination
func (r *feedbackRepository) GetAll(ctx context.Context, limit, offset int) ([]*entities.Feedback, error) {
	var models []models.FeedbackModel
	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&models).Error; err != nil {
		return nil, err
	}

	feedbacks := make([]*entities.Feedback, len(models))
	for i, model := range models {
		feedbacks[i] = model.ToEntity()
	}
	return feedbacks, nil
}

// Update updates an existing feedback
func (r *feedbackRepository) Update(ctx context.Context, id uint, feedback *entities.Feedback) (*entities.Feedback, error) {
	model := &models.FeedbackModel{}
	model.FromEntity(feedback)
	model.ID = id

	if err := r.db.WithContext(ctx).Model(&models.FeedbackModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"product_id": model.ProductID, "user_id": model.UserID, "rating": model.Rating,
		"source_url": model.SourceURL, "content": model.Content, "content_en": model.ContentEN,
		"content_bn": model.ContentBN, "status": model.Status, "created_by": model.CreatedBy,
		"updated_by": model.UpdatedBy,
	}).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}

// Delete soft deletes a feedback
func (r *feedbackRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.FeedbackModel{}, id).Error
}

// GetByProductID retrieves feedback for a specific product (via polymorphic relationship)
func (r *feedbackRepository) GetByProductID(ctx context.Context, productID uint) ([]*entities.Feedback, error) {
	var models []models.FeedbackModel

	if err := r.db.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&models).Error; err != nil {
		return nil, err
	}

	feedbacks := make([]*entities.Feedback, len(models))
	for i, model := range models {
		feedbacks[i] = model.ToEntity()
	}
	return feedbacks, nil
}

// GetByUserID retrieves all feedback by a specific user
func (r *feedbackRepository) GetByUserID(ctx context.Context, userID uint) ([]*entities.Feedback, error) {
	var models []models.FeedbackModel
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}

	feedbacks := make([]*entities.Feedback, len(models))
	for i, model := range models {
		feedbacks[i] = model.ToEntity()
	}
	return feedbacks, nil
}

// CreateTranslation creates a new feedback translation
func (r *feedbackRepository) CreateTranslation(ctx context.Context, translation *entities.FeedbackTranslation) (*entities.FeedbackTranslation, error) {
	model := &models.FeedbackTranslationModel{}
	model.FromEntity(translation)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}

// GetTranslations retrieves all translations for a feedback
func (r *feedbackRepository) GetTranslations(ctx context.Context, feedbackID uint) ([]*entities.FeedbackTranslation, error) {
	var models []models.FeedbackTranslationModel
	if err := r.db.WithContext(ctx).Where("feedback_id = ?", feedbackID).Find(&models).Error; err != nil {
		return nil, err
	}

	translations := make([]*entities.FeedbackTranslation, len(models))
	for i, model := range models {
		translations[i] = model.ToEntity()
	}
	return translations, nil
}

// UpdateStatus updates the status of a feedback
func (r *feedbackRepository) UpdateStatus(ctx context.Context, id uint, status bool) error {
	return r.db.WithContext(ctx).Table("feedback").
		Where("id = ?", id).Update("status", status).Error
}

// GetByStatus retrieves feedback by status with pagination
func (r *feedbackRepository) GetByStatus(ctx context.Context, status bool, limit, offset int) ([]*entities.Feedback, error) {
	var models []models.FeedbackModel
	if err := r.db.WithContext(ctx).Where("status = ?", status).
		Offset(offset).Limit(limit).Find(&models).Error; err != nil {
		return nil, err
	}

	feedbacks := make([]*entities.Feedback, len(models))
	for i, model := range models {
		feedbacks[i] = model.ToEntity()
	}
	return feedbacks, nil
}

// AttachToProduct creates a polymorphic relationship between feedback and product
func (r *feedbackRepository) AttachToProduct(ctx context.Context, feedbackID, productID uint) error {
	return r.db.WithContext(ctx).Model(&models.FeedbackModel{}).Where("id = ?", feedbackID).Update("product_id", productID).Error
}

// DetachFromProduct removes the polymorphic relationship between feedback and product
func (r *feedbackRepository) DetachFromProduct(ctx context.Context, feedbackID, productID uint) error {
	return r.db.WithContext(ctx).Model(&models.FeedbackModel{}).Where("id = ? AND product_id = ?", feedbackID, productID).Update("product_id", 0).Error
}

// GetProductFeedbacks retrieves all feedback for a specific product
func (r *feedbackRepository) GetProductFeedbacks(ctx context.Context, productID uint) ([]*entities.Feedback, error) {
	return r.GetByProductID(ctx, productID)
}
