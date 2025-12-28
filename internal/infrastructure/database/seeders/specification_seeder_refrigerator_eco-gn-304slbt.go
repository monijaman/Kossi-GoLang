package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecogn304slbt seeds specifications/options for product 'eco-gn-304slbt'
type SpecificationSeederRefrigeratorecogn304slbt struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecogn304slbt creates a new seeder instance
func NewSpecificationSeederRefrigeratorecogn304slbt() *SpecificationSeederRefrigeratorecogn304slbt {
	return &SpecificationSeederRefrigeratorecogn304slbt{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-gn-304slbt"},
	}
}

func (s *SpecificationSeederRefrigeratorecogn304slbt) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LG":              "এলজি",
		"LG GN-304SLBT":   "এলজি জিএন-৩০৪এসএলবিটি",
		"Freezer":         "ফ্রিজার",
		"165":             "১৬৫",
		"171":             "১৭১",
		"Frost":           "ফ্রস্ট",
		"Inverter":        "ইনভার্টার",
		"Platinum Silver": "প্লাটিনাম সিলভার",
		"1300*530*600 mm": "১৩০০*৫৩০*৬০০ মিমি",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty": "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Direct Cooling System, Inverter Compressor, Smart Storage":        "ডাইরেক্ট কুলিং সিস্টেম, ইনভার্টার কম্প্রেসার, স্মার্ট স্টোরেজ",
		"Faster Cooling": "ফাস্টার কুলিং",
		"Yes":            "হ্যাঁ",
		"6":              "৬",
		"1":              "১",
		"4":              "৪",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-gn-304slbt'
func (s *SpecificationSeederRefrigeratorecogn304slbt) Seed(db *gorm.DB) error {
	productSlug := "eco-gn-304slbt"
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
		"Brand":              "LG",
		"Model Name":         "LG GN-304SLBT",
		"Door Type":          "Freezer",
		"Capacity":           "165",
		"Freezer Capacity":   "165",
		"Gross Volume":       "171",
		"Net Volume":         "165",
		"Defrost Type":       "Frost",
		"Compressor Type":    "Inverter",
		"Color":              "Platinum Silver",
		"Dimensions":         "1300*530*600 mm",
		"Warranty":           "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Special Features":   "Direct Cooling System, Inverter Compressor, Smart Storage",
		"Cooling Technology": "Faster Cooling",
		"Shelf Material":     "Wire Shelves",
		"Number of Shelves":  "6",
		"Crisper Drawers":    "Yes",
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
