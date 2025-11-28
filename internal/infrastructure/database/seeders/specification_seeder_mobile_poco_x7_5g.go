package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoX75g seeds specifications/options for product 'poco-x7-5g'
type SpecificationSeederMobilePocoX75g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoX75g creates a new seeder instance
func NewSpecificationSeederMobilePocoX75g() *SpecificationSeederMobilePocoX75g {
	return &SpecificationSeederMobilePocoX75g{BaseSeeder: BaseSeeder{name: "Specifications for POCO X7 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoX75g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"16 MP": "১৬ MP",
		"189 g": "১৮৯ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"67W wired": "৬৭W তারযুক্ত",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz": "AMOLED, ১২০Hz",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Black, Blue, Yellow": "কালো, নীল, হলুদ",
		"Dimensity 7050": "Dimensity ৭০৫০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic back, plastic frame": "গ্লাস সামনে (Gorilla Glass ৫), প্লাস্টিক পেছনে, plastic frame",
		"IP54": "IP৫৪",
		"January 2025": "January ২০২৫",
		"Mali-G68 MC4": "Mali-G৬৮ MC৪",
		"Mediatek Dimensity 7050 (6 nm)": "Mediatek Dimensity ৭০৫০ (৬ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-x7-5g'
func (s *SpecificationSeederMobilePocoX75g) Seed(db *gorm.DB) error {
	productSlug := "poco-x7-5g"

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
