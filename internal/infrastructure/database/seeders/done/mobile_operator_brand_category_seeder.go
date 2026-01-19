package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MobileOperatorBrandCategorySeeder handles seeding mobile operator brand-category relationships
type MobileOperatorBrandCategorySeeder struct {
	BaseSeeder
}

// NewMobileOperatorBrandCategorySeeder creates a new mobile operator brand category seeder
func NewMobileOperatorBrandCategorySeeder() *MobileOperatorBrandCategorySeeder {
	return &MobileOperatorBrandCategorySeeder{
		BaseSeeder: BaseSeeder{name: "Mobile Operator Brand Categories"},
	}
}

// Seed implements the Seeder interface
func (mobcs *MobileOperatorBrandCategorySeeder) Seed(db *gorm.DB) error {
	// Mobile operator brand slugs
	mobileOperatorSlugs := []string{
		"banglalink",
		"grameenphone",
		"robi",
		"teletalk",
		"airtel",
	}

	// Find the Mobile Operator category (ID: 80)
	var mobileCategory models.CategoryModel
	if err := db.Where("id = ?", 80).First(&mobileCategory).Error; err != nil {
		// If not found, try by slug
		if err := db.Where("slug = ?", "mobile-operators").First(&mobileCategory).Error; err != nil {
			return nil // Skip if category not found
		}
	}

	for _, slug := range mobileOperatorSlugs {
		// Find the brand by slug
		var brand models.BrandModel
		if err := db.Where("slug = ?", slug).First(&brand).Error; err != nil {
			continue // Skip if brand not found
		}

		// Check if relationship already exists
		var existingRelation models.BrandCategoryModel
		result := db.Where("brand_id = ? AND category_id = ?", brand.ID, mobileCategory.ID).First(&existingRelation)

		if result.Error == gorm.ErrRecordNotFound {
			// Create the relationship
			brandCategory := &models.BrandCategoryModel{
				BrandID:    brand.ID,
				CategoryID: mobileCategory.ID,
			}

			if err := db.Create(brandCategory).Error; err != nil {
				continue
			}
		}
	}

	return nil
}
