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

// GetCommonBanglaTranslations returns a map of common English specification values to their Bangla translations
func (bs *BaseSeeder) GetCommonBanglaTranslations() map[string]string {
	return map[string]string{
		// Screen sizes
		"6.1 inches": "৬.১ ইঞ্চি",
		"6.2 inches": "৬.২ ইঞ্চি",
		"6.3 inches": "৬.৩ ইঞ্চি",
		"6.4 inches": "৬.৪ ইঞ্চি",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"7.0 inches": "৭.০ ইঞ্চি",

		// Display types
		"AMOLED":                "AMOLED",
		"OLED":                  "OLED",
		"LTPO OLED":             "LTPO OLED",
		"Super Retina XDR OLED": "সুপার রেটিনা XDR OLED",
		"Liquid Retina HD":      "লিকুইড রেটিনা HD",
		"IPS LCD":               "IPS LCD",
		"TFT LCD":               "TFT LCD",

		// Display features
		"HDR10+":       "HDR10+",
		"HDR10":        "HDR10",
		"Dolby Vision": "ডলবি ভিশন",
		"120Hz":        "১২০Hz",
		"90Hz":         "৯০Hz",
		"60Hz":         "৬০Hz",
		"1-120Hz":      "১-১২০Hz",

		// Resolutions
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ পিক্সেল",
		"1080 x 2412 pixels": "১০৮০ x ২৪১২ পিক্সেল",
		"1179 x 2556 pixels": "১১৭৯ x ২৫৫৬ পিক্সেল",
		"1284 x 2778 pixels": "১২৮৪ x ২৭৭৮ পিক্সেল",
		"1290 x 2796 pixels": "১২৯০ x ২৭৯৬ পিক্সেল",
		"1344 x 2992 pixels": "১৩৪৪ x ২৯৯২ পিক্সেল",

		// Processors
		"Apple A18":                   "অ্যাপল A18",
		"Apple A18 Pro":               "অ্যাপল A18 প্রো",
		"Apple A19":                   "অ্যাপল A19",
		"Google Tensor G3":            "গুগল টেনসর G3",
		"Google Tensor G4":            "গুগল টেনসর G4",
		"Qualcomm Snapdragon 8 Gen 3": "কোয়ালকম স্ন্যাপড্রাগন ৮ জেন ৩",
		"Qualcomm Snapdragon 8 Gen 4": "কোয়ালকম স্ন্যাপড্রাগন ৮ জেন ৪",
		"MediaTek Dimensity 9300":     "মিডিয়াটেক ডাইমেনসিটি ৯৩০০",
		"MediaTek Dimensity 9400":     "মিডিয়াটেক ডাইমেনসিটি ৯৪০০",

		// CPU types
		"Hexa-core": "হেক্সা-কোর",
		"Octa-core": "অক্টা-কোর",
		"Nona-core": "ননা-কোর",
		"Deca-core": "ডেকা-কোর",

		// RAM
		"4 GB":  "৪ GB",
		"6 GB":  "৬ GB",
		"8 GB":  "৮ GB",
		"12 GB": "১২ GB",
		"16 GB": "১৬ GB",
		"24 GB": "২৪ GB",

		// Storage
		"64 GB":                           "৬৪ GB",
		"128 GB":                          "১২৮ GB",
		"256 GB":                          "২৫৬ GB",
		"512 GB":                          "৫১২ GB",
		"1 TB":                            "১ TB",
		"128 GB / 256 GB":                 "১২৮ GB / ২৫৬ GB",
		"128 GB / 256 GB / 512 GB":        "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"128 GB / 256 GB / 512 GB / 1 TB": "১২৮ GB / ২৫৬ GB / ৫১২ GB / ১ TB",

		// Battery
		"3000 mAh": "৩০০০ mAh",
		"4000 mAh": "৪০০০ mAh",
		"4500 mAh": "৪৫০০ mAh",
		"5000 mAh": "৫০০০ mAh",
		"6000 mAh": "৬০০০ mAh",

		// Operating Systems
		"Android 13": "অ্যান্ড্রয়েড ১৩",
		"Android 14": "অ্যান্ড্রয়েড ১৪",
		"Android 15": "অ্যান্ড্রয়েড ১৫",
		"iOS 17":     "iOS ১৭",
		"iOS 18":     "iOS ১৮",

		// Colors
		"Black":  "কালো",
		"White":  "সাদা",
		"Blue":   "নীল",
		"Red":    "লাল",
		"Green":  "সবুজ",
		"Yellow": "হলুদ",
		"Purple": "বেগুনি",
		"Pink":   "গোলাপী",
		"Gray":   "ধূসর",
		"Silver": "রূপালী",
		"Gold":   "সোনালী",

		// Protection
		"Corning Gorilla Glass":          "কর্নিং গরিলা গ্লাস",
		"Corning Gorilla Glass Victus":   "কর্নিং গরিলা গ্লাস ভিক্টাস",
		"Corning Gorilla Glass Victus 2": "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Ceramic Shield":                 "সিরামিক শিল্ড",

		// Water resistance
		"IP68": "IP68",
		"IP67": "IP67",

		// Network
		"5G":  "5G",
		"4G":  "4G",
		"LTE": "LTE",

		// Status
		"Available": "উপলব্ধ",

		// Months
		"January":   "জানুয়ারি",
		"February":  "ফেব্রুয়ারি",
		"March":     "মার্চ",
		"April":     "এপ্রিল",
		"May":       "মে",
		"June":      "জুন",
		"July":      "জুলাই",
		"August":    "আগস্ট",
		"September": "সেপ্টেম্বর",
		"October":   "অক্টোবর",
		"November":  "নভেম্বর",
		"December":  "ডিসেম্বর",

		// Years
		"2023": "২০২৩",
		"2024": "২০২৪",
		"2025": "২০২৫",

		// Connectivity
		"USB-C":     "ইউএসবি-সি",
		"Wi-Fi":     "ওয়াই-ফাই",
		"Bluetooth": "ব্লুটুথ",
		"GPS":       "জিপিএস",
		"GLONASS":   "গ্লোনাস",
		"GALILEO":   "গ্যালিলিও",
		"BDS":       "বিডিএস",
		"NFC":       "এনএফসি",

		// SIM
		"Nano-SIM": "ন্যানো-সিম",
		"eSIM":     "ই-সিম",

		// Camera
		"OIS":       "ওআইএস",
		"LED flash": "এলইডি ফ্ল্যাশ",
		"4K":        "৪কে",

		// Network bands
		"GSM":   "জিএসএম",
		"HSPA":  "এইচএসপিএ",
		"HSDPA": "এইচএসডিপিএ",
		"SA":    "এসএ",
		"NSA":   "এনএসএ",
		"Sub6":  "সাব৬",

		// UI
		"One UI": "ওয়ান UI",

		// Colors additional
		"Graphite": "গ্রাফাইট",
		"Violet":   "বেগুনি",

		// Storage
		"microSDXC": "মাইক্রো এসডিএক্সসি",

		// Battery
		"Li-Ion": "লি-আয়ন",

		// Features
		"Accelerometer":    "অ্যাক্সেলেরোমিটার",
		"Gyro":             "গাইরো",
		"Compass":          "কম্পাস",
		"Fingerprint":      "ফিঙ্গারপ্রিন্ট",
		"Side fingerprint": "সাইড ফিঙ্গারপ্রিন্ট",
		"Proximity":        "প্রক্সিমিটি",
		"Depth":            "ডেপথ",

		// Display types additional
		"PLS LCD": "পিএলএস LCD",

		// Colors additional
		"Blazing Black": "ব্লেজিং ব্ল্যাক",
		"Sage Green":    "সেজ গ্রীন",

		// Additional screen sizes
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"6.72 inches": "৬.৭২ ইঞ্চি",

		// Resolutions additional
		"720 x 1600 pixels":  "৭২০ × ১৬০০ পিক্সেল",
		"720 × 1604 pixels":  "৭২০ × ১৬০৪ পিক্সেল",
		"1080 × 2400 pixels": "১০৮০ × ২৪০০ পিক্সেল",

		// Processors additional
		"Helio G37":          "হেলিও G37",
		"Unisoc T606":        "ইউনিসক T606",
		"MediaTek Helio G85": "মিডিয়াটেক হেলিও G85",
		"MediaTek Helio G88": "মিডিয়াটেক হেলিও G88",

		// Chipsets
		"Mediatek Helio G37 (12 nm)": "মিডিয়াটেক হেলিও G37 (১২ nm)",
		"Unisoc T606 (12 nm)":        "ইউনিসক T606 (১২ nm)",
		"MediaTek Helio G85 (12 nm)": "মিডিয়াটেক হেলিও G85 (১২ nm)",

		// GPU
		"PowerVR GE8320": "পাওয়ারভিআর GE8320",
		"Mali-G57 MP1":   "মালি-G57 MP1",
		"Mali-G52":       "মালি-G52",
		"Mali-G52 MC2":   "মালি-G52 MC2",

		// Cameras
		"8 MP":                   "৮ মেগাপিক্সেল",
		"16 MP":                  "১৬ মেগাপিক্সেল",
		"32 MP":                  "৩২ মেগাপিক্সেল",
		"50 MP + 2 MP":           "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"108 MP + 2 MP":          "১০৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"108 MP + 2 MP + 0.8 MP": "১০৮ মেগাপিক্সেল + ২ মেগাপিক্সেল + ০.৮ মেগাপিক্সেল",

		// Dimensions
		"164 x 75.8 x 8.9 mm":    "১৬৪ × ৭৫.৮ × ৮.৯ মিলিমিটার",
		"167.4 x 77.6 x 8.54 mm": "১৬৭.৪ × ৭৭.৬ × ৮.৫৪ মিলিমিটার",
		"167 x 77.6 x 8.79 mm":   "১৬৭ × ৭৭.৬ × ৮.৭৯ মিলিমিটার",

		// Weight
		"190 g": "১৯০ গ্রাম",
		"200 g": "২০০ গ্রাম",
		"215 g": "২১৫ গ্রাম",

		// Build materials
		"Plastic body":              "প্লাস্টিক বডি",
		"Glass front, plastic back": "গ্লাস ফ্রন্ট, প্লাস্টিক ব্যাক",

		// Protection
		"Glass front": "গ্লাস ফ্রন্ট",

		// Water resistance
		"No": "না",

		// Battery types
		"Li-Po (non-removable)": "লি-পো (নন-রিমুভেবল)",

		// Charging
		"18 W wired": "১৮ ওয়াট তারযুক্ত",
		"33 W wired": "৩৩ ওয়াট তারযুক্ত",

		// Audio
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ mm",
		"Stereo":      "স্টেরিও",
		"Standard":    "স্ট্যান্ডার্ড",

		// Network
		"GSM / HSPA / LTE": "জিএসএম / এইচএসপিএ / এলটিই",

		// Wi-Fi
		"Wi-Fi 802.11 b/g/n": "ওয়াই-ফাই ৮০২.১১ b/g/n",

		// Bluetooth
		"5.0": "৫.০",
		"5.x": "৫.x",

		// USB
		"USB Type-C 2.0": "ইউএসবি টাইপ-সি ২.০",

		// Positioning
		"GPS, A-GPS": "জিপিএস, এ-জিপিএস",

		// SIM
		"Dual Nano-SIM / Hybrid": "ডুয়াল ন্যানো-সিম / হাইব্রিড",

		// Video resolutions
		"2K@30fps, 1080p@30fps": "২কে@৩০fps, ১০৮০পি@৩০fps",
		"1080p@30fps":           "১০৮০পি@৩০fps",

		// CPU types additional
		"Octa-core, up to 2.0 GHz": "অক্টা-কোর, পর্যন্ত ২.০ GHz",

		// RAM additional
		"6 / 8 GB": "৬ / ৮ GB",

		// Storage additional
		"128 GB + microSD": "১২৮ GB + মাইক্রো এসডি",

		// Features
		"Always-on display":                             "অলওয়েজ-অন ডিসপ্লে",
		"AI, Portrait, Pro Mode, Night, HDR, Macro":     "এআই, পোর্ট্রেট, প্রো মোড, নাইট, এইচডিআর, ম্যাক্রো",
		"AI, RAW, Slow motion, Portrait, HDR":           "এআই, রঅ, স্লো মোশন, পোর্ট্রেট, এইচডিআর",
		"Side fingerprint, proximity, light, gyroscope": "সাইড ফিঙ্গারপ্রিন্ট, প্রক্সিমিটি, লাইট, গাইরোস্কোপ",

		// Colors multiple
		"Green, Black": "সবুজ, কালো",
		"Black, Blue":  "কালো, নীল",
		"Forest Green, Cosmic Gold, Sunshine Green, Twilight Gold, Sky Blue": "ফরেস্ট গ্রীন, কসমিক গোল্ড, সানশাইন গ্রীন, টোয়াইলাইট গোল্ড, স্কাই ব্লু",
		"Radium Green, Cosmic Gold, Honey Dew Green":                         "রেডিয়াম গ্রীন, কসমিক গোল্ড, হানি ডিউ গ্রীন",

		// Announcement dates
		"December 2023": "ডিসেম্বর ২০২৩",
		"January 2024":  "জানুয়ারি ২০২৪",
		"June 2024":     "জুন ২০২৪",

		// Optical zoom
		"None": "কোনোটি না",

		// Units
		"MP":    "মেগাপিক্সেল",
		"Hz":    "হার্জ",
		"mm":    "মিলিমিটার",
		"px":    "পিক্সেল",
		"fps":   "এফপিএস",
		"mAh":   "মিলিঅ্যাম্পিয়ার আওয়ার",
		"ppi":   "পিপিআই",
		"GB":    "জিবি",
		"TB":    "টিবি",
		"1080p": "১০৮০পি",
	}
}

// Helper function to create or find a category
func CreateOrFindCategory(db *gorm.DB, name, slug string) (*entities.Category, error) {
	var categoryModel models.CategoryModel

	// Use FirstOrCreate to handle duplicates
	category := &entities.Category{
		Name: name,
		Slug: slug,
	}

	categoryModel.FromEntity(category)
	if err := db.Where("slug = ?", slug).FirstOrCreate(&categoryModel).Error; err != nil {
		return nil, err
	}

	return categoryModel.ToEntity(), nil
}

// Helper function to create or find a brand
func CreateOrFindBrand(db *gorm.DB, name, slug string) (*entities.Brand, error) {
	var brandModel models.BrandModel

	// Use FirstOrCreate to handle duplicates
	brand := &entities.Brand{
		Name: name,
		Slug: slug,
	}

	brandModel.FromEntity(brand)
	if err := db.Where("slug = ?", slug).FirstOrCreate(&brandModel).Error; err != nil {
		return nil, err
	}

	return brandModel.ToEntity(), nil
}

// Helper function to create brand-category relationship
func CreateBrandCategoryRelation(db *gorm.DB, brandID, categoryID uint) error {
	var brandCategoryModel models.BrandCategoryModel

	// Use FirstOrCreate to handle duplicates
	relation := &entities.BrandCategory{
		BrandID:    brandID,
		CategoryID: categoryID,
	}

	brandCategoryModel.FromEntity(relation)
	if err := db.Where("brand_id = ? AND category_id = ?", brandID, categoryID).FirstOrCreate(&brandCategoryModel).Error; err != nil {
		return err
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

// MobileSpecificationsSeeder implements Seeder for mobile specifications
// type MobileSpecificationsSeeder struct{}

// func (s *MobileSpecificationsSeeder) Seed(db *gorm.DB) error {
// 	return SeedMobileSpecifications(db)
// }

// func (s *MobileSpecificationsSeeder) GetName() string {
// 	return "MobileSpecifications"
// }

// SetupAllSeeders configures all available seeders
func SetupAllSeeders(db *gorm.DB) *SeederManager {
	manager := NewSeederManager(db)
	// Register a minimal, safe set of seeders that are present in the repo.
	// This avoids referencing constructors that may not exist yet while we
	// iteratively add more motorcycle-specific seeders.

	// Core/global seeders - Specification keys MUST run first
	// NOTE: SpecificationKeySeeder is in mobiles/ subdirectory which causes compilation issues
	// The CreateOrFindSpecificationKey function already handles auto-creation of missing keys
	// manager.AddSeeder(NewSpecificationKeySeeder())

	// Core/global seeders
	// manager.AddSeeder(NewCategorySeeder())
	// manager.AddSeeder(NewBrandSeeder())
	// manager.AddSeeder(NewUserSeeder())
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

	// Laptop category seeders (Category ID: 71)
	// Product seeder MUST run first to create products before specification seeders
	manager.AddSeeder(NewLaptopProductSeeder())
	manager.AddSeeder(NewSpecificLaptopProductsSeeder()) // Create specific products for specification seeders

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

	return manager
}
