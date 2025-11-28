package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplus11 seeds specifications/options for product 'oneplus-11'
type SpecificationSeederMobileOneplus11 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus11 creates a new seeder instance
func NewSpecificationSeederMobileOneplus11() *SpecificationSeederMobileOneplus11 {
	return &SpecificationSeederMobileOneplus11{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 11"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus11) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"1440 x 3216 pixels": "১৪৪০ x ৩২১৬ pixels",
		"16 MP": "১৬ MP",
		"163.1 x 74.1 x 8.5 mm": "১৬৩.১ x ৭৪.১ x ৮.৫ মিমি",
		"205 g": "২০৫ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 32 MP + 48 MP": "৫০ MP + ৩২ MP + ৪৮ MP",
		"5G": "৫G",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB / 12 GB / 16 GB": "৮ GB / ১২ GB / ১৬ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, OxygenOS 13": "Android ১৩, OxygenOS ১৩",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP64": "IP৬৪",
		"January 2023": "January ২০২৩",
		"LTPO3 Fluid AMOLED, 120Hz, Dolby Vision, HDR10+": "LTPO৩ Fluid AMOLED, ১২০Hz, Dolby Vision, HDR১০+",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AB Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
		"Titan Black, Eternal Green": "Titan কালো, Eternal সবুজ",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-11'
func (s *SpecificationSeederMobileOneplus11) Seed(db *gorm.DB) error {
	productSlug := "oneplus-11"

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
