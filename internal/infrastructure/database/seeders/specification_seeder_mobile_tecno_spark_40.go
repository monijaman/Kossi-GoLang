package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSpark40 seeds specifications/options for product 'tecno-spark-40'
type SpecificationSeederMobileTecnoSpark40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark40 creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark40() *SpecificationSeederMobileTecnoSpark40 {
	return &SpecificationSeederMobileTecnoSpark40{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark40) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"163.7 x 75.6 x 8.5 mm": "১৬৩.৭ x ৭৫.৬ x ৮.৫ মিমি",
		"185 g": "১৮৫ g",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + AI": "৫০ MP + AI",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 GB": "৮ GB",
		"90Hz": "৯০Hz",
		"Android 13, HIOS 13": "Android ১৩, HIOS ১৩",
		"December 2023": "December ২০২৩",
		"Gravity Black, Cyber White": "Gravity কালো, Cyber সাদা",
		"Helio G85": "Helio G৮৫",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G85 (12 nm)": "Mediatek Helio G৮৫ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-40'
func (s *SpecificationSeederMobileTecnoSpark40) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-40"

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
