// Package repository defines interfaces for data access.
// This package is part of the domain layer in Clean Architecture.
// It contains repository interfaces that define contracts for data access without implementation details.
package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

// ProductReviewRepository defines the interface for product review operations
type ProductReviewRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, review *entities.ProductReview) error
	GetByID(ctx context.Context, id uint) (*entities.ProductReview, error)
	Update(ctx context.Context, review *entities.ProductReview) error
	Delete(ctx context.Context, id uint) error

	// Product-specific operations
	GetByProductID(ctx context.Context, productID uint) ([]*entities.ProductReview, error)
	GetByUserID(ctx context.Context, userID uint) ([]*entities.ProductReview, error)
	GetByProductAndUser(ctx context.Context, productID, userID uint) (*entities.ProductReview, error)

	// Search and filtering operations
	GetAll(ctx context.Context, page, limit int, sortOrder string) ([]*entities.ProductReview, int, error)
	SearchReviews(ctx context.Context, searchTerm string, page, limit int, sortOrder string, categoryID string) ([]*entities.ProductReview, int, error)
	GetPublicReviewsByProduct(ctx context.Context, productID uint, locale string) (*entities.ProductReview, error)

	// Rating operations
	GetAverageRating(ctx context.Context, productID uint) (float64, error)
	GetRatingDistribution(ctx context.Context, productID uint) (map[int]int, error)

	// Status operations
	UpdateStatus(ctx context.Context, id uint, status bool) error
	GetByStatus(ctx context.Context, status bool, page, limit int) ([]*entities.ProductReview, int, error)

	// Translation operations
	CreateTranslation(ctx context.Context, translation *entities.ProductReviewTranslation) error
	UpdateTranslation(ctx context.Context, translation *entities.ProductReviewTranslation) error
	GetTranslation(ctx context.Context, reviewID uint, locale string) (*entities.ProductReviewTranslation, error)
	GetTranslationsByReview(ctx context.Context, reviewID uint) ([]*entities.ProductReviewTranslation, error)
	GetTranslationsByReviewIDsAndLocale(ctx context.Context, reviewIDs []uint, locale string) (map[uint]*entities.ProductReviewTranslation, error)

	// Image operations
	GetReviewImages(ctx context.Context, reviewID uint) ([]*entities.Image, error)
	AttachImage(ctx context.Context, reviewID, imageID uint) error
	DetachImage(ctx context.Context, reviewID, imageID uint) error
	SetDefaultImage(ctx context.Context, imageID uint) error
}
