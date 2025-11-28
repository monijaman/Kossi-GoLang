package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoReno12F5g seeds specifications/options for product 'oppo-reno-12-f-5g'
type SpecificationSeederMobileOppoReno12F5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoReno12F5g creates a new seeder instance
func NewSpecificationSeederMobileOppoReno12F5g() *SpecificationSeederMobileOppoReno12F5g {
	return &SpecificationSeederMobileOppoReno12F5g{BaseSeeder: BaseSeeder{name: "Specifications for Oppo Reno 12 F 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoReno12F5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"185 g": "১৮৫ g",
		"2023": "২০২৩",
		"2412 × 1080 px": "২৪১২ × ১০৮০ px",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 MP + 2 MP": "৬৪ MP + ২ MP",
		"90 Hz": "৯০ Hz",
		"AMOLED, 90 Hz": "AMOLED, ৯০ Hz",
		"Android 14, ColorOS 14": "Android ১৪, ColorOS ১৪",
		"Black, Blue": "কালো, নীল",
		"Dimensity 6100+ (6 nm)": "Dimensity ৬১০০+ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57": "Mali-G৫৭",
		"MediaTek Dimensity 6100+": "MediaTek Dimensity ৬১০০+",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-reno-12-f-5g'
func (s *SpecificationSeederMobileOppoReno12F5g) Seed(db *gorm.DB) error {
	productSlug := "oppo-reno-12-f-5g"

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
