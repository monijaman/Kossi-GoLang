package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS24Fe seeds specifications/options for product 'samsung-galaxy-s24-fe'
type SpecificationSeederMobileSamsungGalaxyS24Fe struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS24Fe creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS24Fe() *SpecificationSeederMobileSamsungGalaxyS24Fe {
	return &SpecificationSeederMobileSamsungGalaxyS24Fe{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S24 FE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS24Fe) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ MP",
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"162 x 77.3 x 8 mm": "১৬২ x ৭৭.৩ x ৮ মিমি",
		"213 g": "২১৩ g",
		"4,700 mAh": "৪,৭০০ এমএএইচ",
		"50 MP + 8 MP + 12 MP": "৫০ MP + ৮ MP + ১২ MP",
		"5G": "৫G",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 14, One UI 6.1": "Android ১৪, One UI ৬.১",
		"Blue, Graphite, Gray, মিন্ট, Yellow": "নীল, Graphite, ধূসর, মিন্ট, হলুদ",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"Exynos 2400e": "Exynos ২৪০০e",
		"Exynos 2400e (4 nm)": "Exynos ২৪০০e (৪ nm)",
		"IP68": "IP৬৮",
		"September 2024": "September ২০২৪",
		"Xclipse 940": "Xclipse ৯৪০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s24-fe'
func (s *SpecificationSeederMobileSamsungGalaxyS24Fe) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s24-fe"

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
