package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoM6Plus5g seeds specifications/options for product 'poco-m6-plus-5g'
type SpecificationSeederMobilePocoM6Plus5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoM6Plus5g creates a new seeder instance
func NewSpecificationSeederMobilePocoM6Plus5g() *SpecificationSeederMobilePocoM6Plus5g {
	return &SpecificationSeederMobilePocoM6Plus5g{BaseSeeder: BaseSeeder{name: "Specifications for POCO M6 Plus 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoM6Plus5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2400 px": "১০৮০ × ২৪০০ px",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"200 g": "২০০ g",
		"2023": "২০২৩",
		"3.5 mm": "৩.৫ মিমি",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ MP",
		"Android 13, MIUI 14": "Android ১৩, MIUI ১৪",
		"Black, Green, Yellow": "কালো, সবুজ, হলুদ",
		"Dimensity 900 (6 nm)": "Dimensity ৯০০ (৬ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"IPS LCD, 120 Hz": "IPS LCD, ১২০ Hz",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68": "Mali-G৬৮",
		"MediaTek Dimensity 900": "MediaTek Dimensity ৯০০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame / back, Glass front": "Plastic frame / back, গ্লাস সামনে",
		"Side fingerprint, Accelerometer, Gyro, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-m6-plus-5g'
func (s *SpecificationSeederMobilePocoM6Plus5g) Seed(db *gorm.DB) error {
	productSlug := "poco-m6-plus-5g"

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
