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

// CameraProductSeeder handles seeding camera products for Bangladesh
type CameraProductSeeder struct {
	BaseSeeder
}

// NewCameraProductSeeder creates a new CameraProductSeeder
func NewCameraProductSeeder() *CameraProductSeeder {
	return &CameraProductSeeder{BaseSeeder: BaseSeeder{name: "Camera Products"}}
}

// Seed implements the Seeder interface
func (cps *CameraProductSeeder) Seed(db *gorm.DB) error {
	// Ensure category exists
	category, err := CreateOrFindCategory(db, "Cameras", "cameras")
	if err != nil {
		return err
	}

	// Load dataset from JSON file for maintainability
	jsonPath := filepath.Join("init-db", "seeders", "cameras.json")
	raw, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read camera dataset %s: %w", jsonPath, err)
	}

	var data struct {
		CamerasInBangladesh struct {
			Brands []struct {
				Brand         string   `json:"brand"`
				Origin        string   `json:"origin"`
				Category      string   `json:"category"`
				Availability  string   `json:"availability"`
				PopularModels []string `json:"popular_models"`
			} `json:"brands"`
		} `json:"cameras_in_bangladesh"`
	}

	if err := json.Unmarshal(raw, &data); err != nil {
		return fmt.Errorf("failed to parse camera dataset: %w", err)
	}

	// Process each brand and its models
	for _, brandData := range data.CamerasInBangladesh.Brands {
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
			if err := cps.createCameraProduct(db, brand.ID, category.ID, brandData.Brand, model); err != nil {
				return fmt.Errorf("failed to create product %s: %w", model, err)
			}
		}
	}

	return nil
}

// createCameraProduct creates a single camera product
func (cps *CameraProductSeeder) createCameraProduct(db *gorm.DB, brandID, categoryID uint, brandName, modelName string) error {
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

	// Check if product already exists
	var existing models.ProductModel
	if err := db.Where("slug = ?", slug).First(&existing).Error; err == nil {
		// Product already exists, skip
		return nil
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	// Create new product
	product := &models.ProductModel{
		Name:       fullName,
		Slug:       slug,
		BrandID:    &brandID,
		CategoryID: &categoryID,
		Status:     1, // Active
	}

	if err := db.Create(product).Error; err != nil {
		return err
	}

	fmt.Printf("   ✅ Created camera product: %s\n", fullName)
	return nil
}
