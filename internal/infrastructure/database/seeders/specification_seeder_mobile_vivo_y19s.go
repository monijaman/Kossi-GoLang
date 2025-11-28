package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoY19s seeds specifications/options for product 'vivo-y19s'
type SpecificationSeederMobileVivoY19s struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY19s creates a new seeder instance
func NewSpecificationSeederMobileVivoY19s() *SpecificationSeederMobileVivoY19s {
	return &SpecificationSeederMobileVivoY19s{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y19s"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY19s) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"15W wired": "১৫W তারযুক্ত",
		"198 g": "১৯৮ g",
		"5 MP": "৫ MP",
		"50 MP + 0.08 MP": "৫০ MP + ০.০৮ MP",
		"5500 mAh": "৫৫০০ এমএএইচ",
		"6 GB": "৬ GB",
		"6.68 inches": "৬.৬৮ ইঞ্চি",
		"720 x 1608 pixels": "৭২০ x ১৬০৮ pixels",
		"90Hz": "৯০Hz",
		"Android 14, ফানটাচ 14": "Android ১৪, ফানটাচ ১৪",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"IP64": "IP৬৪",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G57": "Mali-G৫৭",
		"October 2024": "October ২০২৪",
		"Pearl Silver, Glacier Blue, Diamond Black": "Pearl রূপালী, Glacier নীল, Diamond কালো",
		"Unisoc T612": "Unisoc T৬১২",
		"Unisoc T612 (12 nm)": "Unisoc T৬১২ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-y19s'
func (s *SpecificationSeederMobileVivoY19s) Seed(db *gorm.DB) error {
	productSlug := "vivo-y19s"

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
