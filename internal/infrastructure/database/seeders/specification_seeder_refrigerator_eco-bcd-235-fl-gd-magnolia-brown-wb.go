package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecobcd235flgdmagnoliabrownwb seeds specifications/options for product 'eco-bcd-235-fl-gd-magnolia-brown-wb'
type SpecificationSeederRefrigeratorecobcd235flgdmagnoliabrownwb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecobcd235flgdmagnoliabrownwb creates a new seeder instance
func NewSpecificationSeederRefrigeratorecobcd235flgdmagnoliabrownwb() *SpecificationSeederRefrigeratorecobcd235flgdmagnoliabrownwb {
	return &SpecificationSeederRefrigeratorecobcd235flgdmagnoliabrownwb{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-bcd-235-fl-gd-magnolia-brown-wb"},
	}
}

func (s *SpecificationSeederRefrigeratorecobcd235flgdmagnoliabrownwb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                                 "ইকো+",
		"ECO+ BCD-235 FL GD Magnolia Brown WB": "ECO+ BCD-235 FL GD Magnolia Brown WB",
		"Glass Door":                           "গ্লাস ডোর",
		"235 liter":                            "২৩৫ লিটার",
		"147 liter":                            "১৪৭ লিটার",
		"88 liter":                             "৮৮ লিটার",
		"1613*550*490 mm":                      "১৬১৩*৫৫০*৪৯০ মিমি",
		"Magnolia Brown":                       "ম্যাগনোলিয়া ব্রাউন",
		"R600a":                                "R600a",
		"Manual":                               "ম্যানুয়াল",
		"Mechanical":                           "মেকানিক্যাল",
		"2":                                    "২",
		"1":                                    "১",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty": "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"10": "১০",
		"Frameless Door, R600a Refrigerant, Key & Lock, Interior Light": "ফ্রেমলেস ডোর, R600a রেফ্রিজারেন্ট, কী এবং লক, ইন্টেরিয়র লাইট",
		"Yes":            "হ্যাঁ",
		"No":             "না",
		"Tempered glass": "টেম্পার্ড গ্লাস",
		"220~240":        "২২০~২৪০",
		"50":             "৫০",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-bcd-235-fl-gd-magnolia-brown-wb'
func (s *SpecificationSeederRefrigeratorecobcd235flgdmagnoliabrownwb) Seed(db *gorm.DB) error {
	productSlug := "eco-bcd-235-fl-gd-magnolia-brown-wb"
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
		"Model Name":                  "ECO+ BCD-235 FL GD Magnolia Brown WB",
		"Door Type":                   "Glass Door",
		"Capacity":                    "235 liter",
		"Refrigerator Capacity":       "147 liter",
		"Freezer Capacity":            "88 liter",
		"Dimensions":                  "1613*550*490 mm",
		"Color":                       "Magnolia Brown",
		"Refrigerant":                 "R600a",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Number of Shelves":           "2",
		"Crisper Drawers":             "1",
		"Warranty":                    "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Net Volume":                  "235 liter",
		"Special Features":            "Frameless Door, R600a Refrigerant, Key & Lock, Interior Light",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Shelf Material":              "Tempered glass",
		"Voltage":                     "220~240",
		"Frequency (Hz)":              "50",
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
