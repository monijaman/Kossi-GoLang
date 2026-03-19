package postgres

import (
	"context"
	"errors"
	"fmt"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/infrastructure/database/models"
	"strings"
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

	// Use a single transaction for atomicity and performance
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Collect all unique product IDs to pre-fetch existing specs
		productIDs := make(map[uint]bool)
		for _, spec := range specs {
			productIDs[spec.ProductID] = true
		}

		// Pre-fetch ALL existing specs for these products in ONE query
		var existingModels []models.SpecificationModel
		pids := make([]uint, 0, len(productIDs))
		for pid := range productIDs {
			pids = append(pids, pid)
		}
		if err := tx.Where("product_id IN ?", pids).Find(&existingModels).Error; err != nil {
			return err
		}

		// Build lookup maps for fast access
		// Map by ID for specs that have an ID
		byID := make(map[uint]*models.SpecificationModel)
		// Map by (product_id, specification_key_id) for specs without an ID
		type prodKeyPair struct {
			ProductID uint
			KeyID     uint
		}
		byProdKey := make(map[prodKeyPair]*models.SpecificationModel)

		for i := range existingModels {
			m := &existingModels[i]
			byID[m.ID] = m
			byProdKey[prodKeyPair{m.ProductID, m.SpecificationKeyID}] = m
		}

		now := time.Now()

		for _, spec := range specs {
			if spec.ID != 0 {
				// ID provided: try to update that specific record
				if existing, ok := byID[spec.ID]; ok {
					existing.Value = spec.Value
					existing.SpecificationKeyID = spec.SpecificationKeyID
					existing.ProductID = spec.ProductID
					existing.UpdatedAt = now

					if err := tx.Save(existing).Error; err != nil {
						return err
					}
					result = append(result, existing.ToEntity())
				} else {
					// ID not found, create new
					spec.CreatedAt = now
					spec.UpdatedAt = now

					var specModel models.SpecificationModel
					specModel.FromEntity(spec)

					if err := tx.Create(&specModel).Error; err != nil {
						return err
					}
					result = append(result, specModel.ToEntity())
				}
			} else {
				// No ID provided: upsert by product_id and specification_key_id
				key := prodKeyPair{spec.ProductID, spec.SpecificationKeyID}
				if existing, ok := byProdKey[key]; ok {
					existing.Value = spec.Value
					existing.UpdatedAt = now

					if err := tx.Save(existing).Error; err != nil {
						return err
					}
					result = append(result, existing.ToEntity())
				} else {
					// Create new specification
					spec.CreatedAt = now
					spec.UpdatedAt = now

					var specModel models.SpecificationModel
					specModel.FromEntity(spec)

					if err := tx.Create(&specModel).Error; err != nil {
						return err
					}
					// Add to lookup map in case there are duplicates in the same batch
					byProdKey[key] = &specModel
					result = append(result, specModel.ToEntity())
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
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

// BulkUpsertTranslations creates or updates multiple translations in a single transaction
func (r *PostgresSpecificationRepo) BulkUpsertTranslations(ctx context.Context, translations []*entities.SpecificationTranslation) ([]*entities.SpecificationTranslation, error) {
	if len(translations) == 0 {
		return []*entities.SpecificationTranslation{}, nil
	}

	now := time.Now()

	// Build a single INSERT … ON CONFLICT DO UPDATE statement so the entire
	// batch is handled in ONE database round-trip regardless of how many rows.
	//
	// SQL shape:
	//   INSERT INTO specification_translations (specification_id, locale, value, created_at, updated_at)
	//   VALUES ($1,$2,$3,$4,$5), ($6,$7,$8,$9,$10), ...
	//   ON CONFLICT (specification_id, locale)
	//   DO UPDATE SET value = EXCLUDED.value, updated_at = EXCLUDED.updated_at
	//   RETURNING id, specification_id, locale, value, created_at, updated_at

	placeholders := make([]string, 0, len(translations))
	args := make([]interface{}, 0, len(translations)*5)
	paramIdx := 1

	for _, t := range translations {
		placeholders = append(placeholders,
			fmt.Sprintf("($%d,$%d,$%d,$%d,$%d)", paramIdx, paramIdx+1, paramIdx+2, paramIdx+3, paramIdx+4))
		args = append(args, t.SpecificationID, t.Locale, t.TranslatedValue, now, now)
		paramIdx += 5
	}

	query := fmt.Sprintf(`
		INSERT INTO specification_translations (specification_id, locale, value, created_at, updated_at)
		VALUES %s
		ON CONFLICT (specification_id, locale)
		DO UPDATE SET value = EXCLUDED.value, updated_at = EXCLUDED.updated_at
		RETURNING id, specification_id, locale, value, created_at, updated_at`,
		strings.Join(placeholders, ","))

	rows, err := r.db.WithContext(ctx).Raw(query, args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entities.SpecificationTranslation
	for rows.Next() {
		var m models.SpecificationTranslationModel
		if err := rows.Scan(&m.ID, &m.SpecificationID, &m.Locale, &m.Value, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, m.ToEntity())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
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

// GetPublicSpecsWithTranslations returns all specs for a product with translations in a single query
func (r *PostgresSpecificationRepo) GetPublicSpecsWithTranslations(ctx context.Context, productID uint, locale string) ([]repository.PublicSpecResult, error) {
	type row struct {
		SpecificationKeyID uint   `gorm:"column:specification_key_id"`
		TranslatedKey      string `gorm:"column:translated_key"`
		TranslatedValue    string `gorm:"column:translated_value"`
	}

	var rows []row
	err := r.db.WithContext(ctx).Raw(`
		SELECT
			s.specification_key_id,
			COALESCE(NULLIF(kt.specification_key, ''), sk.specification_key) AS translated_key,
			COALESCE(NULLIF(st.value, ''), s.value) AS translated_value
		FROM specifications s
		JOIN specification_keys sk ON sk.id = s.specification_key_id
		LEFT JOIN specification_key_translations kt
			ON kt.specification_key_id = s.specification_key_id AND kt.locale = ?
		LEFT JOIN specification_translations st
			ON st.specification_id = s.id AND st.locale = ?
		WHERE s.product_id = ?
	`, locale, locale, productID).Scan(&rows).Error

	if err != nil {
		return nil, err
	}

	results := make([]repository.PublicSpecResult, len(rows))
	for i, r := range rows {
		results[i] = repository.PublicSpecResult{
			SpecificationKeyID: r.SpecificationKeyID,
			TranslatedKey:      r.TranslatedKey,
			TranslatedValue:    r.TranslatedValue,
		}
	}
	return results, nil
}

// Ensure the implementation satisfies the interface
var _ repository.SpecificationRepository = (*PostgresSpecificationRepo)(nil)
