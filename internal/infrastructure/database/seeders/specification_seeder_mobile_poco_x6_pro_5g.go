package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobilePocoX6Pro5g seeds specifications/options for product 'poco-x6-pro-5g'
type SpecificationSeederMobilePocoX6Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoX6Pro5g creates a new seeder instance
func NewSpecificationSeederMobilePocoX6Pro5g() *SpecificationSeederMobilePocoX6Pro5g {
	return &SpecificationSeederMobilePocoX6Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for POCO X6 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoX6Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"16 MP": "১৬ MP",
		"186 g": "১৮৬ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"67W wired": "৬৭W তারযুক্ত",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz, Dolby Vision, HDR10+": "AMOLED, ১২০Hz, Dolby Vision, HDR১০+",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Black, Yellow, Gray": "কালো, হলুদ, ধূসর",
		"Dimensity 8300 আল্ট্রা": "Dimensity ৮৩০০ আল্ট্রা",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic back বা সিলিকন পলিমার back, plastic frame": "গ্লাস সামনে (Gorilla Glass ৫), প্লাস্টিক পেছনে বা সিলিকন পলিমার back, plastic frame",
		"IP54": "IP৫৪",
		"January 2024": "January ২০২৪",
		"Mali-G615-MC6": "Mali-G৬১৫-MC৬",
		"Mediatek Dimensity 8300 আল্ট্রা (4 nm)": "Mediatek Dimensity ৮৩০০ আল্ট্রা (৪ nm)",
	}
}

// Seed inserts specification_translations for existing specifications for product 'poco-x6-pro-5g'
func (s *SpecificationSeederMobilePocoX6Pro5g) Seed(db *gorm.DB) error {
	productSlug := "poco-x6-pro-5g"

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
