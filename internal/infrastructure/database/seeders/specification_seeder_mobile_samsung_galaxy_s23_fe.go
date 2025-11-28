package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS23Fe seeds specifications/options for product 'samsung-galaxy-s23-fe'
type SpecificationSeederMobileSamsungGalaxyS23Fe struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS23Fe creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS23Fe() *SpecificationSeederMobileSamsungGalaxyS23Fe {
	return &SpecificationSeederMobileSamsungGalaxyS23Fe{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S23 FE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS23Fe) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ MP",
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"158 x 76.5 x 8.2 mm": "১৫৮ x ৭৬.৫ x ৮.২ মিমি",
		"209 g": "২০৯ g",
		"4,500 mAh": "৪,৫০০ এমএএইচ",
		"50 MP + 8 MP + 12 MP": "৫০ MP + ৮ MP + ১২ MP",
		"5G": "৫G",
		"6.4 inches": "৬.৪ ইঞ্চি",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 13, আপগ্রেডযোগ্য": "Android ১৩, আপগ্রেডযোগ্য",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"Exynos 2200 (4 nm)": "Exynos ২২০০ (৪ nm)",
		"Exynos 2200 / Snapdragon 8 Gen 1": "Exynos ২২০০ / Snapdragon ৮ Gen ১",
		"IP68": "IP৬৮",
		"মিন্ট, Cream, Graphite, Purple, Indigo, Tangerine": "মিন্ট, Cream, Graphite, বেগুনি, Indigo, Tangerine",
		"October 2023": "October ২০২৩",
		"Xclipse 920": "Xclipse ৯২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s23-fe'
func (s *SpecificationSeederMobileSamsungGalaxyS23Fe) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s23-fe"

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
