package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"os"

	"gorm.io/gorm"
)

// CarSeeder handles seeding car brands and products
type CarSeeder struct {
	BaseSeeder
}

// NewCarSeeder creates a new car seeder
func NewCarSeeder() *CarSeeder {
	return &CarSeeder{
		BaseSeeder: BaseSeeder{name: "Car Seeder"},
	}
}

// Seed implements the Seeder interface
func (cs *CarSeeder) Seed(db *gorm.DB) error {
	// Read cars.json
	// Assuming running from project root
	byteValue, err := os.ReadFile("init-db/seeders/cars.json")
	if err != nil {
		return fmt.Errorf("failed to read cars.json: %w", err)
	}

	var carData map[string][]string
	if err := json.Unmarshal(byteValue, &carData); err != nil {
		return fmt.Errorf("failed to unmarshal cars.json: %w", err)
	}

	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	for brandName, modelsList := range carData {
		// 1. Find or Create Brand
		var brand models.BrandModel
		// Simple slug generation: lowercase and replace spaces with hyphens
		slug := generateCarSlug(brandName)
		
		result := db.Where("slug = ?", slug).First(&brand)
		if result.Error == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   brandName,
				Slug:   slug,
				Status: 1,
			}
			if err := db.Create(&brand).Error; err != nil {
				fmt.Printf("Error creating brand %s: %v\n", brandName, err)
				continue
			}
		}

		// 2. Create Brand Translation (bn)
		// Assuming name is same for bn for now, or we could use a map if needed.
		// For now, using the English name as placeholder if no translation available.
		var brandTranslation models.BrandTranslationModel
		if err := db.Where("brand_id = ? AND locale = ?", brand.ID, "bn").First(&brandTranslation).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				brandTranslation = models.BrandTranslationModel{
					BrandID: brand.ID,
					Locale:  "bn",
					Name:    brandName, // Using English name as fallback/default
				}
				db.Create(&brandTranslation)
			}
		}

		// 3. Create Brand Category Relation
		var brandCategory models.BrandCategoryModel
		if err := db.Where("brand_id = ? AND category_id = ?", brand.ID, carCategory.ID).First(&brandCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				brandCategory = models.BrandCategoryModel{
					BrandID:    brand.ID,
					CategoryID: carCategory.ID,
				}
				db.Create(&brandCategory)
			}
		}

		// 4. Create Products
		for _, modelName := range modelsList {
			productSlug := generateCarSlug(fmt.Sprintf("%s %s", brandName, modelName))
			
			var product models.ProductModel
			if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Create Product
					product = models.ProductModel{
						Name:       fmt.Sprintf("%s %s", brandName, modelName),
						Slug:       productSlug,
						CategoryID: &carCategory.ID,
						BrandID:    &brand.ID,
						Status:     1,
					}
					if err := db.Create(&product).Error; err != nil {
						fmt.Printf("Error creating product %s: %v\n", product.Name, err)
						continue
					}

					// Create Product Translation
					productTranslation := models.ProductTranslationModel{
						ProductID:      product.ID,
						Locale:         "bn",
						TranslatedName: product.Name, // Using English name as fallback
					}
					db.Create(&productTranslation)
				}
			}
		}
	}

	return nil
}

func generateCarSlug(s string) string {
	// A simple slug generator. In a real app, might want a robust library.
	// This replaces spaces with hyphens and lowercases.
	// Also removing special chars if any (basic implementation)
	var result []rune
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			result = append(result, r)
		} else if r >= 'A' && r <= 'Z' {
			result = append(result, r+32)
		} else if r == ' ' || r == '-' {
			result = append(result, '-')
		}
	}
	return string(result)
}
