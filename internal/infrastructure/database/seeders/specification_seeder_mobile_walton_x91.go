package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileWaltonX91 seeds specifications/options for product 'walton-x91'
type SpecificationSeederMobileWaltonX91 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonX91 creates a new seeder instance
func NewSpecificationSeederMobileWaltonX91() *SpecificationSeederMobileWaltonX91 {
	return &SpecificationSeederMobileWaltonX91{BaseSeeder: BaseSeeder{name: "Specifications for Walton X91"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonX91) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"128 GB": "১২৮ GB",
		"164.4 x 76.8 x 7.8 mm": "১৬৪.৪ x ৭৬.৮ x ৭.৮ মিমি",
		"185 g": "১৮৫ g",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 5 MP + 2 MP": "৬৪ MP + ৫ MP + ২ MP",
		"8 GB": "৮ GB",
		"90Hz": "৯০Hz",
		"AMOLED, 90Hz": "AMOLED, ৯০Hz",
		"Android 13": "Android ১৩",
		"Black, Blue": "কালো, নীল",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G99": "Helio G৯৯",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"March 2024": "March ২০২৪",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'walton-x91'
func (s *SpecificationSeederMobileWaltonX91) Seed(db *gorm.DB) error {
	productSlug := "walton-x91"

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
