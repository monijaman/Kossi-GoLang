package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyF165g seeds specifications/options for product 'samsung-galaxy-f16-5g'
type SpecificationSeederMobileSamsungGalaxyF165g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF165g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF165g() *SpecificationSeederMobileSamsungGalaxyF165g {
	return &SpecificationSeederMobileSamsungGalaxyF165g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F16 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF165g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2408 px": "১০৮০ × ২৪০৮ px",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"15 W wired": "১৫ W তারযুক্ত",
		"164 × 76 × 8.4 mm": "১৬৪ × ৭৬ × ৮.৪ মিমি",
		"190 g": "১৯০ g",
		"2024": "২০২৪",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.1": "৫.১",
		"50 MP + 2 MP + 2 MP": "৫০ MP + ২ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Accelerometer, Proximity, Compass, Side fingerprint": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস, Side ফিঙ্গারপ্রিন্ট",
		"Android 13, One UI Core 5": "Android ১৩, One UI Core ৫",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dimensity 6100+": "Dimensity ৬১০০+",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 6100+": "MediaTek Dimensity ৬১০০+",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (2×2.0 GHz + 6×1.8 GHz)": "Octa-core (২×২.০ GHz + ৬×১.৮ GHz)",
		"PLS LCD, 90 Hz": "PLS LCD, ৯০ Hz",
		"Plastic frame & back, glass front": "Plastic frame & back, গ্লাস সামনে",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
		"Yes (market dependent)": "হ্যাঁ (market dependent)",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-f16-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF165g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f16-5g"

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
