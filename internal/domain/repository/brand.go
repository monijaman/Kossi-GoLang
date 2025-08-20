package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

// BrandRepository defines the interface for brand-related database operations
type BrandRepository interface {
	// Basic CRUD operations for brands
	Create(ctx context.Context, brand *entities.Brand) (*entities.Brand, error)
	GetByID(ctx context.Context, id uint) (*entities.Brand, error)
	GetBySlug(ctx context.Context, slug string) (*entities.Brand, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entities.Brand, error)
	Update(ctx context.Context, id uint, brand *entities.Brand) (*entities.Brand, error)
	Delete(ctx context.Context, id uint) error
	Search(ctx context.Context, query string, limit, offset int) ([]*entities.Brand, error)
	GetWideBrands(ctx context.Context, limit int) ([]*entities.Brand, error)
	GetPublicBrands(ctx context.Context, limit, offset int) ([]*entities.Brand, error)

	// Translation operations
	CreateTranslation(ctx context.Context, translation *entities.BrandTranslation) (*entities.BrandTranslation, error)
	UpdateTranslation(ctx context.Context, id uint, translation *entities.BrandTranslation) (*entities.BrandTranslation, error)
	GetTranslations(ctx context.Context, brandID uint) ([]*entities.BrandTranslation, error)
	GetTranslationByLocale(ctx context.Context, brandID uint, locale string) (*entities.BrandTranslation, error)

	// Status operations
	UpdateStatus(ctx context.Context, id uint, status int) error
	GetCount(ctx context.Context) (int64, error)
}
