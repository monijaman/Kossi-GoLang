package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecobcd320recflgdblack seeds specifications/options for product 'eco-bcd-320-rec-fl-gd-black'
type SpecificationSeederRefrigeratorecobcd320recflgdblack struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecobcd320recflgdblack creates a new seeder instance
func NewSpecificationSeederRefrigeratorecobcd320recflgdblack() *SpecificationSeederRefrigeratorecobcd320recflgdblack {
	return &SpecificationSeederRefrigeratorecobcd320recflgdblack{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-bcd-320-rec-fl-gd-black"},
	}
}

func (s *SpecificationSeederRefrigeratorecobcd320recflgdblack) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                         "ইকো+",
		"ECO+ BCD-320 REC FL GD Black": "ECO+ BCD-320 REC FL GD Black",
		"VCM":                          "VCM",
		"320":                          "৩২০",
		"203":                          "২০৩",
		"115":                          "১১৫",
		"318":                          "৩১৮",
		"205":                          "২০৫",
		"1793*603*615 mm":              "১৭৯৩*৬০৩*৬১৫ মিমি",
		"Black":                        "ব্ল্যাক",
		"R600a":                        "R600a",
		"Automatic":                    "অটোমেটিক",
		"Mechanical":                   "মেকানিক্যাল",
		"3":                            "৩",
		"2":                            "২",
		"1":                            "১",
		"R600a cooling system, Door Lock, Internal LED Lighting": "R600a কুলিং সিস্টেম, ডোর লক, ইন্টারনাল LED লাইটিং",
		"Yes":            "হ্যাঁ",
		"No":             "না",
		"Tempered glass": "টেম্পার্ড গ্লাস",
		"220~240":        "২২০~২৪০",
		"50":             "৫০",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-bcd-320-rec-fl-gd-black'
func (s *SpecificationSeederRefrigeratorecobcd320recflgdblack) Seed(db *gorm.DB) error {
	productSlug := "eco-bcd-320-rec-fl-gd-black"
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
		"Model Name":                  "ECO+ BCD-320 REC FL GD Black",
		"Door Type":                   "VCM",
		"Capacity":                    "320",
		"Refrigerator Capacity":       "203",
		"Freezer Capacity":            "115",
		"Gross Volume":                "320",
		"Net Volume":                  "318",
		"Dimensions":                  "1793*603*615 mm",
		"Color":                       "Black",
		"Compressor Type":             "R600a",
		"Cooling Technology":          "R600a cooling system",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "1",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Voltage":                     "220~240",
		"Frequency (Hz)":              "50",
		"Warranty":                    "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Special Features":            "R600a cooling system, Door Lock, Internal LED Lighting",
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
