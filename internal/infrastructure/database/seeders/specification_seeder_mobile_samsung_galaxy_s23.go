package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS23 seeds specifications/options for product 'samsung-galaxy-s23'
type SpecificationSeederMobileSamsungGalaxyS23 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS23 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS23() *SpecificationSeederMobileSamsungGalaxyS23 {
	return &SpecificationSeederMobileSamsungGalaxyS23{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S23"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS23) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"146.3 x 70.9 x 7.6 mm": "১৪৬.৩ x ৭০.৯ x ৭.৬ মিমি",
		"168 g": "১৬৮ g",
		"3,900 mAh": "৩,৯০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ MP + ১০ MP + ১২ MP",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"8 GB": "৮ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 13, upgradable": "Android ১৩, upgradable",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"February 2023": "February ২০২৩",
		"IP68": "IP৬৮",
		"Phantom Black, Green, Cream, Lavender": "Phantom কালো, সবুজ, Cream, Lavender",
		"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AC Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s23'
func (s *SpecificationSeederMobileSamsungGalaxyS23) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s23"

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
