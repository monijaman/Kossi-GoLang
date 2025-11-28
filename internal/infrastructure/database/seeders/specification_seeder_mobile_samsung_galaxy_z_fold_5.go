package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyZFold5 seeds specifications/options for product 'samsung-galaxy-z-fold-5'
type SpecificationSeederMobileSamsungGalaxyZFold5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFold5 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFold5() *SpecificationSeederMobileSamsungGalaxyZFold5 {
	return &SpecificationSeederMobileSamsungGalaxyZFold5{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Fold 5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFold5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"1812 x 2176 pixels": "১৮১২ x ২১৭৬ pixels",
		"253 g": "২৫৩ g",
		"256 GB / 512 GB / 1 টিবি": "২৫৬ GB / ৫১২ GB / ১ টিবি",
		"4 MP (Under Display) + 10 MP (Cover)": "৪ MP (ডিসপ্লের নিচে) + ১০ MP (Cover)",
		"4,400 mAh": "৪,৪০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ MP + ১০ MP + ১২ MP",
		"5G": "৫G",
		"7.6 inches (Main) / 6.2 inches (Cover)": "৭.৬ ইঞ্চি (Main) / ৬.২ ইঞ্চি (Cover)",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, আপগ্রেডযোগ্য": "Android ১৩, আপগ্রেডযোগ্য",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Foldable Dynamic AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"Glass front (Gorilla Glass Victus 2), glass back (Victus 2), aluminum frame": "গ্লাস সামনে (Gorilla Glass Victus ২), গ্লাস পেছনে (Victus ২), অ্যালুমিনিয়াম ফ্রেম",
		"IPX8": "IPX৮",
		"Icy Blue, Phantom Black, Cream, Gray, Blue": "Icy নীল, Phantom কালো, Cream, ধূসর, নীল",
		"July 2023": "July ২০২৩",
		"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AC Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
		"Unfolded: 154.9 x 129.9 x 6.1 mm / Folded: 154.9 x 67.1 x 13.4 mm": "Unfolded: ১৫৪.৯ x ১২৯.৯ x ৬.১ মিমি / Folded: ১৫৪.৯ x ৬৭.১ x ১৩.৪ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-z-fold-5'
func (s *SpecificationSeederMobileSamsungGalaxyZFold5) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-fold-5"

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
