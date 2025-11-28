package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoM5 seeds specifications/options for product 'poco-m5'
type SpecificationSeederMobilePocoM5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoM5 creates a new seeder instance
func NewSpecificationSeederMobilePocoM5() *SpecificationSeederMobilePocoM5 {
	return &SpecificationSeederMobilePocoM5{BaseSeeder: BaseSeeder{name: "Specifications for POCO M5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoM5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2408 px": "১০৮০ × ২৪০৮ px",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"18 W wired": "১৮ W তারযুক্ত",
		"201 g": "২০১ g",
		"2022": "২০২২",
		"3.5 mm": "৩.৫ মিমি",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.0": "৫.০",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.58 inches": "৬.৫৮ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Android 12, MIUI 13": "Android ১২, MIUI ১৩",
		"Black, Green, Yellow": "কালো, সবুজ, হলুদ",
		"Helio G99 (6 nm)": "Helio G৯৯ (৬ nm)",
		"IPS LCD, 90 Hz": "IPS LCD, ৯০ Hz",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Helio G99": "MediaTek Helio G৯৯",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame / back, Glass front": "Plastic frame / back, গ্লাস সামনে",
		"Side fingerprint, Accelerometer, Gyro": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-m5'
func (s *SpecificationSeederMobilePocoM5) Seed(db *gorm.DB) error {
	productSlug := "poco-m5"

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
