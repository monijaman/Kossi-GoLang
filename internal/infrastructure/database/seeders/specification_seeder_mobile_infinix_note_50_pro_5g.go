package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixNote50Pro5g seeds specifications/options for product 'infinix-note-50-pro-5g'
type SpecificationSeederMobileInfinixNote50Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixNote50Pro5g creates a new seeder instance
func NewSpecificationSeederMobileInfinixNote50Pro5g() *SpecificationSeederMobileInfinixNote50Pro5g {
	return &SpecificationSeederMobileInfinixNote50Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Note 50 Pro+ 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixNote50Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + 2 MP": "১০৮ MP + ২ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"162.7 x 75.9 x 8.1 mm": "১৬২.৭ x ৭৫.৯ x ৮.১ মিমি",
		"203 g": "২০৩ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"AMOLED, 120Hz, 1B colors": "AMOLED, ১২০Hz, ১B colors",
		"Android 14, XOS 14": "Android ১৪, XOS ১৪",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 8050": "Dimensity ৮০৫০",
		"Glass front, plastic frame, glass back": "গ্লাস সামনে, plastic frame, গ্লাস পেছনে",
		"IP53": "IP৫৩",
		"Magic Black, Variable Gold": "Magic কালো, Variable সোনালী",
		"Mali-G77 MC9": "Mali-G৭৭ MC৯",
		"May 2024": "May ২০২৪",
		"Mediatek Dimensity 8050 (6 nm)": "Mediatek Dimensity ৮০৫০ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-note-50-pro-5g'
func (s *SpecificationSeederMobileInfinixNote50Pro5g) Seed(db *gorm.DB) error {
	productSlug := "infinix-note-50-pro-5g"

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
