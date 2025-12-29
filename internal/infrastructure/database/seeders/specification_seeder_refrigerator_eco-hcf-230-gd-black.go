package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecohcf230gdblack seeds specifications/options for product 'eco-hcf-230-gd-black'
type SpecificationSeederRefrigeratorecohcf230gdblack struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecohcf230gdblack creates a new seeder instance
func NewSpecificationSeederRefrigeratorecohcf230gdblack() *SpecificationSeederRefrigeratorecohcf230gdblack {
	return &SpecificationSeederRefrigeratorecohcf230gdblack{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-hcf-230-gd-black"},
	}
}

func (s *SpecificationSeederRefrigeratorecohcf230gdblack) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HAIER":                               "হাইয়ার",
		"Haier HCF-230 GD Black With Display": "হাইয়ার এইচসিএফ-২৩০ জিডি ব্ল্যাক উইথ ডিসপ্লে",
		"Chest Freezer":                       "চেস্ট ফ্রিজার",
		"200":                                 "২০০",
		"R600a":                               "আর৬০০এ",
		"Black":                               "ব্ল্যাক",
		"820 x 568 x 865 mm":                  "৮২০ x ৫৬৮ x ৮৬৫ মিমি",
		"36 kg":                               "৩৬ কেজি",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty":                            "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Super Fast Cooling up to -30° degrees, 150 Hours Frozen Retention, Electronic Control Panel": "সুপার ফাস্ট কুলিং আপ টু -৩০° ডিগ্রি, ১৫০ ঘণ্টা ফ্রোজেন রিটেনশন, ইলেকট্রনিক কন্ট্রোল প্যানেল",
		"Electronic Control Panel": "ইলেকট্রনিক কন্ট্রোল প্যানেল",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-hcf-230-gd-black'
func (s *SpecificationSeederRefrigeratorecohcf230gdblack) Seed(db *gorm.DB) error {
	productSlug := "eco-hcf-230-gd-black"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID
	existingkeyMapping := map[string]uint{
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Dimensions":                  25,
		"Weight":                      80,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
		"Special Features":            69,
	}

	specs := map[string]string{
		"Brand":               "HAIER",
		"Model Name":          "Haier HCF-230 GD Black With Display",
		"Door Type":           "Chest Freezer",
		"Capacity":            "200",
		"Freezer Capacity":    "200",
		"Gross Volume":        "200",
		"Net Volume":          "200",
		"Defrost Type":        "Manual",
		"Compressor Type":     "Fix",
		"Color":                       "Black",
		"Dimensions":          "820 x 568 x 865 mm",
		"Weight":              "36 kg",
		"Warranty":            "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Special Features":    "Super Fast Cooling up to -30° degrees, 150 Hours Frozen Retention, Electronic Control Panel",
		"Refrigerant":         "R600a",
		"Temperature Control": "Electronic Control Panel",
	}

	banglaTranslations := s.getBanglaTranslations()
for key, val := range specs {
    if len(val) > 500 {
        specs[key] = val[:500]
    }
	}
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
			continue
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: specKeyID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Last(&sModel).Error; err != nil {
					return err
				}

				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					translation := &models.SpecificationTranslationModel{
						SpecificationID: sModel.ID,
						Locale:          "bn",
						Value:           banglaValue,
					}
					if err := db.Create(translation).Error; err != nil {
						return err
					}
				}
			} else {
				return err
			}
		} else {
			// Update the value if different
			if existing.Value != value {
				existing.Value = value
				if err := db.Save(&existing).Error; err != nil {
					return err
				}
			}
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
				} else {
					// Update translation if different
					if existingTranslation.Value != banglaValue {
						existingTranslation.Value = banglaValue
						if err := db.Save(&existingTranslation).Error; err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
