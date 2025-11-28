package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoC75 seeds specifications/options for product 'poco-c75'
type SpecificationSeederMobilePocoC75 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoC75 creates a new seeder instance
func NewSpecificationSeederMobilePocoC75() *SpecificationSeederMobilePocoC75 {
	return &SpecificationSeederMobilePocoC75{BaseSeeder: BaseSeeder{name: "Specifications for POCO C75"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoC75) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2023": "২০২৩",
		"3 / 4 GB": "৩ / ৪ GB",
		"3.5 mm": "৩.৫ মিমি",
		"5.0": "৫.০",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"64 / 128 GB": "৬৪ / ১২৮ GB",
		"720 × 1600 px": "৭২০ × ১৬০০ px",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Accelerometer, Proximity, Compass": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Android 12, MIUI 13": "Android ১২, MIUI ১৩",
		"Black, Blue, Red": "কালো, নীল, লাল",
		"Helio G88 (12 nm)": "Helio G৮৮ (১২ nm)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G52 MC2": "Mali-G৫২ MC২",
		"MediaTek Helio G88": "MediaTek Helio G৮৮",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame / back, Glass front": "Plastic frame / back, গ্লাস সামনে",
		"Wi-Fi 802.11 b/g/n": "Wi-Fi ৮০২.১১ b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-c75'
func (s *SpecificationSeederMobilePocoC75) Seed(db *gorm.DB) error {
	productSlug := "poco-c75"

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
