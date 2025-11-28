package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoV50 seeds specifications/options for product 'vivo-v50'
type SpecificationSeederMobileVivoV50 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV50 creates a new seeder instance
func NewSpecificationSeederMobileVivoV50() *SpecificationSeederMobileVivoV50 {
	return &SpecificationSeederMobileVivoV50{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V50"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV50) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"171 g": "১৭১ g",
		"2376 × 1080 px": "২৩৭৬ × ১০৮০ px",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4315 mAh": "৪৩১৫ এমএএইচ",
		"48 MP + 8 MP + 2 MP": "৪৮ MP + ৮ MP + ২ MP",
		"4K @ 30fps; 1080p @ 30/60fps": "৪K @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"5.1": "৫.১",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"8 GB": "৮ GB",
		"AMOLED display, 5G support": "AMOLED display, ৫G support",
		"AMOLED, 60 Hz": "AMOLED, ৬০ Hz",
		"Adreno 620": "Adreno ৬২০",
		"Android 11, Funtouch OS 11": "Android ১১, Funtouch OS ১১",
		"Black, Blue": "কালো, নীল",
		"December 2020": "December ২০২০",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"SM7250 (7 nm)": "SM৭২৫০ (৭ nm)",
		"Snapdragon 765G": "Snapdragon ৭৬৫G",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-v50'
func (s *SpecificationSeederMobileVivoV50) Seed(db *gorm.DB) error {
	productSlug := "vivo-v50"

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
