package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileTecnoPhantomVFold25g seeds specifications/options for product 'tecno-phantom-v-fold2-5g'
type SpecificationSeederMobileTecnoPhantomVFold25g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPhantomVFold25g creates a new seeder instance
func NewSpecificationSeederMobileTecnoPhantomVFold25g() *SpecificationSeederMobileTecnoPhantomVFold25g {
	return &SpecificationSeederMobileTecnoPhantomVFold25g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno PHANTOM V Fold2 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPhantomVFold25g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"16 MP (Cover) + 32 MP (Inner)": "১৬ MP (Cover) + ৩২ MP (Inner)",
		"2000 x 2296 pixels": "২০০০ x ২২৯৬ pixels",
		"299 g": "২৯৯ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 50 MP + 13 MP": "৫০ MP + ৫০ MP + ১৩ MP",
		"512 GB": "৫১২ GB",
		"5G": "৫G",
		"7.85 inches (Folded: 6.42 inches)": "৭.৮৫ ইঞ্চি (Folded: ৬.৪২ ইঞ্চি)",
		"Android 14, HIOS 14 Fold": "Android ১৪, HIOS ১৪ Fold",
		"Black, White": "কালো, সাদা",
		"Dimensity 9000+": "Dimensity ৯০০০+",
		"Foldable LTPO AMOLED, 120Hz": "Foldable LTPO AMOLED, ১২০Hz",
		"Glass front, glass back, aluminum frame": "গ্লাস সামনে, গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Mali-G710 MC10": "Mali-G৭১০ MC১০",
		"Mediatek Dimensity 9000+ (4 nm)": "Mediatek Dimensity ৯০০০+ (৪ nm)",
		"September 2024": "September ২০২৪",
		"Unfolded: 159.4 x 140.4 x 6.9 mm": "Unfolded: ১৫৯.৪ x ১৪০.৪ x ৬.৯ মিমি",
	}
}

// Seed inserts specification_translations for existing specifications for product 'tecno-phantom-v-fold2-5g'
func (s *SpecificationSeederMobileTecnoPhantomVFold25g) Seed(db *gorm.DB) error {
	productSlug := "tecno-phantom-v-fold2-5g"

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
