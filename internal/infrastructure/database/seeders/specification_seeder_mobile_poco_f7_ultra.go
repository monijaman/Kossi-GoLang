package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoF7Ultra seeds specifications/options for product 'poco-f7-ultra'
type SpecificationSeederMobilePocoF7Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoF7Ultra creates a new seeder instance
func NewSpecificationSeederMobilePocoF7Ultra() *SpecificationSeederMobilePocoF7Ultra {
	return &SpecificationSeederMobilePocoF7Ultra{BaseSeeder: BaseSeeder{name: "Specifications for POCO F7 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoF7Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 60fps": "১০৮০p @ ৬০fps",
		"120 Hz": "১২০ Hz",
		"120 W wired": "১২০ W তারযুক্ত",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"200 MP + 13 MP + 5 MP": "২০০ MP + ১৩ MP + ৫ MP",
		"2024": "২০২৪",
		"210 g": "২১০ g",
		"32 MP": "৩২ MP",
		"3200 × 1440 px": "৩২০০ × ১৪৪০ px",
		"5.3": "৫.৩",
		"5000 mAh": "৫০০০ এমএএইচ",
		"5x": "৫x",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 / 12 GB": "৮ / ১২ GB",
		"8K @ 30fps; 4K @ 60fps; 1080p @ 120fps": "৮K @ ৩০fps; ৪K @ ৬০fps; ১০৮০p @ ১২০fps",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Android 13, MIUI 14": "Android ১৩, MIUI ১৪",
		"Black, Blue": "কালো, নীল",
		"Dimensity 9200+ (4 nm)": "Dimensity ৯২০০+ (৪ nm)",
		"Fingerprint (under-display), Accelerometer, Gyro, Compass": "ফিঙ্গারপ্রিন্ট (under-display), অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G710 MC10": "Mali-G৭১০ MC১০",
		"MediaTek Dimensity 9200+": "MediaTek Dimensity ৯২০০+",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"USB-C 3.1": "USB-C ৩.১",
		"Wi-Fi 802.11 a/b/g/n/ac/6e": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-f7-ultra'
func (s *SpecificationSeederMobilePocoF7Ultra) Seed(db *gorm.DB) error {
	productSlug := "poco-f7-ultra"

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
