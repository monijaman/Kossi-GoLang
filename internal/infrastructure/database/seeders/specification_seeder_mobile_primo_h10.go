package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePrimoH10 seeds specifications/options for product 'primo-h10'
type SpecificationSeederMobilePrimoH10 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoH10 creates a new seeder instance
func NewSpecificationSeederMobilePrimoH10() *SpecificationSeederMobilePrimoH10 {
	return &SpecificationSeederMobilePrimoH10{BaseSeeder: BaseSeeder{name: "Specifications for Primo H10"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoH10) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"165 x 76 x 9 mm": "১৬৫ x ৭৬ x ৯ মিমি",
		"190 g": "১৯০ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB": "৬৪ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 11": "Android ১১",
		"Blue, Black": "নীল, কালো",
		"Helio G35": "Helio G৩৫",
		"Mediatek Helio G35 (12 nm)": "Mediatek Helio G৩৫ (১২ nm)",
		"October 2021": "October ২০২১",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'primo-h10'
func (s *SpecificationSeederMobilePrimoH10) Seed(db *gorm.DB) error {
	productSlug := "primo-h10"

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
