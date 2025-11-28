package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiA3 seeds specifications/options for product 'redmi-a3'
type SpecificationSeederMobileRedmiA3 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiA3 creates a new seeder instance
func NewSpecificationSeederMobileRedmiA3() *SpecificationSeederMobileRedmiA3 {
	return &SpecificationSeederMobileRedmiA3{BaseSeeder: BaseSeeder{name: "Specifications for Redmi A3"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiA3) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10W wired": "১০W তারযুক্ত",
		"193 g": "১৯৩ g",
		"3 GB / 4 GB / 6 GB": "৩ GB / ৪ GB / ৬ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.71 inches": "৬.৭১ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1650 pixels": "৭২০ x ১৬৫০ pixels",
		"8 MP + 0.08 MP": "৮ MP + ০.০৮ MP",
		"90Hz": "৯০Hz",
		"Android 14 (Go edition), MIUI": "Android ১৪ (Go edition), MIUI",
		"February 2024": "February ২০২৪",
		"Glass front (Gorilla Glass 3), glass back or silicone polymer back, plastic frame": "গ্লাস সামনে (Gorilla Glass ৩), গ্লাস পেছনে or silicone polymer back, plastic frame",
		"Helio G36": "Helio G৩৬",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mediatek Helio G36 (12 nm)": "Mediatek Helio G৩৬ (১২ nm)",
		"Midnight Black, Olive Green, Lake Blue": "Midnight কালো, Olive সবুজ, Lake নীল",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-a3'
func (s *SpecificationSeederMobileRedmiA3) Seed(db *gorm.DB) error {
	productSlug := "redmi-a3"

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
