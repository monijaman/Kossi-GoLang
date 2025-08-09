package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

// FormGeneratorRepository defines the interface for form generator operations
type FormGeneratorRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, formGenerator *entities.FormGenerator) (*entities.FormGenerator, error)
	GetByID(ctx context.Context, id uint) (*entities.FormGenerator, error)
	GetByCategoryID(ctx context.Context, categoryID uint) (*entities.FormGenerator, error)
	Update(ctx context.Context, id uint, formGenerator *entities.FormGenerator) (*entities.FormGenerator, error)
	Delete(ctx context.Context, id uint) error

	// Upsert operation (create or update based on category_id)
	UpsertByCategoryID(ctx context.Context, formGenerator *entities.FormGenerator) (*entities.FormGenerator, error)

	// Get category specifications
	GetCategorySpecifications(ctx context.Context, categoryID uint) ([]*entities.Specification, error)
}
