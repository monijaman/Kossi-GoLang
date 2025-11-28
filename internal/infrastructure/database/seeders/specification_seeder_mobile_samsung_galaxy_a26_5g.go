package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA265g seeds specifications/options for product 'samsung-galaxy-a26-5g'
type SpecificationSeederMobileSamsungGalaxyA265g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA265g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA265g() *SpecificationSeederMobileSamsungGalaxyA265g {
	return &SpecificationSeederMobileSamsungGalaxyA265g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A26 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA265g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels (~405 ppi density)": "১০৮০ x ২৪০০ pixels (~৪০৫ ppi density)",
		"120Hz": "১২০Hz",
		"128 GB": "১২৮ GB",
		"13 MP": "১৩ MP",
		"161.3 x 73.9 x 8.4 mm": "১৬১.৩ x ৭৩.৯ x ৮.৪ মিমি",
		"189 g (6.67 oz)": "১৮৯ g (৬.৬৭ oz)",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6 GB / 8 GB": "৬ GB / ৮ GB",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"August 2025": "August ২০২৫",
		"Black, Blue": "কালো, নীল",
		"Corning Gorilla Glass 3": "Corning Gorilla Glass ৩",
		"Dimensity 6020": "Dimensity ৬০২০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek MT6833 Dimensity 6020 (7 nm)": "MediaTek MT৬৮৩৩ Dimensity ৬০২০ (৭ nm)",
		"Octa-core (2x2.2 GHz Cortex-A76 & 6x2.0 GHz Cortex-A55)": "Octa-core (২x২.২ GHz Cortex-A৭৬ & ৬x২.০ GHz Cortex-A৫৫)",
		"PLS LCD, 120Hz": "PLS LCD, ১২০Hz",
		"Plastic frame, plastic back": "Plastic frame, প্লাস্টিক পেছনে",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a26-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA265g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a26-5g"

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
