package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecorvg490p8pbxgr seeds specifications/options for product 'eco-r-vg490p8pb-xgr'
type SpecificationSeederRefrigeratorecorvg490p8pbxgr struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecorvg490p8pbxgr creates a new seeder instance
func NewSpecificationSeederRefrigeratorecorvg490p8pbxgr() *SpecificationSeederRefrigeratorecorvg490p8pbxgr {
	return &SpecificationSeederRefrigeratorecorvg490p8pbxgr{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-r-vg490p8pb-xgr"},
	}
}

func (s *SpecificationSeederRefrigeratorecorvg490p8pbxgr) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HITACHI":                 "হিটাচি",
		"Hitachi R-VG490P8PB XGR": "হিটাচি আর-ভিজি৪৯০পি৮পিবি এক্সজিআর",
		"Top Mount":               "টপ মাউন্ট",
		"407":                     "৪০৭",
		"298":                     "২৯৮",
		"109":                     "১০৯",
		"443":                     "৪৪৩",
		"133":                     "১৩৩",
		"310":                     "৩১০",
		"No-frost":                "নো-ফ্রস্ট",
		"Inverter":                "ইনভার্টার",
		"R600a":                   "আর৬০০এ",
		"Gradation Gray":          "গ্র্যাডেশন গ্রে",
		"1770*680*720 mm":         "১৭৭০*৬৮০*৭২০ মিমি",
		"220-240V":                "২২০-২৪০ভি",
		"50/60":                   "৫০/৬০",
		"Inverter x Dual Fan Cooling, Fresh Select, Powerful Deodorization, Super Stable Operation, Eco-Friendly R600a Gas": "ইনভার্টার এক্স ডুয়াল ফ্যান কুলিং, ফ্রেশ সিলেক্ট, পাওয়ারফুল ডিওডোরাইজেশন, সুপার স্টেবল অপারেশন, ইকো-ফ্রেন্ডলি আর৬০০এ গ্যাস",
		"Tempered Glass":         "টেম্পার্ড গ্লাস",
		"Yes":                    "হ্যাঁ",
		"Powerful Deodorization": "পাওয়ারফুল ডিওডোরাইজেশন",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-r-vg490p8pb-xgr'
func (s *SpecificationSeederRefrigeratorecorvg490p8pbxgr) Seed(db *gorm.DB) error {
	productSlug := "eco-r-vg490p8pb-xgr"
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
		"Model Name":            "Hitachi R-VG490P8PB XGR",
		"Door Type":             "Top Mount",
		"Capacity":              "407",
		"Refrigerator Capacity": "298",
		"Freezer Capacity":      "109",
		"Gross Volume":          "443",
		"Net Volume":            "407",
		"Defrost Type":          "No-frost",
		"Compressor Type":       "Inverter",
		"Refrigerant":           "R600a",
		"Color":                 "Gradation Gray",
		"Dimensions":            "1770*680*720 mm",
		"Voltage":               "220-240V",
		"Frequency (Hz)":        "50/60",
		"Special Features":      "Inverter x Dual Fan Cooling, Fresh Select, Powerful Deodorization, Super Stable Operation, Eco-Friendly R600a Gas",
		"Shelf Material":        "Tempered Glass",
		"Ice Maker":             "Movable Twist Ice Tray",
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
