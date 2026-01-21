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

// LaptopProductSeeder handles seeding laptop products for Bangladesh
type LaptopProductSeeder struct {
	BaseSeeder
}

// NewLaptopProductSeeder creates a new LaptopProductSeeder
func NewLaptopProductSeeder() *LaptopProductSeeder {
	return &LaptopProductSeeder{BaseSeeder: BaseSeeder{name: "Laptop Products"}}
}

// Seed implements the Seeder interface
func (lps *LaptopProductSeeder) Seed(db *gorm.DB) error {
	// Ensure category exists
	category, err := CreateOrFindCategory(db, "Laptops", "laptops")
	if err != nil {
		return err
	}

	// Load dataset from JSON file for maintainability
	jsonPath := filepath.Join("init-db", "seeders", "laptops.json")
	raw, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read laptop dataset %s: %w", jsonPath, err)
	}

	var data struct {
		LaptopsInBangladesh struct {
			Brands []struct {
				Brand         string   `json:"brand"`
				Origin        string   `json:"origin"`
				Category      string   `json:"category"`
				Availability  string   `json:"availability"`
				PopularModels []string `json:"popular_models"`
			} `json:"brands"`
		} `json:"laptops_in_bangladesh"`
	}

	if err := json.Unmarshal(raw, &data); err != nil {
		return fmt.Errorf("failed to parse laptop dataset: %w", err)
	}

	// Process each brand and its models
	for _, brandData := range data.LaptopsInBangladesh.Brands {
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
		for _, model := range brandData.PopularModels {
			if err := lps.createLaptopProduct(db, brand.ID, category.ID, brandData.Brand, model); err != nil {
				return fmt.Errorf("failed to create product %s: %w", model, err)
			}
		}
	}

	return nil
}

// createLaptopProduct creates a single laptop product
func (lps *LaptopProductSeeder) createLaptopProduct(db *gorm.DB, brandID, categoryID uint, brandName, modelName string) error {
	// Generate full product name with brand
	fullName := fmt.Sprintf("%s %s", brandName, modelName)

	// Generate slug from full name
	slug := strings.ToLower(strings.ReplaceAll(fullName, " ", "-"))
	slug = strings.ReplaceAll(slug, "(", "")
	slug = strings.ReplaceAll(slug, ")", "")
	slug = strings.ReplaceAll(slug, "/", "-")
	slug = strings.ReplaceAll(slug, "+", "-plus")
	slug = strings.ReplaceAll(slug, ".", "-")
	slug = strings.ReplaceAll(slug, ",", "")

	// Remove any double dashes
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Trim any trailing dashes
	slug = strings.Trim(slug, "-")

	// Create or find product by name OR slug to avoid constraint violations
	product := &models.ProductModel{
		Name:       fullName,
		Slug:       slug,
		BrandID:    &brandID,
		CategoryID: &categoryID,
		Status:     1, // Active
	}

	// Check both name and slug to handle unique constraints on either field
	if err := db.Where("slug = ? OR name = ?", slug, fullName).FirstOrCreate(product).Error; err != nil {
		return err
	}

	if product.CreatedAt.Equal(product.UpdatedAt) {
		fmt.Printf("   ✅ Created laptop product: %s\n", fullName)
	} else {
		fmt.Printf("   ⚠️  Skipped existing product: %s\n", fullName)
	}
	return nil
}
