// Package seeders provides database seeding functionality for the application.
// This follows Clean Architecture principles by keeping seeding logic in the infrastructure layer.
package seeders

import (
	"fmt"
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// Seeder interface defines the contract for all seeders
type Seeder interface {
	Seed(db *gorm.DB) error
	GetName() string
}

// SeederManager manages all database seeders
type SeederManager struct {
	db      *gorm.DB
	seeders []Seeder
}

// NewSeederManager creates a new seeder manager
func NewSeederManager(db *gorm.DB) *SeederManager {
	return &SeederManager{
		db:      db,
		seeders: make([]Seeder, 0),
	}
}

// AddSeeder adds a seeder to the manager
func (sm *SeederManager) AddSeeder(seeder Seeder) {
	sm.seeders = append(sm.seeders, seeder)
}

// RunAll runs all registered seeders
func (sm *SeederManager) RunAll() error {
	fmt.Println("🌱 Starting database seeding...")

	for _, seeder := range sm.seeders {
		fmt.Printf("   🔄 Running %s seeder...\n", seeder.GetName())

		if err := seeder.Seed(sm.db); err != nil {
			return fmt.Errorf("failed to run %s seeder: %w", seeder.GetName(), err)
		}

		fmt.Printf("   ✅ %s seeder completed successfully\n", seeder.GetName())
	}

	fmt.Println("🎉 All seeders completed successfully!")
	return nil
}

// RunSpecific runs a specific seeder by name
func (sm *SeederManager) RunSpecific(name string) error {
	for _, seeder := range sm.seeders {
		if seeder.GetName() == name {
			fmt.Printf("🔄 Running %s seeder...\n", name)

			if err := seeder.Seed(sm.db); err != nil {
				return fmt.Errorf("failed to run %s seeder: %w", name, err)
			}

			fmt.Printf("✅ %s seeder completed successfully\n", name)
			return nil
		}
	}

	return fmt.Errorf("seeder '%s' not found", name)
}

// BaseSeeder provides common functionality for all seeders
type BaseSeeder struct {
	name string
}

// GetName returns the seeder name
func (bs *BaseSeeder) GetName() string {
	return bs.name
}

// Helper function to create or find a category
func CreateOrFindCategory(db *gorm.DB, name, slug string) (*entities.Category, error) {
	var categoryModel models.CategoryModel

	// Try to find existing category
	if err := db.Where("name = ?", name).First(&categoryModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new category
			category := &entities.Category{
				Name: name,
				Slug: slug,
			}

			categoryModel.FromEntity(category)
			if err := db.Create(&categoryModel).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return categoryModel.ToEntity(), nil
}

// Helper function to create or find a brand
func CreateOrFindBrand(db *gorm.DB, name, slug string) (*entities.Brand, error) {
	var brandModel models.BrandModel

	// Try to find existing brand
	if err := db.Where("name = ?", name).First(&brandModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new brand
			brand := &entities.Brand{
				Name: name,
				Slug: slug,
			}

			brandModel.FromEntity(brand)
			if err := db.Create(&brandModel).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return brandModel.ToEntity(), nil
}

// Helper function to create brand-category relationship
func CreateBrandCategoryRelation(db *gorm.DB, brandID, categoryID uint) error {
	var brandCategoryModel models.BrandCategoryModel

	// Check if relation already exists
	if err := db.Where("brand_id = ? AND category_id = ?", brandID, categoryID).First(&brandCategoryModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new relation
			relation := &entities.BrandCategory{
				BrandID:    brandID,
				CategoryID: categoryID,
			}

			brandCategoryModel.FromEntity(relation)
			if err := db.Create(&brandCategoryModel).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

// Helper function to generate slug from name
func GenerateSlug(name string) string {
	// Simple slug generation - replace spaces with hyphens and convert to lowercase
	slug := ""
	for _, char := range name {
		if char == ' ' || char == '&' {
			slug += "-"
		} else if char >= 'A' && char <= 'Z' {
			slug += string(char + 32) // Convert to lowercase
		} else if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			slug += string(char)
		}
	}
	return slug
}

// Helper function to create or find a specification key
func CreateOrFindSpecificationKey(db *gorm.DB, key string) (*entities.SpecificationKey, error) {
	var specKeyModel models.SpecificationKeyModel

	// Try to find existing specification key
	if err := db.Where("specification_key = ?", key).First(&specKeyModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new specification key
			specKey := &entities.SpecificationKey{
				SpecificationKey: key,
			}

			specKeyModel.FromEntity(specKey)
			if err := db.Create(&specKeyModel).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return specKeyModel.ToEntity(), nil
}

// SetupAllSeeders configures all available seeders
func SetupAllSeeders(db *gorm.DB) *SeederManager {
	manager := NewSeederManager(db)
	// Register a minimal, safe set of seeders that are present in the repo.
	// This avoids referencing constructors that may not exist yet while we
	// iteratively add more motorcycle-specific seeders.

	// Core/global seeders
	// manager.AddSeeder(NewCategorySeeder())
	// manager.AddSeeder(NewBrandSeeder())
	// manager.AddSeeder(NewUserSeeder())
	manager.AddSeeder(NewSpecificationKeySeeder())
	// manager.AddSeeder(NewCategoryTranslationSeeder())
	// manager.AddSeeder(NewBrandTranslationSeeder())
	manager.AddSeeder(NewSpecificationKeyTranslationSeeder())
	// manager.AddSeeder(NewBrandCategorySeeder())
	// manager.AddSeeder(NewProductSeeder())

	// Banking category seeders (Category ID: 10)
	manager.AddSeeder(NewBankBrandSeeder())
	manager.AddSeeder(NewBrandTranslationSeeder())
	manager.AddSeeder(NewBrandCategorySeeder())
	manager.AddSeeder(NewBankProductSeeder())
	manager.AddSeeder(NewBankSpecificationSeeder())
	manager.AddSeeder(NewFormGeneratorSeeder())

	// Mobile Operator category seeders
	manager.AddSeeder(NewMobileOperatorBrandSeeder())
	manager.AddSeeder(NewMobileOperatorBrandCategorySeeder())
	manager.AddSeeder(NewMobileOperatorProductSeeder())
	manager.AddSeeder(NewMobileOperatorSpecificationSeeder())
	manager.AddSeeder(NewMobileOperatorSpecificationKeyTranslationSeeder())
	manager.AddSeeder(NewMobileOperatorFormGeneratorSeeder())

	// Banglalink reviews seeders
	manager.AddSeeder(NewBanglalinkReviewSeeder())
	manager.AddSeeder(NewBanglalinkReviewTranslationSeeder())

	// Form generator (exists)
	// manager.AddSeeder(NewFormGeneratorSeedMotorcycle81())

	// Motorcycle-specific seeders currently implemented in this workspace
	// (only register seeders that have corresponding files/constructors)
	// Bajaj example template

	return manager
}
