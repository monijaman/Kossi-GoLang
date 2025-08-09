package postgres

import (
	"context"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PostgresSpecificationKeyRepo struct {
	db *gorm.DB
}

func NewPostgresSpecificationKeyRepo(db *gorm.DB) *PostgresSpecificationKeyRepo {
	return &PostgresSpecificationKeyRepo{db: db}
}

// Basic CRUD operations for specification keys
func (r *PostgresSpecificationKeyRepo) Create(ctx context.Context, key *entities.SpecificationKey) (*entities.SpecificationKey, error) {
	var keyModel models.SpecificationKeyModel
	keyModel.FromEntity(key)

	if err := r.db.WithContext(ctx).Create(&keyModel).Error; err != nil {
		return nil, err
	}

	return keyModel.ToEntity(), nil
}

func (r *PostgresSpecificationKeyRepo) GetByID(ctx context.Context, id uint) (*entities.SpecificationKey, error) {
	var keyModel models.SpecificationKeyModel

	if err := r.db.WithContext(ctx).First(&keyModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("specification key not found")
		}
		return nil, err
	}

	return keyModel.ToEntity(), nil
}

func (r *PostgresSpecificationKeyRepo) GetAll(ctx context.Context, limit, offset int) ([]*entities.SpecificationKey, error) {
	var keyModels []models.SpecificationKeyModel

	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&keyModels).Error; err != nil {
		return nil, err
	}

	keys := make([]*entities.SpecificationKey, len(keyModels))
	for i, model := range keyModels {
		keys[i] = model.ToEntity()
	}

	return keys, nil
}

func (r *PostgresSpecificationKeyRepo) GetByKey(ctx context.Context, key string) (*entities.SpecificationKey, error) {
	var keyModel models.SpecificationKeyModel

	if err := r.db.WithContext(ctx).
		Where("specification_key = ?", key).
		First(&keyModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("specification key not found")
		}
		return nil, err
	}

	return keyModel.ToEntity(), nil
}

func (r *PostgresSpecificationKeyRepo) Update(ctx context.Context, id uint, key *entities.SpecificationKey) (*entities.SpecificationKey, error) {
	var keyModel models.SpecificationKeyModel
	keyModel.FromEntity(key)
	keyModel.ID = id

	result := r.db.WithContext(ctx).Save(&keyModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("specification key not found")
	}

	return keyModel.ToEntity(), nil
}

func (r *PostgresSpecificationKeyRepo) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.SpecificationKeyModel{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("specification key not found")
	}

	return nil
}

// Translation operations
func (r *PostgresSpecificationKeyRepo) CreateKeyTranslation(ctx context.Context, translation *entities.SpecificationKeyTranslation) (*entities.SpecificationKeyTranslation, error) {
	var translationModel models.SpecificationKeyTranslationModel
	translationModel.FromEntity(translation)

	if err := r.db.WithContext(ctx).Create(&translationModel).Error; err != nil {
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

func (r *PostgresSpecificationKeyRepo) GetKeyTranslations(ctx context.Context, keyID uint) ([]*entities.SpecificationKeyTranslation, error) {
	var translationModels []models.SpecificationKeyTranslationModel

	if err := r.db.WithContext(ctx).
		Where("specification_key_id = ?", keyID).
		Find(&translationModels).Error; err != nil {
		return nil, err
	}

	translations := make([]*entities.SpecificationKeyTranslation, len(translationModels))
	for i, model := range translationModels {
		translations[i] = model.ToEntity()
	}

	return translations, nil
}

func (r *PostgresSpecificationKeyRepo) GetKeyTranslationByLocale(ctx context.Context, keyID uint, locale string) (*entities.SpecificationKeyTranslation, error) {
	var translationModel models.SpecificationKeyTranslationModel

	if err := r.db.WithContext(ctx).
		Where("specification_key_id = ? AND locale = ?", keyID, locale).
		First(&translationModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("key translation not found")
		}
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

func (r *PostgresSpecificationKeyRepo) GetAllKeyTranslations(ctx context.Context) ([]*entities.SpecificationKeyTranslation, error) {
	var translationModels []models.SpecificationKeyTranslationModel

	if err := r.db.WithContext(ctx).Find(&translationModels).Error; err != nil {
		return nil, err
	}

	translations := make([]*entities.SpecificationKeyTranslation, len(translationModels))
	for i, model := range translationModels {
		translations[i] = model.ToEntity()
	}

	return translations, nil
}

// Ensure the implementation satisfies the interface
var _ repository.SpecificationKeyRepository = (*PostgresSpecificationKeyRepo)(nil)
