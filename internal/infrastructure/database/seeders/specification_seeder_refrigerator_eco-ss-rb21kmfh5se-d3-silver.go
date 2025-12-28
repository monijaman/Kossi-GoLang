package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecossrb21kmfh5sed3silver seeds specifications/options for product 'eco-ss-rb21kmfh5se-d3-silver'
type SpecificationSeederRefrigeratorecossrb21kmfh5sed3silver struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecossrb21kmfh5sed3silver creates a new seeder instance
func NewSpecificationSeederRefrigeratorecossrb21kmfh5sed3silver() *SpecificationSeederRefrigeratorecossrb21kmfh5sed3silver {
	return &SpecificationSeederRefrigeratorecossrb21kmfh5sed3silver{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-ss-rb21kmfh5se-d3-silver"},
	}
}

func (s *SpecificationSeederRefrigeratorecossrb21kmfh5sed3silver) getBanglaTranslations() map[string]string {
	return map[string]string{
		"SAMSUNG":                     "স্যামসাং",
		"SS RB21KMFH5SE/D3 Silver":    "এসএস আরবি২১কে এম এফ এইচ৫এসই/ডি৩ সিলভার",
		"VCM":                         "ভিসিএম",
		"212":                         "২১২",
		"142":                         "১৪২",
		"70":                          "৭০",
		"218":                         "২১৮",
		"Frost":                       "ফ্রস্ট",
		"Digital Inverter Compressor": "ডিজিটাল ইনভার্টার কম্প্রেসার",
		"R600a":                       "আর৬০০এ",
		"Electric Silver":             "ইলেকট্রিক সিলভার",
		"1584*546*610 mm":             "১৫৮৪*৫৪৬*৬১০ মিমি",
		"20 Years Compressor Warranty, 1 Year Parts and Service Warranty":                                 "২০ বছর কম্প্রেসার ওয়ারেন্টি, ১ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Digital Inverter Compressor, Power Cool Function, R600a Refrigerant, Mono Cooling, Tact Display": "ডিজিটাল ইনভার্টার কম্প্রেসার, পাওয়ার কুল ফাংশন, আর৬০০এ রেফ্রিজারেন্ট, মনো কুলিং, ট্যাক্ট ডিসপ্লে",
		"Mono Cooling":            "মনো কুলিং",
		"Tact Display (Blue LED)": "ট্যাক্ট ডিসপ্লে (ব্লু এলইডি)",
		"3":                       "৩",
		"Yes":                     "হ্যাঁ",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-ss-rb21kmfh5se-d3-silver'
func (s *SpecificationSeederRefrigeratorecossrb21kmfh5sed3silver) Seed(db *gorm.DB) error {
	productSlug := "eco-ss-rb21kmfh5se-d3-silver"
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
		"Brand":                 "SAMSUNG",
		"Model Name":            "SS RB21KMFH5SE/D3 Silver",
		"Door Type":             "VCM",
		"Capacity":              "212",
		"Refrigerator Capacity": "142",
		"Freezer Capacity":      "70",
		"Gross Volume":          "218",
		"Net Volume":            "212",
		"Defrost Type":          "Frost",
		"Compressor Type":       "Digital Inverter Compressor",
		"Refrigerant":           "R600a",
		"Color":                 "Electric Silver",
		"Dimensions":            "1584*546*610 mm",
		"Warranty":              "20 Years Compressor Warranty, 1 Year Parts and Service Warranty",
		"Special Features":      "Digital Inverter Compressor, Power Cool Function, R600a Refrigerant, Mono Cooling, Tact Display",
		"Cooling Technology":    "Mono Cooling",
		"Temperature Control":   "Tact Display (Blue LED)",
		"Number of Shelves":     "3",
		"Ice Maker":             "Yes",
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
