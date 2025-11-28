package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoY19sPro seeds specifications/options for product 'vivo-y19s-pro'
type SpecificationSeederMobileVivoY19sPro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY19sPro creates a new seeder instance
func NewSpecificationSeederMobileVivoY19sPro() *SpecificationSeederMobileVivoY19sPro {
	return &SpecificationSeederMobileVivoY19sPro{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y19s Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY19sPro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"195 g": "১৯৫ g",
		"256 GB": "২৫৬ GB",
		"44W wired": "৪৪W তারযুক্ত",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5500 mAh": "৫৫০০ এমএএইচ",
		"6.68 inches": "৬.৬৮ ইঞ্চি",
		"720 x 1608 pixels": "৭২০ x ১৬০৮ pixels",
		"8 GB": "৮ GB",
		"8 MP": "৮ MP",
		"90Hz": "৯০Hz",
		"Android 14, ফানটাচ 14": "Android ১৪, ফানটাচ ১৪",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Green, Black": "সবুজ, কালো",
		"Helio G85": "Helio G৮৫",
		"IP64": "IP৬৪",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"Mediatek Helio G85 (12nm)": "Mediatek Helio G৮৫ (১২nm)",
		"September 2024": "September ২০২৪",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-y19s-pro'
func (s *SpecificationSeederMobileVivoY19sPro) Seed(db *gorm.DB) error {
	productSlug := "vivo-y19s-pro"

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
