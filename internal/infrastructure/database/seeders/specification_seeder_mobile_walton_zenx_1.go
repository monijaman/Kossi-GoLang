package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileWaltonZenx1 seeds specifications/options for product 'walton-zenx-1'
type SpecificationSeederMobileWaltonZenx1 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonZenx1 creates a new seeder instance
func NewSpecificationSeederMobileWaltonZenx1() *SpecificationSeederMobileWaltonZenx1 {
	return &SpecificationSeederMobileWaltonZenx1{BaseSeeder: BaseSeeder{name: "Specifications for Walton ZENX 1"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonZenx1) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ pixels",
		"120Hz": "১২০Hz",
		"128 GB": "১২৮ GB",
		"16 MP": "১৬ MP",
		"168.5 x 76.5 x 9.2 mm": "১৬৮.৫ x ৭৬.৫ x ৯.২ মিমি",
		"205 g": "২০৫ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"8 GB": "৮ GB",
		"Adreno 619": "Adreno ৬১৯",
		"Android 13": "Android ১৩",
		"Corning Gorilla Glass 3": "Corning Gorilla Glass ৩",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"IPS LCD, 120Hz": "IPS LCD, ১২০Hz",
		"May 2023": "May ২০২৩",
		"Qualcomm Snapdragon 695 5G (6 nm)": "Qualcomm Snapdragon ৬৯৫ ৫G (৬ nm)",
		"Silver, Black": "রূপালী, কালো",
		"Snapdragon 695": "Snapdragon ৬৯৫",
	}
}

// Seed inserts specification_translations for existing specifications for product 'walton-zenx-1'
func (s *SpecificationSeederMobileWaltonZenx1) Seed(db *gorm.DB) error {
	productSlug := "walton-zenx-1"

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
