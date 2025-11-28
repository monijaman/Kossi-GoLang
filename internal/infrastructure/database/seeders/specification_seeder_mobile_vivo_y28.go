package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoY28 seeds specifications/options for product 'vivo-y28'
type SpecificationSeederMobileVivoY28 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY28 creates a new seeder instance
func NewSpecificationSeederMobileVivoY28() *SpecificationSeederMobileVivoY28 {
	return &SpecificationSeederMobileVivoY28{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y28"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY28) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"199 g": "১৯৯ g",
		"44W wired": "৪৪W তারযুক্ত",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.68 inches": "৬.৬৮ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"720 x 1608 pixels": "৭২০ x ১৬০৮ pixels",
		"8 GB": "৮ GB",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Agate Green, Gleaming Orange": "Agate সবুজ, Gleaming Orange",
		"Android 14, Funtouch 14": "Android ১৪, Funtouch ১৪",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Helio G85": "Helio G৮৫",
		"IP64": "IP৬৪",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"January 2024": "January ২০২৪",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G85 (12nm)": "Mediatek Helio G৮৫ (১২nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-y28'
func (s *SpecificationSeederMobileVivoY28) Seed(db *gorm.DB) error {
	productSlug := "vivo-y28"

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
