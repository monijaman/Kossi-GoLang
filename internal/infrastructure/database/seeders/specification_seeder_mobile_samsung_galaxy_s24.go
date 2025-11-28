package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS24 seeds specifications/options for product 'samsung-galaxy-s24'
type SpecificationSeederMobileSamsungGalaxyS24 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS24 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS24() *SpecificationSeederMobileSamsungGalaxyS24 {
	return &SpecificationSeederMobileSamsungGalaxyS24{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S24"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS24) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"147 x 70.6 x 7.6 mm": "১৪৭ x ৭০.৬ x ৭.৬ মিমি",
		"167 g": "১৬৭ g",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ MP + ১০ MP + ১২ MP",
		"5G": "৫G",
		"6.2 inches": "৬.২ ইঞ্চি",
		"8 GB": "৮ GB",
		"Adreno 750 / Xclipse 940": "Adreno ৭৫০ / Xclipse ৯৪০",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 14, One UI 6.1": "Android ১৪, One UI ৬.১",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Dynamic LTPO AMOLED ২X, ১২০Hz, HDR১০+",
		"IP68": "IP৬৮",
		"January 2024": "January ২০২৪",
		"Onyx Black, Marble Grey, Cobalt Violet, Amber Yellow": "Onyx কালো, Marble ধূসর, Cobalt Violet, Amber হলুদ",
		"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AC Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3 / Exynos 2400": "Snapdragon ৮ Gen ৩ / Exynos ২৪০০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s24'
func (s *SpecificationSeederMobileSamsungGalaxyS24) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s24"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	productID := prod.ID
	banglaTranslations := s.getBanglaTranslations()

	// Get all existing specifications for this product
	var existingSpecs []models.SpecificationModel
	if err := db.Where("product_id = ?", productID).Find(&existingSpecs).Error; err != nil {
		return err
	}

	// Insert translations for all existing specifications
	for _, spec := range existingSpecs {
		banglaValue, exists := banglaTranslations[spec.Value]
		if exists && banglaValue != "" {
			translation := &models.SpecificationTranslationModel{
				SpecificationID: spec.ID,
				Locale:          "bn",
				Value:           banglaValue,
			}
			// Use OnConflict to ignore if translation already exists
			if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(translation).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
