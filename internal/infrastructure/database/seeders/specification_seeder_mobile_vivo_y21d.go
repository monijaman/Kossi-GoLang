package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoY21d seeds specifications/options for product 'vivo-y21d'
type SpecificationSeederMobileVivoY21d struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY21d creates a new seeder instance
func NewSpecificationSeederMobileVivoY21d() *SpecificationSeederMobileVivoY21d {
	return &SpecificationSeederMobileVivoY21d{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y21d"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY21d) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"182 g": "১৮২ g",
		"18W wired": "১৮W তারযুক্ত",
		"4 GB": "৪ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.51 inches": "৬.৫১ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB": "৬৪ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 11, ফানটাচ 11.1": "Android ১১, ফানটাচ ১১.১",
		"August 2021": "August ২০২১",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Helio P35": "Helio P৩৫",
		"Mediatek Helio P35 (12nm)": "Mediatek Helio P৩৫ (১২nm)",
		"মিডনাইট Blue, Diamond Glow": "মিডনাইট নীল, Diamond Glow",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-y21d'
func (s *SpecificationSeederMobileVivoY21d) Seed(db *gorm.DB) error {
	productSlug := "vivo-y21d"

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
