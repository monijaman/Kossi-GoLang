package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyA255g seeds specifications/options for product 'samsung-galaxy-a25-5g'
type SpecificationSeederMobileSamsungGalaxyA255g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA255g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA255g() *SpecificationSeederMobileSamsungGalaxyA255g {
	return &SpecificationSeederMobileSamsungGalaxyA255g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A25 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA255g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD (up to 1 TB)": "১২৮ / ২৫৬ GB + microSD (পর্যন্ত ১ TB)",
		"13 MP": "১৩ MP",
		"161 × 76.5 × 8.3 mm": "১৬১ × ৭৬.৫ × ৮.৩ মিমি",
		"197 g": "১৯৭ g",
		"2340 × 1080 pixels (~396 ppi)": "২৩৪০ × ১০৮০ pixels (~৩৯৬ ppi)",
		"25 W wired": "২৫ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.3": "৫.৩",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"Accelerometer, Gyro, Proximity, Compass": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Blue Black, Blue, Light Blue, Yellow": "নীল কালো, নীল, Light নীল, হলুদ",
		"December 2023": "December ২০২৩",
		"Exynos 1280": "Exynos ১২৮০",
		"Exynos 1280 (5 nm)": "Exynos ১২৮০ (৫ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame / back": "গ্লাস সামনে, plastic frame / back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68": "Mali-G৬৮",
		"Nano-SIM (or hybrid)": "ন্যানো-সিম (or hybrid)",
		"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.৪ GHz Cortex-A৭৮ + ৬×২.০ GHz Cortex-A৫৫)",
		"Super AMOLED, 120 Hz": "Super AMOLED, ১২০ Hz",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-a25-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA255g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a25-5g"

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
