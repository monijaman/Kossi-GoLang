package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileIphone15ProMax seeds specifications/options for product 'iphone-15-pro-max'
type SpecificationSeederMobileIphone15ProMax struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone15ProMax creates a new seeder instance
func NewSpecificationSeederMobileIphone15ProMax() *SpecificationSeederMobileIphone15ProMax {
	return &SpecificationSeederMobileIphone15ProMax{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 15 Pro Max"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone15ProMax) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"1290 x 2796 pixels": "১২৯০ x ২৭৯৬ pixels",
		"159.9 x 76.7 x 8.3 mm": "১৫৯.৯ x ৭৬.৭ x ৮.৩ মিমি",
		"221 g": "২২১ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"4,441 mAh": "৪,৪৪১ এমএএইচ",
		"48 MP + 12 MP + 12 MP": "৪৮ MP + ১২ MP + ১২ MP",
		"5G": "৫G",
		"6-core Apple GPU": "৬-core Apple GPU",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"Apple A17 Pro": "Apple A১৭ Pro",
		"Apple A17 Pro (3 nm)": "Apple A১৭ Pro (৩ nm)",
		"Black Titanium, White Titanium, Blue Titanium, Natural Titanium": "কালো Titanium, সাদা Titanium, নীল Titanium, ন্যাচারাল Titanium",
		"IP68": "IP৬৮",
		"LTPO সুপার Retina XDR OLED, 120Hz": "LTPO সুপার Retina XDR OLED, ১২০Hz",
		"September 2023": "September ২০২৩",
		"Titanium frame, glass front/back": "টাইটানিয়াম ফ্রেম, গ্লাস সামনে/back",
		"iOS 17, আপগ্রেডযোগ্য": "iOS ১৭, আপগ্রেডযোগ্য",
	}
}

// Seed inserts specification_translations for existing specifications for product 'iphone-15-pro-max'
func (s *SpecificationSeederMobileIphone15ProMax) Seed(db *gorm.DB) error {
	productSlug := "iphone-15-pro-max"

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
