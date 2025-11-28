package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePrimoR8 seeds specifications/options for product 'primo-r8'
type SpecificationSeederMobilePrimoR8 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoR8 creates a new seeder instance
func NewSpecificationSeederMobilePrimoR8() *SpecificationSeederMobilePrimoR8 {
	return &SpecificationSeederMobilePrimoR8{BaseSeeder: BaseSeeder{name: "Specifications for Primo R8"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoR8) getBanglaTranslations() map[string]string {
	return map[string]string{
		"164.5 x 76 x 9 mm": "১৬৪.৫ x ৭৬ x ৯ মিমি",
		"188 g": "১৮৮ g",
		"3 GB": "৩ GB",
		"32 GB": "৩২ GB",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"4G": "৪G",
		"5 MP": "৫ MP",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP + 0.3 MP": "৮ MP + ০.৩ MP",
		"Android 11 Go Edition": "Android ১১ Go Edition",
		"August 2021": "August ২০২১",
		"Blue, Green": "নীল, সবুজ",
		"Helio A22": "Helio A২২",
		"Mediatek Helio A22 (12 nm)": "Mediatek Helio A২২ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'primo-r8'
func (s *SpecificationSeederMobilePrimoR8) Seed(db *gorm.DB) error {
	productSlug := "primo-r8"

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
