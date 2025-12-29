package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecohrf680bgblack seeds specifications/options for product 'eco-hrf-680bg-black'
type SpecificationSeederRefrigeratorecohrf680bgblack struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecohrf680bgblack creates a new seeder instance
func NewSpecificationSeederRefrigeratorecohrf680bgblack() *SpecificationSeederRefrigeratorecohrf680bgblack {
	return &SpecificationSeederRefrigeratorecohrf680bgblack{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-hrf-680bg-black"},
	}
}

func (s *SpecificationSeederRefrigeratorecohrf680bgblack) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HAIER":                 "হাইয়ার",
		"Haier HRF-680BG Black": "হাইয়ার এইচআরএফ-৬৮০বিজি ব্ল্যাক",
		"Side by Side":          "সাইড বাই সাইড",
		"630":                   "৬৩০",
		"392":                   "৩৯২",
		"238":                   "২৩৮",
		"No-frost":              "নো-ফ্রস্ট",
		"R600a":                 "আর৬০০এ",
		"Black":                 "ব্ল্যাক",
		"1775*905*697 mm":       "১৭৭৫*৯০৫*৬৯৭ মিমি",
		"103 kg":                "১০৩ কেজি",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty": "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Deo Fresh, Jumbo Ice Maker, Digital LED Door Display":             "ডিও ফ্রেশ, জাম্বো আইস মেকার, ডিজিটাল এলইডি ডোর ডিসপ্লে",
		"Digital LED -Door Display":                                        "ডিজিটাল এলইডি -ডোর ডিসপ্লে",
		"Recess Handle":                                                    "রিসেস হ্যান্ডেল",
		"Yes":                                                              "হ্যাঁ",
		"5":                                                                "৫",
		"1":                                                                "১",
		"3":                                                                "৩",
		"2":                                                                "২",
		"LED":                                                              "এলইডি",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-hrf-680bg-black'
func (s *SpecificationSeederRefrigeratorecohrf680bgblack) Seed(db *gorm.DB) error {
	productSlug := "eco-hrf-680bg-black"
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
		"Brand":                 "HAIER",
		"Model Name":            "Haier HRF-680BG Black",
		"Door Type":             "Side by Side",
		"Capacity":              "630",
		"Refrigerator Capacity": "392",
		"Freezer Capacity":      "238",
		"Gross Volume":          "630",
		"Net Volume":            "630",
		"Defrost Type":          "No-frost",
		"Color":                       "Black",
		"Dimensions":            "1775*905*697 mm",
		"Weight":                "103 kg",
		"Warranty":              "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Special Features":      "Deo Fresh, Jumbo Ice Maker, Digital LED Door Display",
		"Refrigerant":           "R600a",
		"Temperature Control":   "Digital LED -Door Display",
		"Number of Shelves":     "3",
		"Crisper Drawers":       "Yes",
		"Ice Maker":             "Yes",
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
