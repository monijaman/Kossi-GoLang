package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileIphone16ProMax seeds specifications/options for product 'iphone-16-pro-max'
type SpecificationSeederMobileIphone16ProMax struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone16ProMax creates a new seeder instance
func NewSpecificationSeederMobileIphone16ProMax() *SpecificationSeederMobileIphone16ProMax {
	return &SpecificationSeederMobileIphone16ProMax{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 16 Pro Max"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone16ProMax) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"1320 x 2868 pixels": "১৩২০ x ২৮৬৮ pixels",
		"163.0 x 77.6 x 8.25 mm": "১৬৩.০ x ৭৭.৬ x ৮.২৫ মিমি",
		"225 g": "২২৫ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"4,676 mAh": "৪,৬৭৬ এমএএইচ",
		"48 MP + 48 MP + 12 MP": "৪৮ MP + ৪৮ MP + ১২ MP",
		"5G": "৫G",
		"6-core Apple GPU": "৬-core Apple GPU",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"8 GB": "৮ GB",
		"Apple A18 Pro": "Apple A১৮ Pro",
		"Apple A18 Pro (3 nm)": "Apple A১৮ Pro (৩ nm)",
		"Black Titanium, White Titanium, Natural Titanium, Desert Titanium": "কালো Titanium, সাদা Titanium, ন্যাচারাল Titanium, মরুভূমি Titanium",
		"IP68": "IP৬৮",
		"LTPO সুপার Retina XDR OLED, 120Hz": "LTPO সুপার Retina XDR OLED, ১২০Hz",
		"September 2024": "September ২০২৪",
		"Titanium frame, glass front/back": "টাইটানিয়াম ফ্রেম, গ্লাস সামনে/back",
		"iOS 18": "iOS ১৮",
	}
}

// Seed inserts specification_translations for existing specifications for product 'iphone-16-pro-max'
func (s *SpecificationSeederMobileIphone16ProMax) Seed(db *gorm.DB) error {
	productSlug := "iphone-16-pro-max"

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
