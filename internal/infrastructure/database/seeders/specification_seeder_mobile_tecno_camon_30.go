package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoCamon30 seeds specifications/options for product 'tecno-camon-30'
type SpecificationSeederMobileTecnoCamon30 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCamon30 creates a new seeder instance
func NewSpecificationSeederMobileTecnoCamon30() *SpecificationSeederMobileTecnoCamon30 {
	return &SpecificationSeederMobileTecnoCamon30{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 30"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCamon30) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ pixels",
		"120Hz": "১২০Hz",
		"165.3 x 75.3 x 7.7 mm": "১৬৫.৩ x ৭৫.৩ x ৭.৭ মিমি",
		"189 g": "১৮৯ g",
		"256 GB": "২৫৬ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP": "৫০ MP",
		"50 MP + 2 MP + AI": "৫০ MP + ২ MP + AI",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"February 2024": "February ২০২৪",
		"Glass front, glass back": "গ্লাস সামনে, গ্লাস পেছনে",
		"Helio G99 Ultimate": "Helio G৯৯ Ultimate",
		"IP53": "IP৫৩",
		"Iceland Basalt, Uyuni Salt White": "Iceland Basalt, Uyuni Salt সাদা",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Helio G99 Ultimate": "Mediatek Helio G৯৯ Ultimate",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-camon-30'
func (s *SpecificationSeederMobileTecnoCamon30) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-30"

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
