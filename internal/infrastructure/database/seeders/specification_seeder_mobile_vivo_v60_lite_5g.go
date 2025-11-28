package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoV60Lite5g seeds specifications/options for product 'vivo-v60-lite-5g'
type SpecificationSeederMobileVivoV60Lite5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV60Lite5g creates a new seeder instance
func NewSpecificationSeederMobileVivoV60Lite5g() *SpecificationSeederMobileVivoV60Lite5g {
	return &SpecificationSeederMobileVivoV60Lite5g{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V60 Lite 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV60Lite5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"177 g": "১৭৭ g",
		"18 W wired": "১৮ W তারযুক্ত",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"4100 mAh": "৪১০০ এমএএইচ",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.44 inches": "৬.৪৪ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"90 Hz": "৯০ Hz",
		"AMOLED, 90 Hz": "AMOLED, ৯০ Hz",
		"Android 12, Funtouch OS 12": "Android ১২, Funtouch OS ১২",
		"Black, Blue": "কালো, নীল",
		"Dimensity 900 (6 nm)": "Dimensity ৯০০ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"June 2021": "June ২০২১",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "Mali-G৬৮ MC৪",
		"MediaTek Dimensity 900": "MediaTek Dimensity ৯০০",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-v60-lite-5g'
func (s *SpecificationSeederMobileVivoV60Lite5g) Seed(db *gorm.DB) error {
	productSlug := "vivo-v60-lite-5g"

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
