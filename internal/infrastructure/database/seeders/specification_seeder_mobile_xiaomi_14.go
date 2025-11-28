package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi14 seeds specifications/options for product 'xiaomi-14'
type SpecificationSeederMobileXiaomi14 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi14 creates a new seeder instance
func NewSpecificationSeederMobileXiaomi14() *SpecificationSeederMobileXiaomi14 {
	return &SpecificationSeederMobileXiaomi14{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 14"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi14) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1200 x 2670 pixels": "১২০০ x ২৬৭০ pixels",
		"120Hz": "১২০Hz",
		"152.8 x 71.5 x 8.2 mm": "১৫২.৮ x ৭১.৫ x ৮.২ মিমি",
		"188 g": "১৮৮ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"32 MP": "৩২ MP",
		"4,610 mAh": "৪,৬১০ এমএএইচ",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.36 inches": "৬.৩৬ ইঞ্চি",
		"8 GB / 12 GB / 16 GB": "৮ GB / ১২ GB / ১৬ GB",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Black, White, Jade Green, Pink": "কালো, সাদা, Jade সবুজ, গোলাপী",
		"Glass front, glass/সিলিকন পলিমার back, aluminum frame": "গ্লাস সামনে, glass/সিলিকন পলিমার back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO OLED, 120Hz, Dolby Vision, HDR10+, 3000 nits": "LTPO OLED, ১২০Hz, Dolby Vision, HDR১০+, ৩০০০ nits",
		"October 2023": "October ২০২৩",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AB Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-14'
func (s *SpecificationSeederMobileXiaomi14) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-14"

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
