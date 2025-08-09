package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

// SpecificationRepository defines the interface for specification-related database operations
type SpecificationRepository interface {
	// Basic CRUD operations for specifications
	Create(ctx context.Context, spec *entities.Specification) (*entities.Specification, error)
	GetByID(ctx context.Context, id uint) (*entities.Specification, error)
	GetByProductID(ctx context.Context, productID uint) ([]*entities.Specification, error)
	GetByProductAndKey(ctx context.Context, productID, keyID uint) (*entities.Specification, error)
	Update(ctx context.Context, id uint, spec *entities.Specification) (*entities.Specification, error)
	BulkUpsert(ctx context.Context, specs []*entities.Specification) ([]*entities.Specification, error)
	Delete(ctx context.Context, id uint) error
	Search(ctx context.Context, query string, limit, offset int) ([]*entities.Specification, error)

	// Translation operations
	CreateTranslation(ctx context.Context, translation *entities.SpecificationTranslation) (*entities.SpecificationTranslation, error)
	GetTranslations(ctx context.Context, specID uint) ([]*entities.SpecificationTranslation, error)
	GetTranslationByLocale(ctx context.Context, specID uint, locale string) (*entities.SpecificationTranslation, error)
}

// SpecificationKeyRepository defines the interface for specification key operations
type SpecificationKeyRepository interface {
	// Basic CRUD operations for specification keys
	Create(ctx context.Context, key *entities.SpecificationKey) (*entities.SpecificationKey, error)
	GetByID(ctx context.Context, id uint) (*entities.SpecificationKey, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entities.SpecificationKey, error)
	GetByKey(ctx context.Context, key string) (*entities.SpecificationKey, error)
	Update(ctx context.Context, id uint, key *entities.SpecificationKey) (*entities.SpecificationKey, error)
	Delete(ctx context.Context, id uint) error

	// Translation operations
	CreateKeyTranslation(ctx context.Context, translation *entities.SpecificationKeyTranslation) (*entities.SpecificationKeyTranslation, error)
	GetKeyTranslations(ctx context.Context, keyID uint) ([]*entities.SpecificationKeyTranslation, error)
	GetKeyTranslationByLocale(ctx context.Context, keyID uint, locale string) (*entities.SpecificationKeyTranslation, error)
	GetAllKeyTranslations(ctx context.Context) ([]*entities.SpecificationKeyTranslation, error)
}
