package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecobcd218flgdmagnoliamarronwb seeds specifications/options for product 'eco-bcd-218-fl-gd-magnolia-marron-wb'
type SpecificationSeederRefrigeratorecobcd218flgdmagnoliamarronwb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecobcd218flgdmagnoliamarronwb creates a new seeder instance
func NewSpecificationSeederRefrigeratorecobcd218flgdmagnoliamarronwb() *SpecificationSeederRefrigeratorecobcd218flgdmagnoliamarronwb {
	return &SpecificationSeederRefrigeratorecobcd218flgdmagnoliamarronwb{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-bcd-218-fl-gd-magnolia-marron-wb"},
	}
}

func (s *SpecificationSeederRefrigeratorecobcd218flgdmagnoliamarronwb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                                  "ইকো+",
		"ECO+ BCD-218 FL GD Magnolia Marron WB": "ECO+ BCD-218 FL GD Magnolia Marron WB",
		"Glass Door":                            "গ্লাস ডোর",
		"218 liter":                             "২১৮ লিটার",
		"1505*550*490 mm":                       "১৫০৫*৫৫০*৪৯০ মিমি",
		"Magnolia Marron":                       "ম্যাগনোলিয়া ম্যারন",
		"R600a":                                 "R600a",
		"Manual":                                "ম্যানুয়াল",
		"Mechanical":                            "মেকানিক্যাল",
		"2":                                     "২",
		"1":                                     "১",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty": "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"10": "১০",
		"Frameless Design, Interior LED Light, Eco-Friendly Cooling, Key & Lock": "ফ্রেমলেস ডিজাইন, ইন্টেরিয়র LED লাইট, ইকো-ফ্রেন্ডলি কুলিং, কী এবং লক",
		"Yes": "হ্যাঁ",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-bcd-218-fl-gd-magnolia-marron-wb'
func (s *SpecificationSeederRefrigeratorecobcd218flgdmagnoliamarronwb) Seed(db *gorm.DB) error {
	productSlug := "eco-bcd-218-fl-gd-magnolia-marron-wb"
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
		"Model Name":                  "ECO+ BCD-218 FL GD Magnolia Marron WB",
		"Door Type":                   "Glass Door",
		"Capacity":                    "218 liter",
		"Dimensions":                  "1505*550*490 mm",
		"Color":                       "Magnolia Marron",
		"Refrigerant":                 "R600a",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Number of Shelves":           "2",
		"Crisper Drawers":             "1",
		"Warranty":                    "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Net Volume":                  "218 liter",
		"Special Features":            "Frameless Design, Interior LED Light, Eco-Friendly Cooling, Key & Lock",
		"Ice Maker":                   "Yes",
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
