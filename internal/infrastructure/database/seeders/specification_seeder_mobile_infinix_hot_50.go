package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixHot50 seeds specifications/options for product 'infinix-hot-50'
type SpecificationSeederMobileInfinixHot50 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixHot50 creates a new seeder instance
func NewSpecificationSeederMobileInfinixHot50() *SpecificationSeederMobileInfinixHot50 {
	return &SpecificationSeederMobileInfinixHot50{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Hot 50"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixHot50) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"128 GB": "১২৮ GB",
		"168.7 x 76.6 x 8.4 mm": "১৬৮.৭ x ৭৬.৬ x ৮.৪ মিমি",
		"196 g": "১৯৬ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 13, XOS 13.5": "Android ১৩, XOS ১৩.৫",
		"August 2024": "August ২০২৪",
		"Helio G88": "Helio G৮৮",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G88 (12 nm)": "Mediatek Helio G৮৮ (১২ nm)",
		"Racing Black, Surfing Green": "Racing কালো, Surfing সবুজ",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-hot-50'
func (s *SpecificationSeederMobileInfinixHot50) Seed(db *gorm.DB) error {
	productSlug := "infinix-hot-50"

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
