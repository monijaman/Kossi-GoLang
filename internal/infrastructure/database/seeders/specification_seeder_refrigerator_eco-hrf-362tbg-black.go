package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecohrf362tbgblack seeds specifications/options for product 'eco-hrf-362tbg-black'
type SpecificationSeederRefrigeratorecohrf362tbgblack struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecohrf362tbgblack creates a new seeder instance
func NewSpecificationSeederRefrigeratorecohrf362tbgblack() *SpecificationSeederRefrigeratorecohrf362tbgblack {
	return &SpecificationSeederRefrigeratorecohrf362tbgblack{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-hrf-362tbg-black"},
	}
}

func (s *SpecificationSeederRefrigeratorecohrf362tbgblack) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HAIER":                  "হাইয়ার",
		"Haier HRF-362TBG Black": "হাইয়ার এইচআরএফ-৩৬২টিবিজি ব্ল্যাক",
		"Top Mount":              "টপ মাউন্ট",
		"333":                    "৩৩৩",
		"247":                    "২৪৭",
		"86":                     "৮৬",
		"342":                    "৩৪২",
		"No-frost":               "নো-ফ্রস্ট",
		"Twin Inverter":          "টুইন ইনভার্টার",
		"R600a":                  "আর৬০০এ",
		"Black":                  "ব্ল্যাক",
		"710 x 1640 x 598 mm":    "৭১০ x ১৬৪০ x ৫৯৮ মিমি",
		"53 kg":                  "৫৩ কেজি",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty":          "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Twin Inverter Technology, 360º Precise Cooling, Large Space to Store More": "টুইন ইনভার্টার টেকনোলজি, ৩৬০º প্রিসাইজ কুলিং, লার্জ স্পেস টু স্টোর মোর",
		"Electrical":                 "ইলেকট্রিকাল",
		"Recess":                     "রিসেস",
		"Yes":                        "হ্যাঁ",
		"4":                          "৪",
		"Ice Tray + Ice Storage Box": "আইস ট্রে + আইস স্টোরেজ বক্স",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-hrf-362tbg-black'
func (s *SpecificationSeederRefrigeratorecohrf362tbgblack) Seed(db *gorm.DB) error {
	productSlug := "eco-hrf-362tbg-black"
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
		"Brand":                 "HAIER",
		"Model Name":            "Haier HRF-362TBG Black",
		"Door Type":             "Top Mount",
		"Capacity":              "333",
		"Refrigerator Capacity": "247",
		"Freezer Capacity":      "86",
		"Gross Volume":          "342",
		"Net Volume":            "333",
		"Defrost Type":          "No-frost",
		"Compressor Type":       "Twin Inverter",
		"Color":                 "Black",
		"Dimensions":            "710 x 1640 x 598 mm",
		"Weight":                "53 kg",
		"Warranty":              "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Special Features":      "Twin Inverter Technology, 360º Precise Cooling, Large Space to Store More",
		"Refrigerant":           "R600a",
		"Temperature Control":   "Electrical",
		"Number of Shelves":     "4",
		"Ice Maker":             "Ice Tray + Ice Storage Box",
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
