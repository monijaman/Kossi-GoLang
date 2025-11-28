package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoReno14F5g seeds specifications/options for product 'oppo-reno-14-f-5g'
type SpecificationSeederMobileOppoReno14F5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoReno14F5g creates a new seeder instance
func NewSpecificationSeederMobileOppoReno14F5g() *SpecificationSeederMobileOppoReno14F5g {
	return &SpecificationSeederMobileOppoReno14F5g{BaseSeeder: BaseSeeder{name: "Specifications for Oppo Reno 14 F 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoReno14F5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"188 g": "১৮৮ g",
		"2024": "২০২৪",
		"2412 × 1080 px": "২৪১২ × ১০৮০ px",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 14, ColorOS 14": "Android ১৪, ColorOS ১৪",
		"Black, Blue": "কালো, নীল",
		"Dimensity 6100+ (6 nm)": "Dimensity ৬১০০+ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"HDR10+ display": "HDR১০+ display",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 6100+": "MediaTek Dimensity ৬১০০+",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-reno-14-f-5g'
func (s *SpecificationSeederMobileOppoReno14F5g) Seed(db *gorm.DB) error {
	productSlug := "oppo-reno-14-f-5g"

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
