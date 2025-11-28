package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplus13r seeds specifications/options for product 'oneplus-13r'
type SpecificationSeederMobileOneplus13r struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus13r creates a new seeder instance
func NewSpecificationSeederMobileOneplus13r() *SpecificationSeederMobileOneplus13r {
	return &SpecificationSeederMobileOneplus13r{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 13R"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus13r) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1264 x 2780 pixels": "১২৬৪ x ২৭৮০ pixels",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"163.3 x 75.3 x 8.8 mm": "১৬৩.৩ x ৭৫.৩ x ৮.৮ মিমি",
		"207 g": "২০৭ g",
		"5,500 mAh": "৫,৫০০ এমএএইচ",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5G": "৫G",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 16 GB": "৮ GB / ১৬ GB",
		"Adreno 750": "Adreno ৭৫০",
		"Android 15, OxygenOS 15": "Android ১৫, OxygenOS ১৫",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP64": "IP৬৪",
		"Iron Gray, Cool Blue": "Iron ধূসর, Cool নীল",
		"January 2025": "January ২০২৫",
		"LTPO AMOLED, 120Hz, HDR10+, 4500 nits": "LTPO AMOLED, ১২০Hz, HDR১০+, ৪৫০০ nits",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AB Snapdragon ৮ Gen ৩ (৪ nm)",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-13r'
func (s *SpecificationSeederMobileOneplus13r) Seed(db *gorm.DB) error {
	productSlug := "oneplus-13r"

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
