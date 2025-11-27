package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonX91 seeds specifications/options for product 'walton-x91'
type SpecificationSeederMobileWaltonX91 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonX91 creates a new seeder instance
func NewSpecificationSeederMobileWaltonX91() *SpecificationSeederMobileWaltonX91 {
	return &SpecificationSeederMobileWaltonX91{BaseSeeder: BaseSeeder{name: "Specifications for Walton X91"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonX91) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"128 GB": "১২৮ GB",
		"164.4 x 76.8 x 7.8 mm": "১৬৪.৪ x ৭৬.৮ x ৭.৮ মিমি",
		"185 g": "১৮৫ g",
		"32 MP": "৩২ MP",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 5 MP + 2 MP": "৬৪ MP + ৫ MP + ২ MP",
		"8 GB": "৮ GB",
		"90Hz": "৯০Hz",
		"AMOLED, 90Hz": "AMOLED, ৯০Hz",
		"Android 13": "Android ১৩",
		"Black, Blue": "কালো, নীল",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G99": "Helio G৯৯",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"March 2024": "March ২০২৪",
		"Mediatek Helio G99 (6 nm)": "Mediatek Helio G৯৯ (৬ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-x91'
func (s *SpecificationSeederMobileWaltonX91) Seed(db *gorm.DB) error {
	productSlug := "walton-x91"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMobileSpecs()
	banglaTranslations := s.getBanglaTranslations()

	// Override model-specific values for Walton X91
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Helio G99"
	specs["Chipset"] = "Mediatek Helio G99 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "AMOLED, 90Hz"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "185 g"
	specs["Dimensions"] = "164.4 x 76.8 x 7.8 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "64 MP + 5 MP + 2 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "March 2024"
	specs["Device Status"] = "Available"

	for key, value := range specs {
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: sk.ID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}

				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					var existingTranslation models.SpecificationTranslationModel
					if err := db.Where("specification_id = ? AND locale = ?", sModel.ID, "bn").First(&existingTranslation).Error; err != nil {
						if err == gorm.ErrRecordNotFound {
							translation := &models.SpecificationTranslationModel{
								SpecificationID: sModel.ID,
								Locale:          "bn",
								Value:           banglaValue,
							}
							if err := db.Create(translation).Error; err != nil {
								return err
							}
						} else {
							return err
						}
					}
				}
			} else {
				return err
			}
		} else {
			// If specification already exists, check and create Bangla translation if missing
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {
				var existingTranslation models.SpecificationTranslationModel
				if err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						translation := &models.SpecificationTranslationModel{
							SpecificationID: existing.ID,
							Locale:          "bn",
							Value:           banglaValue,
						}
						if err := db.Create(translation).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			}
		}
	}

	return nil
}
