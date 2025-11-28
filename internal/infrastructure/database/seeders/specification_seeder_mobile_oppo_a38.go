package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoA38 seeds specifications/options for product 'oppo-a38'
type SpecificationSeederMobileOppoA38 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA38 creates a new seeder instance
func NewSpecificationSeederMobileOppoA38() *SpecificationSeederMobileOppoA38 {
	return &SpecificationSeederMobileOppoA38{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A38"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA38) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"1600 × 720 px": "১৬০০ × ৭২০ px",
		"166 g": "১৬৬ g",
		"2019": "২০১৯",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"3200 mAh": "৩২০০ এমএএইচ",
		"4 GB": "৪ GB",
		"4.2": "৪.২",
		"6.4 inches": "৬.৪ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"8 MP": "৮ MP",
		"Android 9, ColorOS 6": "Android ৯, ColorOS ৬",
		"Black, Blue": "কালো, নীল",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Helio P35 (12 nm)": "Helio P৩৫ (১২ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"MediaTek Helio P35": "MediaTek Helio P৩৫",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
		"Wi-Fi 802.11 a/b/g/n": "Wi-Fi ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-a38'
func (s *SpecificationSeederMobileOppoA38) Seed(db *gorm.DB) error {
	productSlug := "oppo-a38"

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
