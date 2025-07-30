package postgres

import (
	"context"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"

	"gorm.io/gorm"
)

type PostgresProductRepo struct {
	db *gorm.DB
}

func NewPostgresProductRepo(db *gorm.DB) *PostgresProductRepo {
	return &PostgresProductRepo{db: db}
}

func (r *PostgresProductRepo) GetByID(ctx context.Context, id uint) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *PostgresProductRepo) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

var _ repository.ProductRepository = (*PostgresProductRepo)(nil)
