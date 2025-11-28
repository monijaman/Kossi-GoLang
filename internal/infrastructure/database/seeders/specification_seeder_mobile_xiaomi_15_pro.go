package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi15Pro seeds specifications/options for product 'xiaomi-15-pro'
type SpecificationSeederMobileXiaomi15Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15Pro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15Pro() *SpecificationSeederMobileXiaomi15Pro {
	return &SpecificationSeederMobileXiaomi15Pro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ pixels",
		"161.3 x 75.3 x 8.4 mm": "১৬১.৩ x ৭৫.৩ x ৮.৪ মিমি",
		"213 g": "২১৩ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"32 MP": "৩২ MP",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6,100 mAh": "৬,১০০ এমএএইচ",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, HyperOS": "Android ১৫, HyperOS",
		"Black, White, Green, Silver": "কালো, সাদা, সবুজ, রূপালী",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 3200 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ৩২০০ nits",
		"October 2024": "October ২০২৪",
		"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "Qualcomm SM৮৭৫০ Snapdragon ৮ Elite (৩ nm)",
		"Snapdragon 8 Elite": "Snapdragon ৮ Elite",
		"Xiaomi Dragon Crystal Glass 2.0": "Xiaomi Dragon Crystal Glass ২.০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-15-pro'
func (s *SpecificationSeederMobileXiaomi15Pro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15-pro"

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
