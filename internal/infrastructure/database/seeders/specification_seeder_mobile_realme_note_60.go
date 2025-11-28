package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealmeNote60 seeds specifications/options for product 'realme-note-60'
type SpecificationSeederMobileRealmeNote60 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeNote60 creates a new seeder instance
func NewSpecificationSeederMobileRealmeNote60() *SpecificationSeederMobileRealmeNote60 {
	return &SpecificationSeederMobileRealmeNote60{BaseSeeder: BaseSeeder{name: "Specifications for Realme Note 60"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeNote60) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"190 g": "১৯০ g",
		"2023": "২০২৩",
		"2400 × 1080 px (~405 ppi)": "২৪০০ × ১০৮০ px (~৪০৫ ppi)",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Android 13, Realme UI 4.0": "Android ১৩, Realme UI ৪.০",
		"Black, Blue": "কালো, নীল",
		"Dimensity 6100+ (6 nm)": "Dimensity ৬১০০+ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Gyro, Proximity": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 6100+": "MediaTek Dimensity ৬১০০+",
		"Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "Octa-core (২×২.০ GHz Cortex-A৭৫ + ৬×১.৮ GHz Cortex-A৫৫)",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-note-60'
func (s *SpecificationSeederMobileRealmeNote60) Seed(db *gorm.DB) error {
	productSlug := "realme-note-60"

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
