package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM365g seeds specifications/options for product 'samsung-galaxy-m36-5g'
type SpecificationSeederMobileSamsungGalaxyM365g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM365g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM365g() *SpecificationSeederMobileSamsungGalaxyM365g {
	return &SpecificationSeederMobileSamsungGalaxyM365g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M36 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM365g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ px",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"13 MP": "১৩ MP",
		"162.3 × 78.6 × 9.1 mm": "১৬২.৩ × ৭৮.৬ × ৯.১ মিমি",
		"222 g": "২২২ g",
		"25 W wired": "২৫ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.3": "৫.৩",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6 years of Android / Security updates": "৬ years of Android / Security updates",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"Android 15 (One UI 7)": "Android ১৫ (One UI ৭)",
		"Exynos 1380": "Exynos ১৩৮০",
		"Exynos 1380 (5 nm)": "Exynos ১৩৮০ (৫ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass Victus+), plastic frame, glass back": "গ্লাস সামনে (Gorilla Glass Victus+), plastic frame, গ্লাস পেছনে",
		"July 2025": "July ২০২৫",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G68 MP5": "Mali-G৬৮ MP৫",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)": "Octa-core (৪×২.৪ GHz Cortex-A৭৮ + ৪×২.০ GHz Cortex-A৫৫)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Super AMOLED, 120 Hz": "Super AMOLED, ১২০ Hz",
		"Thunder Grey, DayBreak Blue, Moonlight Blue": "Thunder ধূসর, DayBreak নীল, Moonlight নীল",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m36-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM365g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m36-5g"

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
