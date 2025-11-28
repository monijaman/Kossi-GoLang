package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoY400 seeds specifications/options for product 'vivo-y400'
type SpecificationSeederMobileVivoY400 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY400 creates a new seeder instance
func NewSpecificationSeederMobileVivoY400() *SpecificationSeederMobileVivoY400 {
	return &SpecificationSeederMobileVivoY400{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y400"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY400) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"190 g": "১৯০ g",
		"44W wired": "৪৪W তারযুক্ত",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.68 inches": "৬.৬৮ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"720 x 1608 pixels": "৭২০ x ১৬০৮ pixels",
		"8 GB": "৮ GB",
		"8 MP": "৮ MP",
		"Adreno 613": "Adreno ৬১৩",
		"Android 14, ফানটাচ 14": "Android ১৪, ফানটাচ ১৪",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"IP54": "IP৫৪",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"November 2024": "November ২০২৪",
		"Purple, Green": "বেগুনি, সবুজ",
		"Snapdragon 4 Gen 2": "Snapdragon ৪ Gen ২",
		"Snapdragon 4 Gen 2 (4 nm)": "Snapdragon ৪ Gen ২ (৪ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-y400'
func (s *SpecificationSeederMobileVivoY400) Seed(db *gorm.DB) error {
	productSlug := "vivo-y400"

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
