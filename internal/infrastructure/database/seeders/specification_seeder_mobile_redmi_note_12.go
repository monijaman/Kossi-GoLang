package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiNote12 seeds specifications/options for product 'redmi-note-12'
type SpecificationSeederMobileRedmiNote12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote12 creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote12() *SpecificationSeederMobileRedmiNote12 {
	return &SpecificationSeederMobileRedmiNote12{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"13 MP": "১৩ MP",
		"165.9 x 76.2 x 8 mm": "১৬৫.৯ x ৭৬.২ x ৮ মিমি",
		"188 g": "১৮৮ g",
		"4 GB / 6 GB / 8 GB": "৪ GB / ৬ GB / ৮ GB",
		"48 MP + 8 MP + 2 MP": "৪৮ MP + ৮ MP + ২ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 120Hz, 1200 nits": "AMOLED, ১২০Hz, ১২০০ nits",
		"Adreno 619": "Adreno ৬১৯",
		"Android 12, আপগ্রেডযোগ্য": "Android ১২, আপগ্রেডযোগ্য",
		"Corning Gorilla Glass 3": "Corning Gorilla Glass ৩",
		"Frosted Green, Matte Black, Mystique Blue": "Frosted সবুজ, Matte কালো, Mystique নীল",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP53": "IP৫৩",
		"October 2022": "October ২০২২",
		"Qualcomm SM4375 Snapdragon 4 Gen 1 (6 nm)": "Qualcomm SM৪৩৭৫ Snapdragon ৪ Gen ১ (৬ nm)",
		"Snapdragon 4 Gen 1": "Snapdragon ৪ Gen ১",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-note-12'
func (s *SpecificationSeederMobileRedmiNote12) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-12"

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
