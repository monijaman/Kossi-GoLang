package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi13Ultra seeds specifications/options for product 'xiaomi-13-ultra'
type SpecificationSeederMobileXiaomi13Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi13Ultra creates a new seeder instance
func NewSpecificationSeederMobileXiaomi13Ultra() *SpecificationSeederMobileXiaomi13Ultra {
	return &SpecificationSeederMobileXiaomi13Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 13 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi13Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ pixels",
		"163.2 x 74.6 x 9.1 mm": "১৬৩.২ x ৭৪.৬ x ৯.১ মিমি",
		"227 g": "২২৭ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"32 MP": "৩২ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, upgradable": "Android ১৩, upgradable",
		"April 2023": "April ২০২৩",
		"Black, Olive Green, White": "কালো, Olive সবুজ, সাদা",
		"Glass front, eco leather back, aluminum frame": "গ্লাস সামনে, eco leather back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 2600 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ২৬০০ nits",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AB Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-13-ultra'
func (s *SpecificationSeederMobileXiaomi13Ultra) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-13-ultra"

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
