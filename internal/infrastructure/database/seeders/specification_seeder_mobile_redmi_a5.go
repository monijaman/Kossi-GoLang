package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiA5 seeds specifications/options for product 'redmi-a5'
type SpecificationSeederMobileRedmiA5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiA5 creates a new seeder instance
func NewSpecificationSeederMobileRedmiA5() *SpecificationSeederMobileRedmiA5 {
	return &SpecificationSeederMobileRedmiA5{BaseSeeder: BaseSeeder{name: "Specifications for Redmi A5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiA5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10W wired": "১০W তারযুক্ত",
		"199 g": "১৯৯ g",
		"3 GB / 4 GB": "৩ GB / ৪ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1650 pixels": "৭২০ x ১৬৫০ pixels",
		"8 MP + 0.08 MP": "৮ MP + ০.০৮ MP",
		"90Hz": "৯০Hz",
		"Android 14 (Go এডিশন)": "Android ১৪ (Go এডিশন)",
		"February 2025": "February ২০২৫",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Helio G36": "Helio G৩৬",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mediatek Helio G36 (12 nm)": "Mediatek Helio G৩৬ (১২ nm)",
		"মিডনাইট Black, Lake Blue, Olive Green": "মিডনাইট কালো, Lake নীল, Olive সবুজ",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-a5'
func (s *SpecificationSeederMobileRedmiA5) Seed(db *gorm.DB) error {
	productSlug := "redmi-a5"

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
