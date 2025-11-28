package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyF15 seeds specifications/options for product 'samsung-galaxy-f15'
type SpecificationSeederMobileSamsungGalaxyF15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF15 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF15() *SpecificationSeederMobileSamsungGalaxyF15 {
	return &SpecificationSeederMobileSamsungGalaxyF15{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2408 px": "১০৮০ × ২৪০৮ px",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP": "১৩ MP",
		"15 W wired": "১৫ W তারযুক্ত",
		"164 × 76 × 8.5 mm": "১৬৪ × ৭৬ × ৮.৫ মিমি",
		"195 g": "১৯৫ g",
		"2023": "২০২৩",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.0": "৫.০",
		"50 MP + 2 MP + 2 MP": "৫০ MP + ২ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"90 Hz": "৯০ Hz",
		"Accelerometer, Gyro, Proximity, Compass": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Android 13, One UI Core 5": "Android ১৩, One UI Core ৫",
		"Black, Green, Blue": "কালো, সবুজ, নীল",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Helio G99": "Helio G৯৯",
		"IPS LCD, 90 Hz": "IPS LCD, ৯০ Hz",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Helio G99": "MediaTek Helio G৯৯",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame & back, glass front": "Plastic frame & back, গ্লাস সামনে",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (market dependent)",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-f15'
func (s *SpecificationSeederMobileSamsungGalaxyF15) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f15"

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
