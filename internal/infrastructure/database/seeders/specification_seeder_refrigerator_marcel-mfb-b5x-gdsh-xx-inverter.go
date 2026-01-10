package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5xGdshXxInverter seeds specifications/options for product 'marcel-mfb-b5x-gdsh-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB5xGdshXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5xGdshXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5xGdshXxInverter() *SpecificationSeederRefrigeratorMarcelMfbB5xGdshXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB5xGdshXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b5x-gdsh-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdshXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1675 mm":       "১৬৭৫ mm",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"244 Ltr.":      "২৪৪ লিটার",
		"250 Ltr.":      "২৫০ লিটার",
		"50.5/55.5 ± 2": "৫০.৫/৫৫.৫ ± ২",
		"54.5 ± 2 Kg":   "৫৪.৫ ± ২ কেজি",
		"555 mm":        "৫৫৫ মিমি",
		"580 mm":        "৫৮০ মিমি",
		"59.5 ± 2 Kg":   "৫৯.৫ ± ২ কেজি",
		"630 mm":        "৬৩০ মিমি",
		"645 mm":        "৬৪৫ মিমি",
		"97/ 72/ 36":    "৯৭/ ৭২/ ৩৬",
		"BLDC Inverter": "BLDC Inverter",
		"Copper":        "কপার",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"GPPS/3":          "GPPS/3",
		"Mechanical":      "মেকানিক্যাল",
		"Manual":          "ম্যানুয়াল",
		"No":              "না",
		"Recessed/Grip":   "রিসেসড/গ্রিপ",
		"RoHS Certified":  "RoHS Certified",
		"R600a":           "R600a",
		"T":               "T",
		"V 05.01- 57~125": "V 05.01- 57~125",
		"Wire/2":          "ওয়্যার/2",
		"Yes":             "হ্যাঁ",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b5x-gdsh-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdshXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b5x-gdsh-xx-inverter"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
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
		"Capillary":                               "Copper",
		"Climatic Type (SN, N, ST, T)":            "T",
		"Compressor Input Power (Watt)":           "V 05.01- 57~125",
		"Compressor Type":                         "BLDC Inverter",
		"Cooling Effect":                          "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)":          "Manual",
		"Door Basket":                             "GPPS/3",
		"Drawer":                                  "No",
		"Egg Tray":                                "Yes",
		"Gross Volume":                            "250 Ltr.",
		"Gross Weight":                            "59.5 ± 2 Kg",
		"Handle (Recessed/ Grip)":                 "Recessed/Grip",
		"Interior Lamp":                           "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":      "97/ 72/ 36",
		"Lock":                                    "Yes",
		"Net Depth/mm":                            "630 mm",
		"Net Height/mm":                           "1675 mm",
		"Net Volume":                              "244 Ltr.",
		"Net Weight":                              "54.5 ± 2 Kg",
		"Net Width/mm":                            "555 mm",
		"Packaging Depth/mm":                      "645 mm",
		"Packaging Height/mm":                     "1725 mm",
		"Packaging Width/mm":                      "580 mm",
		"Polyurethane foam blowing agent":         "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Rated Voltage/ Hz":                       "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 05.01-Low Voltage(90~300V) For V05.01-Wide Voltage Range (90Vac - 300Vac). Voltage stabilizer is not required.",
		"Refrigerant":                             "R600a",
		"Shelf (Material/ No.)":                   "Wire/2",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Thermostat":              "RoHS Certified",
		"Type":                    "Direct Cool",
		"Vegetable Crisper":       "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Weight/Kg - Net/Packing": "50.5/55.5 ± 2",
	}

	banglaTranslations := s.getBanglaTranslations()
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
				if err := db.Create(sModel).Error; err != nil {
					return err
				}
				// Ensure the ID is set after creation
				if sModel.ID == 0 {
					if err := db.Where("product_id = ? AND specification_key_id = ? AND value = ?", productID, specKeyID, value).First(sModel).Error; err != nil {
						return err
					}
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
				}
			}
		}
	}

	return nil
}
