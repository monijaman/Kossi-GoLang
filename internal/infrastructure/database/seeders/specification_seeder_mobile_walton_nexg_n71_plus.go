package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileWaltonNexgN71Plus seeds specifications/options for product 'walton-nexg-n71-plus'
type SpecificationSeederMobileWaltonNexgN71Plus struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonNexgN71Plus creates a new seeder instance
func NewSpecificationSeederMobileWaltonNexgN71Plus() *SpecificationSeederMobileWaltonNexgN71Plus {
	return &SpecificationSeederMobileWaltonNexgN71Plus{BaseSeeder: BaseSeeder{name: "Specifications for Walton NEXG N71 PLUS"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonNexgN71Plus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"164 x 75.8 x 8.5 mm": "১৬৪ x ৭৫.৮ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 GB": "৮ GB",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 13": "Android ১৩",
		"April 2023": "April ২০২৩",
		"Blue, Black": "নীল, কালো",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G57 MP1": "Mali-G৫৭ MP১",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'walton-nexg-n71-plus'
func (s *SpecificationSeederMobileWaltonNexgN71Plus) Seed(db *gorm.DB) error {
	productSlug := "walton-nexg-n71-plus"

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
