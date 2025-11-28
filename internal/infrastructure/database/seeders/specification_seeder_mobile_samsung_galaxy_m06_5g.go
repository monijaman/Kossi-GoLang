package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM065g seeds specifications/options for product 'samsung-galaxy-m06-5g'
type SpecificationSeederMobileSamsungGalaxyM065g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM065g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM065g() *SpecificationSeederMobileSamsungGalaxyM065g {
	return &SpecificationSeederMobileSamsungGalaxyM065g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M06 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM065g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"167.4 × 77.4 × 8 mm": "১৬৭.৪ × ৭৭.৪ × ৮ মিমি",
		"2025": "২০২৫",
		"25 W wired (likely)": "২৫ W তারযুক্ত (likely)",
		"3.5 mm (some sources)": "৩.৫ মিমি (some sources)",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.3": "৫.৩",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.74 inches": "৬.৭৪ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"720 × 1600 px (reportedly)": "৭২০ × ১৬০০ px (reportedly)",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Accelerometer, Gyro, Proximity, Compass, Fingerprint (side?)": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (side?)",
		"Blazing Black, Sage Green": "Blazing কালো, Sage সবুজ",
		"Dimensity 6300 (6 nm)": "Dimensity ৬৩০০ (৬ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Dimensity 6300": "MediaTek Dimensity ৬৩০০",
		"Nano-SIM + Nano-SIM or hybrid": "ন্যানো-সিম + ন্যানো-সিম or hybrid",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m06-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM065g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m06-5g"

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
