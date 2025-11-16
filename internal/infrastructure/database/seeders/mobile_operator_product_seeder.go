package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MobileOperatorProductSeeder handles seeding mobile operator products
type MobileOperatorProductSeeder struct {
	BaseSeeder
}

// NewMobileOperatorProductSeeder creates a new mobile operator product seeder
func NewMobileOperatorProductSeeder() *MobileOperatorProductSeeder {
	return &MobileOperatorProductSeeder{
		BaseSeeder: BaseSeeder{name: "Mobile Operator Products"},
	}
}

// Seed implements the Seeder interface
func (mops *MobileOperatorProductSeeder) Seed(db *gorm.DB) error {
	mobileOperators := []struct {
		Name string
		Slug string
	}{
		{Name: "Banglalink", Slug: "banglalink"},
		{Name: "Grameenphone", Slug: "grameenphone"},
		{Name: "Robi", Slug: "robi"},
		{Name: "Teletalk", Slug: "teletalk"},
		{Name: "Airtel", Slug: "airtel"},
	}

	// Find Mobile Operator category (ID: 80)
	var mobileCategory models.CategoryModel
	if err := db.Where("id = ?", 80).First(&mobileCategory).Error; err != nil {
		// If not found by ID, try by slug
		if err := db.Where("slug = ?", "mobile-operators").First(&mobileCategory).Error; err != nil {
			return nil // Skip if category not found
		}
	}

	// Create products for each mobile operator
	for _, operator := range mobileOperators {
		// Find brand by slug
		var brand models.BrandModel
		if err := db.Where("slug = ?", operator.Slug).First(&brand).Error; err != nil {
			continue // Skip if brand not found
		}

		// Check if product already exists
		var existingProduct models.ProductModel
		result := db.Where("brand_id = ? AND slug = ?", brand.ID, operator.Slug).First(&existingProduct)

		if result.Error == gorm.ErrRecordNotFound {
			// Create product
			description := operator.Name + " - Mobile network operator services and information"
			product := models.ProductModel{
				Name:        operator.Name,
				Slug:        operator.Slug,
				Description: &description,
				BrandID:     &brand.ID,
				CategoryID:  &mobileCategory.ID,
				Status:      1,
			}

			if err := db.Create(&product).Error; err != nil {
				continue
			}
		}
	}

	return nil
}
