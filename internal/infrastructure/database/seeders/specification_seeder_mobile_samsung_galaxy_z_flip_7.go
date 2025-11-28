package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyZFlip7 seeds specifications/options for product 'samsung-galaxy-z-flip-7'
type SpecificationSeederMobileSamsungGalaxyZFlip7 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip7 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip7() *SpecificationSeederMobileSamsungGalaxyZFlip7 {
	return &SpecificationSeederMobileSamsungGalaxyZFlip7{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 7"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2640 pixels (Main)": "১০৮০ x ২৬৪০ pixels (Main)",
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"187 g": "১৮৭ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"4,200 mAh": "৪,২০০ এমএএইচ",
		"50 MP + 12 MP": "৫০ MP + ১২ MP",
		"5G": "৫G",
		"6.7 inches (Main) / 3.9 inches (Cover)": "৬.৭ ইঞ্চি (Main) / ৩.৯ ইঞ্চি (Cover)",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, One UI 7.1.1": "Android ১৫, One UI ৭.১.১",
		"Corning Gorilla Armor 2": "Corning Gorilla Armor ২",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic LTPO AMOLED ২X, ১২০Hz, HDR১০+",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IPX8": "IPX৮",
		"July 2025": "July ২০২৫",
		"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "Qualcomm SM৮৭৫০-AC Snapdragon ৮ Elite (৩ nm)",
		"Snapdragon 8 Elite": "Snapdragon ৮ Elite",
		"Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 14.9 mm": "Unfolded: ১৬৫.১ x ৭১.৯ x ৬.৯ মিমি / Folded: ৮৫.১ x ৭১.৯ x ১৪.৯ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-z-flip-7'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-7"

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
