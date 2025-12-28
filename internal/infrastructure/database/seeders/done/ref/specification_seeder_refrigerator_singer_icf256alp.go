package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerICF256ALP seeds specifications/options for product 'singer-icf-256a-lp'
type SpecificationSeederRefrigeratorSingerICF256ALP struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerICF256ALP creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerICF256ALP() *SpecificationSeederRefrigeratorSingerICF256ALP {
	return &SpecificationSeederRefrigeratorSingerICF256ALP{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-icf-256a-lp"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerICF256ALP) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":             "সিঙ্গার",
		"ICF-256A-LP":        "ICF-256A-LP",
		"Display Freezer":    "ডিসপ্লে ফ্রিজার",
		"164 Liters":         "164 লিটার",
		"5 Star":             "৫ তারা",
		"5":                  "৫",
		"250 kWh":            "২৫০ কিলোওয়াট ঘণ্টা",
		"800 x 630 x 840 mm": "800 x 630 x 840 মিমি",
		"204 kg":             "204 kg",
		"Silver":             "সিলভার",
		"Inverter":           "ইনভার্টার",
		"Direct Cool":        "ডাইরেক্ট কুল",
		"Manual":             "ম্যানুয়াল",
		"Mechanical":         "মেকানিক্যাল",
		"Wire Shelves":       "ওয়্যার শেল্ফ",
		"2":                  "২",
		"3":                  "৩",
		"1":                  "১",
		"Yes":                "হ্যাঁ",
		"No":                 "না",
		"40 dB":              "৪০ ডেসিবেল",
		"170 ~ 260V":         "170 ~ 260ভোল্ট",
		"50":                 "৫০",
		"10 Years Compressor & 2 Years Spare Parts & Service Warranty":                            "10 বছর কম্প্রেসর এবং 2 বছর স্পেয়ার পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Display Freezer, Ice Cream Freezer, Fan Cooling, Glass and Aluminum Door and Body Liner": "ডিসপ্লে ফ্রিজার, আইসক্রিম ফ্রিজার, ফ্যান কুলিং, গ্লাস এবং অ্যালুমিনিয়াম ডোর এবং বডি লাইনার",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"256 Ltr.":         "256 লিটার",
		"R290":             "আর290",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-icf-256a-lp'
func (s *SpecificationSeederRefrigeratorSingerICF256ALP) Seed(db *gorm.DB) error {
	productSlug := "singer-icf-256a-lp"
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
		"Brand":                       "Singer",
		"Model Name":                  "ICF-256A-LP",
		"Door Type":                   "Display Freezer",
		"Capacity":                    "164 Liters",
		"Refrigerator Capacity":       "",
		"Freezer Capacity":            "164 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "800 x 630 x 840 mm",
		"Weight":                      "60 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Fan Cooling",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire Shelves",
		"Number of Shelves":           "2",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "170 ~ 260V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor & 2 Years Spare Parts & Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R290",
		"Gross Volume":                "164 Ltr.",
		"Net Volume":                  "164 Ltr.",
		"Special Features":            "Display Freezer, Ice Cream Freezer, Fan Cooling, Glass and Aluminum Door and Body Liner",
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
