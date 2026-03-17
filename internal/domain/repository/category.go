package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

// CategoryRepository defines the interface for category-related database operations
type CategoryRepository interface {
	// Basic CRUD operations for categories
	Create(ctx context.Context, category *entities.Category) (*entities.Category, error)
	GetByID(ctx context.Context, id uint) (*entities.Category, error)
	GetBySlug(ctx context.Context, slug string) (*entities.Category, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entities.Category, error)
	Update(ctx context.Context, id uint, category *entities.Category) (*entities.Category, error)
	Delete(ctx context.Context, id uint) error
	Search(ctx context.Context, query string, limit, offset int) ([]*entities.Category, error)
	GetWideCategories(ctx context.Context, limit int) ([]*entities.Category, error)

	// Translation operations
	CreateTranslation(ctx context.Context, translation *entities.CategoryTranslation) (*entities.CategoryTranslation, error)
	UpdateTranslation(ctx context.Context, translation *entities.CategoryTranslation) (*entities.CategoryTranslation, error)
	GetTranslations(ctx context.Context, categoryID uint) ([]*entities.CategoryTranslation, error)
	GetTranslationByID(ctx context.Context, translationID uint) (*entities.CategoryTranslation, error)
	GetTranslationByLocale(ctx context.Context, categoryID uint, locale string) (*entities.CategoryTranslation, error)

	// Brand-Category relationship operations
	CreateBrandRelation(ctx context.Context, relation *entities.BrandCategory) (*entities.BrandCategory, error)
	GetBrandRelations(ctx context.Context, categoryID uint) ([]*entities.BrandCategory, error)
	GetCategoryBrandRelations(ctx context.Context) ([]*entities.BrandCategory, error)
	DeleteBrandRelation(ctx context.Context, categoryID, brandID uint) error
	DeleteBrandRelationsByCategory(ctx context.Context, categoryID uint) error
	GetBrandsByIDs(ctx context.Context, brandIDs []uint) ([]*entities.Brand, error)

	// Status operations
	UpdateStatus(ctx context.Context, id uint, status int) error
	GetCount(ctx context.Context) (int64, error)
	// Batch-fetches translated names for a list of category IDs; returns map[categoryID]translatedName
	GetTranslatedNamesByCategoryIDs(ctx context.Context, categoryIDs []uint, locale string) (map[uint]string, error)
}
