package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmi13c seeds specifications/options for product 'redmi-13c'
type SpecificationSeederMobileRedmi13c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi13c creates a new seeder instance
func NewSpecificationSeederMobileRedmi13c() *SpecificationSeederMobileRedmi13c {
	return &SpecificationSeederMobileRedmi13c{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 13C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi13c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W": "১০ W",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"1600 × 720 px": "১৬০০ × ৭২০ px",
		"168 × 78 × 8.09 mm": "১৬৮ × ৭৮ × ৮.০৯ মিমি",
		"192 g": "১৯২ g",
		"2023": "২০২৩",
		"3.5 mm": "৩.৫ মিমি",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.3": "৫.৩",
		"50 MP + 2 MP + tiny lens": "৫০ MP + ২ MP + tiny lens",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.74 inches": "৬.৭৪ ইঞ্চি",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Fingerprint (side), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Helio G85": "Helio G৮৫",
		"Li‑Ion (non-removable)": "Li‑Ion (অপসারণযোগ্য নয়)",
		"MIUI 14": "MIUI ১৪",
		"Mali‑G52 MC2": "Mali‑G৫২ MC২",
		"MediaTek Helio G85": "MediaTek Helio G৮৫",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (2×Cortex-A75 + 6×Cortex-A55)": "Octa-core (২×Cortex-A৭৫ + ৬×Cortex-A৫৫)",
		"Plastic frame / back, glass front": "Plastic frame / back, গ্লাস সামনে",
		"Wi‑Fi 802.11 a/b/g/n/ac": "Wi‑Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-13c'
func (s *SpecificationSeederMobileRedmi13c) Seed(db *gorm.DB) error {
	productSlug := "redmi-13c"

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
