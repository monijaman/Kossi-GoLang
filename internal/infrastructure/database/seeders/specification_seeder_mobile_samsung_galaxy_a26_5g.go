package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA265g seeds specifications/options for product 'samsung-galaxy-a26-5g'
type SpecificationSeederMobileSamsungGalaxyA265g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA265g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA265g() *SpecificationSeederMobileSamsungGalaxyA265g {
	return &SpecificationSeederMobileSamsungGalaxyA265g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A26 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA265g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels (~405 ppi density)": "১০৮০ x ২৪০০ pixels (~৪০৫ ppi density)",
		"120Hz": "১২০Hz",
		"128 GB": "১২৮ GB",
		"13 MP": "১৩ MP",
		"161.3 x 73.9 x 8.4 mm": "১৬১.৩ x ৭৩.৯ x ৮.৪ মিমি",
		"189 g (6.67 oz)": "১৮৯ g (৬.৬৭ oz)",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"6 GB / 8 GB": "৬ GB / ৮ GB",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"August 2025": "August ২০২৫",
		"Black, Blue": "কালো, নীল",
		"Corning Gorilla Glass 3": "Corning Gorilla Glass ৩",
		"Dimensity 6020": "Dimensity ৬০২০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"MediaTek MT6833 Dimensity 6020 (7 nm)": "MediaTek MT৬৮৩৩ Dimensity ৬০২০ (৭ nm)",
		"Octa-core (2x2.2 GHz Cortex-A76 & 6x2.0 GHz Cortex-A55)": "Octa-core (২x২.২ GHz Cortex-A৭৬ & ৬x২.০ GHz Cortex-A৫৫)",
		"PLS LCD, 120Hz": "PLS LCD, ১২০Hz",
		"Plastic frame, plastic back": "Plastic frame, প্লাস্টিক পেছনে",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a26-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA265g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a26-5g"

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

	// Override model-specific values for Samsung Galaxy A26 5G
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Dimensity 6020"
	specs["Chipset"] = "MediaTek MT6833 Dimensity 6020 (7 nm)"
	specs["Cpu Type"] = "Octa-core (2x2.2 GHz Cortex-A76 & 6x2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "6 GB / 8 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "PLS LCD, 120Hz"
	specs["Resolution"] = "1080 x 2400 pixels (~405 ppi density)"
	specs["Screen Protection"] = "Corning Gorilla Glass 3"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic frame, plastic back"
	specs["Weight"] = "189 g (6.67 oz)"
	specs["Dimensions"] = "161.3 x 73.9 x 8.4 mm"
	specs["Water Resistance"] = "None"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "13 MP"
	specs["Battery"] = "4500 mAh"
	specs["Operating System"] = "Android 14, One UI 6"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "August 2025"
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
