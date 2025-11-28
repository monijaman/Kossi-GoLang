package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealme15Pro5g seeds specifications/options for product 'realme-15-pro-5g'
type SpecificationSeederMobileRealme15Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealme15Pro5g creates a new seeder instance
func NewSpecificationSeederMobileRealme15Pro5g() *SpecificationSeederMobileRealme15Pro5g {
	return &SpecificationSeederMobileRealme15Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Realme 15 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealme15Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"182 g": "১৮২ g",
		"2023": "২০২৩",
		"2400 × 1080 px (~400 ppi)": "২৪০০ × ১০৮০ px (~৪০০ ppi)",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4800 mAh": "৪৮০০ এমএএইচ",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.1": "৫.১",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"Android 13, Realme UI 4.0": "Android ১৩, Realme UI ৪.০",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 6100+ (6 nm)": "Dimensity ৬১০০+ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 6100+": "MediaTek Dimensity ৬১০০+",
		"Octa-core (2×2.6 GHz Cortex-A77 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.৬ GHz Cortex-A৭৭ + ৬×২.০ GHz Cortex-A৫৫)",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-15-pro-5g'
func (s *SpecificationSeederMobileRealme15Pro5g) Seed(db *gorm.DB) error {
	productSlug := "realme-15-pro-5g"

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
