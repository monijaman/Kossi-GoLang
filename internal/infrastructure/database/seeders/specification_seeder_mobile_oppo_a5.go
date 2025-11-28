package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOppoA5 seeds specifications/options for product 'oppo-a5'
type SpecificationSeederMobileOppoA5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOppoA5 creates a new seeder instance
func NewSpecificationSeederMobileOppoA5() *SpecificationSeederMobileOppoA5 {
	return &SpecificationSeederMobileOppoA5{BaseSeeder: BaseSeeder{name: "Specifications for Oppo A5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOppoA5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W wired": "১০ W তারযুক্ত",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"12 MP + 2 MP + 2 MP + 2 MP": "১২ MP + ২ MP + ২ MP + ২ MP",
		"1600 × 720 px": "১৬০০ × ৭২০ px",
		"192 g": "১৯২ g",
		"2020": "২০২০",
		"3 / 4 GB": "৩ / ৪ GB",
		"3.5 mm headphone jack": "৩.৫ মিমি headphone jack",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ GB + microSD",
		"8 MP": "৮ MP",
		"Adreno 610": "Adreno ৬১০",
		"Android 10, ColorOS 7": "Android ১০, ColorOS ৭",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Dual SIM (Nano-SIM)": "ডুয়াল সিম (ন্যানো-সিম)",
		"Fingerprint (rear), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (rear), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Glass front, plastic frame/back": "গ্লাস সামনে, plastic frame/back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Snapdragon 665": "Snapdragon ৬৬৫",
		"Snapdragon 665 (11 nm)": "Snapdragon ৬৬৫ (১১ nm)",
		"Wi-Fi 802.11 a/b/g/n": "Wi-Fi ৮০২.১১ a/b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oppo-a5'
func (s *SpecificationSeederMobileOppoA5) Seed(db *gorm.DB) error {
	productSlug := "oppo-a5"

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
