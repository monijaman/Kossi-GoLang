package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA24 seeds specifications/options for product 'samsung-galaxy-a24'
type SpecificationSeederMobileSamsungGalaxyA24 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA24 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA24() *SpecificationSeederMobileSamsungGalaxyA24 {
	return &SpecificationSeederMobileSamsungGalaxyA24{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A24"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA24) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB + microSD": "১২৮ GB + microSD",
		"13 MP": "১৩ MP",
		"165.4 × 76.9 × 8.4 mm": "১৬৫.৪ × ৭৬.৯ × ৮.৪ মিমি",
		"195 g": "১৯৫ g",
		"2408 × 1080 pixels (~406 ppi)": "২৪০৮ × ১০৮০ pixels (~৪০৬ ppi)",
		"25 W wired": "২৫ W তারযুক্ত",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.3": "৫.৩",
		"50 MP + 2 MP + 2 MP": "৫০ MP + ২ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"90 Hz": "৯০ Hz",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"April 2023": "April ২০২৩",
		"Black, Dark Red, Light Green, Silver": "কালো, Dark লাল, Light সবুজ, রূপালী",
		"Eye Care Certification (low blue light)": "Eye Care Certification (low নীল light)",
		"GSM / HSPA / LTE (no 5G)": "GSM / HSPA / LTE (no ৫G)",
		"Glass front, plastic frame / back": "গ্লাস সামনে, plastic frame / back",
		"Helio G99 (6 nm)": "Helio G৯৯ (৬ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Helio G99": "MediaTek Helio G৯৯",
		"Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.২ GHz Cortex-A৭৬ + ৬×২.০ GHz Cortex-A৫৫)",
		"Side-mounted fingerprint, Accelerometer, Gyro, Proximity, Compass": "পাশে মাউন্ট ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (market dependent)",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a24'
func (s *SpecificationSeederMobileSamsungGalaxyA24) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a24"

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
