package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoA57 seeds specifications/options for product 'oppo-a57'
type SpecificationSeederMobileOppoA57 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA57 creates a new seeder instance
func NewSpecificationSeederMobileOppoA57() *SpecificationSeederMobileOppoA57 {
	return &SpecificationSeederMobileOppoA57{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A57"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA57) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"1612 × 720 px": "১৬১২ × ৭২০ px",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2022": "২০২২",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"4 / 6 GB": "৪ / ৬ GB",
		"5.0": "৫.০",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"8 MP": "৮ MP",
		"Android 11, ColorOS 11": "Android ১১, ColorOS ১১",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (side), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Helio G35 (12 nm)": "Helio G৩৫ (১২ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Helio G35": "MediaTek Helio G৩৫",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
		"Wi-Fi 802.11 a/b/g/n": "Wi-Fi ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-a57'
func (s *SpecificationSeederMobileOppoA57) Seed(db *gorm.DB) error {
	productSlug := "oppo-a57"

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
