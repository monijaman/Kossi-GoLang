package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratoreco2b312pxcbebonysheen seeds specifications/options for product 'eco-2b312pxcb-ebony-sheen'
type SpecificationSeederRefrigeratoreco2b312pxcbebonysheen struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratoreco2b312pxcbebonysheen creates a new seeder instance
func NewSpecificationSeederRefrigeratoreco2b312pxcbebonysheen() *SpecificationSeederRefrigeratoreco2b312pxcbebonysheen {
	return &SpecificationSeederRefrigeratoreco2b312pxcbebonysheen{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-2b312pxcb-ebony-sheen"},
	}
}

func (s *SpecificationSeederRefrigeratoreco2b312pxcbebonysheen) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                      "ইকো+",
		"2B312PXCB Ebony Sheen":     "2B312PXCB ইবোনি শিন",
		"Top Mount":                 "টপ মাউন্ট",
		"315 Liters":                "315 লিটার",
		"236 Liters":                "236 লিটার",
		"79 Liters":                 "79 লিটার",
		"1640 x 600 x 710 mm":       "1640 x 600 x 710 মিমি",
		"Ebony Sheen":               "ইবোনি শিন",
		"Smart Inverter Compressor": "স্মার্ট ইনভার্টার কম্প্রেসার",
		"Linear Cooling":            "লিনিয়ার কুলিং",
		"No-frost":                  "নো-ফ্রস্ট",
		"Glass":                     "গ্লাস",
		"Yes":                       "হ্যাঁ",
		"No":                        "না",
		"10 Years Motor Warranty":   "10 বছর মোটর ওয়ারেন্টি",
		"10":                        "10",
		"340 Liters":                "340 লিটার",
		"Door Cooling+, Wi-Fi, E-Micom Display, LED Light, Deodorizer, Express Freeze, Pocket Chrome Handle": "ডোর কুলিং+, ওয়াই-ফাই, ই-মাইকম ডিসপ্লে, এলইডি লাইট, ডিওডোরাইজার, এক্সপ্রেস ফ্রিজ, পকেট ক্রোম হ্যান্ডেল",
	}
}

// Seed inserts specification records for the product identified by slug 'eco-2b312pxcb-ebony-sheen'
func (s *SpecificationSeederRefrigeratoreco2b312pxcbebonysheen) Seed(db *gorm.DB) error {
	productSlug := "eco-2b312pxcb-ebony-sheen"
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
		"Brand":                       "ECO+",
		"Model Name":                  "2B312PXCB Ebony Sheen",
		"Door Type":                   "Top Mount",
		"Capacity":                    "315 Liters",
		"Refrigerator Capacity":       "236 Liters",
		"Freezer Capacity":            "79 Liters",
		"Dimensions":                  "1640 x 600 x 710 mm",
		"Color":                       "Ebony Sheen",
		"Compressor Type":             "Smart Inverter Compressor",
		"Cooling Technology":          "Linear Cooling",
		"Defrost Type":                "No-frost",
		"Shelf Material":              "Glass",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Warranty":                    "10 Years Motor Warranty",
		"Compressor Warranty (Years)": "10",
		"Gross Volume":                "340 Liters",
		"Net Volume":                  "315 Liters",
		"Special Features":            "Door Cooling+, Wi-Fi, E-Micom Display, LED Light, Deodorizer, Express Freeze, Pocket Chrome Handle",
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
