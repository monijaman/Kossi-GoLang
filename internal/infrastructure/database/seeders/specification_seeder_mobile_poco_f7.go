package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoF7 seeds specifications/options for product 'poco-f7'
type SpecificationSeederMobilePocoF7 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoF7 creates a new seeder instance
func NewSpecificationSeederMobilePocoF7() *SpecificationSeederMobilePocoF7 {
	return &SpecificationSeederMobilePocoF7{BaseSeeder: BaseSeeder{name: "Specifications for POCO F7"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoF7) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"196 g": "১৯৬ g",
		"20 MP": "২০ MP",
		"2023": "২০২৩",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"4K @ 30/60fps; 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/৬০/১২০fps",
		"5.2": "৫.২",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 12, MIUI 13": "Android ১২, MIUI ১৩",
		"Black, Silver, Blue": "কালো, রূপালী, নীল",
		"Dimensity 1200 (6 nm)": "Dimensity ১২০০ (৬ nm)",
		"Fingerprint (side), Accelerometer, Gyro, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G77 MC9": "Mali-G৭৭ MC৯",
		"MediaTek Dimensity 1200": "MediaTek Dimensity ১২০০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-f7'
func (s *SpecificationSeederMobilePocoF7) Seed(db *gorm.DB) error {
	productSlug := "poco-f7"

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
