package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyZFlip6 seeds specifications/options for product 'samsung-galaxy-z-flip-6'
type SpecificationSeederMobileSamsungGalaxyZFlip6 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip6 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip6() *SpecificationSeederMobileSamsungGalaxyZFlip6 {
	return &SpecificationSeederMobileSamsungGalaxyZFlip6{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 6"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip6) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ MP",
		"1080 x 2640 pixels": "১০৮০ x ২৬৪০ pixels",
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"187 g": "১৮৭ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"50 MP + 12 MP": "৫০ MP + ১২ MP",
		"5G": "৫G",
		"6.7 inches (Main) / 3.4 inches (Cover)": "৬.৭ ইঞ্চি (Main) / ৩.৪ ইঞ্চি (Cover)",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, One UI 6.1.1": "Android ১৪, One UI ৬.১.১",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic LTPO AMOLED ২X, ১২০Hz, HDR১০+",
		"IP48": "IP৪৮",
		"July 2024": "July ২০২৪",
		"Plastic front (open), glass back (Gorilla Glass Victus 2), aluminum frame": "Plastic front (open), গ্লাস পেছনে (Gorilla Glass Victus ২), অ্যালুমিনিয়াম ফ্রেম",
		"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AC Snapdragon ৮ Gen ৩ (৪ nm)",
		"Silver Shadow, Yellow, Blue, মিন্ট, Crafted Black, White, Peach": "রূপালী Shadow, হলুদ, নীল, মিন্ট, Crafted কালো, সাদা, Peach",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
		"Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 14.9 mm": "Unfolded: ১৬৫.১ x ৭১.৯ x ৬.৯ মিমি / Folded: ৮৫.১ x ৭১.৯ x ১৪.৯ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-z-flip-6'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip6) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-6"

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
