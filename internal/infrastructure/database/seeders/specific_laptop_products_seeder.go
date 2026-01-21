package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"strings"

	"gorm.io/gorm"
)

// SpecificLaptopProductsSeeder creates specific laptop products that have specification seeders
type SpecificLaptopProductsSeeder struct {
	BaseSeeder
}

// NewSpecificLaptopProductsSeeder creates a new specific laptop products seeder
func NewSpecificLaptopProductsSeeder() *SpecificLaptopProductsSeeder {
	return &SpecificLaptopProductsSeeder{BaseSeeder: BaseSeeder{name: "Specific Laptop Products"}}
}

// Seed implements the Seeder interface
func (slps *SpecificLaptopProductsSeeder) Seed(db *gorm.DB) error {
	// Ensure category exists
	category, err := CreateOrFindCategory(db, "Laptops", "laptops")
	if err != nil {
		return err
	}

	// Define products with their brands
	products := []struct {
		BrandName string
		ModelName string
		Slug      string
	}{
		// GPD products
		{"GPD", "P2 Max", "gpd-p2-max"},
		{"GPD", "Pocket 3", "gpd-pocket-3"},
		{"GPD", "Win 4", "gpd-win-4"},
		{"GPD", "Win Max 2", "gpd-win-max-2"},

		// iBall products
		{"iBall", "CompBook Premio v2.0", "iball-compbook-premio-v2-0"},

		// MSI products
		{"MSI", "Crosshair 15", "msi-crosshair-15"},
		{"MSI", "Cyborg 15", "msi-cyborg-15"},
		{"MSI", "Katana 15", "msi-katana-15"},
		{"MSI", "Katana 17", "msi-katana-17"},
		{"MSI", "Pulse 15", "msi-pulse-15"},
		{"MSI", "Pulse 17", "msi-pulse-17"},
		{"MSI", "Raider GE68 HX", "msi-raider-ge68-hx"},
		{"MSI", "Raider GE78 HX", "msi-raider-ge78-hx"},
		{"MSI", "Stealth 14 Studio", "msi-stealth-14-studio"},
		{"MSI", "Stealth 16 Studio", "msi-stealth-16-studio"},
		{"MSI", "Titan GT77 HX", "msi-titan-gt77-hx"},
		{"MSI", "Vector GP78", "msi-vector-gp78"},

		// One-Netbook products
		{"One-Netbook", "OneGx1 Pro", "onenetbook-onegx1-pro"},
		{"One-Netbook", "OneMix 4", "onenetbook-onemix-4"},
		{"One-Netbook", "OneXPlayer 2", "onenetbook-onexplayer-2"},
	}

	for _, p := range products {
		// Ensure brand exists
		brand, err := CreateOrFindBrand(db, p.BrandName, strings.ToLower(p.BrandName))
		if err != nil {
			return fmt.Errorf("failed to create/find brand %s: %w", p.BrandName, err)
		}

		// Ensure brand-category relationship exists
		if err := CreateBrandCategoryRelation(db, brand.ID, category.ID); err != nil {
			return fmt.Errorf("failed to create brand-category relationship for %s: %w", p.BrandName, err)
		}

		// Create full product name
		fullName := fmt.Sprintf("%s %s", p.BrandName, p.ModelName)

		// Create or find product
		product := &models.ProductModel{
			Name:       fullName,
			Slug:       p.Slug,
			BrandID:    &brand.ID,
			CategoryID: &category.ID,
			Status:     1, // Active
		}

		if err := db.Where("slug = ? OR name = ?", p.Slug, fullName).FirstOrCreate(product).Error; err != nil {
			return fmt.Errorf("failed to create product %s: %w", fullName, err)
		}

		if product.CreatedAt.Equal(product.UpdatedAt) {
			fmt.Printf("   ✅ Created laptop product: %s\n", fullName)
		} else {
			fmt.Printf("   ⚠️  Skipped existing product: %s\n", fullName)
		}
	}

	return nil
}
