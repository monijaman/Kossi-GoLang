package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoX7Pro5g seeds specifications/options for product 'poco-x7-pro-5g'
type SpecificationSeederMobilePocoX7Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoX7Pro5g creates a new seeder instance
func NewSpecificationSeederMobilePocoX7Pro5g() *SpecificationSeederMobilePocoX7Pro5g {
	return &SpecificationSeederMobilePocoX7Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Poco X7 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoX7Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"160.75 × 75.24 × ~8.3 mm": "১৬০.৭৫ × ৭৫.২৪ × ~৮.৩ মিমি",
		"195 / 198 g": "১৯৫ / ১৯৮ g",
		"2025": "২০২৫",
		"256 / 512 GB (UFS 4.0)": "২৫৬ / ৫১২ GB (UFS ৪.০)",
		"2712 × 1220 px (1.5K)": "২৭১২ × ১২২০ px (১.৫K)",
		"4K, 1080p": "৪K, ১০৮০p",
		"5.4": "৫.৪",
		"50 MP + 8 MP + additional": "৫০ MP + ৮ MP + additional",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 / 12 GB": "৮ / ১২ GB",
		"802.11 a/b/g/n/ac/ax": "৮০২.১১ a/b/g/n/ac/ax",
		"Black, Green, Yellow": "কালো, সবুজ, হলুদ",
		"Dimensity 8400 Ultra (4 nm)": "Dimensity ৮৪০০ Ultra (৪ nm)",
		"Dolby Atmos, 360° ambient light": "Dolby Atmos, ৩৬০° ambient light",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic / eco-leather back": "গ্লাস সামনে, plastic / eco-leather back",
		"HyperOS 2": "HyperOS ২",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali‑G720": "Mali‑G৭২০",
		"MediaTek Dimensity 8400 Ultra": "MediaTek Dimensity ৮৪০০ Ultra",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (up to 3.25 GHz)": "Octa-core (up to ৩.২৫ GHz)",
		"Proximity, Accelerometer, Gyro, Compass, IR": "প্রক্সিমিটি, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস, IR",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-x7-pro-5g'
func (s *SpecificationSeederMobilePocoX7Pro5g) Seed(db *gorm.DB) error {
	productSlug := "poco-x7-pro-5g"

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
