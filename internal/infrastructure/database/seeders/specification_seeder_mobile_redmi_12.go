package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmi12 seeds specifications/options for product 'redmi-12'
type SpecificationSeederMobileRedmi12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi12 creates a new seeder instance
func NewSpecificationSeederMobileRedmi12() *SpecificationSeederMobileRedmi12 {
	return &SpecificationSeederMobileRedmi12{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"1650 × 720 px": "১৬৫০ × ৭২০ px",
		"168.8 × 76.5 × 8.9 mm": "১৬৮.৮ × ৭৬.৫ × ৮.৯ মিমি",
		"18 W wired": "১৮ W তারযুক্ত",
		"198 g": "১৯৮ g",
		"2022": "২০২২",
		"3.5 mm": "৩.৫ মিমি",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.71 inches": "৬.৭১ ইঞ্চি",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Android 13, MIUI 14": "Android ১৩, MIUI ১৪",
		"Black, Sky Blue, Green": "কালো, Sky নীল, সবুজ",
		"Fingerprint (side), Accelerometer, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, কম্পাস",
		"Helio G99": "Helio G৯৯",
		"IPS LCD, 90 Hz": "IPS LCD, ৯০ Hz",
		"Li‑Ion (non-removable)": "Li‑Ion (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Helio G99": "MediaTek Helio G৯৯",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Wi‑Fi 802.11 a/b/g/n/ac": "Wi‑Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-12'
func (s *SpecificationSeederMobileRedmi12) Seed(db *gorm.DB) error {
	productSlug := "redmi-12"

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
