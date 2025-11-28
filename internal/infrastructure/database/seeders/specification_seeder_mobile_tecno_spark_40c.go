package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSpark40c seeds specifications/options for product 'tecno-spark-40c'
type SpecificationSeederMobileTecnoSpark40c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark40c creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark40c() *SpecificationSeederMobileTecnoSpark40c {
	return &SpecificationSeederMobileTecnoSpark40c{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 40C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark40c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"163.7 x 75.6 x 8.6 mm": "১৬৩.৭ x ৭৫.৬ x ৮.৬ মিমি",
		"185 g": "১৮৫ g",
		"4 GB / 8 GB": "৪ GB / ৮ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + AI": "৫০ MP + AI",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 13": "Android ১৩",
		"Gravity Black, Mystery White": "Gravity কালো, Mystery সাদা",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G57 MP1": "Mali-G৫৭ MP১",
		"November 2023": "November ২০২৩",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-40c'
func (s *SpecificationSeederMobileTecnoSpark40c) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-40c"

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
