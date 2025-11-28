package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyHelio90 seeds specifications/options for product 'symphony-helio-90'
type SpecificationSeederMobileSymphonyHelio90 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyHelio90 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyHelio90() *SpecificationSeederMobileSymphonyHelio90 {
	return &SpecificationSeederMobileSymphonyHelio90{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Helio 90"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyHelio90) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2412 pixels": "১০৮০ × ২৪১২ pixels",
		"1080p@30fps": "১০৮০p@৩০fps",
		"120Hz": "১২০Hz",
		"128 GB + microSD": "১২৮ GB + microSD",
		"163.8 x 76.3 x 7.96 mm": "১৬৩.৮ x ৭৬.৩ x ৭.৯৬ মিমি",
		"188 g": "১৮৮ g",
		"2K@30fps, 1080p@30fps": "২K@৩০fps, ১০৮০p@৩০fps",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"5.x": "৫.x",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 MP + 2 MP + 0.8 MP": "৬৪ MP + ২ MP + ০.৮ MP",
		"8 GB": "৮ GB",
		"Android 14": "Android ১৪",
		"Dual Nano-SIM / Hybrid": "Dual ন্যানো-সিম / Hybrid",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"In-display fingerprint, proximity, light, gyroscope": "In-display ফিঙ্গারপ্রিন্ট, প্রক্সিমিটি, light, gyroscope",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Helio G99": "MediaTek Helio G৯৯",
		"MediaTek Helio G99 (6 nm)": "MediaTek Helio G৯৯ (৬ nm)",
		"Octa-core, up to 2.2 GHz": "Octa-core, up to ২.২ GHz",
		"September 2024": "September ২০২৪",
		"Space Black, Thunder White, Cosmic Gold": "Space কালো, Thunder সাদা, Cosmic সোনালী",
		"USB Type-C 2.0": "USB Type-C ২.০",
		"Wi-Fi 2.4 / 5 GHz": "Wi-Fi ২.৪ / ৫ GHz",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-helio-90'
func (s *SpecificationSeederMobileSymphonyHelio90) Seed(db *gorm.DB) error {
	productSlug := "symphony-helio-90"

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
