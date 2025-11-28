package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplusNordN20Se seeds specifications/options for product 'oneplus-nord-n20-se'
type SpecificationSeederMobileOneplusNordN20Se struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNordN20Se creates a new seeder instance
func NewSpecificationSeederMobileOneplusNordN20Se() *SpecificationSeederMobileOneplusNordN20Se {
	return &SpecificationSeederMobileOneplusNordN20Se{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord N20 SE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNordN20Se) getBanglaTranslations() map[string]string {
	return map[string]string{
		"163.8 x 75 x 8 mm": "১৬৩.৮ x ৭৫ x ৮ মিমি",
		"187 g": "১৮৭ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 MP": "৮ MP",
		"Android 12, OxygenOS 12.1": "Android ১২, OxygenOS ১২.১",
		"August 2022": "August ২০২২",
		"Blue Oasis, Celestial Black": "নীল Oasis, Celestial কালো",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G35": "Helio G৩৫",
		"IPS LCD, 600 nits": "IPS LCD, ৬০০ nits",
		"Mediatek MT6765G Helio G35 (12 nm)": "Mediatek MT৬৭৬৫G Helio G৩৫ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-nord-n20-se'
func (s *SpecificationSeederMobileOneplusNordN20Se) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-n20-se"

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
