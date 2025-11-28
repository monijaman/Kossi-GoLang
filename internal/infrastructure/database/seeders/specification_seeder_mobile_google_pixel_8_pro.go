package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileGooglePixel8Pro seeds specifications/options for product 'google-pixel-8-pro'
type SpecificationSeederMobileGooglePixel8Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel8Pro creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel8Pro() *SpecificationSeederMobileGooglePixel8Pro {
	return &SpecificationSeederMobileGooglePixel8Pro{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 8 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel8Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz": "১-১২০Hz",
		"10.5 MP": "১০.৫ MP",
		"12 GB": "১২ GB",
		"128 GB / 256 GB / 512 GB / 1 TB": "১২৮ GB / ২৫৬ GB / ৫১২ GB / ১ TB",
		"1344 x 2992 pixels": "১৩৪৪ x ২৯৯২ pixels",
		"162.6 x 76.5 x 8.8 mm": "১৬২.৬ x ৭৬.৫ x ৮.৮ মিমি",
		"213 g": "২১৩ g",
		"5,050 mAh": "৫,০৫০ এমএএইচ",
		"50 MP + 48 MP + 48 MP": "৫০ MP + ৪৮ MP + ৪৮ MP",
		"5G": "৫G",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"Android 14": "Android ১৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G3": "Google Tensor G৩",
		"Google Tensor G3 (4 nm)": "Google Tensor G৩ (৪ nm)",
		"IP68": "IP৬৮",
		"Immortalis-G715s MC10": "Immortalis-G৭১৫s MC১০",
		"LTPO OLED, 120Hz, HDR10+, 2400 nits": "LTPO OLED, ১২০Hz, HDR১০+, ২৪০০ nits",
		"October 2023": "October ২০২৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'google-pixel-8-pro'
func (s *SpecificationSeederMobileGooglePixel8Pro) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-8-pro"

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
