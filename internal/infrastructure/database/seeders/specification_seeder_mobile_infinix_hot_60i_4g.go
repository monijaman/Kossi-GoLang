package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixHot60i4g seeds specifications/options for product 'infinix-hot-60i-4g'
type SpecificationSeederMobileInfinixHot60i4g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixHot60i4g creates a new seeder instance
func NewSpecificationSeederMobileInfinixHot60i4g() *SpecificationSeederMobileInfinixHot60i4g {
	return &SpecificationSeederMobileInfinixHot60i4g{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Hot 60i 4G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixHot60i4g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"164 x 76 x 8.5 mm": "১৬৪ x ৭৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"2024": "২০২৪",
		"32 MP": "৩২ MP",
		"4 GB / 8 GB": "৪ GB / ৮ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + AI": "৫০ MP + AI",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"90Hz": "৯০Hz",
		"Android 14": "Android ১৪",
		"Black, Blue, Gold": "কালো, নীল, সোনালী",
		"Helio G85": "Helio G৮৫",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52": "Mali-G৫২",
		"Mediatek Helio G85": "Mediatek Helio G৮৫",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-hot-60i-4g'
func (s *SpecificationSeederMobileInfinixHot60i4g) Seed(db *gorm.DB) error {
	productSlug := "infinix-hot-60i-4g"

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
