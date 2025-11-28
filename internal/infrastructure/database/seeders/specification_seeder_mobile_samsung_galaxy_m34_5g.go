package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM345g seeds specifications/options for product 'samsung-galaxy-m34-5g'
type SpecificationSeederMobileSamsungGalaxyM345g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM345g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM345g() *SpecificationSeederMobileSamsungGalaxyM345g {
	return &SpecificationSeederMobileSamsungGalaxyM345g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M34 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM345g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD (up to 1TB)": "১২৮ / ২৫৬ GB + microSD (পর্যন্ত ১TB)",
		"13 MP": "১৩ MP",
		"161.7 × 77.2 × 8.8 mm": "১৬১.৭ × ৭৭.২ × ৮.৮ মিমি",
		"208 g": "২০৮ g",
		"2340 × 1080 px": "২৩৪০ × ১০৮০ px",
		"25 W wired": "২৫ W তারযুক্ত",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"Accelerometer, Gyro, Proximity, Compass, fingerprint (side)": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (side)",
		"Android 13, One UI 5 (upgradable)": "Android ১৩, One UI ৫ (upgradable)",
		"Exynos 1280": "Exynos ১২৮০",
		"Exynos 1280 (5 nm)": "Exynos ১২৮০ (৫ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic frame / back": "গ্লাস সামনে (Gorilla Glass ৫), plastic frame / back",
		"Hybrid Dual SIM (nano + microSD)": "Hybrid ডুয়াল সিম (nano + microSD)",
		"July 2023": "July ২০২৩",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G68 MP4": "Mali-G৬৮ MP৪",
		"Midnight Blue, Prism Silver, Waterfall Blue": "Midnight নীল, Prism রূপালী, Waterfall নীল",
		"Octa-core (2×2.4 GHz + 6×2.0 GHz)": "Octa-core (২×২.৪ GHz + ৬×২.০ GHz)",
		"Super AMOLED, 120 Hz": "Super AMOLED, ১২০ Hz",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m34-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM345g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m34-5g"

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
