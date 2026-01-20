// Package repository contains the implementation of the product repository
package repository

import (
	"context"
	"koshti/gocrit_server/internal/domain/entities"
	"koshti/gocrit_server/internal/interface/models"

	"gorm.io/gorm"
)

// ProductRepository represents the product repository
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

// GetProductsFilter represents the filter for getting products
type GetProductsFilter struct {
	Page       int
	Limit      int
	Category   string
	Brand      string
	SearchTerm string
}

// GetProducts retrieves products based on the filter
func (r *ProductRepository) GetProducts(ctx context.Context, filter *GetProductsFilter) ([]entities.Product, int64, error) {
	var products []models.Product
	query := r.db.WithContext(ctx)

	if filter.Category != "" {
		query = query.Joins("LEFT JOIN categories ON products.category_id = categories.id").
			Where("categories.id = ? OR categories.slug = ?", filter.Category, filter.Category)
	}
	if filter.Brand != "" {
		query = query.Where("brand_id = ?", filter.Brand)
	}
	if filter.SearchTerm != "" {
		query = query.Where("name ILIKE ?", "%"+filter.SearchTerm+"%")
	}

	var total int64
	query.Model(&models.Product{}).Count(&total)

	offset := (filter.Page - 1) * filter.Limit
	err := query.Offset(offset).Limit(filter.Limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	var productEntities []entities.Product
	for _, p := range products {
		productEntities = append(productEntities, entities.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			CategoryID:  p.CategoryID,
			BrandID:     p.BrandID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return productEntities, total, nil
}
