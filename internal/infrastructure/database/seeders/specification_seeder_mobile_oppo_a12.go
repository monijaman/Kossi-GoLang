package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoA12 seeds specifications/options for product 'oppo-a12'
type SpecificationSeederMobileOppoA12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA12 creates a new seeder instance
func NewSpecificationSeederMobileOppoA12() *SpecificationSeederMobileOppoA12 {
	return &SpecificationSeederMobileOppoA12{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"1520 × 720 px (~270 ppi)": "১৫২০ × ৭২০ px (~২৭০ ppi)",
		"165 g": "১৬৫ g",
		"2020": "২০২০",
		"3 / 4 GB": "৩ / ৪ GB",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"32 / 64 GB + microSD": "৩২ / ৬৪ GB + microSD",
		"4.2": "৪.২",
		"4230 mAh": "৪২৩০ এমএএইচ",
		"5 MP": "৫ MP",
		"6.22 inches": "৬.২২ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"Accelerometer, Proximity": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Android 9, ColorOS 6": "Android ৯, ColorOS ৬",
		"Black, Blue, Red": "কালো, নীল, লাল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Helio P35 (12 nm)": "Helio P৩৫ (১২ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Helio P35": "MediaTek Helio P৩৫",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
		"Wi-Fi 802.11 a/b/g/n": "Wi-Fi ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-a12'
func (s *SpecificationSeederMobileOppoA12) Seed(db *gorm.DB) error {
	productSlug := "oppo-a12"

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
