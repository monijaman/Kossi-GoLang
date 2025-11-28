package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA07 seeds specifications/options for product 'samsung-galaxy-a07'
type SpecificationSeederMobileSamsungGalaxyA07 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA07 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA07() *SpecificationSeederMobileSamsungGalaxyA07 {
	return &SpecificationSeederMobileSamsungGalaxyA07{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A07"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA07) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"164.2 x 75.9 x 9.1 mm": "১৬৪.২ x ৭৫.৯ x ৯.১ মিমি",
		"186 g (6.56 oz)": "১৮৬ g (৬.৫৬ oz)",
		"4 GB / 6 GB": "৪ GB / ৬ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1600 pixels (~270 ppi density)": "৭২০ x ১৬০০ pixels (~২৭০ ppi density)",
		"Android 14, One UI Core 6": "Android ১৪, One UI Core ৬",
		"Black, Green, White": "কালো, সবুজ, সাদা",
		"Exynos 850": "Exynos ৮৫০",
		"June 2025": "June ২০২৫",
		"Mali-G52": "Mali-G৫২",
		"Octa-core (4x2.0 GHz Cortex-A55 & 4x2.0 GHz Cortex-A55)": "Octa-core (৪x২.০ GHz Cortex-A৫৫ & ৪x২.০ GHz Cortex-A৫৫)",
		"Plastic frame, plastic back": "Plastic frame, প্লাস্টিক পেছনে",
		"Samsung Exynos 850 (8 nm)": "Samsung Exynos ৮৫০ (৮ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a07'
func (s *SpecificationSeederMobileSamsungGalaxyA07) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a07"

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
