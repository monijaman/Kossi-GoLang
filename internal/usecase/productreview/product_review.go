// Package productreview provides use cases for product review operations.
// This package is part of the use case layer in Clean Architecture.
// It contains business logic for product review management.
package productreview

import (
	"context"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"time"
)

// CreateReview creates a new product review
func CreateReview(ctx context.Context, repo repository.ProductReviewRepository, userID uint, productID uint, rating int, review string, additionalDetails string) (*entities.ProductReview, error) {
	if rating < 1 || rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}

	if review == "" {
		return nil, errors.New("review content is required")
	}

	// Check if user already has a review for this product
	existingReview, err := repo.GetByProductAndUser(ctx, productID, userID)
	if err == nil && existingReview != nil {
		return nil, errors.New("user has already reviewed this product")
	}

	newReview := &entities.ProductReview{
		ProductID: productID,
		UserID:    userID,
		Rating:    rating,
		Review:    &review,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.Create(ctx, newReview); err != nil {
		return nil, err
	}

	return newReview, nil
}

// UpdateReview updates an existing product review
func UpdateReview(ctx context.Context, repo repository.ProductReviewRepository, reviewID uint, userID uint, rating int, review string, additionalDetails string) (*entities.ProductReview, error) {
	if rating < 1 || rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}

	if review == "" {
		return nil, errors.New("review content is required")
	}

	existingReview, err := repo.GetByID(ctx, reviewID)
	if err != nil {
		return nil, errors.New("review not found")
	}

	if existingReview.UserID != userID {
		return nil, errors.New("unauthorized to update this review")
	}

	existingReview.Rating = rating
	existingReview.Review = &review
	existingReview.UpdatedAt = time.Now()

	if err := repo.Update(ctx, existingReview); err != nil {
		return nil, err
	}

	return existingReview, nil
}

// DeleteReview deletes a product review
func DeleteReview(ctx context.Context, repo repository.ProductReviewRepository, reviewID uint, userID uint) error {
	existingReview, err := repo.GetByID(ctx, reviewID)
	if err != nil {
		return errors.New("review not found")
	}

	if existingReview.UserID != userID {
		return errors.New("unauthorized to delete this review")
	}

	return repo.Delete(ctx, reviewID)
}

// GetReviewsByProduct gets all reviews for a specific product
func GetReviewsByProduct(ctx context.Context, repo repository.ProductReviewRepository, productID uint) ([]*entities.ProductReview, error) {
	return repo.GetByProductID(ctx, productID)
}

// GetReviewsByUser gets all reviews by a specific user
func GetReviewsByUser(ctx context.Context, repo repository.ProductReviewRepository, userID uint) ([]*entities.ProductReview, error) {
	return repo.GetByUserID(ctx, userID)
}

// CreateTranslation creates a new translation for a product review
func CreateTranslation(ctx context.Context, repo repository.ProductReviewRepository, reviewID uint, locale string, review string) (*entities.ProductReviewTranslation, error) {
	if locale == "" {
		return nil, errors.New("locale is required")
	}

	if review == "" {
		return nil, errors.New("review content is required")
	}

	// Check if translation already exists
	existingTranslation, err := repo.GetTranslation(ctx, reviewID, locale)
	if err == nil && existingTranslation != nil {
		return nil, errors.New("translation already exists for this locale")
	}

	translation := &entities.ProductReviewTranslation{
		ProductReviewID:  reviewID,
		Locale:           locale,
		TranslatedReview: review,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := repo.CreateTranslation(ctx, translation); err != nil {
		return nil, err
	}

	return translation, nil
}

// UpdateTranslation updates an existing translation
func UpdateTranslation(ctx context.Context, repo repository.ProductReviewRepository, reviewID uint, locale string, review string) (*entities.ProductReviewTranslation, error) {
	if locale == "" {
		return nil, errors.New("locale is required")
	}

	if review == "" {
		return nil, errors.New("review content is required")
	}

	existingTranslation, err := repo.GetTranslation(ctx, reviewID, locale)
	if err != nil {
		return nil, errors.New("translation not found")
	}

	existingTranslation.TranslatedReview = review
	existingTranslation.UpdatedAt = time.Now()

	if err := repo.UpdateTranslation(ctx, existingTranslation); err != nil {
		return nil, err
	}

	return existingTranslation, nil
}

// GetReviewTranslations gets all translations for a specific review
func GetReviewTranslations(ctx context.Context, repo repository.ProductReviewRepository, reviewID uint) ([]*entities.ProductReviewTranslation, error) {
	return repo.GetTranslationsByReview(ctx, reviewID)
}

// UpdateReviewStatus updates the status of a review
func UpdateReviewStatus(ctx context.Context, repo repository.ProductReviewRepository, reviewID uint, status bool) error {
	return repo.UpdateStatus(ctx, reviewID, status)
}

// SearchReviews searches for reviews based on search term
func SearchReviews(ctx context.Context, repo repository.ProductReviewRepository, searchTerm string, page, limit int, sortOrder string) ([]*entities.ProductReview, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	return repo.SearchReviews(ctx, searchTerm, page, limit, sortOrder)
}

// GetAllReviews gets all reviews with pagination
func GetAllReviews(ctx context.Context, repo repository.ProductReviewRepository, page, limit int, sortOrder string) ([]*entities.ProductReview, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	return repo.GetAll(ctx, page, limit, sortOrder)
}
