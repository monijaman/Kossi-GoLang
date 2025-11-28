package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealmeNote70 seeds specifications/options for product 'realme-note-70'
type SpecificationSeederMobileRealmeNote70 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeNote70 creates a new seeder instance
func NewSpecificationSeederMobileRealmeNote70() *SpecificationSeederMobileRealmeNote70 {
	return &SpecificationSeederMobileRealmeNote70{BaseSeeder: BaseSeeder{name: "Specifications for Realme Note 70"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeNote70) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"16 MP": "১৬ MP",
		"195 g": "১৯৫ g",
		"2023": "২০২৩",
		"2412 × 1080 px (~400 ppi)": "২৪১২ × ১০৮০ px (~৪০০ ppi)",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"Android 13, Realme UI 4.0": "Android ১৩, Realme UI ৪.০",
		"Black, Blue, Silver": "কালো, নীল, রূপালী",
		"Dimensity 6080+ (6 nm)": "Dimensity ৬০৮০+ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 6080+": "MediaTek Dimensity ৬০৮০+",
		"Octa-core (2×2.6 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.৬ GHz Cortex-A৭৬ + ৬×২.০ GHz Cortex-A৫৫)",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-note-70'
func (s *SpecificationSeederMobileRealmeNote70) Seed(db *gorm.DB) error {
	productSlug := "realme-note-70"

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
