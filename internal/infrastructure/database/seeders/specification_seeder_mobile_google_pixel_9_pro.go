package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileGooglePixel9Pro seeds specifications/options for product 'google-pixel-9-pro'
type SpecificationSeederMobileGooglePixel9Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel9Pro creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel9Pro() *SpecificationSeederMobileGooglePixel9Pro {
	return &SpecificationSeederMobileGooglePixel9Pro{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 9 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel9Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz": "১-১২০Hz",
		"128 GB / 256 GB / 512 GB / 1 টিবি": "১২৮ GB / ২৫৬ GB / ৫১২ GB / ১ টিবি",
		"1280 x 2856 pixels": "১২৮০ x ২৮৫৬ pixels",
		"152.8 x 72 x 8.5 mm": "১৫২.৮ x ৭২ x ৮.৫ মিমি",
		"16 GB": "১৬ GB",
		"199 g": "১৯৯ g",
		"4,700 mAh": "৪,৭০০ এমএএইচ",
		"42 MP": "৪২ MP",
		"50 MP + 48 MP + 48 MP": "৫০ MP + ৪৮ MP + ৪৮ MP",
		"5G": "৫G",
		"6.3 inches": "৬.৩ ইঞ্চি",
		"Android 15": "Android ১৫",
		"August 2024": "August ২০২৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G4": "Google Tensor G৪",
		"Google Tensor G4 (4 nm)": "Google Tensor G৪ (৪ nm)",
		"IP68": "IP৬৮",
		"LTPO OLED, 120Hz, HDR10+, 3000 nits": "LTPO OLED, ১২০Hz, HDR১০+, ৩০০০ nits",
		"Mali-G715 MC7": "Mali-G৭১৫ MC৭",
	}
}

// Seed inserts specification_translations for existing specifications for product 'google-pixel-9-pro'
func (s *SpecificationSeederMobileGooglePixel9Pro) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-9-pro"

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
