package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyM165g seeds specifications/options for product 'samsung-galaxy-m16-5g'
type SpecificationSeederMobileSamsungGalaxyM165g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM165g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM165g() *SpecificationSeederMobileSamsungGalaxyM165g {
	return &SpecificationSeederMobileSamsungGalaxyM165g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M16 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM165g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 × 2340 px": "১০৮০ × ২৩৪০ px",
		"128 GB (maybe more)": "১২৮ GB (maybe more)",
		"13 MP": "১৩ MP",
		"164.4 × 77.9 × 7.9 mm": "১৬৪.৪ × ৭৭.৯ × ৭.৯ মিমি",
		"191 g (according to 8 GB variant)": "১৯১ g (according to ৮ GB variant)",
		"25 W wired": "২৫ W তারযুক্ত",
		"5.3": "৫.৩",
		"50 MP + 5 MP + 2 MP": "৫০ MP + ৫ MP + ২ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 / 8 GB": "৬ / ৮ GB",
		"6 years of updates (claimed)": "৬ years of updates (claimed)",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"90 Hz": "৯০ Hz",
		"Android 15, One UI 7": "Android ১৫, One UI ৭",
		"Dimensity 6300 (6 nm)": "Dimensity ৬৩০০ (৬ nm)",
		"February 2025": "February ২০২৫",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back / frame": "গ্লাস সামনে, প্লাস্টিক পেছনে / frame",
		"IP54 (splash)": "IP৫৪ (splash)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek Dimensity 6300": "MediaTek Dimensity ৬৩০০",
		"Nano-SIM + Nano-SIM (হাইব্রিড)": "ন্যানো-সিম + ন্যানো-সিম (হাইব্রিড)",
		"Octa-core (2×2.4 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "Octa-core (২×২.৪ GHz Cortex-A৭৬ + ৬×২.০ GHz Cortex-A৫৫)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Thunder Black, মিন্ট Green, Blush Pink": "Thunder কালো, মিন্ট সবুজ, Blush গোলাপী",
		"USB-C (no 3.5mm)": "USB-C (no ৩.৫mm)",
		"Wi-Fi 802.11 a/b/g/n/ac": "Wi-Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-m16-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM165g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m16-5g"

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
