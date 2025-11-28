package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA14 seeds specifications/options for product 'samsung-galaxy-a14'
type SpecificationSeederMobileSamsungGalaxyA14 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA14 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA14() *SpecificationSeederMobileSamsungGalaxyA14 {
	return &SpecificationSeederMobileSamsungGalaxyA14{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A14"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA14) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP": "১৩ MP",
		"15 W wired": "১৫ W তারযুক্ত",
		"167.7 × 78.0 × 9.1 mm": "১৬৭.৭ × ৭৮.০ × ৯.১ মিমি",
		"201 g": "২০১ g",
		"2408 × 1080 px (FHD+)": "২৪০৮ × ১০৮০ px (FHD+)",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.2 (5G) / 5.0 (4G)": "৫.২ (৫G) / ৫.০ (৪G)",
		"50 MP + 5 MP + 2 MP": "৫০ MP + ৫ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB + microSD support": "৬৪ / ১২৮ GB + microSD support",
		"Android 13 (upgradable to Android 14)": "Android ১৩ (upgradable to Android ১৪)",
		"Black, Silver, Green, Red": "কালো, রূপালী, সবুজ, লাল",
		"GSM / HSPA / LTE (4G) / 5G (A14 5G variant)": "GSM / HSPA / LTE (৪G) / ৫G (A১৪ ৫G variant)",
		"Glass front, plastic back & frame": "গ্লাস সামনে, প্লাস্টিক পেছনে & frame",
		"Helio G80 (12 nm) / Dimensity 700 (7 nm)": "Helio G৮০ (১২ nm) / Dimensity ৭০০ (৭ nm)",
		"January 2023": "January ২০২৩",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G52 (4G) / Mali-G57 (5G)": "Mali-G৫২ (৪G) / Mali-G৫৭ (৫G)",
		"MediaTek Helio G80 (4G) / Dimensity 700 (5G)": "MediaTek Helio G৮০ (৪G) / Dimensity ৭০০ (৫G)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (market dependent)",
		"Yes (only A14 5G variant)": "হ্যাঁ (only A১৪ ৫G variant)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a14'
func (s *SpecificationSeederMobileSamsungGalaxyA14) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a14"

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
