package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileGooglePixel8 seeds specifications/options for product 'google-pixel-8'
type SpecificationSeederMobileGooglePixel8 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel8 creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel8() *SpecificationSeederMobileGooglePixel8 {
	return &SpecificationSeederMobileGooglePixel8{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 8"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel8) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10.5 MP":                          "১০.৫ MP",
		"1080 x 2400 pixels":               "১০৮০ x ২৪০০ pixels",
		"128 GB / 256 GB":                  "১২৮ GB / ২৫৬ GB",
		"150.5 x 70.8 x 8.9 mm":            "১৫০.৫ x ৭০.৮ x ৮.৯ মিমি",
		"187 g":                            "১৮৭ g",
		"4,575 mAh":                        "৪,৫৭৫ এমএএইচ",
		"50 MP + 12 MP":                    "৫০ MP + ১২ MP",
		"5G":                               "৫G",
		"6.2 inches":                       "৬.২ ইঞ্চি",
		"60-120Hz":                         "৬০-১২০Hz",
		"8 GB":                             "৮ GB",
		"Android 14":                       "Android ১৪",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G3":                 "Google Tensor G৩",
		"Google Tensor G3 (4 nm)":          "Google Tensor G৩ (৪ nm)",
		"IP68":                             "IP৬৮",
		"Immortalis-G715s MC10":            "Immortalis-G৭১৫s MC১০",
		"OLED, 120Hz, HDR10+, 2000 nits":   "OLED, ১২০Hz, HDR১০+, ২০০০ nits",
		"October 2023":                     "October ২০২৩",
		"Nona-core":                        "নোনা-কোর",
		"Corning Gorilla Glass Victus":     "কর্নিং গরিলা গ্লাস ভিক্টাস",
		"Obsidian, Hazel, Rose, Mint":      "অবসিডিয়ান, হ্যাজেল, রোজ, মিন্ট",
		"Available":                        "উপলব্ধ",
	}
}

// Seed inserts specification_translations for existing specifications for product 'google-pixel-8'
func (s *SpecificationSeederMobileGooglePixel8) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-8"

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
