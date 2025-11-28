package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixNote50 seeds specifications/options for product 'infinix-note-50'
type SpecificationSeederMobileInfinixNote50 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixNote50 creates a new seeder instance
func NewSpecificationSeederMobileInfinixNote50() *SpecificationSeederMobileInfinixNote50 {
	return &SpecificationSeederMobileInfinixNote50{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Note 50"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixNote50) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"168 x 76.6 x 8.5 mm": "১৬৮ x ৭৬.৬ x ৮.৫ মিমি",
		"198 g": "১৯৮ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"8 GB": "৮ GB",
		"Android 14": "Android ১৪",
		"Black, Blue": "কালো, নীল",
		"Helio G99": "Helio G৯৯",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"May 2024": "May ২০২৪",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-note-50'
func (s *SpecificationSeederMobileInfinixNote50) Seed(db *gorm.DB) error {
	productSlug := "infinix-note-50"

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
