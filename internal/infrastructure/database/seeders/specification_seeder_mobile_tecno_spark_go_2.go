package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSparkGo2 seeds specifications/options for product 'tecno-spark-go-2'
type SpecificationSeederMobileTecnoSparkGo2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSparkGo2 creates a new seeder instance
func NewSpecificationSeederMobileTecnoSparkGo2() *SpecificationSeederMobileTecnoSparkGo2 {
	return &SpecificationSeederMobileTecnoSparkGo2{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK Go 2"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSparkGo2) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + AI": "১৩ MP + AI",
		"164.5 x 76 x 9 mm": "১৬৪.৫ x ৭৬ x ৯ মিমি",
		"188 g": "১৮৮ g",
		"2 GB": "২ GB",
		"2020": "২০২০",
		"32 GB": "৩২ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 10 Go Edition": "Android ১০ Go Edition",
		"Helio A22": "Helio A২২",
		"Ice Jadeite, Aqua Blue": "Ice Jadeite, Aqua নীল",
		"Mediatek Helio A22 (12 nm)": "Mediatek Helio A২২ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-go-2'
func (s *SpecificationSeederMobileTecnoSparkGo2) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-go-2"

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
