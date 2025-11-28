package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoCamon40 seeds specifications/options for product 'tecno-camon-40'
type SpecificationSeederMobileTecnoCamon40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCamon40 creates a new seeder instance
func NewSpecificationSeederMobileTecnoCamon40() *SpecificationSeederMobileTecnoCamon40 {
	return &SpecificationSeederMobileTecnoCamon40{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCamon40) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ pixels",
		"120Hz": "১২০Hz",
		"163.4 x 75.9 x 7.8 mm": "১৬৩.৪ x ৭৫.৯ x ৭.৮ মিমি",
		"186 g": "১৮৬ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"Black, Blue": "কালো, নীল",
		"February 2024": "February ২০২৪",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"Helio G99": "Helio G৯৯",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-camon-40'
func (s *SpecificationSeederMobileTecnoCamon40) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-40"

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
