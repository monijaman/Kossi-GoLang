package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA06 seeds specifications/options for product 'samsung-galaxy-a06'
type SpecificationSeederMobileSamsungGalaxyA06 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA06 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA06() *SpecificationSeederMobileSamsungGalaxyA06 {
	return &SpecificationSeederMobileSamsungGalaxyA06{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A06"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA06) getBanglaTranslations() map[string]string {
	return map[string]string{
		"164.0 x 75.8 x 9.2 mm": "১৬৪.০ x ৭৫.৮ x ৯.২ মিমি",
		"190 g (6.70 oz)": "১৯০ g (৬.৭০ oz)",
		"3 GB / 4 GB": "৩ GB / ৪ GB",
		"32 GB / 64 GB": "৩২ GB / ৬৪ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels (~270 ppi density)": "৭২০ x ১৬০০ pixels (~২৭০ ppi density)",
		"8 MP": "৮ MP",
		"Android 14, One UI Core 6": "Android ১৪, One UI Core ৬",
		"Black, Blue": "কালো, নীল",
		"Mali-G57 MP1": "Mali-G৫৭ MP১",
		"May 2025": "May ২০২৫",
		"Octa-core (2x1.6 GHz Cortex-A75 & 6x1.6 GHz Cortex-A55)": "Octa-core (২x১.৬ GHz Cortex-A৭৫ & ৬x১.৬ GHz Cortex-A৫৫)",
		"Plastic frame, plastic back": "Plastic frame, প্লাস্টিক পেছনে",
		"Unisoc T606": "Unisoc T৬০৬",
		"Unisoc T606 (12 nm)": "Unisoc T৬০৬ (১২ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a06'
func (s *SpecificationSeederMobileSamsungGalaxyA06) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a06"

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

	// Override model-specific values for Samsung Galaxy A06
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Unisoc T606"
	specs["Chipset"] = "Unisoc T606 (12 nm)"
	specs["Cpu Type"] = "Octa-core (2x1.6 GHz Cortex-A75 & 6x1.6 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G57 MP1"
	specs["Ram"] = "3 GB / 4 GB"
	specs["Storage"] = "32 GB / 64 GB"
	specs["Display Type"] = "PLS LCD"
	specs["Resolution"] = "720 x 1600 pixels (~270 ppi density)"
	specs["Screen Protection"] = "No"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic frame, plastic back"
	specs["Weight"] = "190 g (6.70 oz)"
	specs["Dimensions"] = "164.0 x 75.8 x 9.2 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "8 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "5000 mAh"
	specs["Operating System"] = "Android 14, One UI Core 6"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "May 2025"
	specs["Device Status"] = "Upcoming"

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
