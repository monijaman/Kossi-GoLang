package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileGooglePixel8a seeds specifications/options for product 'google-pixel-8a'
type SpecificationSeederMobileGooglePixel8a struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel8a creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel8a() *SpecificationSeederMobileGooglePixel8a {
	return &SpecificationSeederMobileGooglePixel8a{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 8a"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel8a) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"13 MP": "১৩ MP",
		"152.1 x 72.7 x 8.9 mm": "১৫২.১ x ৭২.৭ x ৮.৯ মিমি",
		"188 g": "১৮৮ g",
		"4,492 mAh": "৪,৪৯২ এমএএইচ",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"64 MP + 13 MP": "৬৪ MP + ১৩ MP",
		"8 GB": "৮ GB",
		"Android 14": "Android ১৪",
		"Corning Gorilla Glass 3": "Corning Gorilla Glass ৩",
		"Glass front, plastic back, aluminum frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G3": "Google Tensor G৩",
		"Google Tensor G3 (4 nm)": "Google Tensor G৩ (৪ nm)",
		"IP67": "IP৬৭",
		"Immortalis-G715s MC10": "Immortalis-G৭১৫s MC১০",
		"May 2024": "May ২০২৪",
		"OLED, 120Hz, HDR": "OLED, ১২০Hz, HDR",
	}
}

// Seed inserts specification_translations for existing specifications for product 'google-pixel-8a'
func (s *SpecificationSeederMobileGooglePixel8a) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-8a"

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
