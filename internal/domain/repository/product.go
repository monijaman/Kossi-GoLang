package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

type ProductRepository interface {
	GetByID(ctx context.Context, id uint) (*entities.Product, error)
	Create(ctx context.Context, product *entities.Product) (*entities.Product, error)
}
