package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

type ProductRepository interface {
	// Basic CRUD operations
	GetByID(ctx context.Context, id uint) (*entities.Product, error)
	GetBySlug(ctx context.Context, slug string) (*entities.Product, error)
	Create(ctx context.Context, product *entities.Product) (*entities.Product, error)
	Update(ctx context.Context, id uint, product *entities.Product) (*entities.Product, error)
	List(ctx context.Context, limit, offset int) ([]*entities.Product, error)

	// Search and filtering
	Search(ctx context.Context, query string, limit, offset int) ([]*entities.Product, error)
	GetPopular(ctx context.Context, limit int) ([]*entities.Product, error)
	GetByCategory(ctx context.Context, categoryID uint, limit, offset int) ([]*entities.Product, error)
	GetByBrand(ctx context.Context, brandID uint, limit, offset int) ([]*entities.Product, error)

	// View tracking
	IncrementViews(ctx context.Context, id uint) error

	// Count operations
	Count(ctx context.Context) (int64, error)
	CountByCategory(ctx context.Context, categoryID uint) (int64, error)
	CountByBrand(ctx context.Context, brandID uint) (int64, error)

	// Translation operations
	CreateTranslation(ctx context.Context, translation *entities.ProductTranslation) (*entities.ProductTranslation, error)
	GetTranslations(ctx context.Context, productID uint) ([]*entities.ProductTranslation, error)
	GetTranslationByLocale(ctx context.Context, productID uint, locale string) (*entities.ProductTranslation, error)
}
