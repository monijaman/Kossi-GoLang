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
	// manager.AddSeeder(NewSpecificationKeySeeder())
	// // manager.AddSeeder(NewCategoryTranslationSeeder())
	// // manager.AddSeeder(NewBrandTranslationSeeder())
	// manager.AddSeeder(NewSpecificationKeyTranslationSeeder())
	// // manager.AddSeeder(NewBrandCategorySeeder())
	// // manager.AddSeeder(NewProductSeeder())

	// // Banking category seeders (Category ID: 10)
	// manager.AddSeeder(NewBankBrandSeeder())
	// manager.AddSeeder(NewBrandTranslationSeeder())
	// manager.AddSeeder(NewBrandCategorySeeder())
	// manager.AddSeeder(NewBankProductSeeder())
	// manager.AddSeeder(NewBankSpecificationSeeder())
	// manager.AddSeeder(NewFormGeneratorSeeder())

	// // Mobile Operator category seeders
	// manager.AddSeeder(NewMobileOperatorBrandSeeder())
	// manager.AddSeeder(NewMobileOperatorBrandCategorySeeder())
	// manager.AddSeeder(NewMobileOperatorProductSeeder())
	// manager.AddSeeder(NewMobileOperatorSpecificationSeeder())
	// manager.AddSeeder(NewMobileOperatorSpecificationKeyTranslationSeeder())
	// manager.AddSeeder(NewMobileOperatorFormGeneratorSeeder())

	// // Banglalink reviews seeders
	// manager.AddSeeder(NewBanglalinkReviewSeeder())
	// manager.AddSeeder(NewBanglalinkReviewTranslationSeeder())

	// TV category seeders (Category ID: 124)
	// manager.AddSeeder(NewTVSpecificationSeeder())
	// manager.AddSeeder(NewTVSpecificationKeyTranslationSeeder())
	// manager.AddSeeder(NewTVFormGeneratorSeeder())
	// manager.AddSeeder(NewTVBrandCategorySeeder())
	// manager.AddSeeder(NewTVProductSpecificationSeeder())
	// manager.AddSeeder(NewTVProductReviewSeeder())
	// manager.AddSeeder(NewTVDetailedSpecificationSeeder())
	// manager.AddSeeder(NewTVReviewTranslationSeeder())
	// manager.AddSeeder(NewTVComprehensiveReviewSeeder())

	// Form generator (exists)
	// manager.AddSeeder(NewFormGeneratorSeedMotorcycle81())

	// Motorcycle-specific seeders currently implemented in this workspace
	// (only register seeders that have corresponding files/constructors)
	// Bajaj example template

	// Car category seeders (Category ID: 18)
	// manager.AddSeeder(CarBrandTranslationSeeder())
	// manager.AddSeeder(NewBrandTranslationInclusionSeeder())

	// Register the Car Specification Seeder
	// manager.AddSeeder(NewCarSpecificationSeeder())
	// manager.AddSeeder(NewMissingSpecificationsSeeder())
	// manager.AddSeeder(NewMissingSpecificationsSeeder())

	// Mobile Specification Seeders (Category ID: 79)
	// manager.AddSeeder(NewSpecificationSeederMobileGooglePixel8())
	// manager.AddSeeder(NewSpecificationSeederMobileGooglePixel8Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileGooglePixel8a())
	// manager.AddSeeder(NewSpecificationSeederMobileGooglePixel9())
	// manager.AddSeeder(NewSpecificationSeederMobileGooglePixel9Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileGooglePixel9ProFold())
	// manager.AddSeeder(NewSpecificationSeederMobileGooglePixel9ProXl())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixHot50())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixHot60i4g())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixNote40Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixNote50())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixNote50Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixNote50Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixSmart10())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixSmart10Plus())
	// manager.AddSeeder(NewSpecificationSeederMobileInfinixZero305g())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone15())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone15Plus())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone15Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone15ProMax())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone16())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone16Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone16ProMax())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone17())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone17Plus())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone17Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileIphone17ProMax())
	// manager.AddSeeder(NewSpecificationSeederMobileIphoneSe3rdGen())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplus11())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplus12())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplus12r())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplus13())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplus13r())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplusNord4())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplusNord5())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplusNordCe4Lite5g())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplusNordCe5())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplusNordN20Se())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplusNordN30Se5g())
	// manager.AddSeeder(NewSpecificationSeederMobileOneplusOpen())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA12())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA18())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA38())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA3x())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA5())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA555g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA57())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA5Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA5x())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA60())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA745g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoA78())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoFindN4Flip())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoFindN4Fold())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno125g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno12F5g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno12Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno135g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno13F5g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno13Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno145g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno14F5g())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno15())
	// manager.AddSeeder(NewSpecificationSeederMobileOppoReno15Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileOrbitY71())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoC65())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoC71())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoC75())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoF65g())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoF7())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoF7Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoM5())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoM6Plus())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoM6Plus5g())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoM6Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoM7Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoX6Neo5g())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoX6Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoX75g())
	// manager.AddSeeder(NewSpecificationSeederMobilePocoX7Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobilePrimoF10())
	// manager.AddSeeder(NewSpecificationSeederMobilePrimoGh11())
	// manager.AddSeeder(NewSpecificationSeederMobilePrimoH10())
	// manager.AddSeeder(NewSpecificationSeederMobilePrimoR8())
	// manager.AddSeeder(NewSpecificationSeederMobilePrimoS8())
	// manager.AddSeeder(NewSpecificationSeederMobilePrimoZx4())
	// manager.AddSeeder(NewSpecificationSeederMobileRealme12())
	// manager.AddSeeder(NewSpecificationSeederMobileRealme13())
	// manager.AddSeeder(NewSpecificationSeederMobileRealme13Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileRealme145g())
	// manager.AddSeeder(NewSpecificationSeederMobileRealme14t5g())
	// manager.AddSeeder(NewSpecificationSeederMobileRealme155g())
	// manager.AddSeeder(NewSpecificationSeederMobileRealme15Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobileRealme15t5g())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeC61())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeC63())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeC65())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeC71())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeC75x())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeGtMasterEdition())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeGtNeo2())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeGtNeo3())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeNote60())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeNote60x())
	// manager.AddSeeder(NewSpecificationSeederMobileRealmeNote70())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmi12())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmi12c())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmi13c())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmi14c())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmi15())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmi15c())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiA2())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiA3())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiA5())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiNote12())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiNote124g())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiNote13())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiNote14())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiNote14Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiNote15())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiNote15Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileRedmiNote15Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA06())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA065g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA07())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA14())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA15())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA155g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA175g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA23())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA24())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA255g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA265g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA345g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA365g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA545g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyA565g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyF15())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyF165g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyF23())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyF345g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyF365g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyF545g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyF555g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyF565g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM065g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM145g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM155g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM165g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM335g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM345g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM365g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM535g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyM565g())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS22())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS22Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS23())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS23Fe())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS23Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS24())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS24Fe())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS24Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS25())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyS25Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyZFlip5())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyZFlip6())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyZFlip7())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyZFlip7Fe())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyZFold5())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyZFold6())
	// manager.AddSeeder(NewSpecificationSeederMobileSamsungGalaxyZFold7())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyAtom5())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyG27())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyHelio40())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyHelio50())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyHelio90())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyI10())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyInnova30())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyInnova40())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyMax60())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyS72())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyS75())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyXplorerV55())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyXplorerW20())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyXplorerW68())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyXplorerW92())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyXplorerZv())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyZ70())
	// manager.AddSeeder(NewSpecificationSeederMobileSymphonyZ72())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoCamon30())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoCamon30Premier5g())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoCamon30s())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoCamon40())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoCamon40Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoCamon40Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoPhantomVFold25g())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoPova7Pro5g())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoPovaCurve5g())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoPovaSlim5g())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSpark30())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSpark30Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSpark30c())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSpark40())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSpark405g())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSpark40Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSpark40c())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSparkGo1())
	// manager.AddSeeder(NewSpecificationSeederMobileTecnoSparkGo2())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoV40())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoV40Lite())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoV50())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoV50Lite())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoV50Lite5g())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoV605g())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoV60Lite4g())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoV60Lite5g())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoX200Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoX200ProMini())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoX300())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoX300Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoY03t())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoY04())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoY19s())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoY19sPro())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoY21d())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoY28())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoY29())
	// manager.AddSeeder(NewSpecificationSeederMobileVivoY400())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonN26())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonN27())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonN74())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonNexgN10Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonNexgN71Plus())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonNexgN75())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonNexgN9())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonOrbitY11())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonOrbitY12())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonOrbitY13())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonOrbitY70c())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonX91())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonXanonX1Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonXanonX21())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonZenx1())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonZenx1t())
	// manager.AddSeeder(NewSpecificationSeederMobileWaltonZenx2())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi13Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi13Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi13t())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi13tPro())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi14())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi14Pro())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi14Ultra())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi14tPro())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi15())
	// manager.AddSeeder(NewSpecificationSeederMobileXiaomi15Pro())
	manager.AddSeeder(NewMobileSpecificationsExistingKeysSeeder())
	return manager
}
