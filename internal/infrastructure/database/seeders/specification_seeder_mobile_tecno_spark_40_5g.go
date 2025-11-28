package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSpark405g seeds specifications/options for product 'tecno-spark-40-5g'
type SpecificationSeederMobileTecnoSpark405g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark405g creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark405g() *SpecificationSeederMobileTecnoSpark405g {
	return &SpecificationSeederMobileTecnoSpark405g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 40 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark405g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"163.7 x 75.6 x 8.5 mm": "১৬৩.৭ x ৭৫.৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"256 GB": "২৫৬ GB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + AI": "৫০ MP + AI",
		"5G": "৫G",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 GB": "৮ GB",
		"8 MP": "৮ MP",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"Dimensity 6080": "Dimensity ৬০৮০",
		"Gravity Black, Cyber White": "Gravity কালো, Cyber সাদা",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"January 2024": "January ২০২৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6080 (6 nm)": "Mediatek Dimensity ৬০৮০ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-40-5g'
func (s *SpecificationSeederMobileTecnoSpark405g) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-40-5g"

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
