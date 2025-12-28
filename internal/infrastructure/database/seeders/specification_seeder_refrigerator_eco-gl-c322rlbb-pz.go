package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecoglc322rlbbpz seeds specifications/options for product 'eco-gl-c322rlbb-pz'
type SpecificationSeederRefrigeratorecoglc322rlbbpz struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecoglc322rlbbpz creates a new seeder instance
func NewSpecificationSeederRefrigeratorecoglc322rlbbpz() *SpecificationSeederRefrigeratorecoglc322rlbbpz {
	return &SpecificationSeederRefrigeratorecoglc322rlbbpz{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-gl-c322rlbb-pz"},
	}
}

func (s *SpecificationSeederRefrigeratorecoglc322rlbbpz) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LG":                       "এলজি",
		"LG OMEGA3 GL-C322RLBB PZ": "এলজি ওমেগা৩ জিএল-সি৩২২আরএলবিবি পিজেড",
		"Top Mount":                "টপ মাউন্ট",
		"281":                      "২৮১",
		"223":                      "২২৩",
		"58":                       "৫৮",
		"308":                      "৩০৮",
		"No-frost":                 "নো-ফ্রস্ট",
		"Smart Inverter":           "স্মার্ট ইনভার্টার",
		"R-600A":                   "আর-৬০০এ",
		"Shiny steel":              "শাইনি স্টিল",
		"1675*585*703 mm":          "১৬৭৫*৫৮৫*৭০৩ মিমি",
		"51 kg":                    "৫১ কেজি",
		"100-290V":                 "১০০-২৯০ভি",
		"50":                       "৫০",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty":                              "১০ বছর কম্প্রেসার ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Door Cooling+, Multi Air Flow, Moist Balance Crisper, Humidity Controller, Wide Voltage Range": "ডোর কুলিং+, মাল্টি এয়ার ফ্লো, ময়েস্ট ব্যালেন্স ক্রিস্পার, হিউমিডিটি কন্ট্রোলার, ওয়াইড ভোল্টেজ রেঞ্জ",
		"Door Cooling+":   "ডোর কুলিং+",
		"Knob I Micom":    "নব আই মাইকম",
		"Toughened Glass": "টাফেন্ড গ্লাস",
		"Yes":             "হ্যাঁ",
		"Transparent":     "ট্রান্সপারেন্ট",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-gl-c322rlbb-pz'
func (s *SpecificationSeederRefrigeratorecoglc322rlbbpz) Seed(db *gorm.DB) error {
	productSlug := "eco-gl-c322rlbb-pz"
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
		"Brand":                 "LG",
		"Model Name":            "LG OMEGA3 GL-C322RLBB PZ",
		"Door Type":             "Top Mount",
		"Capacity":              "281",
		"Refrigerator Capacity": "223",
		"Freezer Capacity":      "58",
		"Gross Volume":          "308",
		"Net Volume":            "281",
		"Defrost Type":          "No-frost",
		"Compressor Type":       "Smart Inverter",
		"Refrigerant":           "R-600A",
		"Color":                 "Shiny steel",
		"Dimensions":            "1675*585*703 mm",
		"Weight":                "51 kg",
		"Voltage":               "100-290V",
		"Frequency (Hz)":        "50",
		"Warranty":              "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Special Features":      "Door Cooling+, Multi Air Flow, Moist Balance Crisper, Humidity Controller, Wide Voltage Range",
		"Cooling Technology":    "Door Cooling+",
		"Temperature Control":   "Knob I Micom",
		"Shelf Material":        "Toughened Glass",
		"Crisper Drawers":       "Yes",
		"Ice Maker":             "Transparent",
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
