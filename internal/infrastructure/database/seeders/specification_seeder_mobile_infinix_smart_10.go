package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixSmart10 seeds specifications/options for product 'infinix-smart-10'
type SpecificationSeederMobileInfinixSmart10 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixSmart10 creates a new seeder instance
func NewSpecificationSeederMobileInfinixSmart10() *SpecificationSeederMobileInfinixSmart10 {
	return &SpecificationSeederMobileInfinixSmart10{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Smart 10"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixSmart10) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + AI": "১৩ MP + AI",
		"163.6 x 75.6 x 8.5 mm": "১৬৩.৬ x ৭৫.৬ x ৮.৫ মিমি",
		"184 g": "১৮৪ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 13 Go Edition": "Android ১৩ Go Edition",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G57 MP1": "Mali-G৫৭ MP১",
		"November 2023": "November ২০২৩",
		"Timber Black, Shiny Gold": "Timber কালো, Shiny সোনালী",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-smart-10'
func (s *SpecificationSeederMobileInfinixSmart10) Seed(db *gorm.DB) error {
	productSlug := "infinix-smart-10"

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
