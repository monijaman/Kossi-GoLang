package postgres

import (
	"context"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/infrastructure/database/models"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type PostgresProductRepo struct {
	db *gorm.DB
}

func NewPostgresProductRepo(db *gorm.DB) *PostgresProductRepo {
	return &PostgresProductRepo{db: db}
}

// generateSlug creates a URL-friendly slug from a name
func generateSlug(name string) string {
	// Convert to lowercase
	slug := strings.ToLower(name)
	// Replace spaces and special characters with hyphens
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")
	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")
	return slug
}

func (r *PostgresProductRepo) GetByID(ctx context.Context, id uint) (*entities.Product, error) {
	var productModel models.ProductModel
	if err := r.db.WithContext(ctx).First(&productModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return productModel.ToEntity(), nil
}

func (r *PostgresProductRepo) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	// Convert entity to model
	var productModel models.ProductModel
	productModel.FromEntity(product)

	// Generate slug if not provided
	if productModel.Slug == "" {
		productModel.Slug = generateSlug(productModel.Name)
	}

	// Set default category ID if not provided
	if productModel.CategoryID == nil {
		defaultCategoryID := uint(1)
		productModel.CategoryID = &defaultCategoryID
	}

	if err := r.db.WithContext(ctx).Create(&productModel).Error; err != nil {
		return nil, err
	}

	return productModel.ToEntity(), nil
}

func (r *PostgresProductRepo) GetBySlug(ctx context.Context, slug string) (*entities.Product, error) {
	var productModel models.ProductModel
	if err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&productModel).Error; err != nil {
		return nil, err
	}
	return productModel.ToEntity(), nil
}

func (r *PostgresProductRepo) Update(ctx context.Context, id uint, product *entities.Product) (*entities.Product, error) {
	var productModel models.ProductModel

	// First, get the existing product
	if err := r.db.WithContext(ctx).First(&productModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Update fields from entity
	productModel.FromEntity(product)
	productModel.ID = id // Ensure ID is preserved

	// Generate new slug if name changed
	if productModel.Slug == "" {
		productModel.Slug = generateSlug(productModel.Name)
	}

	if err := r.db.WithContext(ctx).Save(&productModel).Error; err != nil {
		return nil, err
	}

	return productModel.ToEntity(), nil
}

func (r *PostgresProductRepo) List(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	var productModels []models.ProductModel
	query := r.db.WithContext(ctx).Where("deleted_at IS NULL")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&productModels).Error; err != nil {
		return nil, err
	}

	products := make([]*entities.Product, len(productModels))
	for i, model := range productModels {
		products[i] = model.ToEntity()
	}

	return products, nil
}

func (r *PostgresProductRepo) Search(ctx context.Context, query string, limit, offset int) ([]*entities.Product, error) {
	var productModels []models.ProductModel
	dbQuery := r.db.WithContext(ctx).Where("deleted_at IS NULL")

	if query != "" {
		searchTerm := "%" + query + "%"
		dbQuery = dbQuery.Where("name ILIKE ? OR description ILIKE ?", searchTerm, searchTerm)
	}

	if limit > 0 {
		dbQuery = dbQuery.Limit(limit)
	}
	if offset > 0 {
		dbQuery = dbQuery.Offset(offset)
	}

	if err := dbQuery.Find(&productModels).Error; err != nil {
		return nil, err
	}

	products := make([]*entities.Product, len(productModels))
	for i, model := range productModels {
		products[i] = model.ToEntity()
	}

	return products, nil
}

func (r *PostgresProductRepo) GetPopular(ctx context.Context, limit int) ([]*entities.Product, error) {
	var productModels []models.ProductModel
	query := r.db.WithContext(ctx).Where("deleted_at IS NULL").Order("views_count DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&productModels).Error; err != nil {
		return nil, err
	}

	products := make([]*entities.Product, len(productModels))
	for i, model := range productModels {
		products[i] = model.ToEntity()
	}

	return products, nil
}

func (r *PostgresProductRepo) GetByCategory(ctx context.Context, categoryID uint, limit, offset int) ([]*entities.Product, error) {
	var productModels []models.ProductModel
	query := r.db.WithContext(ctx).Where("deleted_at IS NULL AND category_id = ?", categoryID)

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&productModels).Error; err != nil {
		return nil, err
	}

	products := make([]*entities.Product, len(productModels))
	for i, model := range productModels {
		products[i] = model.ToEntity()
	}

	return products, nil
}

func (r *PostgresProductRepo) GetByBrand(ctx context.Context, brandID uint, limit, offset int) ([]*entities.Product, error) {
	var productModels []models.ProductModel
	query := r.db.WithContext(ctx).Where("deleted_at IS NULL AND brand_id = ?", brandID)

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&productModels).Error; err != nil {
		return nil, err
	}

	products := make([]*entities.Product, len(productModels))
	for i, model := range productModels {
		products[i] = model.ToEntity()
	}

	return products, nil
}

func (r *PostgresProductRepo) IncrementViews(ctx context.Context, id uint) error {
	// First check if the product exists
	var productModel models.ProductModel
	if err := r.db.WithContext(ctx).Select("id").First(&productModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	// If product exists, increment the views count
	return r.db.WithContext(ctx).Model(&models.ProductModel{}).Where("id = ?", id).
		UpdateColumn("views_count", gorm.Expr("views_count + 1")).Error
}

func (r *PostgresProductRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ProductModel{}).Where("deleted_at IS NULL").Count(&count).Error
	return count, err
}

func (r *PostgresProductRepo) CountByCategory(ctx context.Context, categoryID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ProductModel{}).
		Where("deleted_at IS NULL AND category_id = ?", categoryID).Count(&count).Error
	return count, err
}

func (r *PostgresProductRepo) CountByBrand(ctx context.Context, brandID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ProductModel{}).
		Where("deleted_at IS NULL AND brand_id = ?", brandID).Count(&count).Error
	return count, err
}

// Translation methods
func (r *PostgresProductRepo) CreateTranslation(ctx context.Context, translation *entities.ProductTranslation) (*entities.ProductTranslation, error) {
	var translationModel models.ProductTranslationModel
	translationModel.FromEntity(translation)

	if err := r.db.WithContext(ctx).Create(&translationModel).Error; err != nil {
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

func (r *PostgresProductRepo) GetTranslations(ctx context.Context, productID uint) ([]*entities.ProductTranslation, error) {
	var translationModels []models.ProductTranslationModel

	if err := r.db.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&translationModels).Error; err != nil {
		return nil, err
	}

	translations := make([]*entities.ProductTranslation, len(translationModels))
	for i, model := range translationModels {
		translations[i] = model.ToEntity()
	}

	return translations, nil
}

func (r *PostgresProductRepo) GetTranslationByLocale(ctx context.Context, productID uint, locale string) (*entities.ProductTranslation, error) {
	var translationModel models.ProductTranslationModel

	if err := r.db.WithContext(ctx).
		Where("product_id = ? AND locale = ?", productID, locale).
		First(&translationModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("translation not found")
		}
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

var _ repository.ProductRepository = (*PostgresProductRepo)(nil)
