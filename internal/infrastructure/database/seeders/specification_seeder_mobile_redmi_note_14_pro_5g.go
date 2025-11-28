package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileRedmiNote14Pro5g seeds specifications/options for product 'redmi-note-14-pro-5g'
type SpecificationSeederMobileRedmiNote14Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote14Pro5g creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote14Pro5g() *SpecificationSeederMobileRedmiNote14Pro5g {
	return &SpecificationSeederMobileRedmiNote14Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 14 Pro+ 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote14Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"162.5 x 74.7 x 8.7 mm": "১৬২.৫ x ৭৪.৭ x ৮.৭ মিমি",
		"20 MP": "২০ MP",
		"210 g": "২১০ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"50 MP + 50 MP + 8 MP": "৫০ MP + ৫০ MP + ৮ MP",
		"5G": "৫G",
		"6,200 mAh": "৬,২০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 68B colors, 120Hz, Dolby Vision, 3000 nits": "AMOLED, ৬৮B colors, ১২০Hz, Dolby Vision, ৩০০০ nits",
		"Adreno 710": "Adreno ৭১০",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"IP68": "IP৬৮",
		"Qualcomm Snapdragon 7s Gen 3 (4 nm)": "Qualcomm Snapdragon ৭s Gen ৩ (৪ nm)",
		"September 2024": "September ২০২৪",
		"Snapdragon 7s Gen 3": "Snapdragon ৭s Gen ৩",
		"Starry White, মিডনাইট Black, Mirror Porcelain White": "Starry সাদা, মিডনাইট কালো, Mirror Porcelain সাদা",
	}
}

// Seed inserts specification_translations for existing specifications for product 'redmi-note-14-pro-5g'
func (s *SpecificationSeederMobileRedmiNote14Pro5g) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-14-pro-5g"

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
