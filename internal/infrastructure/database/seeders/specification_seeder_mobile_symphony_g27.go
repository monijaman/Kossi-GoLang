package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyG27 seeds specifications/options for product 'symphony-g27'
type SpecificationSeederMobileSymphonyG27 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyG27 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyG27() *SpecificationSeederMobileSymphonyG27 {
	return &SpecificationSeederMobileSymphonyG27{BaseSeeder: BaseSeeder{name: "Specifications for Symphony G27"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyG27) getBanglaTranslations() map[string]string {
	return map[string]string{
		"164 x 75.8 x 8.9 mm": "১৬৪ x ৭৫.৮ x ৮.৯ মিমি",
		"190 g": "১৯০ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB": "৬৪ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 13": "Android ১৩",
		"Black, Blue": "কালো, নীল",
		"January 2024": "January ২০২৪",
		"Mali-G57 MP1": "Mali-G৫৭ MP১",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-g27'
func (s *SpecificationSeederMobileSymphonyG27) Seed(db *gorm.DB) error {
	productSlug := "symphony-g27"

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
