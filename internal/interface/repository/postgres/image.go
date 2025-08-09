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
	result := r.db.WithContext(ctx).Delete(&models.ImageModel{}, id)
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

// Ensure the implementation satisfies the interface
var _ repository.ImageRepository = (*PostgresImageRepo)(nil)
