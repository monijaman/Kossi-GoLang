package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileWaltonXanonX21 seeds specifications/options for product 'walton-xanon-x21'
type SpecificationSeederMobileWaltonXanonX21 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonXanonX21 creates a new seeder instance
func NewSpecificationSeederMobileWaltonXanonX21() *SpecificationSeederMobileWaltonXanonX21 {
	return &SpecificationSeederMobileWaltonXanonX21{BaseSeeder: BaseSeeder{name: "Specifications for Walton XANON X21"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonXanonX21) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"163 x 76 x 8 mm": "১৬৩ x ৭৬ x ৮ মিমি",
		"180 g": "১৮০ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 13": "Android ১৩",
		"Black, Gold": "কালো, সোনালী",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G99": "Helio G৯৯",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
		"October 2023": "October ২০২৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'walton-xanon-x21'
func (s *SpecificationSeederMobileWaltonXanonX21) Seed(db *gorm.DB) error {
	productSlug := "walton-xanon-x21"

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
