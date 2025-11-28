package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileVivoV40Lite seeds specifications/options for product 'vivo-v40-lite'
type SpecificationSeederMobileVivoV40Lite struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV40Lite creates a new seeder instance
func NewSpecificationSeederMobileVivoV40Lite() *SpecificationSeederMobileVivoV40Lite {
	return &SpecificationSeederMobileVivoV40Lite{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V40 Lite"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV40Lite) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 GB": "১২৮ GB",
		"16 MP": "১৬ MP",
		"176 g": "১৭৬ g",
		"18 W wired": "১৮ W তারযুক্ত",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"4100 mAh": "৪১০০ এমএএইচ",
		"48 MP + 2 MP": "৪৮ MP + ২ MP",
		"5.1": "৫.১",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.38 inches": "৬.৩৮ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"AMOLED, 60 Hz": "AMOLED, ৬০ Hz",
		"Android 12, Funtouch OS 12": "Android ১২, Funtouch OS ১২",
		"Black, Blue, Pink": "কালো, নীল, গোলাপী",
		"Dimensity 810 (6 nm)": "Dimensity ৮১০ (৬ nm)",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side-mounted), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (পাশে মাউন্ট), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back/frame": "গ্লাস সামনে, প্লাস্টিক পেছনে/frame",
		"July 2022": "July ২০২২",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 810": "MediaTek Dimensity ৮১০",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'vivo-v40-lite'
func (s *SpecificationSeederMobileVivoV40Lite) Seed(db *gorm.DB) error {
	productSlug := "vivo-v40-lite"

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
