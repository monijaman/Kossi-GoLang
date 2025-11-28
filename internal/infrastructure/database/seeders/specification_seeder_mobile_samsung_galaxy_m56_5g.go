package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM565g seeds specifications/options for product 'samsung-galaxy-m56-5g'
type SpecificationSeederMobileSamsungGalaxyM565g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM565g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM565g() *SpecificationSeederMobileSamsungGalaxyM565g {
	return &SpecificationSeederMobileSamsungGalaxyM565g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M56 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM565g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ px",
		"12 MP": "১২ MP",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSDXC": "১২৮ / ২৫৬ GB + microSDXC",
		"162 × 77.3 × 7.2 mm": "১৬২ × ৭৭.৩ × ৭.২ মিমি",
		"180 g": "১৮০ g",
		"45 W wired": "৪৫ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.3": "৫.৩",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.74 inches": "৬.৭৪ ইঞ্চি",
		"8 GB": "৮ GB",
		"Accelerometer, Gyro, Proximity, Compass, fingerprint (side?)": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (side?)",
		"Android 15, One UI 7": "Android ১৫, One UI ৭",
		"April 2025": "April ২০২৫",
		"Black, Light Green": "কালো, Light সবুজ",
		"Exynos 1480": "Exynos ১৪৮০",
		"Exynos 1480 (4 nm)": "Exynos ১৪৮০ (৪ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass Victus+), plastic frame, glass back": "গ্লাস সামনে (Gorilla Glass Victus+), plastic frame, গ্লাস পেছনে",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (4×2.75 GHz + 4×2.0 GHz)": "Octa-core (৪×২.৭৫ GHz + ৪×২.০ GHz)",
		"Super AMOLED+ 120 Hz": "Super AMOLED+ ১২০ Hz",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
		"Xclipse 530": "Xclipse ৫৩০",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m56-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM565g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m56-5g"

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
