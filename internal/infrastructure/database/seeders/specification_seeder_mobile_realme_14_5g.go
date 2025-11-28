package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealme145g seeds specifications/options for product 'realme-14-5g'
type SpecificationSeederMobileRealme145g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealme145g creates a new seeder instance
func NewSpecificationSeederMobileRealme145g() *SpecificationSeederMobileRealme145g {
	return &SpecificationSeederMobileRealme145g{BaseSeeder: BaseSeeder{name: "Specifications for Realme 14 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealme145g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"18 W wired": "১৮ W তারযুক্ত",
		"180 g": "১৮০ g",
		"2023": "২০২৩",
		"2400 × 1080 px (~405 ppi)": "২৪০০ × ১০৮০ px (~৪০৫ ppi)",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Android 13, Realme UI 4.0": "Android ১৩, Realme UI ৪.০",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 6020 (6 nm)": "Dimensity ৬০২০ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 6020": "MediaTek Dimensity ৬০২০",
		"Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.২ GHz Cortex-A৭৬ + ৬×২.০ GHz Cortex-A৫৫)",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-14-5g'
func (s *SpecificationSeederMobileRealme145g) Seed(db *gorm.DB) error {
	productSlug := "realme-14-5g"

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
