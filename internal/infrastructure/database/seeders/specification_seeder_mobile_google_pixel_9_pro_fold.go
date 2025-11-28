package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileGooglePixel9ProFold seeds specifications/options for product 'google-pixel-9-pro-fold'
type SpecificationSeederMobileGooglePixel9ProFold struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel9ProFold creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel9ProFold() *SpecificationSeederMobileGooglePixel9ProFold {
	return &SpecificationSeederMobileGooglePixel9ProFold{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 9 Pro Fold"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel9ProFold) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz": "১-১২০Hz",
		"10 MP (Cover) + 10 MP (Inner)": "১০ MP (Cover) + ১০ MP (Inner)",
		"16 GB": "১৬ GB",
		"2076 x 2152 pixels": "২০৭৬ x ২১৫২ pixels",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"257 g": "২৫৭ g",
		"4,650 mAh": "৪,৬৫০ এমএএইচ",
		"48 MP + 10.5 MP + 10.8 MP": "৪৮ MP + ১০.৫ MP + ১০.৮ MP",
		"5G": "৫G",
		"8.0 inches (Main) / 6.3 inches (Cover)": "৮.০ ইঞ্চি (Main) / ৬.৩ ইঞ্চি (Cover)",
		"Android 15": "Android ১৫",
		"August 2024": "August ২০২৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Foldable LTPO OLED, 120Hz, HDR10+, 2700 nits": "Foldable LTPO OLED, ১২০Hz, HDR১০+, ২৭০০ nits",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G4": "Google Tensor G৪",
		"Google Tensor G4 (4 nm)": "Google Tensor G৪ (৪ nm)",
		"IPX8": "IPX৮",
		"Mali-G715 MC7": "Mali-G৭১৫ MC৭",
		"Unfolded: 155.2 x 150.2 x 5.1 mm / Folded: 155.2 x 77.1 x 10.5 mm": "Unfolded: ১৫৫.২ x ১৫০.২ x ৫.১ মিমি / Folded: ১৫৫.২ x ৭৭.১ x ১০.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'google-pixel-9-pro-fold'
func (s *SpecificationSeederMobileGooglePixel9ProFold) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-9-pro-fold"

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
