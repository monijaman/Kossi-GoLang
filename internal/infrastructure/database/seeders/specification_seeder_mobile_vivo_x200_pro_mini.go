package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoX200ProMini seeds specifications/options for product 'vivo-x200-pro-mini'
type SpecificationSeederMobileVivoX200ProMini struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX200ProMini creates a new seeder instance
func NewSpecificationSeederMobileVivoX200ProMini() *SpecificationSeederMobileVivoX200ProMini {
	return &SpecificationSeederMobileVivoX200ProMini{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X200 Pro Mini"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX200ProMini) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB": "১২৮ GB",
		"185 g": "১৮৫ g",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"2x": "২x",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4000 mAh": "৪০০০ এমএএইচ",
		"48 MP + 8 MP + 2 MP": "৪৮ MP + ৮ MP + ২ MP",
		"4K @ 30fps; 1080p @ 30fps": "৪K @ ৩০fps; ১০৮০p @ ৩০fps",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.44 inches": "৬.৪৪ ইঞ্চি",
		"90 Hz": "৯০ Hz",
		"AMOLED, 90 Hz": "AMOLED, ৯০ Hz",
		"Android 11, ফানটাচ OS 11": "Android ১১, ফানটাচ OS ১১",
		"Black, Blue": "কালো, নীল",
		"Dimensity 1100 (6 nm)": "Dimensity ১১০০ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G77 MC9": "Mali-G৭৭ MC৯",
		"MediaTek Dimensity 1100": "MediaTek Dimensity ১১০০",
		"November 2022": "November ২০২২",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-x200-pro-mini'
func (s *SpecificationSeederMobileVivoX200ProMini) Seed(db *gorm.DB) error {
	productSlug := "vivo-x200-pro-mini"

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
