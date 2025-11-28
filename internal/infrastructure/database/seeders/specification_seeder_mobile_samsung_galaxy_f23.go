package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyF23 seeds specifications/options for product 'samsung-galaxy-f23'
type SpecificationSeederMobileSamsungGalaxyF23 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF23 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF23() *SpecificationSeederMobileSamsungGalaxyF23 {
	return &SpecificationSeederMobileSamsungGalaxyF23{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F23"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF23) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2408 px": "১০৮০ × ২৪০৮ px",
		"1080p @ 30/60fps": "১০৮০p @ ৩০/৬০fps",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 GB + microSD": "১২৮ GB + microSD",
		"13 MP": "১৩ MP",
		"164 × 76 × 8.4 mm": "১৬৪ × ৭৬ × ৮.৪ মিমি",
		"198 g": "১৯৮ g",
		"2023": "২০২৩",
		"25 W wired": "২৫ W তারযুক্ত",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.0": "৫.০",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"Adreno 619": "Adreno ৬১৯",
		"Android 13, One UI 5.1": "Android ১৩, One UI ৫.১",
		"Black, Green, Copper": "কালো, সবুজ, Copper",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"IPS LCD, 120 Hz": "IPS LCD, ১২০ Hz",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Side fingerprint, Accelerometer, Gyro, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"Snapdragon 750G": "Snapdragon ৭৫০G",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-f23'
func (s *SpecificationSeederMobileSamsungGalaxyF23) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f23"

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
