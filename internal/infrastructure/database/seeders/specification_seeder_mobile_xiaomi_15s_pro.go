package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi15sPro seeds specifications/options for product 'xiaomi-15s-pro'
type SpecificationSeederMobileXiaomi15sPro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15sPro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15sPro() *SpecificationSeederMobileXiaomi15sPro {
	return &SpecificationSeederMobileXiaomi15sPro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15S Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15sPro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"120W wired, 50W wireless": "১২০W তারযুক্ত, ৫০W বেতার",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ pixels",
		"215 g": "২১৫ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"32 MP": "৩২ MP",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, HyperOS": "Android ১৫, HyperOS",
		"Black, White, Green": "কালো, সাদা, সবুজ",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, glass back or eco leather back, aluminum frame": "গ্লাস সামনে, গ্লাস পেছনে or eco leather back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+",
		"May 2025": "May ২০২৫",
		"Qualcomm Snapdragon 8 Gen 4": "Qualcomm Snapdragon ৮ Gen ৪",
		"Snapdragon 8 Gen 4": "Snapdragon ৮ Gen ৪",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-15s-pro'
func (s *SpecificationSeederMobileXiaomi15sPro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15s-pro"

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
