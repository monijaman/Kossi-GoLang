package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoF65g seeds specifications/options for product 'poco-f6-5g'
type SpecificationSeederMobilePocoF65g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoF65g creates a new seeder instance
func NewSpecificationSeederMobilePocoF65g() *SpecificationSeederMobilePocoF65g {
	return &SpecificationSeederMobilePocoF65g{BaseSeeder: BaseSeeder{name: "Specifications for POCO F6 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoF65g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"198 g": "১৯৮ g",
		"20 MP": "২০ MP",
		"2022": "২০২২",
		"2400 × 1080 px": "২৪০০ × ১০৮০ px",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4700 mAh": "৪৭০০ এমএএইচ",
		"4K @ 30/60fps; 1080p @ 30/60/120fps": "৪K @ ৩০/৬০fps; ১০৮০p @ ৩০/৬০/১২০fps",
		"5.1": "৫.১",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"AMOLED, 120 Hz": "AMOLED, ১২০ Hz",
		"Adreno 650": "Adreno ৬৫০",
		"Android 12, MIUI 13": "Android ১২, MIUI ১৩",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Fingerprint (side), Accelerometer, Gyro, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (1×3.2 GHz Cortex-A78 + 3×2.42 GHz Cortex-A78 + 4×1.8 GHz Cortex-A55)": "Octa-core (১×৩.২ GHz Cortex-A৭৮ + ৩×২.৪২ GHz Cortex-A৭৮ + ৪×১.৮ GHz Cortex-A৫৫)",
		"Snapdragon 870": "Snapdragon ৮৭০",
		"Snapdragon 870 (7 nm)": "Snapdragon ৮৭০ (৭ nm)",
		"USB-C 2.0": "USB-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-f6-5g'
func (s *SpecificationSeederMobilePocoF65g) Seed(db *gorm.DB) error {
	productSlug := "poco-f6-5g"

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
