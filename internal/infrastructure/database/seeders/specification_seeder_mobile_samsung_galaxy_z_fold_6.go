package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyZFold6 seeds specifications/options for product 'samsung-galaxy-z-fold-6'
type SpecificationSeederMobileSamsungGalaxyZFold6 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFold6 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFold6() *SpecificationSeederMobileSamsungGalaxyZFold6 {
	return &SpecificationSeederMobileSamsungGalaxyZFold6{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Fold 6"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFold6) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"1856 x 2160 pixels": "১৮৫৬ x ২১৬০ pixels",
		"239 g": "২৩৯ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"4 MP (Under Display) + 10 MP (Cover)": "৪ MP (ডিসপ্লের নিচে) + ১০ MP (Cover)",
		"4,400 mAh": "৪,৪০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ MP + ১০ MP + ১২ MP",
		"5G": "৫G",
		"7.6 inches (Main) / 6.3 inches (Cover)": "৭.৬ ইঞ্চি (Main) / ৬.৩ ইঞ্চি (Cover)",
		"Adreno 750": "Adreno ৭৫০",
		"Android 14, One UI 6.1.1": "Android ১৪, One UI ৬.১.১",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic LTPO AMOLED ২X, ১২০Hz, HDR১০+",
		"Glass front (Gorilla Glass Victus 2), glass back (Victus 2), aluminum frame": "গ্লাস সামনে (Gorilla Glass Victus ২), গ্লাস পেছনে (Victus ২), অ্যালুমিনিয়াম ফ্রেম",
		"IP48": "IP৪৮",
		"July 2024": "July ২০২৪",
		"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM৮৬৫০-AC Snapdragon ৮ Gen ৩ (৪ nm)",
		"Silver Shadow, Pink, Navy, Crafted Black, White": "রূপালী Shadow, গোলাপী, Navy, Crafted কালো, সাদা",
		"Snapdragon 8 Gen 3": "Snapdragon ৮ Gen ৩",
		"Unfolded: 153.5 x 132.6 x 5.6 mm / Folded: 153.5 x 68.1 x 12.1 mm": "Unfolded: ১৫৩.৫ x ১৩২.৬ x ৫.৬ মিমি / Folded: ১৫৩.৫ x ৬৮.১ x ১২.১ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-z-fold-6'
func (s *SpecificationSeederMobileSamsungGalaxyZFold6) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-fold-6"

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
