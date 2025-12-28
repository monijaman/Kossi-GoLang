package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratoreco2b502hxhlebonysheen seeds specifications/options for product 'eco-2b502hxhl-ebony-sheen'
type SpecificationSeederRefrigeratoreco2b502hxhlebonysheen struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratoreco2b502hxhlebonysheen creates a new seeder instance
func NewSpecificationSeederRefrigeratoreco2b502hxhlebonysheen() *SpecificationSeederRefrigeratoreco2b502hxhlebonysheen {
	return &SpecificationSeederRefrigeratoreco2b502hxhlebonysheen{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-2b502hxhl-ebony-sheen"},
	}
}

func (s *SpecificationSeederRefrigeratoreco2b502hxhlebonysheen) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                    "ইকো+",
		"2B502HXHL":               "2B502HXHL",
		"Top Mount":               "টপ মাউন্ট",
		"438 Liters":              "438 লিটার",
		"321 Liters":              "321 লিটার",
		"117 Liters":              "117 লিটার",
		"1780 x 700 x 730 mm":     "1780 x 700 x 730 মিমি",
		"Ebony Sheen":             "ইবোনি শিন",
		"Smart Inverter":          "স্মার্ট ইনভার্টার",
		"Linear Cooling":          "লিনিয়ার কুলিং",
		"No-frost":                "নো-ফ্রস্ট",
		"Moving Twist":            "মুভিং টুইস্ট",
		"No":                      "না",
		"Yes":                     "হ্যাঁ",
		"10 Years Motor Warranty": "10 বছর মোটর ওয়ারেন্টি",
		"10":                      "10",
		"471 Liters":              "471 লিটার",
		"335 Liters":              "335 লিটার",
		"136 Liters":              "136 লিটার",
		"3 Piece (Cr)":            "3 পিস (ক্রোম)",
		"Tray":                    "ট্রে",
		"1 Box":                   "1 বক্স",
		"Full Basket (2EA)":       "ফুল বাস্কেট (2EA)",
		"Top (Point)":             "টপ (পয়েন্ট)",
		"E-Micom":                 "ই-মাইকম",
		"Hygiene Fresh+®, Door Cooling+, Multi-Air Flow, Moving Ice Maker, Antibacterial Gasket, Moist Balance Crisper": "হাইজিন ফ্রেশ+®, ডোর কুলিং+, মাল্টি-এয়ার ফ্লো, মুভিং আইস মেকার, অ্যান্টিব্যাকটেরিয়াল গ্যাসকেট, ময়েস্ট ব্যালেন্স ক্রিসপার",
	}
}

// Seed inserts specification records for the product identified by slug 'eco-2b502hxhl-ebony-sheen'
func (s *SpecificationSeederRefrigeratoreco2b502hxhlebonysheen) Seed(db *gorm.DB) error {
	productSlug := "eco-2b502hxhl-ebony-sheen"
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
		"Model Name":                  "2B502HXHL",
		"Door Type":                   "Top Mount",
		"Capacity":                    "438 Liters",
		"Refrigerator Capacity":       "321 Liters",
		"Freezer Capacity":            "117 Liters",
		"Dimensions":                  "1780 x 700 x 730 mm",
		"Color":                       "Ebony Sheen",
		"Compressor Type":             "Smart Inverter",
		"Cooling Technology":          "Linear Cooling",
		"Defrost Type":                "No-frost",
		"Ice Maker":                   "Moving Twist",
		"Water Dispenser":             "No",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Motor Warranty",
		"Compressor Warranty (Years)": "10",
		"Gross Volume":                "471 Liters",
		"Net Volume":                  "438 Liters",
		"Display":                     "E-Micom",
		"LED Light":                   "Top (Point)",
		"Handle Type":                 "3 Piece (Cr)",
		"Egg Tray":                    "Tray",
		"Crisper Drawers":             "1 Box",
		"Shelf":                       "Yes",
		"Basket":                      "Full Basket (2EA)",
		"Special Features":            "Hygiene Fresh+®, Door Cooling+, Multi-Air Flow, Moving Ice Maker, Antibacterial Gasket, Moist Balance Crisper",
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
