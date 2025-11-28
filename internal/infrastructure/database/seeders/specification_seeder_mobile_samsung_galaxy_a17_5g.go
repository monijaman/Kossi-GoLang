package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA175g seeds specifications/options for product 'samsung-galaxy-a17-5g'
type SpecificationSeederMobileSamsungGalaxyA175g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA175g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA175g() *SpecificationSeederMobileSamsungGalaxyA175g {
	return &SpecificationSeederMobileSamsungGalaxyA175g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A17 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA175g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2408 pixels (~400 ppi density)": "১০৮০ x ২৪০৮ pixels (~৪০০ ppi density)",
		"120Hz": "১২০Hz",
		"128 GB": "১২৮ GB",
		"164.7 x 76.7 x 8.6 mm": "১৬৪.৭ x ৭৬.৭ x ৮.৬ মিমি",
		"194 g (6.84 oz)": "১৯৪ g (৬.৮৪ oz)",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 GB / 8 GB": "৬ GB / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"8 MP": "৮ MP",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Black, Blue": "কালো, নীল",
		"Corning Gorilla Glass 3": "Corning Gorilla Glass ৩",
		"Dimensity 6100+": "Dimensity ৬১০০+",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"July 2025": "July ২০২৫",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek MT6833 Dimensity 6100+ (6 nm)": "MediaTek MT৬৮৩৩ Dimensity ৬১০০+ (৬ nm)",
		"Octa-core (2x2.2 GHz Cortex-A76 & 6x2.0 GHz Cortex-A55)": "Octa-core (২x২.২ GHz Cortex-A৭৬ & ৬x২.০ GHz Cortex-A৫৫)",
		"PLS LCD, 120Hz": "PLS LCD, ১২০Hz",
		"Plastic frame, plastic back": "Plastic frame, প্লাস্টিক পেছনে",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a17-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA175g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a17-5g"

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
