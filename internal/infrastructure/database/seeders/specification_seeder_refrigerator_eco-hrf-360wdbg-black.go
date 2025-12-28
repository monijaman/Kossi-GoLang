package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecohrf360wdbgblack seeds specifications/options for product 'eco-hrf-360wdbg-black'
type SpecificationSeederRefrigeratorecohrf360wdbgblack struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecohrf360wdbgblack creates a new seeder instance
func NewSpecificationSeederRefrigeratorecohrf360wdbgblack() *SpecificationSeederRefrigeratorecohrf360wdbgblack {
	return &SpecificationSeederRefrigeratorecohrf360wdbgblack{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-hrf-360wdbg-black"},
	}
}

func (s *SpecificationSeederRefrigeratorecohrf360wdbgblack) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HAIER":                   "হাইয়ার",
		"Haier HRF-360WDBG Black": "হাইয়ার এইচআরএফ-৩৬০ডব্লিউডিবিজি ব্ল্যাক",
		"Glass Door":              "গ্লাস ডোর",
		"340":                     "৩৪০",
		"96":                      "৯৬",
		"205":                     "২০৫",
		"39":                      "৩৯",
		"No-frost":                "নো-ফ্রস্ট",
		"Inverter":                "ইনভার্টার",
		"R600a":                   "আর৬০০এ",
		"Black":                   "ব্ল্যাক",
		"698*1735*598 mm":         "৬৯৮*১৭৩৫*৫৯৮ মিমি",
		"66.5 kg":                 "৬৬.৫ কেজি",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty":                                                                                       "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Twin Inverter Technology, Magic Cooling 3rd Space, 360° Airflow System, Large Freezer Space To Store More, Deo Fresh Technology, Digital Control Panel": "টুইন ইনভার্টার টেকনোলজি, ম্যাজিক কুলিং থার্ড স্পেস, ৩৬০° এয়ারফ্লো সিস্টেম, লার্জ ফ্রিজার স্পেস টু স্টোর মোর, ডিও ফ্রেশ টেকনোলজি, ডিজিটাল কন্ট্রোল প্যানেল",
		"Recess type": "রিসেস টাইপ",
		"Yes":         "হ্যাঁ",
		"4":           "৪",
		"1":           "১",
		"2":           "২",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-hrf-360wdbg-black'
func (s *SpecificationSeederRefrigeratorecohrf360wdbgblack) Seed(db *gorm.DB) error {
	productSlug := "eco-hrf-360wdbg-black"
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
		"Model Name":            "Haier HRF-360WDBG Black",
		"Door Type":             "Glass Door",
		"Capacity":              "340",
		"Refrigerator Capacity": "205",
		"Freezer Capacity":      "96",
		"Gross Volume":          "340",
		"Net Volume":            "340",
		"Defrost Type":          "No-frost",
		"Compressor Type":       "Inverter",
		"Color":                 "Black",
		"Dimensions":            "698*1735*598 mm",
		"Weight":                "66.5 kg",
		"Warranty":              "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Special Features":      "Twin Inverter Technology, Magic Cooling 3rd Space, 360° Airflow System, Large Freezer Space To Store More, Deo Fresh Technology, Digital Control Panel",
		"Refrigerant":           "R600a",
		"Number of Shelves":     "4",
		"Crisper Drawers":       "Yes",
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
