package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoM7Pro5g seeds specifications/options for product 'poco-m7-pro-5g'
type SpecificationSeederMobilePocoM7Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoM7Pro5g creates a new seeder instance
func NewSpecificationSeederMobilePocoM7Pro5g() *SpecificationSeederMobilePocoM7Pro5g {
	return &SpecificationSeederMobilePocoM7Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for POCO M7 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoM7Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2400 px": "১০৮০ × ২৪০০ px",
		"1080p @ 30/60fps": "১০৮০p @ ৩০/৬০fps",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"190 g": "১৯০ g",
		"2024": "২০২৪",
		"3.5 mm": "৩.৫ মিমি",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ MP",
		"Android 13, MIUI 14": "Android ১৩, MIUI ১৪",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 930 (6 nm)": "Dimensity ৯৩০ (৬ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"IPS LCD, 120 Hz": "IPS LCD, ১২০ Hz",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G68": "Mali-G৬৮",
		"MediaTek Dimensity 930": "MediaTek Dimensity ৯৩০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (2×2.2 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.২ GHz Cortex-A৭৮ + ৬×২.০ GHz Cortex-A৫৫)",
		"Plastic frame / back, Glass front": "Plastic frame / back, গ্লাস সামনে",
		"Side fingerprint, Accelerometer, Gyro, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-m7-pro-5g'
func (s *SpecificationSeederMobilePocoM7Pro5g) Seed(db *gorm.DB) error {
	productSlug := "poco-m7-pro-5g"

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
