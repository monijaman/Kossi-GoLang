package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi13Pro seeds specifications/options for product 'xiaomi-13-pro'
type SpecificationSeederMobileXiaomi13Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi13Pro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi13Pro() *SpecificationSeederMobileXiaomi13Pro {
	return &SpecificationSeederMobileXiaomi13Pro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 13 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi13Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ pixels",
		"162.9 x 74.6 x 8.4 mm": "১৬২.৯ x ৭৪.৬ x ৮.৪ মিমি",
		"210 g": "২১০ g",
		"32 MP": "৩২ MP",
		"4,820 mAh": "৪,৮২০ এমএএইচ",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, আপগ্রেডযোগ্য": "Android ১৩, আপগ্রেডযোগ্য",
		"Ceramic White, Ceramic Black, Ceramic Flora Green, Mountain Blue": "Ceramic সাদা, Ceramic কালো, Ceramic Flora সবুজ, Mountain নীল",
		"December 2022": "December ২০২২",
		"Glass front, সিরামিক/পলিমার back, aluminum frame": "গ্লাস সামনে, সিরামিক/পলিমার back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 1900 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ১৯০০ nits",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AB Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-13-pro'
func (s *SpecificationSeederMobileXiaomi13Pro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-13-pro"

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
