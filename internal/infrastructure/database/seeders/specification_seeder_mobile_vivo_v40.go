package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoV40 seeds specifications/options for product 'vivo-v40'
type SpecificationSeederMobileVivoV40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV40 creates a new seeder instance
func NewSpecificationSeederMobileVivoV40() *SpecificationSeederMobileVivoV40 {
	return &SpecificationSeederMobileVivoV40{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV40) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"179 g": "১৭৯ g",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"32 MP": "৩২ MP",
		"4200 mAh": "৪২০০ এমএএইচ",
		"44 W wired": "৪৪ W তারযুক্ত",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.2": "৫.২",
		"6.44 inches": "৬.৪৪ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED display, 120 Hz refresh rate": "AMOLED display, ১২০ Hz refresh rate",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 12, Funtouch OS 12": "Android ১২, Funtouch OS ১২",
		"Black, Blue, Silver": "কালো, নীল, রূপালী",
		"Dimensity 920 (6 nm)": "Dimensity ৯২০ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68 MC4": "Mali-G৬৮ MC৪",
		"May 2022": "May ২০২২",
		"MediaTek Dimensity 920": "MediaTek Dimensity ৯২০",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-v40'
func (s *SpecificationSeederMobileVivoV40) Seed(db *gorm.DB) error {
	productSlug := "vivo-v40"

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
