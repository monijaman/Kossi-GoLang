package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplus12 seeds specifications/options for product 'oneplus-12'
type SpecificationSeederMobileOneplus12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus12 creates a new seeder instance
func NewSpecificationSeederMobileOneplus12() *SpecificationSeederMobileOneplus12 {
	return &SpecificationSeederMobileOneplus12{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB / 24 GB": "১২ GB / ১৬ GB / ২৪ GB",
		"120Hz": "১২০Hz",
		"1440 x 3168 pixels": "১৪৪০ x ৩১৬৮ pixels",
		"164.3 x 75.8 x 9.2 mm": "১৬৪.৩ x ৭৫.৮ x ৯.২ মিমি",
		"220 g": "২২০ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"32 MP": "৩২ MP",
		"5,400 mAh": "৫,৪০০ এমএএইচ",
		"50 MP + 64 MP + 48 MP": "৫০ MP + ৬৪ MP + ৪৮ MP",
		"5G": "৫G",
		"6.82 inches": "৬.৮২ ইঞ্চি",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, OxygenOS 14": "Android ১৪, OxygenOS ১৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"December 2023": "December ২০২৩",
		"Flowy Emerald, Silky Black, Silver": "Flowy Emerald, Silky কালো, রূপালী",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP65": "IP৬৫",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 4500 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ৪৫০০ nits",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AB Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-12'
func (s *SpecificationSeederMobileOneplus12) Seed(db *gorm.DB) error {
	productSlug := "oneplus-12"

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
