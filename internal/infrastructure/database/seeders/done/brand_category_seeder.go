package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BrandCategorySeeder handles seeding brand-category relationships
type BrandCategorySeeder struct {
	BaseSeeder
}

// NewBrandCategorySeeder creates a new brand category seeder
func NewBrandCategorySeeder() *BrandCategorySeeder {
	return &BrandCategorySeeder{
		BaseSeeder: BaseSeeder{name: "Brand Categories"},
	}
}

// Seed implements the Seeder interface
func (bcs *BrandCategorySeeder) Seed(db *gorm.DB) error {
	// Bangladeshi banks only (removed: Citibank, Commercial Bank of Ceylon, Habib Bank, HSBC Bank, National Bank of Pakistan, State Bank of India)
	brandSlugs := []string{
		"ab-bank",
		"agrani-bank",
		"al-arafah-islami-bank",
		"bangladesh-commerce-bank",
		"bangladesh-development-bank",
		"bangladesh-krishi-bank",
		"bank-alfalah",
		"bank-asia",
		"basic-bank",
		"bengal-bank",
		"brac-bank",
		"citizens-bank",
		"city-bank",
		"community-bank",
		"dhaka-bank",
		"dutch-bangla-bank",
		"eastern-bank",
		"exim-bank",
		"first-security-islami-bank",
		"global-islami-bank",
		"icb-islamic-bank",
		"ific-bank",
		"islami-bank",
		"jamuna-bank",
		"janata-bank",
		"meghna-bank",
		"mercantile-bank",
		"midland-bank",
		"modhumoti-bank",
		"mutual-trust-bank",
		"national-bank",
		"ncc-bank",
		"nrb-bank",
		"nrbc-bank",
		"one-bank",
		"padma-bank",
		"premier-bank",
		"prime-bank",
		"probashi-kallyan-bank",
		"pubali-bank",
		"rajshahi-krishi-unnayan-bank",
		"rupali-bank",
		"sbac-bank",
		"shahjalal-islami-bank",
		"shimanto-bank",
		"social-islami-bank",
		"sonali-bank",
		"southeast-bank",
		"standard-bank",
		"standard-chartered-bank",
		"trust-bank",
		"union-bank",
		"united-commercial-bank",
		"uttara-bank",
		"woori-bank",
	}

	// Find the Banking category (ID: 10)
	var bankingCategory models.CategoryModel
	if err := db.Where("id = ?", 10).First(&bankingCategory).Error; err != nil {
		// If not found, try by slug
		if err := db.Where("slug = ?", "banking").First(&bankingCategory).Error; err != nil {
			return err
		}
	}

	for _, slug := range brandSlugs {
		// Find the brand by slug
		var brand models.BrandModel
		if err := db.Where("slug = ?", slug).First(&brand).Error; err != nil {
			continue // Skip if brand not found
		}

		// Check if relationship already exists
		var existingRelation models.BrandCategoryModel
		result := db.Where("brand_id = ? AND category_id = ?", brand.ID, bankingCategory.ID).First(&existingRelation)

		if result.Error == gorm.ErrRecordNotFound {
			// Create the relationship
			brandCategory := &models.BrandCategoryModel{
				BrandID:    brand.ID,
				CategoryID: bankingCategory.ID,
			}

			if err := db.Create(brandCategory).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
