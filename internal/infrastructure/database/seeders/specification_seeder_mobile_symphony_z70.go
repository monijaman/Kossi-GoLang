package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyZ70 seeds specifications/options for product 'symphony-z70'
type SpecificationSeederMobileSymphonyZ70 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyZ70 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyZ70() *SpecificationSeederMobileSymphonyZ70 {
	return &SpecificationSeederMobileSymphonyZ70{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Z70"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyZ70) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p@30fps": "১০৮০p@৩০fps",
		"10W wired": "১০W তারযুক্ত",
		"164.27 x 76.02 x 8.45 mm": "১৬৪.২৭ x ৭৬.০২ x ৮.৪৫ মিমি",
		"193 g": "১৯৩ g",
		"4 GB": "৪ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"52 MP + 2 MP + 2 MP": "৫২ MP + ২ MP + ২ MP",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"64 GB": "৬৪ GB",
		"650 MHz GPU": "৬৫০ MHz GPU",
		"720 × 1612 px": "৭২০ × ১৬১২ px",
		"8 MP": "৮ MP",
		"Android 13": "Android ১৩",
		"November 2024": "November ২০২৪",
		"Octa-core 1.6 GHz": "Octa-core ১.৬ GHz",
		"Reflective Green, Electric Blue, Honey Dew Green, Fusion Gold": "Reflective সবুজ, Electric নীল, Honey Dew সবুজ, Fusion সোনালী",
		"Side fingerprint, accelerometer, proximity, compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
		"Wi-Fi, Bluetooth 5.0, USB-C, OTG": "Wi-Fi, Bluetooth ৫.০, USB-C, OTG",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-z70'
func (s *SpecificationSeederMobileSymphonyZ70) Seed(db *gorm.DB) error {
	productSlug := "symphony-z70"

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
