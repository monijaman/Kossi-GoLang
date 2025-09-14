package postgres

import (
	"context"
	"errors"
	"fmt"
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

	// Debug: Log what's being sent to database
	fmt.Printf("Repository CreateTranslation - Model before DB: ProductID=%d, Locale='%s', TranslatedName='%s' (len=%d)\n",
		translationModel.ProductID, translationModel.Locale, translationModel.TranslatedName, len(translationModel.TranslatedName))

	// Validate that the translated_name field is not empty before database call
	if translationModel.TranslatedName == "" {
		fmt.Printf("ERROR: Model.TranslatedName is empty before database call!\n")
		return nil, fmt.Errorf("translated_name field cannot be empty")
	}

	if err := r.db.WithContext(ctx).Create(&translationModel).Error; err != nil {
		fmt.Printf("Database create error: %v\n", err)
		return nil, err
	}

	fmt.Printf("Translation created successfully in database with ID: %d\n", translationModel.ID)
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

func (r *PostgresProductRepo) UpdateTranslation(ctx context.Context, translation *entities.ProductTranslation) (*entities.ProductTranslation, error) {
	var translationModel models.ProductTranslationModel
	translationModel.FromEntity(translation)

	if err := r.db.WithContext(ctx).Save(&translationModel).Error; err != nil {
		return nil, err
	}

	return translationModel.ToEntity(), nil
}

// GetWithFilters implements Laravel-compatible filtering for products
func (r *PostgresProductRepo) GetWithFilters(ctx context.Context, filters *repository.ProductFilters) ([]*entities.Product, int64, error) {
	var productModels []models.ProductModel
	var totalCount int64

	// Start building the query
	query := r.db.WithContext(ctx).Model(&models.ProductModel{})

	// Apply filters
	r.applyFilters(query, filters)

	// Get total count for pagination
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	r.applySorting(query, filters.SortBy)

	// Apply pagination
	offset := (filters.Page - 1) * filters.Limit
	if err := query.Limit(filters.Limit).Offset(offset).Find(&productModels).Error; err != nil {
		return nil, 0, err
	}

	// Convert models to entities
	products := make([]*entities.Product, len(productModels))
	for i, model := range productModels {
		products[i] = model.ToEntity()
	}

	return products, totalCount, nil
}

// applyFilters applies all filters to the query
func (r *PostgresProductRepo) applyFilters(query *gorm.DB, filters *repository.ProductFilters) {
	// Search term filter
	if filters.SearchTerm != "" {
		query.Where("name ILIKE ?", "%"+filters.SearchTerm+"%")
	}

	// Category filter
	if filters.CategorySlug != "" {
		query.Joins("JOIN categories ON products.category_id = categories.id").
			Where("categories.slug = ?", filters.CategorySlug)
	}

	// Brand filter (multiple brands supported)
	if len(filters.BrandSlugs) > 0 {
		query.Joins("JOIN brands ON products.brand_id = brands.id").
			Where("brands.slug IN ?", filters.BrandSlugs)
	}

	// Price range filter
	if filters.MinPrice != nil {
		query.Where("price >= ?", *filters.MinPrice)
	}
	if filters.MaxPrice != nil {
		query.Where("price <= ?", *filters.MaxPrice)
	}
}

// applySorting applies sorting based on sortby parameter
func (r *PostgresProductRepo) applySorting(query *gorm.DB, sortBy string) {
	switch sortBy {
	case "popular":
		query.Order("priority ASC, views_count DESC")
	case "price_asc":
		query.Order("priority ASC, price ASC")
	case "price_desc":
		query.Order("priority ASC, price DESC")
	default:
		// Default sorting by priority
		query.Order("priority ASC")
	}
}

var _ repository.ProductRepository = (*PostgresProductRepo)(nil)
