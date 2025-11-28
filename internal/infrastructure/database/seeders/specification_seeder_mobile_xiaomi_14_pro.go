package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi14Pro seeds specifications/options for product 'xiaomi-14-pro'
type SpecificationSeederMobileXiaomi14Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi14Pro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi14Pro() *SpecificationSeederMobileXiaomi14Pro {
	return &SpecificationSeederMobileXiaomi14Pro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 14 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi14Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ pixels",
		"161.4 x 75.3 x 8.5 mm": "১৬১.৪ x ৭৫.৩ x ৮.৫ মিমি",
		"223 g": "২২৩ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"32 MP": "৩২ MP",
		"4,880 mAh": "৪,৮৮০ এমএএইচ",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Black, Silver, Titanium, Green": "কালো, রূপালী, Titanium, সবুজ",
		"Glass front/back, titanium/aluminum frame": "গ্লাস সামনে/back, titanium/অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 3000 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ৩০০০ nits",
		"October 2023": "October ২০২৩",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AB Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-14-pro'
func (s *SpecificationSeederMobileXiaomi14Pro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-14-pro"

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
