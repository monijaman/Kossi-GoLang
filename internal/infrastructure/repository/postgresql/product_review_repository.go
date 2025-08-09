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

// productReviewRepository implements the ProductReviewRepository interface using PostgreSQL
type productReviewRepository struct {
	db *gorm.DB
}

// NewProductReviewRepository creates a new PostgreSQL product review repository
func NewProductReviewRepository(db *gorm.DB) repository.ProductReviewRepository {
	return &productReviewRepository{
		db: db,
	}
}

// Create creates a new product review
func (r *productReviewRepository) Create(ctx context.Context, review *entities.ProductReview) error {
	model := &models.ProductReviewModel{}
	model.FromEntity(review)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}

	review.ID = model.ID
	review.CreatedAt = model.CreatedAt
	review.UpdatedAt = model.UpdatedAt
	return nil
}

// GetByID retrieves a product review by ID
func (r *productReviewRepository) GetByID(ctx context.Context, id uint) (*entities.ProductReview, error) {
	var model models.ProductReviewModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// Update updates an existing product review
func (r *productReviewRepository) Update(ctx context.Context, review *entities.ProductReview) error {
	model := &models.ProductReviewModel{}
	model.FromEntity(review)

	return r.db.WithContext(ctx).Save(model).Error
}

// Delete soft deletes a product review
func (r *productReviewRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ProductReviewModel{}, id).Error
}

// GetByProductID retrieves all reviews for a specific product
func (r *productReviewRepository) GetByProductID(ctx context.Context, productID uint) ([]*entities.ProductReview, error) {
	// Note: This would need a proper relationship mapping in the models
	// For now, we'll return empty slice as a placeholder
	return []*entities.ProductReview{}, nil
}

// GetByUserID retrieves all reviews by a specific user
func (r *productReviewRepository) GetByUserID(ctx context.Context, userID uint) ([]*entities.ProductReview, error) {
	var models []models.ProductReviewModel
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}

	reviews := make([]*entities.ProductReview, len(models))
	for i, model := range models {
		reviews[i] = model.ToEntity()
	}
	return reviews, nil
}

// GetByProductAndUser retrieves a review by a specific user for a specific product
func (r *productReviewRepository) GetByProductAndUser(ctx context.Context, productID, userID uint) (*entities.ProductReview, error) {
	var model models.ProductReviewModel
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&model).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// GetAll retrieves all product reviews with pagination
func (r *productReviewRepository) GetAll(ctx context.Context, page, limit int, sortOrder string) ([]*entities.ProductReview, int, error) {
	var models []models.ProductReviewModel
	var total int64

	// Count total records
	if err := r.db.WithContext(ctx).Table("product_reviews").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and sorting
	offset := (page - 1) * limit
	query := r.db.WithContext(ctx).Offset(offset).Limit(limit)

	if sortOrder == "asc" {
		query = query.Order("updated_at ASC")
	} else {
		query = query.Order("updated_at DESC")
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, 0, err
	}

	reviews := make([]*entities.ProductReview, len(models))
	for i, model := range models {
		reviews[i] = model.ToEntity()
	}

	return reviews, int(total), nil
}

// SearchReviews searches reviews by content with pagination
func (r *productReviewRepository) SearchReviews(ctx context.Context, searchTerm string, page, limit int, sortOrder string) ([]*entities.ProductReview, int, error) {
	var models []models.ProductReviewModel
	var total int64

	searchPattern := "%" + searchTerm + "%"

	// Count total matching records
	if err := r.db.WithContext(ctx).Table("product_reviews").
		Where("reviews ILIKE ?", searchPattern).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and sorting
	offset := (page - 1) * limit
	query := r.db.WithContext(ctx).Where("reviews ILIKE ?", searchPattern).
		Offset(offset).Limit(limit)

	if sortOrder == "asc" {
		query = query.Order("updated_at ASC")
	} else {
		query = query.Order("updated_at DESC")
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, 0, err
	}

	reviews := make([]*entities.ProductReview, len(models))
	for i, model := range models {
		reviews[i] = model.ToEntity()
	}

	return reviews, int(total), nil
}

// GetPublicReviewsByProduct retrieves public reviews for a product with optional locale
func (r *productReviewRepository) GetPublicReviewsByProduct(ctx context.Context, productID uint, locale string) (*entities.ProductReview, error) {
	// This would need proper implementation based on product-review relationships
	// For now, return nil as placeholder
	return nil, nil
}

// GetAverageRating calculates average rating for a product
func (r *productReviewRepository) GetAverageRating(ctx context.Context, productID uint) (float64, error) {
	var avg float64
	// This would need proper product-review relationship
	// For now, return 0 as placeholder
	return avg, nil
}

// GetRatingDistribution gets rating distribution for a product
func (r *productReviewRepository) GetRatingDistribution(ctx context.Context, productID uint) (map[int]int, error) {
	// This would need proper implementation
	// For now, return empty map as placeholder
	return make(map[int]int), nil
}

// UpdateStatus updates the status of a product review
func (r *productReviewRepository) UpdateStatus(ctx context.Context, id uint, status bool) error {
	return r.db.WithContext(ctx).Table("product_reviews").
		Where("id = ?", id).Update("status", status).Error
}

// GetByStatus retrieves reviews by status with pagination
func (r *productReviewRepository) GetByStatus(ctx context.Context, status bool, page, limit int) ([]*entities.ProductReview, int, error) {
	var models []models.ProductReviewModel
	var total int64

	// Count total matching records
	if err := r.db.WithContext(ctx).Table("product_reviews").
		Where("status = ?", status).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * limit
	if err := r.db.WithContext(ctx).Where("status = ?", status).
		Offset(offset).Limit(limit).Order("updated_at DESC").Find(&models).Error; err != nil {
		return nil, 0, err
	}

	reviews := make([]*entities.ProductReview, len(models))
	for i, model := range models {
		reviews[i] = model.ToEntity()
	}

	return reviews, int(total), nil
}

// CreateTranslation creates a new product review translation
func (r *productReviewRepository) CreateTranslation(ctx context.Context, translation *entities.ProductReviewTranslation) error {
	model := &models.ProductReviewTranslationModel{}
	model.FromEntity(translation)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}

	translation.ID = model.ID
	translation.CreatedAt = model.CreatedAt
	translation.UpdatedAt = model.UpdatedAt
	return nil
}

// UpdateTranslation updates an existing product review translation
func (r *productReviewRepository) UpdateTranslation(ctx context.Context, translation *entities.ProductReviewTranslation) error {
	model := &models.ProductReviewTranslationModel{}
	model.FromEntity(translation)

	return r.db.WithContext(ctx).Save(model).Error
}

// GetTranslation retrieves a specific translation by review ID and locale
func (r *productReviewRepository) GetTranslation(ctx context.Context, reviewID uint, locale string) (*entities.ProductReviewTranslation, error) {
	var model models.ProductReviewTranslationModel
	if err := r.db.WithContext(ctx).Where("product_review_id = ? AND locale = ?", reviewID, locale).
		First(&model).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// GetTranslationsByReview retrieves all translations for a specific review
func (r *productReviewRepository) GetTranslationsByReview(ctx context.Context, reviewID uint) ([]*entities.ProductReviewTranslation, error) {
	var models []models.ProductReviewTranslationModel
	if err := r.db.WithContext(ctx).Where("product_review_id = ?", reviewID).Find(&models).Error; err != nil {
		return nil, err
	}

	translations := make([]*entities.ProductReviewTranslation, len(models))
	for i, model := range models {
		translations[i] = model.ToEntity()
	}
	return translations, nil
}

// GetReviewImages retrieves all images for a specific review
func (r *productReviewRepository) GetReviewImages(ctx context.Context, reviewID uint) ([]*entities.Image, error) {
	// This would need proper implementation based on review-image relationships
	// For now, return empty slice as placeholder
	return []*entities.Image{}, nil
}

// AttachImage attaches an image to a review
func (r *productReviewRepository) AttachImage(ctx context.Context, reviewID, imageID uint) error {
	// This would need proper implementation based on review-image relationships
	// For now, return nil as placeholder
	return nil
}

// DetachImage detaches an image from a review
func (r *productReviewRepository) DetachImage(ctx context.Context, reviewID, imageID uint) error {
	// This would need proper implementation based on review-image relationships
	// For now, return nil as placeholder
	return nil
}

// SetDefaultImage sets an image as default for a review
func (r *productReviewRepository) SetDefaultImage(ctx context.Context, imageID uint) error {
	// This would need proper implementation based on image model
	// For now, return nil as placeholder
	return nil
}
