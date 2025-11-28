package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA06 seeds specifications/options for product 'samsung-galaxy-a06'
type SpecificationSeederMobileSamsungGalaxyA06 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA06 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA06() *SpecificationSeederMobileSamsungGalaxyA06 {
	return &SpecificationSeederMobileSamsungGalaxyA06{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A06"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA06) getBanglaTranslations() map[string]string {
	return map[string]string{
		"164.0 x 75.8 x 9.2 mm": "১৬৪.০ x ৭৫.৮ x ৯.২ মিমি",
		"190 g (6.70 oz)": "১৯০ g (৬.৭০ oz)",
		"3 GB / 4 GB": "৩ GB / ৪ GB",
		"32 GB / 64 GB": "৩২ GB / ৬৪ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels (~270 ppi density)": "৭২০ x ১৬০০ pixels (~২৭০ ppi density)",
		"8 MP": "৮ MP",
		"Android 14, One UI Core 6": "Android ১৪, One UI Core ৬",
		"Black, Blue": "কালো, নীল",
		"Mali-G57 MP1": "Mali-G৫৭ MP১",
		"May 2025": "May ২০২৫",
		"Octa-core (2x1.6 GHz Cortex-A75 & 6x1.6 GHz Cortex-A55)": "Octa-core (২x১.৬ GHz Cortex-A৭৫ & ৬x১.৬ GHz Cortex-A৫৫)",
		"Plastic frame, plastic back": "Plastic frame, প্লাস্টিক পেছনে",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a06'
func (s *SpecificationSeederMobileSamsungGalaxyA06) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a06"

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
