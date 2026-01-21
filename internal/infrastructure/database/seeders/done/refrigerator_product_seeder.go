package seeders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RefrigeratorProductSeeder handles seeding refrigerator products for Bangladesh
type RefrigeratorProductSeeder struct {
	BaseSeeder
}

// NewRefrigeratorProductSeeder creates a new RefrigeratorProductSeeder
func NewRefrigeratorProductSeeder() *RefrigeratorProductSeeder {
	return &RefrigeratorProductSeeder{BaseSeeder: BaseSeeder{name: "Refrigerator Products"}}
}

// Seed implements the Seeder interface
func (rps *RefrigeratorProductSeeder) Seed(db *gorm.DB) error {
	// Ensure category exists
	category, err := CreateOrFindCategory(db, "Refrigerators", "refrigerators")
	if err != nil {
		return err
	}

	// Load dataset from JSON file for maintainability
	jsonPath := filepath.Join("init-db", "seeders", "refrigerator.json")
	raw, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read refrigerator dataset %s: %w", jsonPath, err)
	}

	var brands []struct {
		Brand  string   `json:"brand"`
		Origin string   `json:"origin"`
		Models []string `json:"models"`
	}

	if err := json.Unmarshal(raw, &brands); err != nil {
		return fmt.Errorf("failed to parse refrigerator dataset: %w", err)
	}

	// Process each brand and its models
	for _, brandData := range brands {
		// Ensure brand exists
		brand, err := CreateOrFindBrand(db, brandData.Brand, strings.ToLower(brandData.Brand))
		if err != nil {
			return fmt.Errorf("failed to create/find brand %s: %w", brandData.Brand, err)
		}

		// Ensure brand-category relationship exists
		if err := CreateBrandCategoryRelation(db, brand.ID, category.ID); err != nil {
			return fmt.Errorf("failed to create brand-category relationship for %s: %w", brandData.Brand, err)
		}

		// Create products for each model
		for _, model := range brandData.Models {
			if err := rps.createRefrigeratorProduct(db, brand.ID, category.ID, model); err != nil {
				return fmt.Errorf("failed to create product %s: %w", model, err)
			}
		}
	}

	return nil
}

// createRefrigeratorProduct creates a single refrigerator product
func (rps *RefrigeratorProductSeeder) createRefrigeratorProduct(db *gorm.DB, brandID, categoryID uint, modelName string) error {
	// Generate slug from model name
	slug := strings.ToLower(strings.ReplaceAll(modelName, " ", "-"))
	slug = strings.ReplaceAll(slug, "(", "")
	slug = strings.ReplaceAll(slug, ")", "")
	slug = strings.ReplaceAll(slug, "/", "-")

	// Create or find product by name OR slug to avoid constraint violations
	product := &models.ProductModel{
		Name:       modelName,
		Slug:       slug,
		BrandID:    &brandID,
		CategoryID: &categoryID,
		Status:     1, // Active
	}

	// Check both name and slug to handle unique constraints on either field
	if err := db.Where("slug = ? OR name = ?", slug, modelName).FirstOrCreate(product).Error; err != nil {
		return err
	}

	if product.CreatedAt.Equal(product.UpdatedAt) {
		fmt.Printf("   ✅ Created refrigerator product: %s\n", modelName)
	} else {
		fmt.Printf("   ⚠️  Skipped existing product: %s\n", modelName)
	}
	return nil
}
