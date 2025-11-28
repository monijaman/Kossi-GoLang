package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyZFlip5 seeds specifications/options for product 'samsung-galaxy-z-flip-5'
type SpecificationSeederMobileSamsungGalaxyZFlip5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip5 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip5() *SpecificationSeederMobileSamsungGalaxyZFlip5 {
	return &SpecificationSeederMobileSamsungGalaxyZFlip5{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ MP",
		"1080 x 2640 pixels": "১০৮০ x ২৬৪০ pixels",
		"12 MP + 12 MP": "১২ MP + ১২ MP",
		"120Hz": "১২০Hz",
		"187 g": "১৮৭ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"3,700 mAh": "৩,৭০০ এমএএইচ",
		"5G": "৫G",
		"6.7 inches (Main) / 3.4 inches (Cover)": "৬.৭ ইঞ্চি (Main) / ৩.৪ ইঞ্চি (Cover)",
		"8 GB": "৮ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, আপগ্রেডযোগ্য": "Android ১৩, আপগ্রেডযোগ্য",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Foldable Dynamic AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"IPX8": "IPX৮",
		"July 2023": "July ২০২৩",
		"মিন্ট, Graphite, Cream, Lavender, Gray, Blue, Green, Yellow": "মিন্ট, Graphite, Cream, Lavender, ধূসর, নীল, সবুজ, হলুদ",
		"Plastic front (open), glass back (Gorilla Glass Victus 2), aluminum frame": "Plastic front (open), গ্লাস পেছনে (Gorilla Glass Victus ২), অ্যালুমিনিয়াম ফ্রেম",
		"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AC Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
		"Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 15.1 mm": "Unfolded: ১৬৫.১ x ৭১.৯ x ৬.৯ মিমি / Folded: ৮৫.১ x ৭১.৯ x ১৫.১ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-z-flip-5'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip5) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-5"

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
