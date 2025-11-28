package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyAtom5 seeds specifications/options for product 'symphony-atom-5'
type SpecificationSeederMobileSymphonyAtom5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyAtom5 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyAtom5() *SpecificationSeederMobileSymphonyAtom5 {
	return &SpecificationSeederMobileSymphonyAtom5{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Atom 5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyAtom5) getBanglaTranslations() map[string]string {
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
		"December 2023": "December ২০২৩",
		"Green, Black": "সবুজ, কালো",
		"Helio G37": "Helio G৩৭",
		"Mediatek Helio G37 (12 nm)": "Mediatek Helio G৩৭ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-atom-5'
func (s *SpecificationSeederMobileSymphonyAtom5) Seed(db *gorm.DB) error {
	productSlug := "symphony-atom-5"

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
