package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoSparkGo2 seeds specifications/options for product 'tecno-spark-go-2'
type SpecificationSeederMobileTecnoSparkGo2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSparkGo2 creates a new seeder instance
func NewSpecificationSeederMobileTecnoSparkGo2() *SpecificationSeederMobileTecnoSparkGo2 {
	return &SpecificationSeederMobileTecnoSparkGo2{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK Go 2"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSparkGo2) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + AI": "১৩ MP + AI",
		"164.5 x 76 x 9 mm": "১৬৪.৫ x ৭৬ x ৯ মিমি",
		"188 g": "১৮৮ g",
		"2 GB": "২ GB",
		"2020": "২০২০",
		"32 GB": "৩২ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.52 inches": "৬.৫২ ইঞ্চি",
		"60Hz": "৬০Hz",
		"720 x 1600 pixels": "৭২০ x ১৬০০ pixels",
		"8 MP": "৮ MP",
		"Android 10 Go Edition": "Android ১০ Go Edition",
		"Helio A22": "Helio A২২",
		"Ice Jadeite, Aqua Blue": "Ice Jadeite, Aqua নীল",
		"Mediatek Helio A22 (12 nm)": "Mediatek Helio A২২ (১২ nm)",
		"PowerVR GE8320": "PowerVR GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-spark-go-2'
func (s *SpecificationSeederMobileTecnoSparkGo2) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-go-2"

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

	// Override model-specific values for Tecno SPARK Go 2
	specs["Display Size"] = "6.52 inches"
	specs["Processor"] = "Helio A22"
	specs["Chipset"] = "Mediatek Helio A22 (12 nm)"
	specs["Cpu Type"] = "Quad-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "2 GB"
	specs["Storage"] = "32 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "188 g"
	specs["Dimensions"] = "164.5 x 76 x 9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "13 MP + AI"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 10 Go Edition"
	specs["Available Colors"] = "Ice Jadeite, Aqua Blue"
	specs["Announcement Date"] = "2020"
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
