package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplusNordN30Se5g seeds specifications/options for product 'oneplus-nord-n30-se-5g'
type SpecificationSeederMobileOneplusNordN30Se5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNordN30Se5g creates a new seeder instance
func NewSpecificationSeederMobileOneplusNordN30Se5g() *SpecificationSeederMobileOneplusNordN30Se5g {
	return &SpecificationSeederMobileOneplusNordN30Se5g{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord N30 SE 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNordN30Se5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"128 GB": "১২৮ GB",
		"165.6 x 76 x 8 mm": "১৬৫.৬ x ৭৬ x ৮ মিমি",
		"193 g": "১৯৩ g",
		"4 GB": "৪ GB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6.72 inches": "৬.৭২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 MP": "৮ MP",
		"Android 13, OxygenOS 13.1": "Android ১৩, OxygenOS ১৩.১",
		"Black Satin, Cyan Sparkle": "কালো Satin, Cyan Sparkle",
		"Dimensity 6020": "Dimensity ৬০২০",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"January 2024": "January ২০২৪",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"Mediatek Dimensity 6020 (7 nm)": "Mediatek Dimensity ৬০২০ (৭ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-nord-n30-se-5g'
func (s *SpecificationSeederMobileOneplusNordN30Se5g) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-n30-se-5g"

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

	// Override model-specific values for OnePlus Nord N30 SE 5G
	specs["Display Size"] = "6.72 inches"
	specs["Processor"] = "Dimensity 6020"
	specs["Chipset"] = "Mediatek Dimensity 6020 (7 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "193 g"
	specs["Dimensions"] = "165.6 x 76 x 8 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13, OxygenOS 13.1"
	specs["Available Colors"] = "Black Satin, Cyan Sparkle"
	specs["Announcement Date"] = "January 2024"
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
