package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecobcd170flgdpurpleazalea seeds specifications/options for product 'eco-bcd-170-fl-gd-purple-azalea'
type SpecificationSeederRefrigeratorecobcd170flgdpurpleazalea struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecobcd170flgdpurpleazalea creates a new seeder instance
func NewSpecificationSeederRefrigeratorecobcd170flgdpurpleazalea() *SpecificationSeederRefrigeratorecobcd170flgdpurpleazalea {
	return &SpecificationSeederRefrigeratorecobcd170flgdpurpleazalea{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-bcd-170-fl-gd-purple-azalea"},
	}
}

func (s *SpecificationSeederRefrigeratorecobcd170flgdpurpleazalea) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                        "ইকো+",
		"BCD-170 FL GD Purple Azalea": "BCD-170 FL GD Purple Azalea",
		"Glass Door":                  "গ্লাস ডোর",
		"170 Liters":                  "170 লিটার",
		"101 Liters":                  "101 লিটার",
		"69 Liters":                   "69 লিটার",
		"1250 x 535 x 579 mm":         "1250 x 535 x 579 মিমি",
		"Purple Azalea":               "পার্পল অ্যাজালিয়া",
		"R600a":                       "R600a",
		"Mechanical":                  "মেকানিক্যাল",
		"Frost":                       "ফ্রস্ট",
		"Manual":                      "ম্যানুয়াল",
		"Yes":                         "হ্যাঁ",
		"No":                          "না",
		"10 Years Motor Warranty":     "10 বছর মোটর ওয়ারেন্টি",
		"10":                          "10",
		"Integrated":                  "ইন্টিগ্রেটেড",
		"Flat":                        "ফ্ল্যাট",
		"LED lighting":                "এলইডি লাইটিং",
		"2":                           "2",
		"Tempered glass":              "টেম্পার্ড গ্লাস",
		"1":                           "1",
		"Frameless Design, Eco-Friendly Cooling with R600a Gas, LED Lighting, Big Freezer Space, Elegant Exterior Design, Colorful Glass Panel": "ফ্রেমলেস ডিজাইন, R600a গ্যাস দিয়ে ইকো-ফ্রেন্ডলি কুলিং, এলইডি লাইটিং, বড় ফ্রিজার স্পেস, এলিগ্যান্ট এক্সটেরিয়র ডিজাইন, কালারফুল গ্লাস প্যানেল",
	}
}

// Seed inserts specification records for the product identified by slug 'eco-bcd-170-fl-gd-purple-azalea'
func (s *SpecificationSeederRefrigeratorecobcd170flgdpurpleazalea) Seed(db *gorm.DB) error {
	productSlug := "eco-bcd-170-fl-gd-purple-azalea"
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
		"Model Name":                  "BCD-170 FL GD Purple Azalea",
		"Door Type":                   "Glass Door",
		"Capacity":                    "170 Liters",
		"Refrigerator Capacity":       "101 Liters",
		"Freezer Capacity":            "69 Liters",
		"Dimensions":                  "1250 x 535 x 579 mm",
		"Color":                       "Purple Azalea",
		"Compressor Type":             "R600a",
		"Defrost Type":                "Frost",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Warranty":                    "10 Years Motor Warranty",
		"Compressor Warranty (Years)": "10",
		"Net Volume":                  "170 Liters",
		"Display":                     "Mechanical",
		"LED Light":                   "LED lighting",
		"Handle Type":                 "Integrated",
		"Egg Tray":                    "Yes",
		"Crisper Drawers":             "1",
		"Shelf Material":              "Tempered glass",
		"Number of Shelves":           "2",
		"Voltage":                     "220~240 / 50Hz",
		"Frequency (Hz)":              "50",
		"Refrigerant":                 "R600a",
		"Special Features":            "Frameless Design, Eco-Friendly Cooling with R600a Gas, LED Lighting, Big Freezer Space, Elegant Exterior Design, Colorful Glass Panel",
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
