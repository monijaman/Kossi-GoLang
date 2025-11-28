package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM155g seeds specifications/options for product 'samsung-galaxy-m15-5g'
type SpecificationSeederMobileSamsungGalaxyM155g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM155g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM155g() *SpecificationSeederMobileSamsungGalaxyM155g {
	return &SpecificationSeederMobileSamsungGalaxyM155g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M15 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM155g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"13 MP": "১৩ MP",
		"160.1 × 76.8 × 9.3 mm": "১৬০.১ × ৭৬.৮ × ৯.৩ মিমি",
		"217 g": "২১৭ g",
		"2340 × 1080 px": "২৩৪০ × ১০৮০ px",
		"3.5 mm": "৩.৫ মিমি",
		"5.3": "৫.৩",
		"50 MP + 5 MP + 2 MP": "৫০ MP + ৫ MP + ২ MP",
		"6 / 8 GB": "৬ / ৮ GB",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"90 Hz": "৯০ Hz",
		"Accelerometer, Gyro, Proximity, Compass, Fingerprint": "অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট",
		"Available / Unofficial": "উপলব্ধ / Unofficial",
		"Dimensity 810 (7 nm)": "Dimensity ৮১০ (৭ nm)",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame / back": "গ্লাস সামনে, plastic frame / back",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 810": "MediaTek Dimensity ৮১০",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m15-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM155g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m15-5g"

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
