package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

// ImageRepository defines the interface for image-related database operations
type ImageRepository interface {
	// Create creates a new image
	Create(ctx context.Context, image *entities.Image) (*entities.Image, error)

	// GetByImageableID gets images for a specific imageable entity (e.g., product)
	GetByImageableID(ctx context.Context, imageableType string, imageableID uint) ([]*entities.Image, error)

	// GetByID gets an image by ID
	GetByID(ctx context.Context, id uint) (*entities.Image, error)

	// Delete deletes an image by ID
	Delete(ctx context.Context, id uint) error

	// Update updates an existing image
	Update(ctx context.Context, image *entities.Image) (*entities.Image, error)
}
