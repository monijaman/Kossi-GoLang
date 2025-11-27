package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoY21d seeds specifications/options for product 'vivo-y21d'
type SpecificationSeederMobileVivoY21d struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY21d creates a new seeder instance
func NewSpecificationSeederMobileVivoY21d() *SpecificationSeederMobileVivoY21d {
	return &SpecificationSeederMobileVivoY21d{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y21d"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY21d) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"182 g": "১৮২ g",
		"18W wired": "১৮W তারযুক্ত",
		"4 GB": "৪ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.51 inches": "৬.৫১ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB": "৬৪ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 11, Funtouch 11.1": "Android ১১, Funtouch ১১.১",
		"August 2021": "August ২০২১",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Helio P35": "Helio P৩৫",
		"Mediatek Helio P35 (12nm)": "Mediatek Helio P৩৫ (১২nm)",
		"Midnight Blue, Diamond Glow": "Midnight নীল, Diamond Glow",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-y21d'
func (s *SpecificationSeederMobileVivoY21d) Seed(db *gorm.DB) error {
	productSlug := "vivo-y21d"

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

	// Override model-specific values for Vivo Y21d
	specs["Display Size"] = "6.51 inches"
	specs["Processor"] = "Helio P35"
	specs["Chipset"] = "Mediatek Helio P35 (12nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "182 g"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "18W wired"
	specs["Operating System"] = "Android 11, Funtouch 11.1"
	specs["Available Colors"] = "Midnight Blue, Diamond Glow"
	specs["Announcement Date"] = "August 2021"
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
