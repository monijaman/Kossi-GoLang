package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSymphonyMax60 seeds specifications/options for product 'symphony-max-60'
type SpecificationSeederMobileSymphonyMax60 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyMax60 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyMax60() *SpecificationSeederMobileSymphonyMax60 {
	return &SpecificationSeederMobileSymphonyMax60{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Max 60"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyMax60) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p@30fps": "১০৮০p@৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"168.75 x 78.1 x 8.79 mm": "১৬৮.৭৫ x ৭৮.১ x ৮.৭৯ মিমি",
		"210 g": "২১০ g",
		"4 GB": "৪ GB",
		"5 MP": "৫ MP",
		"6.75 inches": "৬.৭৫ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"64 GB (expandable up to 256GB)": "৬৪ GB (expandable পর্যন্ত ২৫৬GB)",
		"720 × 1600 px": "৭২০ × ১৬০০ px",
		"90 Hz": "৯০ Hz",
		"Android 15": "Android ১৫",
		"IMG GE8322": "IMG GE৮৩২২",
		"IP64 dust & splash resistant": "IP৬৪ dust & splash resistant",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Octa-core 1.6 GHz": "Octa-core ১.৬ GHz",
		"Plastic body, glass front": "Plastic body, গ্লাস সামনে",
		"September 2025": "September ২০২৫",
		"Side fingerprint, accelerometer, proximity": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Space Black, Frost Blue, Classic Navy Blue, Titanium Gray, Divine Gold": "Space কালো, Frost নীল, Classic Navy নীল, Titanium ধূসর, Divine সোনালী",
		"Unisoc SC9863A1": "Unisoc SC৯৮৬৩A১",
		"Unisoc SC9863A1 (28 nm)": "Unisoc SC৯৮৬৩A১ (২৮ nm)",
		"Wi-Fi, Bluetooth 5.0, OTG": "Wi-Fi, Bluetooth ৫.০, OTG",
	}
}

// Seed inserts specification_translations for existing specifications for product 'symphony-max-60'
func (s *SpecificationSeederMobileSymphonyMax60) Seed(db *gorm.DB) error {
	productSlug := "symphony-max-60"

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
