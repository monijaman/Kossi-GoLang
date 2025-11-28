package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA565g seeds specifications/options for product 'samsung-galaxy-a56-5g'
type SpecificationSeederMobileSamsungGalaxyA565g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA565g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA565g() *SpecificationSeederMobileSamsungGalaxyA565g {
	return &SpecificationSeederMobileSamsungGalaxyA565g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A56 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA565g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels (~393 ppi density)": "১০৮০ x ২৪০০ pixels (~৩৯৩ ppi density)",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"165.1 x 76.4 x 8.4 mm": "১৬৫.১ x ৭৬.৪ x ৮.৪ মিমি",
		"202 g (7.13 oz)": "২০২ g (৭.১৩ oz)",
		"32 MP": "৩২ MP",
		"50 MP + 12 MP + 5 MP": "৫০ MP + ১২ MP + ৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Black, White, Blue": "কালো, সাদা, নীল",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Exynos 1380": "Exynos ১৩৮০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP67 dust/water resistant (up to 1m for 30 min)": "IP৬৭ dust/water resistant (up to ১m for ৩০ min)",
		"Mali-G68 MP5": "Mali-G৬৮ MP৫",
		"Octa-core (4x2.4 GHz Cortex-A78 & 4x2.0 GHz Cortex-A55)": "Octa-core (৪x২.৪ GHz Cortex-A৭৮ & ৪x২.০ GHz Cortex-A৫৫)",
		"October 2025": "October ২০২৫",
		"Samsung Exynos 1380 (5 nm)": "Samsung Exynos ১৩৮০ (৫ nm)",
		"সুপার AMOLED, 120Hz": "সুপার AMOLED, ১২০Hz",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a56-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA565g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a56-5g"

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
