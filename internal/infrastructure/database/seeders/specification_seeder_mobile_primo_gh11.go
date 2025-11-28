package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePrimoGh11 seeds specifications/options for product 'primo-gh11'
type SpecificationSeederMobilePrimoGh11 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoGh11 creates a new seeder instance
func NewSpecificationSeederMobilePrimoGh11() *SpecificationSeederMobilePrimoGh11 {
	return &SpecificationSeederMobilePrimoGh11{BaseSeeder: BaseSeeder{name: "Specifications for Primo GH11"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoGh11) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"165 x 76 x 9 mm": "১৬৫ x ৭৬ x ৯ মিমি",
		"190 g": "১৯০ g",
		"2 GB": "২ GB",
		"32 GB": "৩২ GB",
		"4,200 mAh": "৪,২০০ এমএএইচ",
		"4G": "৪G",
		"5 MP": "৫ MP",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"Android 11 Go Edition": "Android ১১ Go Edition",
		"Blue, Black": "নীল, কালো",
		"Helio A25": "Helio A২৫",
		"Mediatek Helio A25 (12 nm)": "Mediatek Helio A২৫ (১২ nm)",
		"November 2021": "November ২০২১",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'primo-gh11'
func (s *SpecificationSeederMobilePrimoGh11) Seed(db *gorm.DB) error {
	productSlug := "primo-gh11"

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
