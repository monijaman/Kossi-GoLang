package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

// FeedbackRepository defines the interface for feedback operations
type FeedbackRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, feedback *entities.Feedback) (*entities.Feedback, error)
	GetByID(ctx context.Context, id uint) (*entities.Feedback, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entities.Feedback, error)
	Update(ctx context.Context, id uint, feedback *entities.Feedback) (*entities.Feedback, error)
	Delete(ctx context.Context, id uint) error

	// Product-specific operations
	GetByProductID(ctx context.Context, productID uint) ([]*entities.Feedback, error)
	GetByUserID(ctx context.Context, userID uint) ([]*entities.Feedback, error)

	// Translation operations
	CreateTranslation(ctx context.Context, translation *entities.FeedbackTranslation) (*entities.FeedbackTranslation, error)
	GetTranslations(ctx context.Context, feedbackID uint) ([]*entities.FeedbackTranslation, error)

	// Status operations
	UpdateStatus(ctx context.Context, id uint, status bool) error
	GetByStatus(ctx context.Context, status bool, limit, offset int) ([]*entities.Feedback, error)

	// Product-feedback relationship
	AttachToProduct(ctx context.Context, feedbackID, productID uint) error
	DetachFromProduct(ctx context.Context, feedbackID, productID uint) error
	GetProductFeedbacks(ctx context.Context, productID uint) ([]*entities.Feedback, error)
}
