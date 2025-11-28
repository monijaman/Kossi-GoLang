package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM335g seeds specifications/options for product 'samsung-galaxy-m33-5g'
type SpecificationSeederMobileSamsungGalaxyM335g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM335g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM335g() *SpecificationSeederMobileSamsungGalaxyM335g {
	return &SpecificationSeederMobileSamsungGalaxyM335g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M33 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM335g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"128 GB + microSD": "১২৮ GB + microSD",
		"165.4 × 76.9 × 9.4 mm": "১৬৫.৪ × ৭৬.৯ × ৯.৪ মিমি",
		"198 g": "১৯৮ g",
		"2408 × 1080 px": "২৪০৮ × ১০৮০ px",
		"25 W wired": "২৫ W তারযুক্ত",
		"3.5 mm": "৩.৫ মিমি",
		"4K @ 30fps; 1080p @ 30/60/720fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০/৭২০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.1": "৫.১",
		"50 MP + 5 MP + 2 MP + 2 MP": "৫০ MP + ৫ MP + ২ MP + ২ MP",
		"5000 mAh (some markets 6000 mAh)": "৫০০০ এমএএইচ (some markets ৬০০০ এমএএইচ)",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"8 MP": "৮ MP",
		"Android 12 (upgradable)": "Android ১২ (upgradable)",
		"April 2022": "April ২০২২",
		"Deep Ocean Blue, Mystique Green, Emerald Brown": "Deep Ocean নীল, Mystique সবুজ, Emerald Brown",
		"Exynos 1280": "Exynos ১২৮০",
		"Exynos 1280 (5 nm)": "Exynos ১২৮০ (৫ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic frame / back": "গ্লাস সামনে (Gorilla Glass ৫), plastic frame / back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MP4": "Mali-G৬৮ MP৪",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (2×2.4 GHz + 6×2.0 GHz)": "Octa-core (২×২.৪ GHz + ৬×২.০ GHz)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"TFT LCD, 120 Hz": "TFT LCD, ১২০ Hz",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band": "Wi-Fi ৮০২.১১ a/b/g/n/ac, dual-band",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m33-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM335g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m33-5g"

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
