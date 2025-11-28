package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyHelio50 seeds specifications/options for product 'symphony-helio-50'
type SpecificationSeederMobileSymphonyHelio50 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyHelio50 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyHelio50() *SpecificationSeederMobileSymphonyHelio50 {
	return &SpecificationSeederMobileSymphonyHelio50{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Helio 50"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyHelio50) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP + 0.8 MP": "১০৮ MP + ২ MP + ০.৮ MP",
		"1080 × 2400 pixels": "১০৮০ × ২৪০০ pixels",
		"1080p@30fps": "১০৮০p@৩০fps",
		"128 GB + microSD": "১২৮ GB + microSD",
		"167 x 77.6 x 8.79 mm": "১৬৭ x ৭৭.৬ x ৮.৭৯ মিমি",
		"215 g": "২১৫ g",
		"2K@30fps, 1080p@30fps": "২K@৩০fps, ১০৮০p@৩০fps",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"5.x": "৫.x",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 GB": "৬ GB",
		"6.72 inches": "৬.৭২ ইঞ্চি",
		"90Hz": "৯০Hz",
		"Android 13": "Android ১৩",
		"Dual Nano-SIM / হাইব্রিড": "Dual ন্যানো-সিম / হাইব্রিড",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"June 2024": "June ২০২৪",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G52": "Mali-G৫২",
		"MediaTek Helio G88": "MediaTek Helio G৮৮",
		"MediaTek Helio G88 (12 nm)": "MediaTek Helio G৮৮ (১২ nm)",
		"Octa-core, up to 2.0 GHz": "Octa-core, up to ২.০ GHz",
		"Radium Green, Cosmic Gold, Honey Dew Green": "Radium সবুজ, Cosmic সোনালী, Honey Dew সবুজ",
		"Side fingerprint, proximity, light, gyroscope": "Side ফিঙ্গারপ্রিন্ট, প্রক্সিমিটি, light, gyroscope",
		"USB Type-C 2.0": "USB Type-C ২.০",
		"Wi-Fi 802.11 b/g/n": "Wi-Fi ৮০২.১১ b/g/n",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-helio-50'
func (s *SpecificationSeederMobileSymphonyHelio50) Seed(db *gorm.DB) error {
	productSlug := "symphony-helio-50"

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
