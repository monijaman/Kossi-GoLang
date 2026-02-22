// Package repository contains the implementation of the product repository
package repository

import (
	"context"
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

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
	var products []models.ProductModel
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
	query.Model(&models.ProductModel{}).Count(&total)

	offset := (filter.Page - 1) * filter.Limit
	err := query.Offset(offset).Limit(filter.Limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	var productEntities []entities.Product
	for _, p := range products {
		// Map start/end price from model
		var startPrice *float64
		var endPrice *float64
		if p.StartPrice != nil {
			startPrice = p.StartPrice
		}
		if p.EndPrice != nil {
			endPrice = p.EndPrice
		}
		productEntities = append(productEntities, entities.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       0.0,
			StartPrice:  startPrice,
			EndPrice:    endPrice,
			CategoryID:  p.CategoryID,
			BrandID:     p.BrandID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return productEntities, total, nil
}
