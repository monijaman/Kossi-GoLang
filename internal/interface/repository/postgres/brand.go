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

type PostgresBrandRepo struct {
	db *gorm.DB
}

func NewPostgresBrandRepo(db *gorm.DB) *PostgresBrandRepo {
	return &PostgresBrandRepo{db: db}
}

// Basic CRUD operations for brands
func (r *PostgresBrandRepo) Create(ctx context.Context, brand *entities.Brand) (*entities.Brand, error) {
	var brandModel models.BrandModel
	brandModel.FromEntity(brand)

	if err := r.db.WithContext(ctx).Create(&brandModel).Error; err != nil {
		return nil, err
	}

	return brandModel.ToEntity(), nil
}

func (r *PostgresBrandRepo) GetByID(ctx context.Context, id uint) (*entities.Brand, error) {
	var brandModel models.BrandModel

	if err := r.db.WithContext(ctx).First(&brandModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("brand not found")
		}
		return nil, err
	}

	return brandModel.ToEntity(), nil
}

func (r *PostgresBrandRepo) GetBySlug(ctx context.Context, slug string) (*entities.Brand, error) {
	var brandModel models.BrandModel

	if err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&brandModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("brand not found")
		}
		return nil, err
	}

	return brandModel.ToEntity(), nil
}

func (r *PostgresBrandRepo) GetAll(ctx context.Context, limit, offset int) ([]*entities.Brand, error) {
	var brandModels []models.BrandModel

	query := r.db.WithContext(ctx)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&brandModels).Error; err != nil {
		return nil, err
	}

	brands := make([]*entities.Brand, len(brandModels))
	for i, model := range brandModels {
		brands[i] = model.ToEntity()
	}

	return brands, nil
}

func (r *PostgresBrandRepo) Update(ctx context.Context, id uint, brand *entities.Brand) (*entities.Brand, error) {
	var brandModel models.BrandModel
	brandModel.FromEntity(brand)
	brandModel.ID = id
	brandModel.UpdatedAt = time.Now()

	result := r.db.WithContext(ctx).Save(&brandModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("brand not found")
	}

	return brandModel.ToEntity(), nil
}

func (r *PostgresBrandRepo) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.BrandModel{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("brand not found")
	}

	return nil
}

func (r *PostgresBrandRepo) Search(ctx context.Context, query string, limit, offset int) ([]*entities.Brand, error) {
	var brandModels []models.BrandModel

	dbQuery := r.db.WithContext(ctx).Where("name ILIKE ? OR slug ILIKE ?", "%"+query+"%", "%"+query+"%")

	if limit > 0 {
		dbQuery = dbQuery.Limit(limit)
	}
	if offset > 0 {
		dbQuery = dbQuery.Offset(offset)
	}

	if err := dbQuery.Find(&brandModels).Error; err != nil {
		return nil, err
	}

	brands := make([]*entities.Brand, len(brandModels))
	for i, model := range brandModels {
		brands[i] = model.ToEntity()
	}

	return brands, nil
}

func (r *PostgresBrandRepo) GetWideBrands(ctx context.Context, limit int) ([]*entities.Brand, error) {
	var brandModels []models.BrandModel

	query := r.db.WithContext(ctx).Order("name ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&brandModels).Error; err != nil {
		return nil, err
	}

	brands := make([]*entities.Brand, len(brandModels))
	for i, model := range brandModels {
		brands[i] = model.ToEntity()
	}

	return brands, nil
}

func (r *PostgresBrandRepo) GetPublicBrands(ctx context.Context, limit, offset int) ([]*entities.Brand, error) {
	var brandModels []models.BrandModel

	query := r.db.WithContext(ctx).Where("deleted_at IS NULL")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&brandModels).Error; err != nil {
		return nil, err
	}

	brands := make([]*entities.Brand, len(brandModels))
	for i, model := range brandModels {
		brands[i] = model.ToEntity()
	}

	return brands, nil
}

// Translation operations
func (r *PostgresBrandRepo) CreateTranslation(ctx context.Context, translation *entities.BrandTranslation) (*entities.BrandTranslation, error) {
	var translationModel models.BrandTranslationModel
	translationModel.FromEntity(translation)

	if err := r.db.WithContext(ctx).Create(&translationModel).Error; err != nil {
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

func (r *PostgresBrandRepo) GetTranslations(ctx context.Context, brandID uint) ([]*entities.BrandTranslation, error) {
	var translationModels []models.BrandTranslationModel

	if err := r.db.WithContext(ctx).
		Where("brand_id = ?", brandID).
		Find(&translationModels).Error; err != nil {
		return nil, err
	}

	translations := make([]*entities.BrandTranslation, len(translationModels))
	for i, model := range translationModels {
		translations[i] = model.ToEntity()
	}

	return translations, nil
}

func (r *PostgresBrandRepo) GetTranslationByLocale(ctx context.Context, brandID uint, locale string) (*entities.BrandTranslation, error) {
	var translationModel models.BrandTranslationModel

	if err := r.db.WithContext(ctx).
		Where("brand_id = ? AND locale = ?", brandID, locale).
		First(&translationModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("translation not found")
		}
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

// Status operations
func (r *PostgresBrandRepo) UpdateStatus(ctx context.Context, id uint, status int) error {
	result := r.db.WithContext(ctx).
		Model(&models.BrandModel{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("brand not found")
	}

	return nil
}

func (r *PostgresBrandRepo) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.BrandModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Ensure the implementation satisfies the interface
var _ repository.BrandRepository = (*PostgresBrandRepo)(nil)
