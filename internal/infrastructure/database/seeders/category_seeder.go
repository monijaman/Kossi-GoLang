package seeders

import (
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// CategorySeeder handles seeding categories
type CategorySeeder struct {
	BaseSeeder
}

// NewCategorySeeder creates a new category seeder
func NewCategorySeeder() *CategorySeeder {
	return &CategorySeeder{
		BaseSeeder: BaseSeeder{name: "Categories"},
	}
}

// Seed implements the Seeder interface
func (cs *CategorySeeder) Seed(db *gorm.DB) error {
	categories := []struct {
		Name string
		Slug string
	}{
		{"Footwear", "footwear"},
		{"Clothing", "clothing"},
		{"Sports", "sports"},
		{"Automotive", "automotive"},
		{"Books", "books"},
		{"Electronics", "electronics"},
		{"Groceries", "groceries"},
		{"Mobile Phones", "mobile-phones"},
		{"Laptops", "laptops"},
		{"Office Supplies", "office-supplies"},
		{"Home Appliances", "home-appliances"},
		{"Cameras", "cameras"},
		{"Cosmetics", "cosmetics"},
		{"Perfumes", "perfumes"},
		{"Personal Care", "personal-care"},
		{"Toys", "toys"},
		{"Airline", "airline"},
		{"Watches", "watches"},
		{"Health & Beauty", "health-beauty"},
		{"Bags & Luggage", "bags-luggage"},
		{"Gaming", "gaming"},
		{"Fitness Equipment", "fitness-equipment"},
		{"Jewelry", "jewelry"},
		{"Furniture", "furniture"},
		{"Music Instruments", "music-instruments"},
		{"Glass", "glass"},
		{"Melamine", "melamine"},
		{"Printing", "printing"},
		{"Packaging", "packaging"},
		{"Light", "light"},
		{"Bulbs", "bulbs"},
		{"Bakery", "bakery"},
		{"Beer", "beer"},
	}

	for _, cat := range categories {
		category := &entities.Category{
			Name: cat.Name,
			Slug: cat.Slug,
		}

		var categoryModel models.CategoryModel
		// Check if category already exists
		if err := db.Where("name = ?", cat.Name).First(&categoryModel).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new category
				categoryModel.FromEntity(category)
				if err := db.Create(&categoryModel).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
