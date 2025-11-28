package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileIphone17Pro seeds specifications/options for product 'iphone-17-pro'
type SpecificationSeederMobileIphone17Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone17Pro creates a new seeder instance
func NewSpecificationSeederMobileIphone17Pro() *SpecificationSeederMobileIphone17Pro {
	return &SpecificationSeederMobileIphone17Pro{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 17 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone17Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"12 MP, f/2.2": "১২ MP, f/২.২",
		"120Hz ProMotion": "১২০Hz ProMotion",
		"153.7 x 73.6 x 8.2 mm (6.05 x 2.90 x 0.32 in)": "১৫৩.৭ x ৭৩.৬ x ৮.২ মিমি (৬.০৫ x ২.৯০ x ০.৩২ in)",
		"199 g (7.02 oz)": "১৯৯ g (৭.০২ oz)",
		"2556 x 1179 pixels (~460 ppi density)": "২৫৫৬ x ১১৭৯ pixels (~৪৬০ ppi density)",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"2x + 5x Optical Zoom": "২x + ৫x অপটিক্যাল জুম",
		"3,800 mAh (typical)": "৩,৮০০ এমএএইচ (সাধারণ)",
		"45W wired; 25W MagSafe wireless; 5W reverse wireless": "৪৫W wired; ২৫W MagSafe wireless; ৫W reverse wireless",
		"48 MP + 12 MP + 12 MP": "৪৮ MP + ১২ MP + ১২ MP",
		"4K@24/25/30/60fps, 1080p@30/60/120/240fps, ProRes video recording": "৪K@২৪/২৫/৩০/৬০fps, ১০৮০p@৩০/৬০/১২০/২৪০fps, ProRes video recording",
		"4K@24/25/30/60fps, 1080p@30/60fps": "৪K@২৪/২৫/৩০/৬০fps, ১০৮০p@৩০/৬০fps",
		"5.4, A2DP, LE, EDR": "৫.৪, A২DP, LE, EDR",
		"6-core Apple GPU": "৬-core Apple GPU",
		"6.3 inches": "৬.৩ ইঞ্চি",
		"Apple A19 Pro": "Apple A১৯ Pro",
		"Black, White, Gold, Rose Gold, Blue": "কালো, সাদা, সোনালী, Rose সোনালী, নীল",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "GSM / CDMA / HSPA / EVDO / LTE / ৫G",
		"Hexa-core (2x Performance + 4x Efficiency)": "Hexa-core (২x Performance + ৪x Efficiency)",
		"IP68 dust tight and water resistant (immersible up to 6m for 30 min)": "IP৬৮ dust tight and water resistant (immersible up to ৬m for ৩০ min)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Nano-SIM + eSIM": "ন্যানো-সিম + ই-সিম",
		"September 2024": "September ২০২৪",
		"Super Retina XDR LTPO OLED, 120Hz, HDR10, Dolby Vision": "Super Retina XDR LTPO OLED, ১২০Hz, HDR১০, Dolby Vision",
		"USB Type-C with USB 3.1 speeds": "USB Type-C with USB ৩.১ speeds",
		"Wi-Fi 802.11 a/b/g/n/ac/6e (802.11ax)": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e (৮০২.১১ax)",
		"Yes - 25W MagSafe wireless; 5W reverse wireless": "Yes - ২৫W MagSafe wireless; ৫W reverse wireless",
		"iOS 18": "iOS ১৮",
	}
}

// Seed inserts specification_translations for existing specifications for product 'iphone-17-pro'
func (s *SpecificationSeederMobileIphone17Pro) Seed(db *gorm.DB) error {
	productSlug := "iphone-17-pro"

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
