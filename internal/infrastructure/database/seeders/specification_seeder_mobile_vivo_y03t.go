package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoY03t seeds specifications/options for product 'vivo-y03t'
type SpecificationSeederMobileVivoY03t struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY03t creates a new seeder instance
func NewSpecificationSeederMobileVivoY03t() *SpecificationSeederMobileVivoY03t {
	return &SpecificationSeederMobileVivoY03t{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y03t"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY03t) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 0.08 MP": "১৩ MP + ০.০৮ MP",
		"15W wired": "১৫W তারযুক্ত",
		"185 g": "১৮৫ g",
		"4 GB": "৪ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"90Hz": "৯০Hz",
		"Android 14, Funtouch 14": "Android ১৪, Funtouch ১৪",
		"August 2024": "August ২০২৪",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"IP54": "IP৫৪",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G57": "Mali-G৫৭",
		"Space Black, Gem Green": "Space কালো, Gem সবুজ",
		"Unisoc T612": "Unisoc T৬১২",
		"Unisoc T612 (12 nm)": "Unisoc T৬১২ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-y03t'
func (s *SpecificationSeederMobileVivoY03t) Seed(db *gorm.DB) error {
	productSlug := "vivo-y03t"

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
