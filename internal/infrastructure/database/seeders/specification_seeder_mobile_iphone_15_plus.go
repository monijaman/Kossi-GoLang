package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileIphone15Plus seeds specifications/options for product 'iphone-15-plus'
type SpecificationSeederMobileIphone15Plus struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone15Plus creates a new seeder instance
func NewSpecificationSeederMobileIphone15Plus() *SpecificationSeederMobileIphone15Plus {
	return &SpecificationSeederMobileIphone15Plus{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 15 Plus"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone15Plus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ MP",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"1290 x 2796 pixels": "১২৯০ x ২৭৯৬ pixels",
		"160.9 x 77.8 x 7.8 mm": "১৬০.৯ x ৭৭.৮ x ৭.৮ মিমি",
		"201 g": "২০১ g",
		"4,383 mAh": "৪,৩৮৩ এমএএইচ",
		"48 MP + 12 MP": "৪৮ MP + ১২ MP",
		"5-core Apple GPU": "৫-core Apple GPU",
		"5G": "৫G",
		"6 GB": "৬ GB",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"60Hz": "৬০Hz",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Apple A16 Bionic": "Apple A১৬ Bionic",
		"Apple A16 Bionic (4 nm)": "Apple A১৬ Bionic (৪ nm)",
		"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপী",
		"IP68": "IP৬৮",
		"September 2023": "September ২০২৩",
		"iOS 17, আপগ্রেডযোগ্য": "iOS ১৭, আপগ্রেডযোগ্য",
	}
}

// Seed inserts specification_translations for existing specifications for product 'iphone-15-plus'
func (s *SpecificationSeederMobileIphone15Plus) Seed(db *gorm.DB) error {
	productSlug := "iphone-15-plus"

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
