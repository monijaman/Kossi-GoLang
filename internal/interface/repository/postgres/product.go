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
	if err := r.db.WithContext(ctx).Preload("Category").Preload("Brand").First(&productModel, id).Error; err != nil {
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
	if err := r.db.WithContext(ctx).Preload("Category").Preload("Brand").Where("slug = ?", slug).First(&productModel).Error; err != nil {
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
	query := r.db.WithContext(ctx).Where("deleted_at IS NULL AND status >= 1").Preload("Category").Preload("Brand").Order("priority ASC, id DESC")

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
	dbQuery := r.db.WithContext(ctx).Where("deleted_at IS NULL AND status >= 1").Preload("Category").Preload("Brand")

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
	query := r.db.WithContext(ctx).Preload("Category").Preload("Brand").Where("deleted_at IS NULL AND status >= 1").Order("views_count DESC")

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
	query := r.db.WithContext(ctx).Where("deleted_at IS NULL AND status >= 1 AND category_id = ?", categoryID).Preload("Category").Preload("Brand")

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
	query := r.db.WithContext(ctx).Where("deleted_at IS NULL AND status >= 1 AND brand_id = ?", brandID).Preload("Category").Preload("Brand")

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

func (r *PostgresProductRepo) GetSimilarProducts(ctx context.Context, product *entities.Product, limit int) ([]*entities.Product, error) {
	var productModels []models.ProductModel

	if product.CategoryID == nil {
		return nil, nil // No category to match against
	}

	// Determine representative price from start/end for similarity calculation
	var centerPrice float64
	if product.StartPrice != nil {
		centerPrice = *product.StartPrice
	} else if product.EndPrice != nil {
		centerPrice = *product.EndPrice
	} else {
		centerPrice = product.Price
	}
	// Calculate price range (+/- 10%)
	minPrice := centerPrice * 0.9
	maxPrice := centerPrice * 1.1

	query := r.db.WithContext(ctx).
		Preload("Category").Preload("Brand").
		Where("deleted_at IS NULL AND status >= 1").
		Where("category_id = ?", *product.CategoryID).
		Where("id != ?", product.ID).
		Where("COALESCE(start_price, end_price, 0) BETWEEN ? AND ?", minPrice, maxPrice)

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
	err := r.db.WithContext(ctx).Model(&models.ProductModel{}).Where("deleted_at IS NULL AND status >= 1").Count(&count).Error
	return count, err
}

func (r *PostgresProductRepo) CountByCategory(ctx context.Context, categoryID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ProductModel{}).
		Where("deleted_at IS NULL AND status >= 1 AND category_id = ?", categoryID).Count(&count).Error
	return count, err
}

func (r *PostgresProductRepo) CountByBrand(ctx context.Context, brandID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ProductModel{}).
		Where("deleted_at IS NULL AND status >= 1 AND brand_id = ?", brandID).Count(&count).Error
	return count, err
}

// Translation methods
func (r *PostgresProductRepo) CreateTranslation(ctx context.Context, translation *entities.ProductTranslation) (*entities.ProductTranslation, error) {
	var translationModel models.ProductTranslationModel
	translationModel.FromEntity(translation)

	// Debug: Log what's being sent to database
	fmt.Printf("Repository CreateTranslation - Model before DB: ProductID=%d, Locale='%s', TranslatedName='%s', StartPrice=%v, EndPrice=%v\n",
		translationModel.ProductID, translationModel.Locale, translationModel.TranslatedName, translationModel.StartPrice, translationModel.EndPrice)

	// Validate that the translated_name field is not empty before database call
	if translationModel.TranslatedName == "" {
		fmt.Printf("ERROR: Model.TranslatedName is empty before database call!\n")
		return nil, fmt.Errorf("translated_name field cannot be empty")
	}

	// Create using raw SQL to ensure proper handling of nullable columns
	if err := r.db.WithContext(ctx).Create(&translationModel).Error; err != nil {
		fmt.Printf("ERROR creating translation: %v\n", err)
		return nil, err
	}

	fmt.Printf("Translation created successfully in database with ID: %d, StartPrice=%v, EndPrice=%v\n",
		translationModel.ID, translationModel.StartPrice, translationModel.EndPrice)
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

	fmt.Printf("Repository UpdateTranslation - Updating ID=%d with StartPrice=%v, EndPrice=%v\n",
		translationModel.ID, translationModel.StartPrice, translationModel.EndPrice)

	// Use raw SQL to ensure nullable columns are properly updated (including NULLs)
	// This bypasses GORM's default behavior of skipping nil/zero values
	sql := `UPDATE product_translations 
	        SET translated_name = ?, 
	            start_price = ?, 
	            end_price = ?, 
	            updated_at = ? 
	        WHERE id = ?`

	if err := r.db.WithContext(ctx).Exec(sql,
		translationModel.TranslatedName,
		translationModel.StartPrice,
		translationModel.EndPrice,
		translationModel.UpdatedAt,
		translationModel.ID,
	).Error; err != nil {
		fmt.Printf("ERROR updating translation via raw SQL: %v\n", err)
		return nil, err
	}

	// Reload to get the latest data from database
	if err := r.db.WithContext(ctx).Where("id = ?", translationModel.ID).First(&translationModel).Error; err != nil {
		fmt.Printf("ERROR reloading translation after update: %v\n", err)
		return nil, err
	}

	fmt.Printf("Translation updated successfully - StartPrice=%v, EndPrice=%v\n",
		translationModel.StartPrice, translationModel.EndPrice)
	return translationModel.ToEntity(), nil
}

// GetWithFilters implements Laravel-compatible filtering for products
func (r *PostgresProductRepo) GetWithFilters(ctx context.Context, filters *repository.ProductFilters) ([]*entities.Product, int64, error) {
	var productModels []models.ProductModel
	var totalCount int64

	// Start building the query (only include non-deleted, active products by default)
	query := r.db.WithContext(ctx).Model(&models.ProductModel{}).Preload("Category").Preload("Brand").Where("deleted_at IS NULL AND status >= 1")

	// Apply filters
	query = r.applyFilters(query, filters)

	// Get total count for pagination
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	query = r.applySorting(query, filters.SortBy)

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

// applyFilters applies all filters to the query and returns the modified query
// IMPORTANT: GORM's Where/Joins methods return a new *gorm.DB, so we must reassign
func (r *PostgresProductRepo) applyFilters(query *gorm.DB, filters *repository.ProductFilters) *gorm.DB {
	// Search term filter
	if filters.SearchTerm != "" {
		query = query.Where("products.name ILIKE ?", "%"+filters.SearchTerm+"%")
	}

	// Category filter - supports both ID (numeric) and slug (alphanumeric)
	if filters.CategorySlug != "" {
		categoryValue := filters.CategorySlug
		if isNumeric(categoryValue) {
			// It's an ID - search by ID directly
			query = query.Where("products.category_id = ?", categoryValue)
		} else {
			// It's a slug - search by slug
			query = query.Joins("LEFT JOIN categories ON products.category_id = categories.id").
				Where("categories.slug = ?", categoryValue)
		}
	}

	// Brand filter (multiple brands supported)
	if len(filters.BrandSlugs) > 0 {
		query = query.Joins("JOIN brands ON products.brand_id = brands.id").
			Where("brands.slug IN ?", filters.BrandSlugs)
	}

	// Price range filter - use COALESCE of start/end price for comparisons
	if filters.MinPrice != nil {
		query = query.Where("COALESCE(products.start_price, products.end_price, 0) >= ?", *filters.MinPrice)
	}
	if filters.MaxPrice != nil {
		query = query.Where("COALESCE(products.start_price, products.end_price, 0) <= ?", *filters.MaxPrice)
	}

	return query
}

// isNumeric checks if a string represents a numeric value
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// applySorting applies sorting based on sortby parameter and returns the modified query
func (r *PostgresProductRepo) applySorting(query *gorm.DB, sortBy string) *gorm.DB {
	switch sortBy {
	case "popular":
		query = query.Order("priority ASC, views_count DESC")
	case "price_asc":
		query = query.Order("priority ASC, COALESCE(start_price, end_price) ASC")
	case "price_desc":
		query = query.Order("priority ASC, COALESCE(start_price, end_price) DESC")
	default:
		// Default sorting by priority
		query = query.Order("priority ASC")
	}
	return query
}

// DeleteTranslation deletes a product translation by its ID
func (r *PostgresProductRepo) DeleteTranslation(ctx context.Context, translationID uint) error {
	err := r.db.WithContext(ctx).Delete(&models.ProductTranslationModel{}, translationID).Error
	if err != nil {
		return fmt.Errorf("failed to delete product translation: %w", err)
	}
	return nil
}

var _ repository.ProductRepository = (*PostgresProductRepo)(nil)
