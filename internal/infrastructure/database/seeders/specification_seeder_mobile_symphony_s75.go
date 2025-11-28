package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyS75 seeds specifications/options for product 'symphony-s75'
type SpecificationSeederMobileSymphonyS75 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyS75 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyS75() *SpecificationSeederMobileSymphonyS75 {
	return &SpecificationSeederMobileSymphonyS75{BaseSeeder: BaseSeeder{name: "Specifications for Symphony S75"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyS75) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"164.2 x 76 x 8.5 mm": "১৬৪.২ x ৭৬ x ৮.৫ মিমি",
		"192 g": "১৯২ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6 GB": "৬ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 13": "Android ১৩",
		"Black, Gold": "কালো, সোনালী",
		"Helio G85": "Helio G৮৫",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G85 (12 nm)": "Mediatek Helio G৮৫ (১২ nm)",
		"November 2023": "November ২০২৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-s75'
func (s *SpecificationSeederMobileSymphonyS75) Seed(db *gorm.DB) error {
	productSlug := "symphony-s75"

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
