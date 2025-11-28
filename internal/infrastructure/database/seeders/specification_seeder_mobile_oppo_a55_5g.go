package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoA555g seeds specifications/options for product 'oppo-a55-5g'
type SpecificationSeederMobileOppoA555g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA555g creates a new seeder instance
func NewSpecificationSeederMobileOppoA555g() *SpecificationSeederMobileOppoA555g {
	return &SpecificationSeederMobileOppoA555g{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A55 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA555g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"1600 × 720 px (~269 ppi)": "১৬০০ × ৭২০ px (~২৬৯ ppi)",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2021": "২০২১",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Android 11, ColorOS 11": "Android ১১, ColorOS ১১",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 700 (7 nm)": "Dimensity ৭০০ (৭ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 700": "MediaTek Dimensity ৭০০",
		"Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.২ GHz Cortex-A৭৬ + ৬×২.০ GHz Cortex-A৫৫)",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-a55-5g'
func (s *SpecificationSeederMobileOppoA555g) Seed(db *gorm.DB) error {
	productSlug := "oppo-a55-5g"

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
