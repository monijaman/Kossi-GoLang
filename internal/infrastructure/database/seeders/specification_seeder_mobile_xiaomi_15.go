package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi15 seeds specifications/options for product 'xiaomi-15'
type SpecificationSeederMobileXiaomi15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15 creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15() *SpecificationSeederMobileXiaomi15 {
	return &SpecificationSeederMobileXiaomi15{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"1200 x 2670 pixels": "১২০০ x ২৬৭০ pixels",
		"120Hz": "১২০Hz",
		"152.8 x 71.5 x 8.2 mm": "১৫২.৮ x ৭১.৫ x ৮.২ মিমি",
		"191 g": "১৯১ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"32 MP": "৩২ MP",
		"5,400 mAh": "৫,৪০০ এমএএইচ",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.36 inches": "৬.৩৬ ইঞ্চি",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, HyperOS": "Android ১৫, HyperOS",
		"Black, White, Green, Purple, Silver": "কালো, সাদা, সবুজ, বেগুনি, রূপালী",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO OLED, 120Hz, Dolby Vision, HDR10+, 3200 nits": "LTPO OLED, ১২০Hz, Dolby Vision, HDR১০+, ৩২০০ nits",
		"October 2024": "October ২০২৪",
		"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "Qualcomm SM৮৭৫০ Snapdragon ৮ Elite (৩ nm)",
		"Snapdragon 8 Elite": "Snapdragon ৮ Elite",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-15'
func (s *SpecificationSeederMobileXiaomi15) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15"

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
