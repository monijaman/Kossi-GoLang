package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSparkGo1 seeds specifications/options for product 'tecno-spark-go-1'
type SpecificationSeederMobileTecnoSparkGo1 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSparkGo1 creates a new seeder instance
func NewSpecificationSeederMobileTecnoSparkGo1() *SpecificationSeederMobileTecnoSparkGo1 {
	return &SpecificationSeederMobileTecnoSparkGo1{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK Go 1"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSparkGo1) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"13 MP": "১৩ MP",
		"165.6 x 76.3 x 8.4 mm": "১৬৫.৬ x ৭৬.৩ x ৮.৪ মিমি",
		"180 g": "১৮০ g",
		"3 GB / 4 GB": "৩ GB / ৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 14 Go Edition": "Android ১৪ Go Edition",
		"August 2024": "August ২০২৪",
		"IP54": "IP৫৪",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"Mali-G57": "Mali-G৫৭",
		"Startrail Black, Glittery White": "Startrail কালো, Glittery সাদা",
		"Unisoc T615": "Unisoc T৬১৫",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-go-1'
func (s *SpecificationSeederMobileTecnoSparkGo1) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-go-1"

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
