package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplusNordN20Se seeds specifications/options for product 'oneplus-nord-n20-se'
type SpecificationSeederMobileOneplusNordN20Se struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNordN20Se creates a new seeder instance
func NewSpecificationSeederMobileOneplusNordN20Se() *SpecificationSeederMobileOneplusNordN20Se {
	return &SpecificationSeederMobileOneplusNordN20Se{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord N20 SE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNordN20Se) getBanglaTranslations() map[string]string {
	return map[string]string{
		"163.8 x 75 x 8 mm": "১৬৩.৮ x ৭৫ x ৮ মিমি",
		"187 g": "১৮৭ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6.56 inches": "৬.৫৬ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1612 pixels": "৭২০ x ১৬১২ pixels",
		"8 MP": "৮ MP",
		"Android 12, OxygenOS 12.1": "Android ১২, OxygenOS ১২.১",
		"August 2022": "August ২০২২",
		"Blue Oasis, Celestial Black": "নীল Oasis, Celestial কালো",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"Helio G35": "Helio G৩৫",
		"IPS LCD, 600 nits": "IPS LCD, ৬০০ nits",
		"Mediatek MT6765G Helio G35 (12 nm)": "Mediatek MT৬৭৬৫G Helio G৩৫ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-nord-n20-se'
func (s *SpecificationSeederMobileOneplusNordN20Se) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-n20-se"

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

	// Override model-specific values for OnePlus Nord N20 SE
	specs["Display Size"] = "6.56 inches"
	specs["Processor"] = "Helio G35"
	specs["Chipset"] = "Mediatek MT6765G Helio G35 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB / 128 GB"
	specs["Display Type"] = "IPS LCD, 600 nits"
	specs["Resolution"] = "720 x 1612 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "187 g"
	specs["Dimensions"] = "163.8 x 75 x 8 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 12, OxygenOS 12.1"
	specs["Available Colors"] = "Blue Oasis, Celestial Black"
	specs["Announcement Date"] = "August 2022"
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
