package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSpark30c seeds specifications/options for product 'tecno-spark-30c'
type SpecificationSeederMobileTecnoSpark30c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark30c creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark30c() *SpecificationSeederMobileTecnoSpark30c {
	return &SpecificationSeederMobileTecnoSpark30c{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 30C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark30c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"165 x 76 x 8 mm": "১৬৫ x ৭৬ x ৮ মিমি",
		"182 g": "১৮২ g",
		"4 GB / 6 GB / 8 GB": "৪ GB / ৬ GB / ৮ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + AI": "৫০ MP + AI",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 14": "Android ১৪",
		"Helio G81": "Helio G৮১",
		"IP54": "IP৫৪",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"Mali-G52": "Mali-G৫২",
		"Mediatek Helio G81": "Mediatek Helio G৮১",
		"Orbit Black, Orbit White, Magic Skin": "Orbit কালো, Orbit সাদা, Magic Skin",
		"September 2024": "September ২০২৪",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-30c'
func (s *SpecificationSeederMobileTecnoSpark30c) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-30c"

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
