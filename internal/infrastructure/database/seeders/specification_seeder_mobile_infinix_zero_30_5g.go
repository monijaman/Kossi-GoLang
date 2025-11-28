package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileInfinixZero305g seeds specifications/options for product 'infinix-zero-30-5g'
type SpecificationSeederMobileInfinixZero305g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixZero305g creates a new seeder instance
func NewSpecificationSeederMobileInfinixZero305g() *SpecificationSeederMobileInfinixZero305g {
	return &SpecificationSeederMobileInfinixZero305g{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Zero 30 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixZero305g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 13 MP + 2 MP": "১০৮ MP + ১৩ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"144Hz": "১৪৪Hz",
		"164.5 x 75 x 7.9 mm": "১৬৪.৫ x ৭৫ x ৭.৯ মিমি",
		"185 g": "১৮৫ g",
		"256 GB": "২৫৬ GB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP": "৫০ MP",
		"5G": "৫G",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 144Hz, Curved": "AMOLED, ১৪৪Hz, Curved",
		"Android 13, XOS 13": "Android ১৩, XOS ১৩",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 8020": "Dimensity ৮০২০",
		"Glass front, glass back, plastic frame": "গ্লাস সামনে, গ্লাস পেছনে, plastic frame",
		"IP53": "IP৫৩",
		"Mali-G77 MC9": "Mali-G৭৭ MC৯",
		"Mediatek Dimensity 8020 (6 nm)": "Mediatek Dimensity ৮০২০ (৬ nm)",
		"Rome Green, Golden Hour, Fantasy Purple": "Rome সবুজ, Golden Hour, Fantasy বেগুনি",
		"September 2023": "September ২০২৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'infinix-zero-30-5g'
func (s *SpecificationSeederMobileInfinixZero305g) Seed(db *gorm.DB) error {
	productSlug := "infinix-zero-30-5g"

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
