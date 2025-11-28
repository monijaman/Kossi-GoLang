package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixSmart10Plus seeds specifications/options for product 'infinix-smart-10-plus'
type SpecificationSeederMobileInfinixSmart10Plus struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixSmart10Plus creates a new seeder instance
func NewSpecificationSeederMobileInfinixSmart10Plus() *SpecificationSeederMobileInfinixSmart10Plus {
	return &SpecificationSeederMobileInfinixSmart10Plus{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Smart 10 Plus"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixSmart10Plus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"164 x 76 x 8.5 mm": "১৬৪ x ৭৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"50 MP + AI": "৫০ MP + AI",
		"6,000 mAh": "৬,০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 13 Go Edition": "Android ১৩ Go Edition",
		"Black, Green": "কালো, সবুজ",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"January 2024": "January ২০২৪",
		"Mali-G57 MP1": "Mali-G৫৭ MP১",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-smart-10-plus'
func (s *SpecificationSeederMobileInfinixSmart10Plus) Seed(db *gorm.DB) error {
	productSlug := "infinix-smart-10-plus"

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
