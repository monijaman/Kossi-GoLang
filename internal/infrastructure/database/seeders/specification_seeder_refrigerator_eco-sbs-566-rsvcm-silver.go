package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecosbs566rsvcmsilver seeds specifications/options for product 'eco-sbs-566-rsvcm-silver'
type SpecificationSeederRefrigeratorecosbs566rsvcmsilver struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecosbs566rsvcmsilver creates a new seeder instance
func NewSpecificationSeederRefrigeratorecosbs566rsvcmsilver() *SpecificationSeederRefrigeratorecosbs566rsvcmsilver {
	return &SpecificationSeederRefrigeratorecosbs566rsvcmsilver{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-sbs-566-rsvcm-silver"},
	}
}

func (s *SpecificationSeederRefrigeratorecosbs566rsvcmsilver) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                      "ইকো+",
		"Eco+ SBS-566-RSVCM Silver": "ইকো+ এসবিএস-৫৬৬-আরএসভিসিএম সিলভার",
		"Side by side":              "সাইড বাই সাইড",
		"518":                       "৫১৮",
		"341":                       "৩৪১",
		"177":                       "১৭৭",
		"566":                       "৫৬৬",
		"350":                       "৩৫০",
		"216":                       "২১৬",
		"No-frost":                  "নো-ফ্রস্ট",
		"R600a":                     "আর৬০০এ",
		"Silver":                    "সিলভার",
		"1786*910*643 mm":           "১৭৮৬*৯১০*৬৪৩ মিমি",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty":                                                                   "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Total No Frost, Multi Air Flow, Fast Freeze, Surge Protection Technology, Wide Voltage Range, Door Self Closing and Alarm Function": "টোটাল নো ফ্রস্ট, মাল্টি এয়ার ফ্লো, ফাস্ট ফ্রিজ, সার্জ প্রোটেকশন টেকনোলজি, ওয়াইড ভোল্টেজ রেঞ্জ, ডোর সেল্ফ ক্লোজিং এবং অ্যালার্ম ফাংশন",
		"Electronic":    "ইলেকট্রনিক",
		"5":             "৫",
		"3":             "৩",
		"Yes":           "হ্যাঁ",
		"Soft LED":      "সফ্ট এলইডি",
		"Movable Twist": "মুভেবল টুইস্ট",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-sbs-566-rsvcm-silver'
func (s *SpecificationSeederRefrigeratorecosbs566rsvcmsilver) Seed(db *gorm.DB) error {
	productSlug := "eco-sbs-566-rsvcm-silver"
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
		"Brand":                 "ECO+",
		"Model Name":            "Eco+ SBS-566-RSVCM Silver",
		"Door Type":             "Side by side",
		"Capacity":              "518",
		"Refrigerator Capacity": "341",
		"Freezer Capacity":      "177",
		"Gross Volume":          "566",
		"Net Volume":            "518",
		"Defrost Type":          "No-frost",
		"Refrigerant":           "R600a",
		"Color":                 "Silver",
		"Dimensions":            "1786*910*643 mm",
		"Warranty":              "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Special Features":      "Total No Frost, Multi Air Flow, Fast Freeze, Surge Protection Technology, Wide Voltage Range, Door Self Closing and Alarm Function",
		"Temperature Control":   "Electronic",
		"Shelf Material":        "Transparent",
		"Number of Shelves":     "5",
		"Crisper Drawers":       "Yes",
		"Ice Maker":             "Movable Twist",
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
