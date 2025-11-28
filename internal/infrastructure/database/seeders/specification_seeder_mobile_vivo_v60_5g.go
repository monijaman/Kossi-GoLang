package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoV605g seeds specifications/options for product 'vivo-v60-5g'
type SpecificationSeederMobileVivoV605g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV605g creates a new seeder instance
func NewSpecificationSeederMobileVivoV605g() *SpecificationSeederMobileVivoV605g {
	return &SpecificationSeederMobileVivoV605g{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V60 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV605g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"176 g": "১৭৬ g",
		"2376 × 1080 px": "২৩৭৬ × ১০৮০ px",
		"2x": "২x",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4300 mAh": "৪৩০০ এমএএইচ",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.1": "৫.১",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 11, ফানটাচ OS 11": "Android ১১, ফানটাচ OS ১১",
		"April 2021": "April ২০২১",
		"Black, Blue": "কালো, নীল",
		"Dimensity 800U (7 nm)": "Dimensity ৮০০U (৭ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC3": "Mali-G৫৭ MC৩",
		"MediaTek Dimensity 800U": "MediaTek Dimensity ৮০০U",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-v60-5g'
func (s *SpecificationSeederMobileVivoV605g) Seed(db *gorm.DB) error {
	productSlug := "vivo-v60-5g"

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
