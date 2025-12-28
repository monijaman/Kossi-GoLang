package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecossrt37k5532s8d3silver seeds specifications/options for product 'eco-ss-rt37k5532s8-d3-silver'
type SpecificationSeederRefrigeratorecossrt37k5532s8d3silver struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecossrt37k5532s8d3silver creates a new seeder instance
func NewSpecificationSeederRefrigeratorecossrt37k5532s8d3silver() *SpecificationSeederRefrigeratorecossrt37k5532s8d3silver {
	return &SpecificationSeederRefrigeratorecossrt37k5532s8d3silver{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-ss-rt37k5532s8-d3-silver"},
	}
}

func (s *SpecificationSeederRefrigeratorecossrt37k5532s8d3silver) getBanglaTranslations() map[string]string {
	return map[string]string{
		"SAMSUNG":                     "স্যামসাং",
		"SS RT37K5532S8/D3 Silver":    "এসএস আরটি৩৭কে৫৫৩২এস৮/ডি৩ সিলভার",
		"VCM":                         "ভিসিএম",
		"345":                         "৩৪৫",
		"No-frost":                    "নো-ফ্রস্ট",
		"Digital Inverter Compressor": "ডিজিটাল ইনভার্টার কম্প্রেসর",
		"R600a":                       "আর৬০০এ",
		"Silver":                      "সিলভার",
		"1715*600*605":                "১৭১৫*৬০০*৬০৫",
		"20 Years Compressor Warranty, 1 Year Parts and Service Warranty": "২০ বছর কম্প্রেসর ওয়ারেন্টি, ১ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Twin Cooling Plus, Multi flow, Deodorizer, Frost Free":           "টুইন কুলিং প্লাস, মাল্টি ফ্লো, ডিওডোরাইজার, ফ্রস্ট ফ্রি",
		"Twin Cooling Plus": "টুইন কুলিং প্লাস",
		"External":          "এক্সটার্নাল",
		"3":                 "৩",
		"4":                 "৪",
		"1":                 "১",
		"2":                 "২",
		"Yes":               "হ্যাঁ",
		"59":                "৫৯",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-ss-rt37k5532s8-d3-silver'
func (s *SpecificationSeederRefrigeratorecossrt37k5532s8d3silver) Seed(db *gorm.DB) error {
	productSlug := "eco-ss-rt37k5532s8-d3-silver"
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
		"Brand":               "SAMSUNG",
		"Model Name":          "SS RT37K5532S8/D3 Silver",
		"Door Type":           "VCM",
		"Capacity":            "345",
		"Gross Volume":        "345",
		"Defrost Type":        "No-frost",
		"Compressor Type":     "Digital Inverter Compressor",
		"Refrigerant":         "R600a",
		"Color":               "Silver",
		"Dimensions":          "1715*600*605",
		"Weight":              "59",
		"Warranty":            "20 Years Compressor Warranty, 1 Year Parts and Service Warranty",
		"Special Features":    "Twin Cooling Plus, Multi flow, Deodorizer, Frost Free",
		"Cooling Technology":  "Twin Cooling Plus",
		"Temperature Control": "External",
		"Number of Shelves":   "3",
		"Door Bins":           "4",
		"Crisper Drawers":     "1",
		"Ice Maker":           "Yes",
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
