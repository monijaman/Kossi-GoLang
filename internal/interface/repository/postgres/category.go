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

type PostgresCategoryRepo struct {
	db *gorm.DB
}

func NewPostgresCategoryRepo(db *gorm.DB) *PostgresCategoryRepo {
	return &PostgresCategoryRepo{db: db}
}

// Basic CRUD operations for categories
func (r *PostgresCategoryRepo) Create(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	var categoryModel models.CategoryModel
	categoryModel.FromEntity(category)

	if err := r.db.WithContext(ctx).Create(&categoryModel).Error; err != nil {
		return nil, err
	}

	return categoryModel.ToEntity(), nil
}

func (r *PostgresCategoryRepo) GetByID(ctx context.Context, id uint) (*entities.Category, error) {
	var categoryModel models.CategoryModel

	if err := r.db.WithContext(ctx).First(&categoryModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return categoryModel.ToEntity(), nil
}

func (r *PostgresCategoryRepo) GetBySlug(ctx context.Context, slug string) (*entities.Category, error) {
	var categoryModel models.CategoryModel

	if err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&categoryModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return categoryModel.ToEntity(), nil
}

func (r *PostgresCategoryRepo) GetAll(ctx context.Context, limit, offset int) ([]*entities.Category, error) {
	var categoryModels []models.CategoryModel

	query := r.db.WithContext(ctx)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&categoryModels).Error; err != nil {
		return nil, err
	}

	categories := make([]*entities.Category, len(categoryModels))
	for i, model := range categoryModels {
		categories[i] = model.ToEntity()
	}

	return categories, nil
}

func (r *PostgresCategoryRepo) Update(ctx context.Context, id uint, category *entities.Category) (*entities.Category, error) {
	var categoryModel models.CategoryModel
	categoryModel.FromEntity(category)
	categoryModel.ID = id
	categoryModel.UpdatedAt = time.Now()

	result := r.db.WithContext(ctx).Save(&categoryModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("category not found")
	}

	return categoryModel.ToEntity(), nil
}

func (r *PostgresCategoryRepo) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.CategoryModel{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}

func (r *PostgresCategoryRepo) Search(ctx context.Context, query string, limit, offset int) ([]*entities.Category, error) {
	var categoryModels []models.CategoryModel

	dbQuery := r.db.WithContext(ctx).Where("name ILIKE ? OR slug ILIKE ?", "%"+query+"%", "%"+query+"%")

	if limit > 0 {
		dbQuery = dbQuery.Limit(limit)
	}
	if offset > 0 {
		dbQuery = dbQuery.Offset(offset)
	}

	if err := dbQuery.Find(&categoryModels).Error; err != nil {
		return nil, err
	}

	categories := make([]*entities.Category, len(categoryModels))
	for i, model := range categoryModels {
		categories[i] = model.ToEntity()
	}

	return categories, nil
}

func (r *PostgresCategoryRepo) GetWideCategories(ctx context.Context, limit int) ([]*entities.Category, error) {
	var categoryModels []models.CategoryModel

	query := r.db.WithContext(ctx).Order("name ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&categoryModels).Error; err != nil {
		return nil, err
	}

	categories := make([]*entities.Category, len(categoryModels))
	for i, model := range categoryModels {
		categories[i] = model.ToEntity()
	}

	return categories, nil
}

// Translation operations
func (r *PostgresCategoryRepo) CreateTranslation(ctx context.Context, translation *entities.CategoryTranslation) (*entities.CategoryTranslation, error) {
	var translationModel models.CategoryTranslationModel
	translationModel.FromEntity(translation)

	if err := r.db.WithContext(ctx).Create(&translationModel).Error; err != nil {
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

func (r *PostgresCategoryRepo) UpdateTranslation(ctx context.Context, translation *entities.CategoryTranslation) (*entities.CategoryTranslation, error) {
	var translationModel models.CategoryTranslationModel
	translationModel.FromEntity(translation)
	translationModel.UpdatedAt = time.Now()

	result := r.db.WithContext(ctx).Save(&translationModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("translation not found")
	}

	return translationModel.ToEntity(), nil
}

func (r *PostgresCategoryRepo) GetTranslations(ctx context.Context, categoryID uint) ([]*entities.CategoryTranslation, error) {
	var translationModels []models.CategoryTranslationModel

	if err := r.db.WithContext(ctx).
		Where("category_id = ?", categoryID).
		Find(&translationModels).Error; err != nil {
		return nil, err
	}

	translations := make([]*entities.CategoryTranslation, len(translationModels))
	for i, model := range translationModels {
		translations[i] = model.ToEntity()
	}

	return translations, nil
}

func (r *PostgresCategoryRepo) GetTranslationByLocale(ctx context.Context, categoryID uint, locale string) (*entities.CategoryTranslation, error) {
	var translationModel models.CategoryTranslationModel

	if err := r.db.WithContext(ctx).
		Where("category_id = ? AND locale = ?", categoryID, locale).
		First(&translationModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("translation not found")
		}
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

// Brand-Category relationship operations
func (r *PostgresCategoryRepo) CreateBrandRelation(ctx context.Context, relation *entities.BrandCategory) (*entities.BrandCategory, error) {
	var relationModel models.BrandCategoryModel
	relationModel.FromEntity(relation)

	if err := r.db.WithContext(ctx).Create(&relationModel).Error; err != nil {
		return nil, err
	}

	return relationModel.ToEntity(), nil
}

func (r *PostgresCategoryRepo) GetBrandRelations(ctx context.Context, categoryID uint) ([]*entities.BrandCategory, error) {
	var relationModels []models.BrandCategoryModel

	if err := r.db.WithContext(ctx).
		Where("category_id = ?", categoryID).
		Find(&relationModels).Error; err != nil {
		return nil, err
	}

	relations := make([]*entities.BrandCategory, len(relationModels))
	for i, model := range relationModels {
		relations[i] = model.ToEntity()
	}

	return relations, nil
}

func (r *PostgresCategoryRepo) GetCategoryBrandRelations(ctx context.Context) ([]*entities.BrandCategory, error) {
	var relationModels []models.BrandCategoryModel

	if err := r.db.WithContext(ctx).Find(&relationModels).Error; err != nil {
		return nil, err
	}

	relations := make([]*entities.BrandCategory, len(relationModels))
	for i, model := range relationModels {
		relations[i] = model.ToEntity()
	}

	return relations, nil
}

func (r *PostgresCategoryRepo) DeleteBrandRelation(ctx context.Context, categoryID, brandID uint) error {
	result := r.db.WithContext(ctx).
		Where("category_id = ? AND brand_id = ?", categoryID, brandID).
		Delete(&models.BrandCategoryModel{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("brand-category relation not found")
	}

	return nil
}

// Status operations (assuming we add a status field to Category)
func (r *PostgresCategoryRepo) UpdateStatus(ctx context.Context, id uint, status int) error {
	result := r.db.WithContext(ctx).
		Model(&models.CategoryModel{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}

func (r *PostgresCategoryRepo) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.CategoryModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Ensure the implementation satisfies the interface
var _ repository.CategoryRepository = (*PostgresCategoryRepo)(nil)
