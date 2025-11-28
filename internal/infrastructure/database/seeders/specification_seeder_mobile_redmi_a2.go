package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiA2 seeds specifications/options for product 'redmi-a2'
type SpecificationSeederMobileRedmiA2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiA2 creates a new seeder instance
func NewSpecificationSeederMobileRedmiA2() *SpecificationSeederMobileRedmiA2 {
	return &SpecificationSeederMobileRedmiA2{BaseSeeder: BaseSeeder{name: "Specifications for Redmi A2+"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiA2) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10W wired": "১০W তারযুক্ত",
		"192 g": "১৯২ g",
		"2 GB / 3 GB / 4 GB": "২ GB / ৩ GB / ৪ GB",
		"32 GB / 64 GB": "৩২ GB / ৬৪ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP + 0.08 MP": "৮ MP + ০.০৮ MP",
		"Android 12 or 13 (Go edition), MIUI": "Android ১২ or ১৩ (Go edition), MIUI",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Helio G36": "Helio G৩৬",
		"Light Blue, Light Green, Black": "Light নীল, Light সবুজ, কালো",
		"March 2023": "March ২০২৩",
		"Mediatek Helio G36 (12 nm)": "Mediatek Helio G৩৬ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-a2'
func (s *SpecificationSeederMobileRedmiA2) Seed(db *gorm.DB) error {
	productSlug := "redmi-a2"

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
