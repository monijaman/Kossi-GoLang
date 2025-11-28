package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealmeC71 seeds specifications/options for product 'realme-c71'
type SpecificationSeederMobileRealmeC71 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeC71 creates a new seeder instance
func NewSpecificationSeederMobileRealmeC71() *SpecificationSeederMobileRealmeC71 {
	return &SpecificationSeederMobileRealmeC71{BaseSeeder: BaseSeeder{name: "Specifications for Realme C71"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeC71) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"1600 × 720 px (~269 ppi)": "১৬০০ × ৭২০ px (~২৬৯ ppi)",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2023": "২০২৩",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"4 GB": "৪ GB",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"8 MP": "৮ MP",
		"Android 12, Realme UI 3.0": "Android ১২, Realme UI ৩.০",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Helio G88 (12 nm)": "Helio G৮৮ (১২ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"MediaTek Helio G88": "MediaTek Helio G৮৮",
		"Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "Octa-core (২×২.০ GHz Cortex-A৭৫ + ৬×১.৮ GHz Cortex-A৫৫)",
		"Wi-Fi 802.11 a/b/g/n": "Wi-Fi ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-c71'
func (s *SpecificationSeederMobileRealmeC71) Seed(db *gorm.DB) error {
	productSlug := "realme-c71"

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
