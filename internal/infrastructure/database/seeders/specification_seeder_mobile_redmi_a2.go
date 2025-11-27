package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiA2 seeds specifications/options for product 'redmi-a2'
type SpecificationSeederMobileRedmiA2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiA2 creates a new seeder instance
func NewSpecificationSeederMobileRedmiA2() *SpecificationSeederMobileRedmiA2 {
	return &SpecificationSeederMobileRedmiA2{BaseSeeder: BaseSeeder{name: "Specifications for Redmi A2+"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiA2) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10W wired": "১০W তারযুক্ত",
		"192 g": "১৯২ g",
		"2 GB / 3 GB / 4 GB": "২ GB / ৩ GB / ৪ GB",
		"32 GB / 64 GB": "৩২ GB / ৬৪ GB",
		"5 MP": "৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP + 0.08 MP": "৮ MP + ০.০৮ MP",
		"Android 12 or 13 (Go edition), MIUI": "Android ১২ or ১৩ (Go edition), MIUI",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"Helio G36": "Helio G৩৬",
		"Light Blue, Light Green, Black": "Light নীল, Light সবুজ, কালো",
		"March 2023": "March ২০২৩",
		"Mediatek Helio G36 (12 nm)": "Mediatek Helio G৩৬ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-a2'
func (s *SpecificationSeederMobileRedmiA2) Seed(db *gorm.DB) error {
	productSlug := "redmi-a2"

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

	// Override model-specific values for Redmi A2+
	specs["Display Size"] = "6.52 inches"
	specs["Processor"] = "Helio G36"
	specs["Chipset"] = "Mediatek Helio G36 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "2 GB / 3 GB / 4 GB"
	specs["Storage"] = "32 GB / 64 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "192 g"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "8 MP + 0.08 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "10W wired"
	specs["Operating System"] = "Android 12 or 13 (Go edition), MIUI"
	specs["Available Colors"] = "Light Blue, Light Green, Black"
	specs["Announcement Date"] = "March 2023"
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
