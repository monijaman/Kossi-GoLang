package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileWaltonN27 seeds specifications/options for product 'walton-n27'
type SpecificationSeederMobileWaltonN27 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonN27 creates a new seeder instance
func NewSpecificationSeederMobileWaltonN27() *SpecificationSeederMobileWaltonN27 {
	return &SpecificationSeederMobileWaltonN27{BaseSeeder: BaseSeeder{name: "Specifications for Walton N27"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonN27) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"164 x 75.8 x 8.9 mm": "১৬৪ x ৭৫.৮ x ৮.৯ মিমি",
		"192 g": "১৯২ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5 MP": "৫ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB": "৬৪ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"Android 13 Go Edition": "Android ১৩ Go Edition",
		"Blue, Black": "নীল, কালো",
		"Helio G36": "Helio G৩৬",
		"January 2024": "January ২০২৪",
		"Mediatek Helio G36 (12 nm)": "Mediatek Helio G৩৬ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'walton-n27'
func (s *SpecificationSeederMobileWaltonN27) Seed(db *gorm.DB) error {
	productSlug := "walton-n27"

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
