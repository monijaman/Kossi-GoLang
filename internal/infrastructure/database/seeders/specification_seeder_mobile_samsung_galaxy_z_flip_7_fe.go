package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyZFlip7Fe seeds specifications/options for product 'samsung-galaxy-z-flip-7-fe'
type SpecificationSeederMobileSamsungGalaxyZFlip7Fe struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip7Fe creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip7Fe() *SpecificationSeederMobileSamsungGalaxyZFlip7Fe {
	return &SpecificationSeederMobileSamsungGalaxyZFlip7Fe{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 7 FE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7Fe) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ MP",
		"1080 x 2640 pixels": "১০৮০ x ২৬৪০ pixels",
		"12 MP + 12 MP": "১২ MP + ১২ MP",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"190 g": "১৯০ g",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"5G": "৫G",
		"6.7 inches (Main) / 3.4 inches (Cover)": "৬.৭ ইঞ্চি (Main) / ৩.৪ ইঞ্চি (Cover)",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 15, One UI 7": "Android ১৫, One UI ৭",
		"Exynos 2400e": "Exynos ২৪০০e",
		"Exynos 2400e (4 nm)": "Exynos ২৪০০e (৪ nm)",
		"Foldable Dynamic AMOLED 2X, 120Hz": "Foldable Dynamic AMOLED ২X, ১২০Hz",
		"IPX8": "IPX৮",
		"July 2025": "July ২০২৫",
		"Unfolded: 165.1 x 71.9 x 6.9 mm": "Unfolded: ১৬৫.১ x ৭১.৯ x ৬.৯ মিমি",
		"Xclipse 940": "Xclipse ৯৪০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-z-flip-7-fe'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7Fe) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-7-fe"

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
