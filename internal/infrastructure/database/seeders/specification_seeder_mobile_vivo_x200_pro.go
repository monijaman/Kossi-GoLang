package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoX200Pro seeds specifications/options for product 'vivo-x200-pro'
type SpecificationSeederMobileVivoX200Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX200Pro creates a new seeder instance
func NewSpecificationSeederMobileVivoX200Pro() *SpecificationSeederMobileVivoX200Pro {
	return &SpecificationSeederMobileVivoX200Pro{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X200 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX200Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"196 g": "১৯৬ g",
		"2376 × 1080 px": "২৩৭৬ × ১০৮০ px",
		"2x": "২x",
		"32 MP": "৩২ MP",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4200 mAh": "৪২০০ এমএএইচ",
		"4K @ 30/60fps; 1080p @ 30/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/১২০fps",
		"5.2": "৫.২",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"8 / 12 GB": "৮ / ১২ GB",
		"AMOLED HDR10+": "AMOLED HDR১০+",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 11, Funtouch OS 11": "Android ১১, Funtouch OS ১১",
		"Black, Silver": "কালো, রূপালী",
		"December 2022": "December ২০২২",
		"Dimensity 1200 (6 nm)": "Dimensity ১২০০ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G77 MC9": "Mali-G৭৭ MC৯",
		"MediaTek Dimensity 1200": "MediaTek Dimensity ১২০০",
		"USB-C 3.1": "USB-C ৩.১",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-x200-pro'
func (s *SpecificationSeederMobileVivoX200Pro) Seed(db *gorm.DB) error {
	productSlug := "vivo-x200-pro"

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
