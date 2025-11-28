package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA545g seeds specifications/options for product 'samsung-galaxy-a54-5g'
type SpecificationSeederMobileSamsungGalaxyA545g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA545g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA545g() *SpecificationSeederMobileSamsungGalaxyA545g {
	return &SpecificationSeederMobileSamsungGalaxyA545g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A54 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA545g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"158.2 × 76.7 × 8.2 mm": "১৫৮.২ × ৭৬.৭ × ৮.২ মিমি",
		"202 g": "২০২ g",
		"2340 × 1080 px (19.5:9)": "২৩৪০ × ১০৮০ px (১৯.৫:৯)",
		"25 W wired": "২৫ W তারযুক্ত",
		"32 MP": "৩২ MP",
		"4 OS upgrades, 5 years of security updates": "৪ OS upgrades, ৫ years of security updates",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.2": "৫.২",
		"50 MP + 8 MP + 5 MP": "৫০ MP + ৮ MP + ৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.4 inches": "৬.৪ ইঞ্চি",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"Exynos 1380": "Exynos ১৩৮০",
		"Exynos 1380 (5 nm)": "Exynos ১৩৮০ (৫ nm)",
		"Fingerprint (in-display, optical), Accelerometer, Gyro, Compass, Barometer": "ফিঙ্গারপ্রিন্ট (in-display, optical), অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস, ব্যারোমিটার",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic frame / back": "গ্লাস সামনে (Gorilla Glass ৫), plastic frame / back",
		"IP67 (up to 1 m for 30 min)": "IP৬৭ (up to ১ m for ৩০ min)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Lime, Graphite, Violet, White": "Lime, Graphite, Violet, সাদা",
		"Mali-G68 (MP5)": "Mali-G৬৮ (MP৫)",
		"March 2023": "March ২০২৩",
		"Nano-SIM বা হাইব্রিড Dual SIM (বাজার নির্ভরশীল)": "ন্যানো-সিম বা হাইব্রিড ডুয়াল সিম (বাজার নির্ভরশীল)",
		"Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)": "Octa-core (৪×২.৪ GHz Cortex-A৭৮ + ৪×২.০ GHz Cortex-A৫৫)",
		"সুপার AMOLED, 120 Hz": "সুপার AMOLED, ১২০ Hz",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band": "Wi-Fi ৮০২.১১ a/b/g/n/ac, dual-band",
		"Yes (region নির্ভরশীল)": "হ্যাঁ (region নির্ভরশীল)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a54-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA545g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a54-5g"

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
