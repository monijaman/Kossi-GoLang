package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoReno12Pro5g seeds specifications/options for product 'oppo-reno-12-pro-5g'
type SpecificationSeederMobileOppoReno12Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoReno12Pro5g creates a new seeder instance
func NewSpecificationSeederMobileOppoReno12Pro5g() *SpecificationSeederMobileOppoReno12Pro5g {
	return &SpecificationSeederMobileOppoReno12Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Oppo Reno 12 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoReno12Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 8 MP + 2 MP": "১০৮ MP + ৮ MP + ২ MP",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"195 g": "১৯৫ g",
		"2023": "২০২৩",
		"2412 × 1080 px": "২৪১২ × ১০৮০ px",
		"2x": "২x",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.3": "৫.৩",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 14, ColorOS 14": "Android ১৪, ColorOS ১৪",
		"Black, Silver, Blue": "কালো, রূপালী, নীল",
		"Dimensity 6100+ (6 nm)": "Dimensity ৬১০০+ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, aluminum frame/back": "গ্লাস সামনে, অ্যালুমিনিয়াম ফ্রেম/back",
		"HDR10+ display, Dolby Atmos": "HDR১০+ display, Dolby Atmos",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57": "Mali-G৫৭",
		"MediaTek Dimensity 6100+": "MediaTek Dimensity ৬১০০+",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-reno-12-pro-5g'
func (s *SpecificationSeederMobileOppoReno12Pro5g) Seed(db *gorm.DB) error {
	productSlug := "oppo-reno-12-pro-5g"

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
