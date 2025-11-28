package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileIphone17Plus seeds specifications/options for product 'iphone-17-plus'
type SpecificationSeederMobileIphone17Plus struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone17Plus creates a new seeder instance
func NewSpecificationSeederMobileIphone17Plus() *SpecificationSeederMobileIphone17Plus {
	return &SpecificationSeederMobileIphone17Plus{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 17 Plus"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone17Plus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP, f/1.9": "১২ MP, f/১.৯",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"160.9 x 77.8 x 7.8 mm": "১৬০.৯ x ৭৭.৮ x ৭.৮ মিমি",
		"203 g (7.16 oz)": "২০৩ g (৭.১৬ oz)",
		"20W wired, 15W MagSafe wireless": "২০W wired, ১৫W MagSafe wireless",
		"2796 x 1290 pixels (~460 ppi density)": "২৭৯৬ x ১২৯০ pixels (~৪৬০ ppi density)",
		"2x Optical Zoom (via crop)": "২x অপটিক্যাল জুম (ক্রপের মাধ্যমে)",
		"4,383 mAh": "৪,৩৮৩ এমএএইচ",
		"48 MP + 12 MP": "৪৮ MP + ১২ MP",
		"4K@24/25/30/60fps, 1080p@25/30/60/120/240fps, HDR, Dolby Vision HDR": "৪K@২৪/২৫/৩০/৬০fps, ১০৮০p@২৫/৩০/৬০/১২০/২৪০fps, HDR, Dolby Vision HDR",
		"4K@24/25/30/60fps, 1080p@25/30/60/120fps": "৪K@২৪/২৫/৩০/৬০fps, ১০৮০p@২৫/৩০/৬০/১২০fps",
		"5-core Apple GPU": "৫-core Apple GPU",
		"5.3, A2DP, LE": "৫.৩, A২DP, LE",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 GB": "৮ GB",
		"Apple A19": "Apple A১৯",
		"Apple A19 (3 nm)": "Apple A১৯ (৩ nm)",
		"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপী",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "GSM / CDMA / HSPA / EVDO / LTE / ৫G",
		"Hexa-core (2x Performance + 4x Efficiency)": "Hexa-core (২x Performance + ৪x Efficiency)",
		"IP68 dust/water resistant (up to 6m for 30 min)": "IP৬৮ dust/water resistant (up to ৬m for ৩০ min)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"No 3.5mm jack": "না ৩.৫mm jack",
		"September 2024": "September ২০২৪",
		"Super Retina XDR OLED, HDR10, Dolby Vision, 2000 nits (HBM)": "Super Retina XDR OLED, HDR১০, Dolby Vision, ২০০০ nits (HBM)",
		"USB Type-C 2.0": "USB Type-C ২.০",
		"Wi-Fi 802.11 a/b/g/n/ac/6": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬",
		"iOS 18": "iOS ১৮",
	}
}

// Seed inserts specification_translations for existing specifications for product 'iphone-17-plus'
func (s *SpecificationSeederMobileIphone17Plus) Seed(db *gorm.DB) error {
	productSlug := "iphone-17-plus"

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
