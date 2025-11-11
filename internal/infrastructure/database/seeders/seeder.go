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
	// manager.AddSeeder(NewCategoryTranslationSeeder())
	// manager.AddSeeder(NewBrandTranslationSeeder())
	// manager.AddSeeder(NewSpecificationKeyTranslationSeeder())
	// manager.AddSeeder(NewBrandCategorySeeder())
	// manager.AddSeeder(NewProductSeeder())

	// Form generator (exists)
	// manager.AddSeeder(NewFormGeneratorSeederMotorcycle81())

	// Motorcycle-specific seeders currently implemented in this workspace
	// (only register seeders that have corresponding files/constructors)
	// Bajaj example template

	// Product review seeders (16 total)
	manager.AddSeeder(NewHondaCD70ReviewSeeder())
	manager.AddSeeder(NewHondaLivoReviewSeeder())
	manager.AddSeeder(NewHondaCB125FReviewSeeder())
	manager.AddSeeder(NewHondaCBShineReviewSeeder())
	manager.AddSeeder(NewHondaDioReviewSeeder())
	manager.AddSeeder(NewYamahaSalutoRXReviewSeeder())
	manager.AddSeeder(NewYamahaFZSReviewSeeder())
	manager.AddSeeder(NewYamahaFazerReviewSeeder())
	manager.AddSeeder(NewYamahaYBR125ReviewSeeder())
	manager.AddSeeder(NewSuzukiGixxerReviewSeeder())
	manager.AddSeeder(NewSuzukiGixxerSFReviewSeeder())
	manager.AddSeeder(NewSuzukiHayateReviewSeeder())
	manager.AddSeeder(NewSuzukiAccess125ReviewSeeder())
	// 	manager.AddSeeder(NewBajajPulsar150ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewBajajPulsar125ReviewSeeder())
	manager.AddSeeder(NewBajajDiscoverReviewSeeder())
	// 	manager.AddSeeder(NewTVSApacheRTR160ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewTVSApacheRTR180ReviewSeeder())
	manager.AddSeeder(NewTVSMetroReviewSeeder())
	manager.AddSeeder(NewTVSJupiterReviewSeeder())
	// 	manager.AddSeeder(NewHeroSplendorPlusReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewHeroPassionProReviewSeeder())
	// 	manager.AddSeeder(NewHeroHFDeluxeReviewSeeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewRoyalEnfieldClassic350ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewRoyalEnfieldMeteor350ReviewSeeder())
	manager.AddSeeder(NewRunnerHunterReviewSeeder())
	manager.AddSeeder(NewRunnerBoxerReviewSeeder())
	manager.AddSeeder(NewWaltonPrimoReviewSeeder())
	manager.AddSeeder(NewHondaActivaReviewSeeder())
	manager.AddSeeder(NewYamahaFascinoReviewSeeder())
	manager.AddSeeder(NewSuzukiBurgmanStreetReviewSeeder())
	// 	manager.AddSeeder(NewKTMDuke200ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewKawasakiNinja300ReviewSeeder())
	manager.AddSeeder(NewHondaXBladeReviewSeeder())
	manager.AddSeeder(NewHondaSP125ReviewSeeder())
	// 	manager.AddSeeder(NewBajajCT100ReviewSeeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewBajajPlatina100ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewYamahaR15ReviewSeeder())
	// 	manager.AddSeeder(NewSuzukiGS150RReviewSeeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewRoadPrinceWegoReviewSeeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewRoadPrinceClassicReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewSYMSymphonySTReviewSeeder())
	manager.AddSeeder(NewBenelliTNT25ReviewSeeder())
	manager.AddSeeder(NewSuzukiGSXR150ABSReviewSeeder())
	manager.AddSeeder(NewSuzukiGixxerSF250SpecialEditionReviewSeeder())
	// 	manager.AddSeeder(NewSuzukiGixxerSF250ReviewSeeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewSuzukiGixxer250ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewSuzukiNewGixxerSFFiABSReviewSeeder())
	manager.AddSeeder(NewSuzukiGixxerSFFiDiscReviewSeeder())
	manager.AddSeeder(NewSuzukiIntruderFiABSReviewSeeder())
	manager.AddSeeder(NewNewSuzukiGixxerFiABSReviewSeeder())
	manager.AddSeeder(NewSuzukiGixxerFiDiscReviewSeeder())
	manager.AddSeeder(NewYamahaR15MBS7ReviewSeeder())
	manager.AddSeeder(NewYamahaR15V4BS7ReviewSeeder())
	manager.AddSeeder(NewYamahaR15MReviewSeeder())
	// 	manager.AddSeeder(NewYamahaR15V4ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewYamahaXSR155ReviewSeeder())
	manager.AddSeeder(NewYamahaMT15V2ReviewSeeder())
	manager.AddSeeder(NewYamahaAerox155ReviewSeeder())
	manager.AddSeeder(NewYamahaMT15ReviewSeeder())
	manager.AddSeeder(NewYamahaFazerFiV2ReviewSeeder())
	manager.AddSeeder(NewYamahaFzsFiDoubleDiscReviewSeeder())
	manager.AddSeeder(NewHondaCb150RExmotionReviewSeeder())
	manager.AddSeeder(NewHondaCbr150RReviewSeeder())
	manager.AddSeeder(NewHondaCb150XReviewSeeder())
	manager.AddSeeder(NewHondaPcx150ReviewSeeder())
	manager.AddSeeder(NewHondaCrf150LReviewSeeder())
	manager.AddSeeder(NewHondaAdv150ReviewSeeder())
	manager.AddSeeder(NewHondaSp160DoubleDiscReviewSeeder())
	manager.AddSeeder(NewHondaActiva125FiBs6ReviewSeeder())
	manager.AddSeeder(NewNewHondaXblade160ReviewSeeder())
	manager.AddSeeder(NewHondaSp125ReviewSeeder())
	manager.AddSeeder(NewBajajPulsarF250ReviewSeeder())
	manager.AddSeeder(NewBajajPulsarN250ReviewSeeder())
	manager.AddSeeder(NewBajajAvenger160AbsReviewSeeder())
	manager.AddSeeder(NewNewBajajPulsarN160ReviewSeeder())

	// Adding Batch 9 seeders
	manager.AddSeeder(NewBajajPulsar150TwinDiscReviewSeeder())
	manager.AddSeeder(NewBajajPulsar150SingleDiscABSReviewSeeder())
	// 	manager.AddSeeder(NewBajajPulsarNS125ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewBajajDiscover125DiscReviewSeeder())

	// Adding Batch 10 seeders
	manager.AddSeeder(NewBajajCT100ESReviewSeeder())
	manager.AddSeeder(NewBajajPlatinaCB100ReviewSeeder())
	manager.AddSeeder(NewBajajRX100ReviewSeeder())
	manager.AddSeeder(NewBajajDominar400ReviewSeeder())
	manager.AddSeeder(NewBajajPulsarAS200ReviewSeeder())

	// Adding Batch 11 seeders
	// 	manager.AddSeeder(NewBajajPulsarSingleDiscReviewSeeder())
	// 	manager.AddSeeder(NewBajajAvenger220ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewBajajDiscover150ReviewSeeder())
	manager.AddSeeder(NewBajajPulsar200TwinDiscReviewSeeder())
	manager.AddSeeder(NewBajajPlatina110ReviewSeeder())

	// Adding Batch 12 seeders - Multi-brand
	manager.AddSeeder(NewHondaCB110ReviewSeeder())
	manager.AddSeeder(NewYamahaYBR125FiReviewSeeder())
	// 	manager.AddSeeder(NewSuzukiGixxer155ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewHeroMaster125ReviewSeeder())
	// 	manager.AddSeeder(NewTVSApacheRTR200ReviewSeeder()) // Moved to non-batch or not in JSON

	// Adding Batch 13 seeders - Continued multi-brand
	manager.AddSeeder(NewSuzukiBurgman125FiReviewSeeder())
	// 	manager.AddSeeder(NewYamahaR15V3ReviewSeeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewHeroHFDeluxeIXReviewSeeder())

	// Adding Batch 14 seeders - Multi-brand diverse
	manager.AddSeeder(NewTVSJupiter125ReviewSeeder())
	manager.AddSeeder(NewBajajAvenger200DTSiReviewSeeder())
	manager.AddSeeder(NewSuzukiIntruder150ReviewSeeder())
	manager.AddSeeder(NewYamahaFZSFiDTSReviewSeeder())
	manager.AddSeeder(NewHondaCBTriggerReviewSeeder())

	// Adding Batch 15 seeders - Multi-brand diverse (10 seeders)
	manager.AddSeeder(NewTVSApacheRTR160TwoWheelerReviewSeeder())
	manager.AddSeeder(NewHeroSplendorPlusDTSiReviewSeeder())
	manager.AddSeeder(NewHeroMotocorpXPulseReviewSeeder())
	manager.AddSeeder(NewBajajPulsarN160ReviewSeeder())
	manager.AddSeeder(NewHondaCD110DreamReviewSeeder())
	// 	manager.AddSeeder(NewYamahaFZ25ReviewSeeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewSuzukiSmashReviewSeeder())
	manager.AddSeeder(NewBajajPulsarAS250ReviewSeeder())
	manager.AddSeeder(NewTVSApacheRTR160ReviewBatch15ReviewSeeder())
	manager.AddSeeder(NewSuzukiGixxerSF250Batch15Seeder())

	// Adding Batch 16 seeders - Premium & diverse (10 seeders)
	// 	manager.AddSeeder(NewKTMDuke390ReviewSeeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewBajajPulsarF150ReviewSeeder())
	manager.AddSeeder(NewHeroMotocorpExtaReviewSeeder())
	manager.AddSeeder(NewHondaCB350ReviewBatch16Seeder())
	manager.AddSeeder(NewYamahaXSR160ReviewBatch16Seeder())
	manager.AddSeeder(NewSuzukiGSXS1000ReviewBatch16Seeder())
	manager.AddSeeder(NewKawasakiNinja650ReviewBatch16Seeder())
	manager.AddSeeder(NewRoyalEnfieldClassic350ReviewBatch16Seeder())
	manager.AddSeeder(NewBenelliTNT600ReviewBatch16Seeder())

	// Adding Batch 17 seeders - Scooters & diverse mix (10 seeders)
	manager.AddSeeder(NewVespaLX125ReviewSeeder())
	manager.AddSeeder(NewKTMDuke200ReviewBatch17Seeder())

	// Adding Batch 18 seeders - Premium & comprehensive (10 seeders)
	manager.AddSeeder(NewHondaCRF450RallyReviewBatch18Seeder())
	manager.AddSeeder(NewYamahaXMax300ReviewBatch18Seeder())
	manager.AddSeeder(NewKawasakiNinja1000SXReviewBatch18Seeder())
	manager.AddSeeder(NewTVSApacheRTR180ReviewBatch18Seeder())
	manager.AddSeeder(NewSuzukiVStrom650ReviewBatch18Seeder())
	manager.AddSeeder(NewHeroHFDeluxeIXReviewBatch18Seeder())
	manager.AddSeeder(NewBajajPulsarNS200ReviewBatch18Seeder())
	manager.AddSeeder(NewYamahaFZ300ReviewBatch18Seeder())

	// Adding Batch 19 seeders - Diverse market coverage (10 seeders)
	manager.AddSeeder(NewHondaCB500FReviewBatch19Seeder())
	manager.AddSeeder(NewKawasakiNinja400ReviewBatch19Seeder())
	manager.AddSeeder(NewSuzukiGixxer250ReviewBatch19Seeder())
	manager.AddSeeder(NewBajajAvenger160ReviewBatch19Seeder())
	manager.AddSeeder(NewHondaActiva6GReviewBatch19Seeder())

	// Batch 21
	manager.AddSeeder(NewHeroSplendorPlusBatch21Seeder())
	manager.AddSeeder(NewBajajPulsar150Batch21Seeder())
	manager.AddSeeder(NewTVSApacheBatch21Seeder())
	// 	manager.AddSeeder(NewHondaCB200Batch21Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewSuzukiGSX150RBatch21Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewYamahaSZ150Batch21Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewKawasakiNinja250Batch21Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewUniversalMotorBike100Batch21Seeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewVespaLX300Batch21Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewPiggiBike125Batch21Seeder()) // Moved to non-batch or not in JSON

	// Batch 22
	// 	manager.AddSeeder(NewSymBike100Batch22Seeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewCebiBike125Batch22Seeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewBajajPlatinaBatch22Seeder())
	// 	manager.AddSeeder(NewHondaCG125Batch22Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewYamahaYZF150Batch22Seeder()) // Not in JSON
	manager.AddSeeder(NewSuzukiGixxer155Batch22Seeder())
	// 	manager.AddSeeder(NewHeroHF100Batch22Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewKawasakiNinja400Batch22Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewUniversalMotorbike75Batch22Seeder()) // Moved to non-batch or not in JSON

	// Batch 23
	manager.AddSeeder(NewTVSNtorq125Batch23Seeder())
	manager.AddSeeder(NewHondaActiva6Batch23Seeder())
	// 	manager.AddSeeder(NewBajajPulsar220FBatch23Seeder()) // Not in JSON
	manager.AddSeeder(NewRoyalEnfieldClassic350Batch23Seeder())
	// 	manager.AddSeeder(NewHeroHeroicBatch23Seeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewTVSApacheRTR200Batch23Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewSuzukiStorm125Batch23Seeder()) // Not in JSON
	manager.AddSeeder(NewYamahaR15V3Batch23Seeder())
	// 	manager.AddSeeder(NewBajajDominoBatch23Seeder()) // Moved to non-batch or not in JSON
	manager.AddSeeder(NewKTMDuke125Batch23Seeder())
	// 	manager.AddSeeder(NewUniversalMotorbike110Batch23Seeder()) // Moved to non-batch or not in JSON

	// Batch 24
	// 	manager.AddSeeder(NewBajajDualDo125Batch24Seeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewHoneyWellMaxBatch24Seeder()) // Moved to non-batch or not in JSON
	// 	manager.AddSeeder(NewVespaSprintBatch24Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewRoyalEnfieldSignetBatch24Seeder()) // Not in JSON
	manager.AddSeeder(NewKTMDuke200Batch24Seeder())
	// 	manager.AddSeeder(NewYamahaR3Batch24Seeder()) // Not in JSON
	manager.AddSeeder(NewSuzukiGixxer250Batch24Seeder())
	manager.AddSeeder(NewHeroPulseBatch24Seeder())
	// 	manager.AddSeeder(NewBajajAvenger220Batch24Seeder()) // Not in JSON

	// Batch 25
	// 	manager.AddSeeder(NewHondaX1XBatch25Seeder()) // Not in JSON
	manager.AddSeeder(NewBajajPlatina100Batch25Seeder())
	// 	manager.AddSeeder(NewTVSApache150Batch25Seeder()) // Not in JSON
	manager.AddSeeder(NewYamahaFZS150Batch25Seeder())
	manager.AddSeeder(NewHeroSplendorPlusBatch25Seeder())
	manager.AddSeeder(NewSuzukiGS150RBatch25Seeder())
	manager.AddSeeder(NewBajajCT100Batch25Seeder())
	// 	manager.AddSeeder(NewKawasakiNinja650Batch25Seeder()) // Not in JSON
	// 	manager.AddSeeder(NewTVSScootypBatch25Seeder()) // Not in JSON
	manager.AddSeeder(NewBajajPulsarNS125Batch25Seeder())

	// Batch 26
	manager.AddSeeder(NewHondaCG125Batch26Seeder())
	manager.AddSeeder(NewBajajPulsar150Batch26Seeder())
	manager.AddSeeder(NewTVSNtorq125Batch26Seeder())
	manager.AddSeeder(NewYamahaR15V4Batch26Seeder())
	manager.AddSeeder(NewKTMDuke390Batch26Seeder())
	manager.AddSeeder(NewHeroHFDeluxeBatch26Seeder())
	manager.AddSeeder(NewSuzukiGixxer155Batch26Seeder())
	manager.AddSeeder(NewRoyalEnfieldClassic350Batch26Seeder())
	manager.AddSeeder(NewHondaActiva6GBatch26Seeder())
	manager.AddSeeder(NewBajajDominarD400Batch26Seeder())

	// Batch 27
	manager.AddSeeder(NewTVSApacheRTR160Batch27Seeder())
	manager.AddSeeder(NewYamahaFZ25Batch27Seeder())

	// Batch 28 - Lifan Complete Series (14 seeders)
	manager.AddSeeder(NewLifanKPV150RaceEditionBatch28Seeder())
	manager.AddSeeder(NewLifanXPECT150V2Batch28Seeder())
	manager.AddSeeder(NewLifanK19Batch28Seeder())
	manager.AddSeeder(NewLifanKPV150Batch28Seeder())
	manager.AddSeeder(NewLifanKPT1504VBatch28Seeder())
	manager.AddSeeder(NewLifanKP1654VBatch28Seeder())
	manager.AddSeeder(NewLifanKPR165REFIBatch28Seeder())
	manager.AddSeeder(NewLifanKPR165RCarburetorBatch28Seeder())
	manager.AddSeeder(NewLifanKPR150Batch28Seeder())
	manager.AddSeeder(NewLifanKP165Batch28Seeder())
	manager.AddSeeder(NewLifan150T13Batch28Seeder())
	manager.AddSeeder(NewLifanBlink125Batch28Seeder())
	manager.AddSeeder(NewLifanRazor100Batch28Seeder())
	manager.AddSeeder(NewLifanGlint100Batch28Seeder())

	// Batch 29 - Revoo Complete Series (11 seeders)
	manager.AddSeeder(NewRevooE52Batch29Seeder())
	manager.AddSeeder(NewRevooC03Batch29Seeder())
	manager.AddSeeder(NewRevooC32Batch29Seeder())
	manager.AddSeeder(NewRevooC00Batch29Seeder())
	manager.AddSeeder(NewRevooC32YBatch29Seeder())
	manager.AddSeeder(NewRevooA12SBatch29Seeder())
	manager.AddSeeder(NewRevooA12Batch29Seeder())
	manager.AddSeeder(NewRevooA01Batch29Seeder())
	manager.AddSeeder(NewRevooA11Batch29Seeder())
	manager.AddSeeder(NewRevooA10Batch29Seeder())
	manager.AddSeeder(NewRevooA04Batch29Seeder())

	// Batch 30 - Zongshen Complete Series (5 seeders)
	manager.AddSeeder(NewZongshenSparkZS12570Batch30Seeder())
	manager.AddSeeder(NewZongshenCG1254Batch30Seeder())
	manager.AddSeeder(NewZongshenStreetBikeCG125Batch30Seeder())
	manager.AddSeeder(NewZongshenZS1004Batch30Seeder())
	manager.AddSeeder(NewZongshenED80Batch30Seeder())

	// Batch 31 - Akij Complete Series (9 seeders)
	manager.AddSeeder(NewAkijDurdantoV6Batch31Seeder())
	manager.AddSeeder(NewAkijBondhuBatch31Seeder())
	manager.AddSeeder(NewAkijDurbarBatch31Seeder())
	manager.AddSeeder(NewAkijRomeoBatch31Seeder())
	manager.AddSeeder(NewAkijSathiBatch31Seeder())
	manager.AddSeeder(NewAkijDurontoBatch31Seeder())
	manager.AddSeeder(NewAkijDurjoyBatch31Seeder())
	manager.AddSeeder(NewAkijSamratBatch31Seeder())
	manager.AddSeeder(NewAkijPonkhirajBatch31Seeder())

	// Batch 32 - Roadmaster Complete Series (8 seeders)
	manager.AddSeeder(NewRoadmasterRapido165Batch32Seeder())
	manager.AddSeeder(NewRoadmasterRapidoBatch32Seeder())
	manager.AddSeeder(NewRoadmasterVelocity125Batch32Seeder())
	manager.AddSeeder(NewRoadmasterVelocityBatch32Seeder())
	manager.AddSeeder(NewRoadmasterDelightBatch32Seeder())
	manager.AddSeeder(NewRoadmasterPrime100Batch32Seeder())
	manager.AddSeeder(NewRoadmasterRexBatch32Seeder())
	manager.AddSeeder(NewRoadmasterPrime80Batch32Seeder())

	// Batch 33 - PHP Complete Series (5 seeders)
	manager.AddSeeder(NewPHPMerkabaPlus150Batch33Seeder())
	manager.AddSeeder(NewPHPCommando150Batch33Seeder())
	manager.AddSeeder(NewPHPMerkaba150Batch33Seeder())
	manager.AddSeeder(NewPHPPride125Batch33Seeder())
	manager.AddSeeder(NewPHPSuper125Batch33Seeder())

	// Batch 34 - Victor-R Complete Series (3 seeders)
	manager.AddSeeder(NewVictorRCafeRacer125Batch34Seeder())
	manager.AddSeeder(NewVictorRClassic100Batch34Seeder())
	manager.AddSeeder(NewVictorRV80XpressBatch34Seeder())

	// Batch 35 - Road Prince Complete Series (2 seeders)
	manager.AddSeeder(NewRoadPrinceClassicBatch35Seeder())
	manager.AddSeeder(NewRoadPrinceWegoBatch35Seeder())

	// Batch 36 - Atlas Zongshen Complete Series (8 seeders)
	manager.AddSeeder(NewAtlasZongshenZOneTBatch36Seeder())
	manager.AddSeeder(NewAtlasZongshenZS15058Batch36Seeder())
	manager.AddSeeder(NewAtlasZongshenZOneBatch36Seeder())
	manager.AddSeeder(NewAtlasZongshenZS12568Batch36Seeder())
	manager.AddSeeder(NewAtlasZongshenZS11056Batch36Seeder())
	manager.AddSeeder(NewAtlasZongshenZS11072Batch36Seeder())
	manager.AddSeeder(NewAtlasZongshenZS10027Batch36Seeder())
	manager.AddSeeder(NewAtlasZongshenZS80Batch36Seeder())

	// Batch 37 - Haojue Complete Series (4 seeders)
	manager.AddSeeder(NewHaojueDR160Batch37Seeder())
	manager.AddSeeder(NewHaojueTRBatch37Seeder())
	manager.AddSeeder(NewHaojueTZBatch37Seeder())
	manager.AddSeeder(NewHaojueKABatch37Seeder())

	// Batch 38 - Walton Complete Series (5 seeders - excluding Primo which exists as review)
	manager.AddSeeder(NewWaltonTakyonFusion25FZBatch38Seeder())
	manager.AddSeeder(NewWaltonTakyon100Batch38Seeder())
	manager.AddSeeder(NewWaltonTakyon120Batch38Seeder())
	manager.AddSeeder(NewWaltonTakyonLEO25T1Batch38Seeder())
	manager.AddSeeder(NewWaltonTakyonLeo12AHBatch38Seeder())

	// Batch 39 - Runner Complete Series (16 seeders - excluding Hunter & Boxer which exist as reviews)
	manager.AddSeeder(NewRunnerBolt165RBatch39Seeder())
	manager.AddSeeder(NewRunnerSkooty125Batch39Seeder())
	manager.AddSeeder(NewRunnerKnightRiderV2Batch39Seeder())
	manager.AddSeeder(NewRunnerKnightRiderBatch39Seeder())
	manager.AddSeeder(NewRunnerTurbo125MattBatch39Seeder())
	manager.AddSeeder(NewRunnerSkooty110Batch39Seeder())
	manager.AddSeeder(NewRunnerTurbo125V2MattBatch39Seeder())
	manager.AddSeeder(NewRunnerKitePlusBatch39Seeder())
	manager.AddSeeder(NewRunnerBullet100V2MattBatch39Seeder())
	manager.AddSeeder(NewRunnerRoyalPlusBatch39Seeder())
	manager.AddSeeder(NewRunnerBullet100Batch39Seeder())
	manager.AddSeeder(NewRunnerF1006ABatch39Seeder())
	manager.AddSeeder(NewRunnerCheetaBatch39Seeder())
	manager.AddSeeder(NewRunnerRoyalPlusV2Batch39Seeder())
	manager.AddSeeder(NewRunnerAD80sDeluxeRedBatch39Seeder())
	manager.AddSeeder(NewRunnerAD80SAlloyBatch39Seeder())

	return manager
}
