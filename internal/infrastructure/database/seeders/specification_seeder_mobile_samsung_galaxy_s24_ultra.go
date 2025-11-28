package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS24Ultra seeds specifications/options for product 'samsung-galaxy-s24-ultra'
type SpecificationSeederMobileSamsungGalaxyS24Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS24Ultra creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS24Ultra() *SpecificationSeederMobileSamsungGalaxyS24Ultra {
	return &SpecificationSeederMobileSamsungGalaxyS24Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S24 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS24Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"1440 x 3120 pixels": "১৪৪০ x ৩১২০ pixels",
		"162.3 x 79 x 8.6 mm": "১৬২.৩ x ৭৯ x ৮.৬ মিমি",
		"200 MP + 50 MP + 10 MP + 12 MP": "২০০ MP + ৫০ MP + ১০ MP + ১২ MP",
		"232 g": "২৩২ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, One UI 6.1": "Android ১৪, One UI ৬.১",
		"Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Dynamic LTPO AMOLED ২X, ১২০Hz, HDR১০+",
		"IP68": "IP৬৮",
		"January 2024": "January ২০২৪",
		"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AC Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
		"Titanium Gray, Titanium Black, Titanium Violet, Titanium Yellow": "Titanium ধূসর, Titanium কালো, Titanium Violet, Titanium হলুদ",
		"Titanium frame, glass front/back": "টাইটানিয়াম ফ্রেম, গ্লাস সামনে/back",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s24-ultra'
func (s *SpecificationSeederMobileSamsungGalaxyS24Ultra) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s24-ultra"

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
