package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileWaltonOrbitY11 seeds specifications/options for product 'walton-orbit-y11'
type SpecificationSeederMobileWaltonOrbitY11 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonOrbitY11 creates a new seeder instance
func NewSpecificationSeederMobileWaltonOrbitY11() *SpecificationSeederMobileWaltonOrbitY11 {
	return &SpecificationSeederMobileWaltonOrbitY11{BaseSeeder: BaseSeeder{name: "Specifications for Walton ORBIT Y11"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonOrbitY11) getBanglaTranslations() map[string]string {
	return map[string]string{
		"164.5 x 76 x 9 mm": "১৬৪.৫ x ৭৬ x ৯ মিমি",
		"188 g": "১৮৮ g",
		"2 GB": "২ GB",
		"32 GB": "৩২ GB",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"4G": "৪G",
		"5 MP": "৫ MP",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP + 0.3 MP": "৮ MP + ০.৩ MP",
		"Android 12 Go Edition": "Android ১২ Go Edition",
		"Blue, Green": "নীল, সবুজ",
		"Helio A22": "Helio A২২",
		"January 2023": "January ২০২৩",
		"Mediatek Helio A22 (12 nm)": "Mediatek Helio A২২ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'walton-orbit-y11'
func (s *SpecificationSeederMobileWaltonOrbitY11) Seed(db *gorm.DB) error {
	productSlug := "walton-orbit-y11"

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
