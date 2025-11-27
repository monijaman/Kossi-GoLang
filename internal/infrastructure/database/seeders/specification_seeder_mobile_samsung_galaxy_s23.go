package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS23 seeds specifications/options for product 'samsung-galaxy-s23'
type SpecificationSeederMobileSamsungGalaxyS23 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS23 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS23() *SpecificationSeederMobileSamsungGalaxyS23 {
	return &SpecificationSeederMobileSamsungGalaxyS23{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S23"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS23) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"12 MP": "১২ MP",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"146.3 x 70.9 x 7.6 mm": "১৪৬.৩ x ৭০.৯ x ৭.৬ মিমি",
		"168 g": "১৬৮ g",
		"3,900 mAh": "৩,৯০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ MP + ১০ MP + ১২ MP",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"8 GB": "৮ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 13, upgradable": "Android ১৩, upgradable",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"February 2023": "February ২০২৩",
		"IP68": "IP৬৮",
		"Phantom Black, Green, Cream, Lavender": "Phantom কালো, সবুজ, Cream, Lavender",
		"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AC Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s23'
func (s *SpecificationSeederMobileSamsungGalaxyS23) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s23"

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

	// Override model-specific values for Samsung Galaxy S23
	specs["Display Size"] = "6.1 inches"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Dynamic AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2340 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "168 g"
	specs["Dimensions"] = "146.3 x 70.9 x 7.6 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 10 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "3,900 mAh"
	specs["Operating System"] = "Android 13, upgradable"
	specs["Available Colors"] = "Phantom Black, Green, Cream, Lavender"
	specs["Announcement Date"] = "February 2023"
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
