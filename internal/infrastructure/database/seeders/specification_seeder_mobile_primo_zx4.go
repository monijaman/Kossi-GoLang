package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePrimoZx4 seeds specifications/options for product 'primo-zx4'
type SpecificationSeederMobilePrimoZx4 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoZx4 creates a new seeder instance
func NewSpecificationSeederMobilePrimoZx4() *SpecificationSeederMobilePrimoZx4 {
	return &SpecificationSeederMobilePrimoZx4{BaseSeeder: BaseSeeder{name: "Specifications for Primo ZX4"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoZx4) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"165 x 76 x 8.9 mm": "১৬৫ x ৭৬ x ৮.৯ মিমি",
		"200 g": "২০০ g",
		"4 GB": "৪ GB",
		"48 MP + 5 MP + 2 MP": "৪৮ MP + ৫ MP + ২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 GB": "৬৪ GB",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 11": "Android ১১",
		"Black, Blue": "কালো, নীল",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G88": "Helio G৮৮",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G88 (12 nm)": "Mediatek Helio G৮৮ (১২ nm)",
		"September 2021": "September ২০২১",
	}
}

// Seed inserts specification_translations for existing specifications for product 'primo-zx4'
func (s *SpecificationSeederMobilePrimoZx4) Seed(db *gorm.DB) error {
	productSlug := "primo-zx4"

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
