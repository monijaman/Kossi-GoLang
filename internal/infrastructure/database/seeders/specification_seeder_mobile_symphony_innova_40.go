package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyInnova40 seeds specifications/options for product 'symphony-innova-40'
type SpecificationSeederMobileSymphonyInnova40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyInnova40 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyInnova40() *SpecificationSeederMobileSymphonyInnova40 {
	return &SpecificationSeederMobileSymphonyInnova40{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Innova 40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyInnova40) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p@30fps": "১০৮০p@৩০fps",
		"120 Hz": "১২০ Hz",
		"128 GB (uMCP)": "১২৮ GB (uMCP)",
		"169 x 78 x 9 mm": "১৬৯ x ৭৮ x ৯ মিমি",
		"18W": "১৮W",
		"210 g": "২১০ g",
		"50 MP + 0.08 MP + 0.08 MP": "৫০ MP + ০.০৮ MP + ০.০৮ MP",
		"6 GB / 12 GB / 16 GB (expandable)": "৬ GB / ১২ GB / ১৬ GB (expandable)",
		"6.75 inches": "৬.৭৫ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"720 × 1600 px": "৭২০ × ১৬০০ px",
		"8 MP": "৮ MP",
		"Android 15": "Android ১৫",
		"August 2025": "August ২০২৫",
		"IP52 rating": "IP৫২ rating",
		"Mali-G57": "Mali-G৫৭",
		"Octa-core (A75 + A55)": "Octa-core (A৭৫ + A৫৫)",
		"Side fingerprint, face unlock, accelerometer, proximity, light, compass": "Side ফিঙ্গারপ্রিন্ট, face unlock, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, light, কম্পাস",
		"Space Black, Titanium Silver, Cosmic Gold, Ice Blue, Radium Green": "Space কালো, Titanium রূপালী, Cosmic সোনালী, Ice নীল, Radium সবুজ",
		"Unisoc T615": "Unisoc T৬১৫",
		"Unisoc T615 (12 nm)": "Unisoc T৬১৫ (১২ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-innova-40'
func (s *SpecificationSeederMobileSymphonyInnova40) Seed(db *gorm.DB) error {
	productSlug := "symphony-innova-40"

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
