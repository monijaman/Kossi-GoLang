package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi14Ultra seeds specifications/options for product 'xiaomi-14-ultra'
type SpecificationSeederMobileXiaomi14Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi14Ultra creates a new seeder instance
func NewSpecificationSeederMobileXiaomi14Ultra() *SpecificationSeederMobileXiaomi14Ultra {
	return &SpecificationSeederMobileXiaomi14Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 14 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi14Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ pixels",
		"161.4 x 75.3 x 9.2 mm": "১৬১.৪ x ৭৫.৩ x ৯.২ মিমি",
		"219 g": "২১৯ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"32 MP": "৩২ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Black, White, Blue, Titanium Gray": "কালো, সাদা, নীল, Titanium ধূসর",
		"February 2024": "February ২০২৪",
		"Glass front, eco leather back, titanium frame": "গ্লাস সামনে, eco leather back, টাইটানিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 3000 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ৩০০০ nits",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AB Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-14-ultra'
func (s *SpecificationSeederMobileXiaomi14Ultra) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-14-ultra"

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
