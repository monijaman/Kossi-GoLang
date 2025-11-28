package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixNote40Pro seeds specifications/options for product 'infinix-note-40-pro'
type SpecificationSeederMobileInfinixNote40Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixNote40Pro creates a new seeder instance
func NewSpecificationSeederMobileInfinixNote40Pro() *SpecificationSeederMobileInfinixNote40Pro {
	return &SpecificationSeederMobileInfinixNote40Pro{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Note 40 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixNote40Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + 2 MP": "১০৮ MP + ২ MP + ২ MP",
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ pixels",
		"120Hz": "১২০Hz",
		"164.4 x 74.6 x 7.8 mm": "১৬৪.৪ x ৭৪.৬ x ৭.৮ মিমি",
		"190 g": "১৯০ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz, 1300 nits": "AMOLED, ১২০Hz, ১৩০০ nits",
		"Android 14, XOS 14": "Android ১৪, XOS ১৪",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G99 Ultimate": "Helio G৯৯ Ultimate",
		"IP54": "IP৫৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"March 2024": "March ২০২৪",
		"Mediatek Helio G99 Ultimate": "Mediatek Helio G৯৯ Ultimate",
		"Vintage Green, Titan Gold": "Vintage সবুজ, Titan সোনালী",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-note-40-pro'
func (s *SpecificationSeederMobileInfinixNote40Pro) Seed(db *gorm.DB) error {
	productSlug := "infinix-note-40-pro"

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
