package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecorbg410p6pbxxgr seeds specifications/options for product 'eco-r-bg410p6pbx-xgr'
type SpecificationSeederRefrigeratorecorbg410p6pbxxgr struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecorbg410p6pbxxgr creates a new seeder instance
func NewSpecificationSeederRefrigeratorecorbg410p6pbxxgr() *SpecificationSeederRefrigeratorecorbg410p6pbxxgr {
	return &SpecificationSeederRefrigeratorecorbg410p6pbxxgr{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-r-bg410p6pbx-xgr"},
	}
}

func (s *SpecificationSeederRefrigeratorecorbg410p6pbxxgr) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HITACHI":                  "হিটাচি",
		"Hitachi R-BG410P6PBX XGR": "হিটাচি আর-বিজি৪১০পি৬পিবিএক্স এক্সজিআর",
		"Bottom Mount":             "বটম মাউন্ট",
		"330":                      "৩৩০",
		"215":                      "২১৫",
		"115":                      "১১৫",
		"366":                      "৩৬৬",
		"139":                      "১৩৯",
		"227":                      "২২৭",
		"No-frost":                 "নো-ফ্রস্ট",
		"Inverter Control":         "ইনভার্টার কন্ট্রোল",
		"Gradation Gray":           "গ্র্যাডেশন গ্রে",
		"1900*595*650 mm":          "১৯০০*৫৯৫*৬৫০ মিমি",
		"220-240V":                 "২২০-২৪০ভি",
		"50/60":                    "৫০/৬০",
		"Inverter x Dual Fan Cooling, Powerful Cooling, Flat Glass Door, Touch Screen Controller, Nano Titanium Filter": "ইনভার্টার এক্স ডুয়াল ফ্যান কুলিং, পাওয়ারফুল কুলিং, ফ্ল্যাট গ্লাস ডোর, টাচ স্ক্রিন কন্ট্রোলার, ন্যানো টাইটানিয়াম ফিল্টার",
		"Touch Screen Controller": "টাচ স্ক্রিন কন্ট্রোলার",
		"Tempered Glass":          "টেম্পার্ড গ্লাস",
		"Yes":                     "হ্যাঁ",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-r-bg410p6pbx-xgr'
func (s *SpecificationSeederRefrigeratorecorbg410p6pbxxgr) Seed(db *gorm.DB) error {
	productSlug := "eco-r-bg410p6pbx-xgr"
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
		"Brand":                 "HITACHI",
		"Model Name":            "Hitachi R-BG410P6PBX XGR",
		"Door Type":             "Bottom Mount",
		"Capacity":              "330",
		"Refrigerator Capacity": "215",
		"Freezer Capacity":      "115",
		"Gross Volume":          "366",
		"Net Volume":            "330",
		"Defrost Type":          "No-frost",
		"Compressor Type":       "Inverter Control",
		"Color":                 "Gradation Gray",
		"Dimensions":            "1900*595*650 mm",
		"Voltage":               "220-240V",
		"Frequency (Hz)":        "50/60",
		"Special Features":      "Inverter x Dual Fan Cooling, Powerful Cooling, Flat Glass Door, Touch Screen Controller, Nano Titanium Filter",
		"Temperature Control":   "Touch Screen Controller",
		"Shelf Material":        "Tempered Glass",
		"Crisper Drawers":       "Yes",
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
