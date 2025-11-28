package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoM6Plus seeds specifications/options for product 'poco-m6-plus'
type SpecificationSeederMobilePocoM6Plus struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoM6Plus creates a new seeder instance
func NewSpecificationSeederMobilePocoM6Plus() *SpecificationSeederMobilePocoM6Plus {
	return &SpecificationSeederMobilePocoM6Plus{BaseSeeder: BaseSeeder{name: "Specifications for POCO M6 Plus"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoM6Plus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2400 px": "১০৮০ × ২৪০০ px",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"16 MP": "১৬ MP",
		"202 g": "২০২ g",
		"2022": "২০২২",
		"3.5 mm": "৩.৫ মিমি",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 8 MP": "৫০ MP + ৮ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"Android 12, MIUI 13": "Android ১২, MIUI ১৩",
		"Black, Green": "কালো, সবুজ",
		"Helio G96 (12 nm)": "Helio G৯৬ (১২ nm)",
		"IPS LCD, 120 Hz": "IPS LCD, ১২০ Hz",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Helio G96": "MediaTek Helio G৯৬",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame / back, Glass front": "Plastic frame / back, গ্লাস সামনে",
		"Side fingerprint, Accelerometer, Gyro": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-m6-plus'
func (s *SpecificationSeederMobilePocoM6Plus) Seed(db *gorm.DB) error {
	productSlug := "poco-m6-plus"

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
