package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePrimoS8 seeds specifications/options for product 'primo-s8'
type SpecificationSeederMobilePrimoS8 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoS8 creates a new seeder instance
func NewSpecificationSeederMobilePrimoS8() *SpecificationSeederMobilePrimoS8 {
	return &SpecificationSeederMobilePrimoS8{BaseSeeder: BaseSeeder{name: "Specifications for Primo S8"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoS8) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"128 GB": "১২৮ GB",
		"169 x 76.5 x 8.9 mm": "১৬৯ x ৭৬.৫ x ৮.৯ মিমি",
		"205 g": "২০৫ g",
		"48 MP + 5 MP + 2 MP + 2 MP": "৪৮ MP + ৫ MP + ২ MP + ২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6 GB": "৬ GB",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 11": "Android ১১",
		"Black, Blue": "কালো, নীল",
		"December 2021": "December ২০২১",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G88": "Helio G৮৮",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G88 (12 nm)": "Mediatek Helio G৮৮ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'primo-s8'
func (s *SpecificationSeederMobilePrimoS8) Seed(db *gorm.DB) error {
	productSlug := "primo-s8"

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
