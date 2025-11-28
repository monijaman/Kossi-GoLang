package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS23Ultra seeds specifications/options for product 'samsung-galaxy-s23-ultra'
type SpecificationSeederMobileSamsungGalaxyS23Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS23Ultra creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS23Ultra() *SpecificationSeederMobileSamsungGalaxyS23Ultra {
	return &SpecificationSeederMobileSamsungGalaxyS23Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S23 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS23Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"1440 x 3088 pixels": "১৪৪০ x ৩০৮৮ pixels",
		"163.4 x 78.1 x 8.9 mm": "১৬৩.৪ x ৭৮.১ x ৮.৯ মিমি",
		"200 MP + 10 MP + 10 MP + 12 MP": "২০০ MP + ১০ MP + ১০ MP + ১২ MP",
		"234 g": "২৩৪ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
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

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s23-ultra'
func (s *SpecificationSeederMobileSamsungGalaxyS23Ultra) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s23-ultra"

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
