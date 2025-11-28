package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoX300 seeds specifications/options for product 'vivo-x300'
type SpecificationSeederMobileVivoX300 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX300 creates a new seeder instance
func NewSpecificationSeederMobileVivoX300() *SpecificationSeederMobileVivoX300 {
	return &SpecificationSeederMobileVivoX300{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X300"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX300) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"202 g": "২০২ g",
		"2x + 5x": "২x + ৫x",
		"32 MP": "৩২ MP",
		"3200 × 1440 px": "৩২০০ × ১৪৪০ px",
		"4400 mAh": "৪৪০০ এমএএইচ",
		"48 MP + 13 MP + 13 MP": "৪৮ MP + ১৩ MP + ১৩ MP",
		"4K @ 30/60fps; 1080p @ 30/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/১২০fps",
		"5.2": "৫.২",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"66 W wired": "৬৬ W তারযুক্ত",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Adreno 650": "Adreno ৬৫০",
		"Android 12, ফানটাচ OS 12": "Android ১২, ফানটাচ OS ১২",
		"Black, Blue, Silver": "কালো, নীল, রূপালী",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"HDR10+, AMOLED display": "HDR১০+, AMOLED display",
		"January 2023": "January ২০২৩",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Snapdragon 870": "Snapdragon ৮৭০",
		"Snapdragon 870 (7 nm)": "Snapdragon ৮৭০ (৭ nm)",
		"USB-C 3.1": "USB-C ৩.১",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-x300'
func (s *SpecificationSeederMobileVivoX300) Seed(db *gorm.DB) error {
	productSlug := "vivo-x300"

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
