package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileIphone17 seeds specifications/options for product 'iphone-17'
type SpecificationSeederMobileIphone17 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone17 creates a new seeder instance
func NewSpecificationSeederMobileIphone17() *SpecificationSeederMobileIphone17 {
	return &SpecificationSeederMobileIphone17{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 17"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone17) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ MP",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"147.6 x 71.6 x 7.8 mm": "১৪৭.৬ x ৭১.৬ x ৭.৮ মিমি",
		"171 g": "১৭১ g",
		"2556 x 1179 pixels (~461 ppi density)": "২৫৫৬ x ১১৭৯ pixels (~৪৬১ ppi density)",
		"3,349 mAh": "৩,৩৪৯ এমএএইচ",
		"48 MP + 12 MP": "৪৮ MP + ১২ MP",
		"5-core Apple GPU": "৫-core Apple GPU",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 GB": "৮ GB",
		"Apple A19": "Apple A১৯",
		"Apple A19 (3 nm)": "Apple A১৯ (৩ nm)",
		"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপী",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"September 2024": "September ২০২৪",
		"সুপার Retina XDR OLED, HDR10, Dolby Vision": "সুপার Retina XDR OLED, HDR১০, Dolby Vision",
		"iOS 18": "iOS ১৮",
	}
}

// Seed inserts specification_translations for existing specifications for product 'iphone-17'
func (s *SpecificationSeederMobileIphone17) Seed(db *gorm.DB) error {
	productSlug := "iphone-17"

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
