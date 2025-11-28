package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS22 seeds specifications/options for product 'samsung-galaxy-s22'
type SpecificationSeederMobileSamsungGalaxyS22 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS22 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS22() *SpecificationSeederMobileSamsungGalaxyS22 {
	return &SpecificationSeederMobileSamsungGalaxyS22{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S22"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS22) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ MP",
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"146 x 70.6 x 7.6 mm": "১৪৬ x ৭০.৬ x ৭.৬ মিমি",
		"167 g": "১৬৭ g",
		"3,700 mAh": "৩,৭০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ MP + ১০ MP + ১২ MP",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 12, আপগ্রেডযোগ্য": "Android ১২, আপগ্রেডযোগ্য",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"Exynos 2200 (4 nm)": "Exynos ২২০০ (৪ nm)",
		"Exynos 2200 / Snapdragon 8 Gen 1": "Exynos ২২০০ / Snapdragon ৮ Gen ১",
		"February 2022": "February ২০২২",
		"IP68": "IP৬৮",
		"Phantom Black, White, Pink Gold, Green": "Phantom কালো, সাদা, গোলাপী সোনালী, সবুজ",
		"Xclipse 920": "Xclipse ৯২০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s22'
func (s *SpecificationSeederMobileSamsungGalaxyS22) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s22"

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
