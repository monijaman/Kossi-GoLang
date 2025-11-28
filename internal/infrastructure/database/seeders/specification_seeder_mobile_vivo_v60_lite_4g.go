package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoV60Lite4g seeds specifications/options for product 'vivo-v60-lite-4g'
type SpecificationSeederMobileVivoV60Lite4g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV60Lite4g creates a new seeder instance
func NewSpecificationSeederMobileVivoV60Lite4g() *SpecificationSeederMobileVivoV60Lite4g {
	return &SpecificationSeederMobileVivoV60Lite4g{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V60 Lite 4G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV60Lite4g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB": "১২৮ GB",
		"16 MP": "১৬ MP",
		"175 g": "১৭৫ g",
		"18 W wired": "১৮ W তারযুক্ত",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"4100 mAh": "৪১০০ এমএএইচ",
		"5.0": "৫.০",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.44 inches": "৬.৪৪ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"AMOLED, 60 Hz": "AMOLED, ৬০ Hz",
		"Android 11, Funtouch OS 11": "Android ১১, Funtouch OS ১১",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE (4G)": "GSM / HSPA / LTE (৪G)",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"Helio G95 (12 nm)": "Helio G৯৫ (১২ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G76 MC4": "Mali-G৭৬ MC৪",
		"May 2021": "May ২০২১",
		"MediaTek Helio G95": "MediaTek Helio G৯৫",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-v60-lite-4g'
func (s *SpecificationSeederMobileVivoV60Lite4g) Seed(db *gorm.DB) error {
	productSlug := "vivo-v60-lite-4g"

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
