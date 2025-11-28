package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOrbitY71 seeds specifications/options for product 'orbit-y71'
type SpecificationSeederMobileOrbitY71 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOrbitY71 creates a new seeder instance
func NewSpecificationSeederMobileOrbitY71() *SpecificationSeederMobileOrbitY71 {
	return &SpecificationSeederMobileOrbitY71{BaseSeeder: BaseSeeder{name: "Specifications for Orbit Y71"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOrbitY71) getBanglaTranslations() map[string]string {
	return map[string]string{
		"164 x 76 x 8.5 mm": "১৬৪ x ৭৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"2023": "২০২৩",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 GB": "৬৪ GB",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 13": "Android ১৩",
		"Black, Blue": "কালো, নীল",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G57 MP1": "Mali-G৫৭ MP১",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'orbit-y71'
func (s *SpecificationSeederMobileOrbitY71) Seed(db *gorm.DB) error {
	productSlug := "orbit-y71"

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
