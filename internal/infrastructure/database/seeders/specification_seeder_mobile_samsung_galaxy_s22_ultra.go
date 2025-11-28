package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS22Ultra seeds specifications/options for product 'samsung-galaxy-s22-ultra'
type SpecificationSeederMobileSamsungGalaxyS22Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS22Ultra creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS22Ultra() *SpecificationSeederMobileSamsungGalaxyS22Ultra {
	return &SpecificationSeederMobileSamsungGalaxyS22Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S22 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS22Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 10 MP + 10 MP + 12 MP": "১০৮ MP + ১০ MP + ১০ MP + ১২ MP",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB / 1 টিবি": "১২৮ GB / ২৫৬ GB / ৫১২ GB / ১ টিবি",
		"1440 x 3088 pixels": "১৪৪০ x ৩০৮৮ pixels",
		"163.3 x 77.9 x 8.9 mm": "১৬৩.৩ x ৭৭.৯ x ৮.৯ মিমি",
		"228 g": "২২৮ g",
		"40 MP": "৪০ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 12, আপগ্রেডযোগ্য": "Android ১২, আপগ্রেডযোগ্য",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"Exynos 2200 (4 nm)": "Exynos ২২০০ (৪ nm)",
		"Exynos 2200 / Snapdragon 8 Gen 1": "Exynos ২২০০ / Snapdragon ৮ Gen ১",
		"February 2022": "February ২০২২",
		"IP68": "IP৬৮",
		"Phantom Black, White, Burgundy, Green": "Phantom কালো, সাদা, Burgundy, সবুজ",
		"Xclipse 920": "Xclipse ৯২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s22-ultra'
func (s *SpecificationSeederMobileSamsungGalaxyS22Ultra) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s22-ultra"

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
