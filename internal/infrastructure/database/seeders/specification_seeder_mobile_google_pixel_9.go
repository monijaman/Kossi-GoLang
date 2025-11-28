package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileGooglePixel9 seeds specifications/options for product 'google-pixel-9'
type SpecificationSeederMobileGooglePixel9 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel9 creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel9() *SpecificationSeederMobileGooglePixel9 {
	return &SpecificationSeederMobileGooglePixel9{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 9"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel9) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10.5 MP": "১০.৫ MP",
		"1080 x 2424 pixels": "১০৮০ x ২৪২৪ pixels",
		"12 GB": "১২ GB",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"152.8 x 72 x 8.5 mm": "১৫২.৮ x ৭২ x ৮.৫ মিমি",
		"198 g": "১৯৮ g",
		"4,700 mAh": "৪,৭০০ এমএএইচ",
		"50 MP + 48 MP": "৫০ MP + ৪৮ MP",
		"5G": "৫G",
		"6.3 inches": "৬.৩ ইঞ্চি",
		"60-120Hz": "৬০-১২০Hz",
		"Android 15": "Android ১৫",
		"August 2024": "August ২০২৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G4": "Google Tensor G৪",
		"Google Tensor G4 (4 nm)": "Google Tensor G৪ (৪ nm)",
		"IP68": "IP৬৮",
		"Mali-G715 MC7": "Mali-G৭১৫ MC৭",
		"OLED, 120Hz, HDR10+, 2700 nits": "OLED, ১২০Hz, HDR১০+, ২৭০০ nits",
	}
}

// Seed inserts specification_translations for existing specifications for product 'google-pixel-9'
func (s *SpecificationSeederMobileGooglePixel9) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-9"

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
