package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSpark30Pro seeds specifications/options for product 'tecno-spark-30-pro'
type SpecificationSeederMobileTecnoSpark30Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark30Pro creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark30Pro() *SpecificationSeederMobileTecnoSpark30Pro {
	return &SpecificationSeederMobileTecnoSpark30Pro{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 30 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark30Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + AI": "১০৮ MP + AI",
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"168 x 76 x 8.5 mm": "১৬৮ x ৭৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14": "Android ১৪",
		"Black, White": "কালো, সাদা",
		"Helio G99": "Helio G৯৯",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
		"September 2024": "September ২০২৪",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-30-pro'
func (s *SpecificationSeederMobileTecnoSpark30Pro) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-30-pro"

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
