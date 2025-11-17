package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVBrandCategorySeeder handles seeding TV brand-category relationships
type TVBrandCategorySeeder struct {
	BaseSeeder
}

// NewTVBrandCategorySeeder creates a new TV brand category seeder
func NewTVBrandCategorySeeder() *TVBrandCategorySeeder {
	return &TVBrandCategorySeeder{
		BaseSeeder: BaseSeeder{name: "TV Brand Categories"},
	}
}

// Seed implements the Seeder interface
func (tvbcs *TVBrandCategorySeeder) Seed(db *gorm.DB) error {
	// TV brand slugs (from the cleaned list)
	tvBrandSlugs := []string{
		"samsung",
		"lg",
		"sony",
		"xiaomi",
		"hisense",
		"haier",
		"tcl",
		"panasonic",
		"singer",
		"walton",
		"vision",
		"minister",
		"rangs",
		"myOne",
		"toshiba",
		"sharp",
		"konka",
		"eco+",
		"haiko",
	}

	// Find the TV category (ID: 124)
	var tvCategory models.CategoryModel
	if err := db.Where("id = ?", 124).First(&tvCategory).Error; err != nil {
		// If not found, try by slug
		if err := db.Where("slug = ?", "tv").First(&tvCategory).Error; err != nil {
			return nil // Skip if category not found
		}
	}

	for _, slug := range tvBrandSlugs {
		// Find or create the brand by slug
		var brand models.BrandModel
		result := db.Where("slug = ?", slug).First(&brand)

		if result.Error == gorm.ErrRecordNotFound {
			// Create brand if it doesn't exist
			// Map slug to display name
			brandNameMap := map[string]string{
				"samsung":   "Samsung",
				"lg":        "LG",
				"sony":      "Sony",
				"xiaomi":    "Xiaomi",
				"hisense":   "Hisense",
				"haier":     "Haier",
				"tcl":       "TCL",
				"panasonic": "Panasonic",
				"singer":    "Singer",
				"walton":    "Walton",
				"vision":    "Vision",
				"minister":  "Minister",
				"rangs":     "Rangs",
				"myOne":     "MyOne",
				"toshiba":   "Toshiba",
				"sharp":     "Sharp",
				"konka":     "Konka",
				"eco+":      "ECO+",
				"haiko":     "Haiko",
			}

			brandName := brandNameMap[slug]
			if brandName == "" {
				brandName = slug
			}

			brand = models.BrandModel{
				Name:   brandName,
				Slug:   slug,
				Status: 1,
			}
			if err := db.Create(&brand).Error; err != nil {
				continue
			}
		}

		// Check if relationship already exists
		var existingRelation models.BrandCategoryModel
		result = db.Where("brand_id = ? AND category_id = ?", brand.ID, tvCategory.ID).First(&existingRelation)

		if result.Error == gorm.ErrRecordNotFound {
			// Create the relationship
			brandCategory := &models.BrandCategoryModel{
				BrandID:    brand.ID,
				CategoryID: tvCategory.ID,
			}

			if err := db.Create(brandCategory).Error; err != nil {
				continue
			}
		}
	}

	return nil
}
