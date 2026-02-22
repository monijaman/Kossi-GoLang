package postgres

import (
	"context"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PostgresImageRepo struct {
	db *gorm.DB
}

func NewPostgresImageRepo(db *gorm.DB) *PostgresImageRepo {
	return &PostgresImageRepo{db: db}
}

func (r *PostgresImageRepo) Create(ctx context.Context, image *entities.Image) (*entities.Image, error) {
	var imageModel models.ImageModel
	imageModel.FromEntity(image)

	if err := r.db.WithContext(ctx).Create(&imageModel).Error; err != nil {
		return nil, err
	}

	return imageModel.ToEntity(), nil
}

func (r *PostgresImageRepo) GetByImageableID(ctx context.Context, imageableType string, imageableID uint) ([]*entities.Image, error) {
	var imageModels []models.ImageModel

	if err := r.db.WithContext(ctx).
		Where("imageable_type = ? AND imageable_id = ? AND deleted_at IS NULL", imageableType, imageableID).
		Find(&imageModels).Error; err != nil {
		return nil, err
	}

	images := make([]*entities.Image, len(imageModels))
	for i, model := range imageModels {
		images[i] = model.ToEntity()
	}

	return images, nil
}

func (r *PostgresImageRepo) GetByID(ctx context.Context, id uint) (*entities.Image, error) {
	var imageModel models.ImageModel

	if err := r.db.WithContext(ctx).First(&imageModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("image not found")
		}
		return nil, err
	}

	return imageModel.ToEntity(), nil
}

func (r *PostgresImageRepo) Delete(ctx context.Context, id uint) error {
	// Use Unscoped() to perform a hard delete (not soft delete)
	result := r.db.WithContext(ctx).Unscoped().Delete(&models.ImageModel{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("image not found")
	}

	return nil
}

func (r *PostgresImageRepo) Update(ctx context.Context, image *entities.Image) (*entities.Image, error) {
	var imageModel models.ImageModel
	imageModel.FromEntity(image)

	result := r.db.WithContext(ctx).Save(&imageModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("image not found")
	}

	return imageModel.ToEntity(), nil
}

// GetDefaultImagesByProductIDs fetches default images for multiple products in a single query
// Returns a map of productID -> image, avoiding N+1 queries
func (r *PostgresImageRepo) GetDefaultImagesByProductIDs(ctx context.Context, productIDs []uint) (map[uint]*entities.Image, error) {
	result := make(map[uint]*entities.Image)
	
	if len(productIDs) == 0 {
		return result, nil
	}

	var imageModels []models.ImageModel

	// Fetch all images for these products, preferring default_photo = 1
	if err := r.db.WithContext(ctx).
		Where("imageable_type = ? AND imageable_id IN ? AND deleted_at IS NULL", "Product", productIDs).
		Order("default_photo DESC"). // Default photo first
		Find(&imageModels).Error; err != nil {
		return result, err
	}

	// For each product, keep only the first (default) image
	for _, model := range imageModels {
		if _, exists := result[model.ImageableID]; !exists {
			// Only add if we haven't already added an image for this product
			result[model.ImageableID] = model.ToEntity()
		}
	}

	return result, nil
}

// Ensure the implementation satisfies the interface
var _ repository.ImageRepository = (*PostgresImageRepo)(nil)
