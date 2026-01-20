// Package usecase contains application logic and use cases
package usecase

import (
	"context"
	"go-crit/internal/domain/entities"
)

// ProductUsecase represents the product use case
type ProductUsecase struct {
	productRepository ProductRepository
}

// NewProductUsecase creates a new instance of ProductUsecase
func NewProductUsecase(productRepo ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		productRepository: productRepo,
	}
}

// GetProductsFilter defines the filter for getting products
type GetProductsFilter struct {
	Page       int
	Limit      int
	Category   string
	Brand      string
	PriceRange string
	SearchTerm string
	Locale     string
}

// GetProducts retrieves products based on the filter
func (u *ProductUsecase) GetProducts(ctx context.Context, filter *GetProductsFilter) ([]entities.Product, int64, error) {
	products, total, err := u.productRepository.GetProducts(ctx, filter)
	return products, total, err
}

// ProductRepository defines the methods that a product repository should implement
type ProductRepository interface {
	GetProducts(ctx context.Context, filter *GetProductsFilter) ([]entities.Product, int64, error)
}
