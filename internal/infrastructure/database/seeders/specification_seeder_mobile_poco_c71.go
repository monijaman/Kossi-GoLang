package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoC71 seeds specifications/options for product 'poco-c71'
type SpecificationSeederMobilePocoC71 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoC71 creates a new seeder instance
func NewSpecificationSeederMobilePocoC71() *SpecificationSeederMobilePocoC71 {
	return &SpecificationSeederMobilePocoC71{BaseSeeder: BaseSeeder{name: "Specifications for POCO C71"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoC71) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"18 W wired": "১৮ W তারযুক্ত",
		"190 g": "১৯০ g",
		"2 / 3 GB": "২ / ৩ GB",
		"2023": "২০২৩",
		"3.5 mm": "৩.৫ মিমি",
		"32 / 64 GB": "৩২ / ৬৪ GB",
		"5 MP": "৫ MP",
		"5.0": "৫.০",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60 Hz": "৬০ Hz",
		"720 × 1600 px": "৭২০ × ১৬০০ px",
		"Accelerometer, Proximity, Compass": "অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Android 12, MIUI 13": "Android ১২, MIUI ১৩",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57": "Mali-G৫৭",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Plastic frame / back, Glass front": "Plastic frame / back, গ্লাস সামনে",
		"T606 (12 nm)": "T৬০৬ (১২ nm)",
		"Unisoc T606": "Unisoc T৬০৬",
		"Wi-Fi 802.11 b/g/n": "Wi-Fi ৮০২.১১ b/g/n",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-c71'
func (s *SpecificationSeederMobilePocoC71) Seed(db *gorm.DB) error {
	productSlug := "poco-c71"

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
