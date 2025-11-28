package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoCamon30s seeds specifications/options for product 'tecno-camon-30s'
type SpecificationSeederMobileTecnoCamon30s struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCamon30s creates a new seeder instance
func NewSpecificationSeederMobileTecnoCamon30s() *SpecificationSeederMobileTecnoCamon30s {
	return &SpecificationSeederMobileTecnoCamon30s{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 30S"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCamon30s) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ pixels",
		"120Hz": "১২০Hz",
		"164 x 75 x 7.6 mm": "১৬৪ x ৭৫ x ৭.৬ মিমি",
		"180 g": "১৮০ g",
		"256 GB": "২৫৬ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP": "৫০ MP",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz, Curved": "AMOLED, ১২০Hz, Curved",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"Helio G99 Ultimate": "Helio G৯৯ Ultimate",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"May 2024": "May ২০২৪",
		"Mediatek Helio G99 Ultimate": "Mediatek Helio G৯৯ Ultimate",
		"Nebula Violet, Celestial Black": "Nebula Violet, Celestial কালো",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-camon-30s'
func (s *SpecificationSeederMobileTecnoCamon30s) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-30s"

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
