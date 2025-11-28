package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA365g seeds specifications/options for product 'samsung-galaxy-a36-5g'
type SpecificationSeederMobileSamsungGalaxyA365g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA365g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA365g() *SpecificationSeederMobileSamsungGalaxyA365g {
	return &SpecificationSeederMobileSamsungGalaxyA365g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A36 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA365g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels (~399 ppi density)": "১০৮০ x ২৪০০ pixels (~৩৯৯ ppi density)",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"162.5 x 75.8 x 8.1 mm": "১৬২.৫ x ৭৫.৮ x ৮.১ মিমি",
		"195 g (6.88 oz)": "১৯৫ g (৬.৮৮ oz)",
		"20 MP": "২০ MP",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"48 MP + 8 MP + 5 MP": "৪৮ MP + ৮ MP + ৫ MP",
		"6 GB / 8 GB": "৬ GB / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"Adreno 644": "Adreno ৬৪৪",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Black, White, Green": "কালো, সাদা, সবুজ",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP67 dust/water resistant (up to 1m for 30 min)": "IP৬৭ dust/water resistant (up to ১m for ৩০ min)",
		"Octa-core (1x2.4 GHz Cortex-A710 & 3x2.36 GHz Cortex-A710 & 4x1.8 GHz Cortex-A510)": "Octa-core (১x২.৪ GHz Cortex-A৭১০ & ৩x২.৩৬ GHz Cortex-A৭১০ & ৪x১.৮ GHz Cortex-A৫১০)",
		"Qualcomm SM7450-AB Snapdragon 7 Gen 1 (4 nm)": "Qualcomm SM৭৪৫০-AB Snapdragon ৭ Gen ১ (৪ nm)",
		"September 2025": "September ২০২৫",
		"Snapdragon 7 Gen 1": "Snapdragon ৭ Gen ১",
		"Super AMOLED, 120Hz": "Super AMOLED, ১২০Hz",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a36-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA365g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a36-5g"

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
