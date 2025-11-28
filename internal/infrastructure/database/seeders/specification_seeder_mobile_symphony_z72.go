package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyZ72 seeds specifications/options for product 'symphony-z72'
type SpecificationSeederMobileSymphonyZ72 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyZ72 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyZ72() *SpecificationSeederMobileSymphonyZ72 {
	return &SpecificationSeederMobileSymphonyZ72{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Z72"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyZ72) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p@30fps": "১০৮০p@৩০fps",
		"10W wired": "১০W তারযুক্ত",
		"171.45 x 78.07 x 8.35 mm": "১৭১.৪৫ x ৭৮.০৭ x ৮.৩৫ মিমি",
		"195 g": "১৯৫ g",
		"4 GB": "৪ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"52 MP + 2 MP": "৫২ MP + ২ MP",
		"6.88 inches": "৬.৮৮ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"650 MHz GPU": "৬৫০ MHz GPU",
		"720 × 1640 px": "৭২০ × ১৬৪০ px",
		"8 MP": "৮ MP",
		"Android 15": "Android ১৫",
		"April 2025": "April ২০২৫",
		"IP54 water & dust resistance": "IP৫৪ water & dust resistance",
		"Imperial Purple, Cosmic Gold, Titanium Gray, Graphite Black, Swamp Green": "Imperial বেগুনি, Cosmic সোনালী, Titanium ধূসর, Graphite কালো, Swamp সবুজ",
		"Octa-core 1.6 GHz": "Octa-core ১.৬ GHz",
		"Side fingerprint, accelerometer, proximity, compass, gyroscope": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস, gyroscope",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-z72'
func (s *SpecificationSeederMobileSymphonyZ72) Seed(db *gorm.DB) error {
	productSlug := "symphony-z72"

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
