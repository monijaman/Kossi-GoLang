package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoSpark40Pro seeds specifications/options for product 'tecno-spark-40-pro'
type SpecificationSeederMobileTecnoSpark40Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark40Pro creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark40Pro() *SpecificationSeederMobileTecnoSpark40Pro {
	return &SpecificationSeederMobileTecnoSpark40Pro{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 40 Pro+"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark40Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + AI": "১০৮ MP + ২ MP + AI",
		"1080 x 2436 pixels": "১০৮০ x ২৪৩৬ pixels",
		"120Hz": "১২০Hz",
		"164.7 x 75 x 7.6 mm": "১৬৪.৭ x ৭৫ x ৭.৬ মিমি",
		"179 g": "১৭৯ g",
		"256 GB": "২৫৬ GB",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz, Curved": "AMOLED, ১২০Hz, Curved",
		"Android 14, HIOS 14": "Android ১৪, HIOS ১৪",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"Helio G99 Ultimate": "Helio G৯৯ Ultimate",
		"IP53": "IP৫৩",
		"January 2024": "January ২০২৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Helio G99 Ultimate": "Mediatek Helio G৯৯ Ultimate",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-spark-40-pro'
func (s *SpecificationSeederMobileTecnoSpark40Pro) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-40-pro"

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
