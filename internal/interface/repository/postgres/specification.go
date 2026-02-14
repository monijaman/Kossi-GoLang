package postgres

import (
	"context"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/infrastructure/database/models"
	"time"

	"gorm.io/gorm"
)

type PostgresSpecificationRepo struct {
	db *gorm.DB
}

func NewPostgresSpecificationRepo(db *gorm.DB) *PostgresSpecificationRepo {
	return &PostgresSpecificationRepo{db: db}
}

// Basic CRUD operations for specifications
func (r *PostgresSpecificationRepo) Create(ctx context.Context, spec *entities.Specification) (*entities.Specification, error) {
	// Check for existing specification with same product_id and specification_key_id
	existingSpec, err := r.GetByProductAndKey(ctx, spec.ProductID, spec.SpecificationKeyID)
	if err != nil {
		return nil, err
	}

	if existingSpec != nil {
		// Update existing specification
		existingSpec.Value = spec.Value
		existingSpec.UpdatedAt = time.Now()

		var specModel models.SpecificationModel
		specModel.FromEntity(existingSpec)

		if err := r.db.WithContext(ctx).Save(&specModel).Error; err != nil {
			return nil, err
		}

		return specModel.ToEntity(), nil
	}

	// Create new specification
	var specModel models.SpecificationModel
	specModel.FromEntity(spec)

	if err := r.db.WithContext(ctx).Create(&specModel).Error; err != nil {
		return nil, err
	}

	return specModel.ToEntity(), nil
}

func (r *PostgresSpecificationRepo) GetByID(ctx context.Context, id uint) (*entities.Specification, error) {
	var specModel models.SpecificationModel

	if err := r.db.WithContext(ctx).First(&specModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("specification not found")
		}
		return nil, err
	}

	return specModel.ToEntity(), nil
}

func (r *PostgresSpecificationRepo) GetByProductID(ctx context.Context, productID uint) ([]*entities.Specification, error) {
	var results []struct {
		models.SpecificationModel
		SpecificationKeyName string `gorm:"column:specification_key_name"`
	}

	if err := r.db.WithContext(ctx).
		Table("specifications").
		Select("specifications.*, specification_keys.specification_key as specification_key_name").
		Joins("LEFT JOIN specification_keys ON specifications.specification_key_id = specification_keys.id").
		Where("specifications.product_id = ?", productID).
		Find(&results).Error; err != nil {
		return nil, err
	}

	specs := make([]*entities.Specification, len(results))
	for i, result := range results {
		spec := result.SpecificationModel.ToEntity()
		spec.SpecificationKey = result.SpecificationKeyName
		specs[i] = spec
	}

	return specs, nil
}

func (r *PostgresSpecificationRepo) GetByProductAndKey(ctx context.Context, productID, keyID uint) (*entities.Specification, error) {
	var specModel models.SpecificationModel

	if err := r.db.WithContext(ctx).
		Where("product_id = ? AND specification_key_id = ?", productID, keyID).
		First(&specModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil without error if not found (this is expected for checking duplicates)
		}
		return nil, err
	}

	return specModel.ToEntity(), nil
}

func (r *PostgresSpecificationRepo) Update(ctx context.Context, id uint, spec *entities.Specification) (*entities.Specification, error) {
	var specModel models.SpecificationModel
	specModel.FromEntity(spec)
	specModel.ID = id

	result := r.db.WithContext(ctx).Save(&specModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("specification not found")
	}

	return specModel.ToEntity(), nil
}

func (r *PostgresSpecificationRepo) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.SpecificationModel{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("specification not found")
	}

	return nil
}

func (r *PostgresSpecificationRepo) BulkUpsert(ctx context.Context, specs []*entities.Specification) ([]*entities.Specification, error) {
	if len(specs) == 0 {
		return []*entities.Specification{}, nil
	}

	var result []*entities.Specification

	// Process each specification
	for _, spec := range specs {
		var existingSpec *entities.Specification
		var err error

		if spec.ID != 0 {
			// ID provided: try to update that specific record
			existingSpec, err = r.GetByID(ctx, spec.ID)
			if err != nil && err.Error() != "specification not found" {
				return nil, err
			}
			if existingSpec != nil {
				// Update existing specification
				existingSpec.Value = spec.Value
				existingSpec.SpecificationKeyID = spec.SpecificationKeyID
				existingSpec.ProductID = spec.ProductID
				existingSpec.UpdatedAt = time.Now()

				var specModel models.SpecificationModel
				specModel.FromEntity(existingSpec)

				if err := r.db.WithContext(ctx).Save(&specModel).Error; err != nil {
					return nil, err
				}

				result = append(result, specModel.ToEntity())
			} else {
				// ID not found, create new with this ID
				spec.CreatedAt = time.Now()
				spec.UpdatedAt = time.Now()

				var specModel models.SpecificationModel
				specModel.FromEntity(spec)

				if err := r.db.WithContext(ctx).Create(&specModel).Error; err != nil {
					return nil, err
				}

				result = append(result, specModel.ToEntity())
			}
		} else {
			// No ID provided: upsert by product_id and specification_key_id
			existingSpec, err = r.GetByProductAndKey(ctx, spec.ProductID, spec.SpecificationKeyID)
			if err != nil {
				return nil, err
			}

			if existingSpec != nil {
				// Update existing specification
				existingSpec.Value = spec.Value
				existingSpec.UpdatedAt = time.Now()

				var specModel models.SpecificationModel
				specModel.FromEntity(existingSpec)

				if err := r.db.WithContext(ctx).Save(&specModel).Error; err != nil {
					return nil, err
				}

				result = append(result, specModel.ToEntity())
			} else {
				// Create new specification
				spec.CreatedAt = time.Now()
				spec.UpdatedAt = time.Now()

				var specModel models.SpecificationModel
				specModel.FromEntity(spec)

				if err := r.db.WithContext(ctx).Create(&specModel).Error; err != nil {
					return nil, err
				}

				result = append(result, specModel.ToEntity())
			}
		}
	}

	return result, nil
}

func (r *PostgresSpecificationRepo) Search(ctx context.Context, query string, limit, offset int) ([]*entities.Specification, error) {
	var specModels []models.SpecificationModel

	if err := r.db.WithContext(ctx).
		Where("value ILIKE ?", "%"+query+"%").
		Limit(limit).
		Offset(offset).
		Find(&specModels).Error; err != nil {
		return nil, err
	}

	specs := make([]*entities.Specification, len(specModels))
	for i, model := range specModels {
		specs[i] = model.ToEntity()
	}

	return specs, nil
}

// Translation operations
func (r *PostgresSpecificationRepo) CreateTranslation(ctx context.Context, translation *entities.SpecificationTranslation) (*entities.SpecificationTranslation, error) {
	var translationModel models.SpecificationTranslationModel
	translationModel.FromEntity(translation)

	if err := r.db.WithContext(ctx).Create(&translationModel).Error; err != nil {
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

// UpsertTranslation creates or updates a translation
func (r *PostgresSpecificationRepo) UpsertTranslation(ctx context.Context, translation *entities.SpecificationTranslation) (*entities.SpecificationTranslation, error) {
	var translationModel models.SpecificationTranslationModel

	// Try to find existing translation
	err := r.db.WithContext(ctx).
		Where("specification_id = ? AND locale = ?", translation.SpecificationID, translation.Locale).
		First(&translationModel).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new translation
		translationModel.FromEntity(translation)
		translationModel.ID = 0 // Ensure it's treated as new record
	} else {
		// Update existing translation
		translationModel.Value = translation.TranslatedValue
		translationModel.UpdatedAt = time.Now()
	}

	if err := r.db.WithContext(ctx).Save(&translationModel).Error; err != nil {
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

func (r *PostgresSpecificationRepo) GetTranslations(ctx context.Context, specID uint) ([]*entities.SpecificationTranslation, error) {
	var translationModels []models.SpecificationTranslationModel

	if err := r.db.WithContext(ctx).
		Where("specification_id = ?", specID).
		Find(&translationModels).Error; err != nil {
		return nil, err
	}

	translations := make([]*entities.SpecificationTranslation, len(translationModels))
	for i, model := range translationModels {
		translations[i] = model.ToEntity()
	}

	return translations, nil
}

func (r *PostgresSpecificationRepo) GetTranslationByLocale(ctx context.Context, specID uint, locale string) (*entities.SpecificationTranslation, error) {
	var translationModel models.SpecificationTranslationModel

	if err := r.db.WithContext(ctx).
		Where("specification_id = ? AND locale = ?", specID, locale).
		First(&translationModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("translation not found")
		}
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

// Ensure the implementation satisfies the interface
var _ repository.SpecificationRepository = (*PostgresSpecificationRepo)(nil)
