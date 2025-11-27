package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiA5 seeds specifications/options for product 'redmi-a5'
type SpecificationSeederMobileRedmiA5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiA5 creates a new seeder instance
func NewSpecificationSeederMobileRedmiA5() *SpecificationSeederMobileRedmiA5 {
	return &SpecificationSeederMobileRedmiA5{BaseSeeder: BaseSeeder{name: "Specifications for Redmi A5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiA5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10W wired": "১০W তারযুক্ত",
		"199 g": "১৯৯ g",
		"3 GB / 4 GB": "৩ GB / ৪ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ GB / ১২৮ GB",
		"720 x 1650 pixels": "৭২০ x ১৬৫০ pixels",
		"8 MP + 0.08 MP": "৮ MP + ০.০৮ MP",
		"90Hz": "৯০Hz",
		"Android 14 (Go edition)": "Android ১৪ (Go edition)",
		"February 2025": "February ২০২৫",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Helio G36": "Helio G৩৬",
		"IPS LCD, 90Hz": "IPS LCD, ৯০Hz",
		"Mediatek Helio G36 (12 nm)": "Mediatek Helio G৩৬ (১২ nm)",
		"Midnight Black, Lake Blue, Olive Green": "Midnight কালো, Lake নীল, Olive সবুজ",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-a5'
func (s *SpecificationSeederMobileRedmiA5) Seed(db *gorm.DB) error {
	productSlug := "redmi-a5"

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

	// Override model-specific values for Redmi A5
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Helio G36"
	specs["Chipset"] = "Mediatek Helio G36 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "3 GB / 4 GB"
	specs["Storage"] = "64 GB / 128 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "720 x 1650 pixels"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "199 g"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "8 MP + 0.08 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "10W wired"
	specs["Operating System"] = "Android 14 (Go edition)"
	specs["Available Colors"] = "Midnight Black, Lake Blue, Olive Green"
	specs["Announcement Date"] = "February 2025"
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
