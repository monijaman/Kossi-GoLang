package postgres

import (
	"context"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/infrastructure/database/models"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type PostgresProductRepo struct {
	db *gorm.DB
}

func NewPostgresProductRepo(db *gorm.DB) *PostgresProductRepo {
	return &PostgresProductRepo{db: db}
}

// generateSlug creates a URL-friendly slug from a name
func generateSlug(name string) string {
	// Convert to lowercase
	slug := strings.ToLower(name)
	// Replace spaces and special characters with hyphens
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")
	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")
	return slug
}

func (r *PostgresProductRepo) GetByID(ctx context.Context, id uint) (*entities.Product, error) {
	var productModel models.ProductModel
	if err := r.db.WithContext(ctx).First(&productModel, id).Error; err != nil {
		return nil, err
	}
	return productModel.ToEntity(), nil
}

func (r *PostgresProductRepo) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	// Convert entity to model
	var productModel models.ProductModel
	productModel.FromEntity(product)

	// Generate slug if not provided
	if productModel.Slug == "" {
		productModel.Slug = generateSlug(productModel.Name)
	}

	// Set default category ID if not provided
	if productModel.CategoryID == "" {
		productModel.CategoryID = "1" // Default category
	}

	if err := r.db.WithContext(ctx).Create(&productModel).Error; err != nil {
		return nil, err
	}

	return productModel.ToEntity(), nil
}

var _ repository.ProductRepository = (*PostgresProductRepo)(nil)
