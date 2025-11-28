package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoA3x seeds specifications/options for product 'oppo-a3x'
type SpecificationSeederMobileOppoA3x struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA3x creates a new seeder instance
func NewSpecificationSeederMobileOppoA3x() *SpecificationSeederMobileOppoA3x {
	return &SpecificationSeederMobileOppoA3x{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A3x"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA3x) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"1520 × 720 px": "১৫২০ × ৭২০ px",
		"170 g": "১৭০ g",
		"2019": "২০১৯",
		"3 / 4 GB": "৩ / ৪ GB",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"32 / 64 GB + microSD": "৩২ / ৬৪ GB + microSD",
		"4.2": "৪.২",
		"4230 mAh": "৪২৩০ এমএএইচ",
		"6.2 inches": "৬.২ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"8 MP": "৮ MP",
		"Android 9, ColorOS 6": "Android ৯, ColorOS ৬",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Helio P22 (12 nm)": "Helio P২২ (১২ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Helio P22": "MediaTek Helio P২২",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
		"Wi-Fi 802.11 a/b/g/n": "Wi-Fi ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-a3x'
func (s *SpecificationSeederMobileOppoA3x) Seed(db *gorm.DB) error {
	productSlug := "oppo-a3x"

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
