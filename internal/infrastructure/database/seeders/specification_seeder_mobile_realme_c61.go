package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRealmeC61 seeds specifications/options for product 'realme-c61'
type SpecificationSeederMobileRealmeC61 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealmeC61 creates a new seeder instance
func NewSpecificationSeederMobileRealmeC61() *SpecificationSeederMobileRealmeC61 {
	return &SpecificationSeederMobileRealmeC61{BaseSeeder: BaseSeeder{name: "Specifications for Realme C61"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealmeC61) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"1600 × 720 px (~270 ppi)": "১৬০০ × ৭২০ px (~২৭০ ppi)",
		"182 g": "১৮২ g",
		"2022": "২০২২",
		"3 / 4 GB": "৩ / ৪ GB",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"32 / 64 GB": "৩২ / ৬৪ GB",
		"5 MP": "৫ MP",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"Android 12, Realme UI Go Edition": "Android ১২, Realme UI Go Edition",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57": "Mali-G৫৭",
		"Octa-core (2×1.8 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "Octa-core (২×১.৮ GHz Cortex-A৭৫ + ৬×১.৮ GHz Cortex-A৫৫)",
		"T612 (12 nm)": "T৬১২ (১২ nm)",
		"Unisoc T612": "Unisoc T৬১২",
		"Wi-Fi 802.11 a/b/g/n": "Wi-Fi ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'realme-c61'
func (s *SpecificationSeederMobileRealmeC61) Seed(db *gorm.DB) error {
	productSlug := "realme-c61"

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
