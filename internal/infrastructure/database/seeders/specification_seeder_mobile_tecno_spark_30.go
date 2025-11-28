package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSpark30 seeds specifications/options for product 'tecno-spark-30'
type SpecificationSeederMobileTecnoSpark30 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark30 creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark30() *SpecificationSeederMobileTecnoSpark30 {
	return &SpecificationSeederMobileTecnoSpark30{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 30"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark30) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"164 x 75 x 8 mm": "১৬৪ x ৭৫ x ৮ মিমি",
		"186 g": "১৮৬ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 MP + AI": "৬৪ MP + AI",
		"8 GB": "৮ GB",
		"90Hz": "৯০Hz",
		"Android 14": "Android ১৪",
		"Helio G91": "Helio G৯১",
		"IP53": "IP৫৩",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52": "Mali-G৫২",
		"Mediatek Helio G91": "Mediatek Helio G৯১",
		"Orbit Black, Orbit White": "Orbit কালো, Orbit সাদা",
		"September 2024": "September ২০২৪",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-30'
func (s *SpecificationSeederMobileTecnoSpark30) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-30"

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
