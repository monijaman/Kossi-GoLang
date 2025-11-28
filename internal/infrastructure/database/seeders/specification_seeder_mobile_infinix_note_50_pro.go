package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixNote50Pro seeds specifications/options for product 'infinix-note-50-pro'
type SpecificationSeederMobileInfinixNote50Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixNote50Pro creates a new seeder instance
func NewSpecificationSeederMobileInfinixNote50Pro() *SpecificationSeederMobileInfinixNote50Pro {
	return &SpecificationSeederMobileInfinixNote50Pro{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Note 50 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixNote50Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ MP + ২ MP",
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"16 MP": "১৬ MP",
		"168.6 x 76.6 x 8.6 mm": "১৬৮.৬ x ৭৬.৬ x ৮.৬ মিমি",
		"203 g": "২০৩ g",
		"256 GB": "২৫৬ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, XOS 14": "Android ১৪, XOS ১৪",
		"Black, Gold": "কালো, সোনালী",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G99": "Helio G৯৯",
		"IP53": "IP৫৩",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"May 2024": "May ২০২৪",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-note-50-pro'
func (s *SpecificationSeederMobileInfinixNote50Pro) Seed(db *gorm.DB) error {
	productSlug := "infinix-note-50-pro"

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
