package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileIphone15 seeds specifications/options for product 'iphone-15'
type SpecificationSeederMobileIphone15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone15 creates a new seeder instance
func NewSpecificationSeederMobileIphone15() *SpecificationSeederMobileIphone15 {
	return &SpecificationSeederMobileIphone15{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1179 x 2556 pixels": "১১৭৯ x ২৫৫৬ pixels",
		"12 MP": "১২ MP",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"147.6 x 71.6 x 7.8 mm": "১৪৭.৬ x ৭১.৬ x ৭.৮ মিমি",
		"171 g": "১৭১ g",
		"3,349 mAh": "৩,৩৪৯ এমএএইচ",
		"48 MP + 12 MP": "৪৮ MP + ১২ MP",
		"5-core Apple GPU": "৫-core Apple GPU",
		"5G": "৫G",
		"6 GB": "৬ GB",
		"6.1 inches": "৬.১ ইঞ্চি",
		"60Hz": "৬০Hz",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Apple A16 Bionic": "Apple A১৬ Bionic",
		"Apple A16 Bionic (4 nm)": "Apple A১৬ Bionic (৪ nm)",
		"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপী",
		"IP68": "IP৬৮",
		"September 2023": "September ২০২৩",
		"iOS 17, upgradable": "iOS ১৭, upgradable",
	}
}

// Seed inserts specification_translations for existing specifications for product 'iphone-15'
func (s *SpecificationSeederMobileIphone15) Seed(db *gorm.DB) error {
	productSlug := "iphone-15"

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
