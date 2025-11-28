package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplusNordN30Se5g seeds specifications/options for product 'oneplus-nord-n30-se-5g'
type SpecificationSeederMobileOneplusNordN30Se5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNordN30Se5g creates a new seeder instance
func NewSpecificationSeederMobileOneplusNordN30Se5g() *SpecificationSeederMobileOneplusNordN30Se5g {
	return &SpecificationSeederMobileOneplusNordN30Se5g{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord N30 SE 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNordN30Se5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"128 GB": "১২৮ GB",
		"165.6 x 76 x 8 mm": "১৬৫.৬ x ৭৬ x ৮ মিমি",
		"193 g": "১৯৩ g",
		"4 GB": "৪ GB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6.72 inches": "৬.৭২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 MP": "৮ MP",
		"Android 13, OxygenOS 13.1": "Android ১৩, OxygenOS ১৩.১",
		"Black Satin, Cyan Sparkle": "কালো Satin, Cyan Sparkle",
		"Dimensity 6020": "Dimensity ৬০২০",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"January 2024": "January ২০২৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6020 (7 nm)": "Mediatek Dimensity ৬০২০ (৭ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-nord-n30-se-5g'
func (s *SpecificationSeederMobileOneplusNordN30Se5g) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-n30-se-5g"

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
