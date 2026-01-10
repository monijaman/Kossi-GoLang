package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMnhD3xHdxxXxInverter seeds specifications/options for product 'marcel-mnh-d3x-hdxx-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMnhD3xHdxxXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMnhD3xHdxxXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMnhD3xHdxxXxInverter() *SpecificationSeederRefrigeratorMarcelMnhD3xHdxxXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMnhD3xHdxxXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mnh-d3x-hdxx-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMnhD3xHdxxXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mnh-d3x-hdxx-xx-inverter": "মার্সেল-mnh-d3x-hdxx-xx-inverter",
		"MNH-D3X-HDXX-XX-INVERTER":        "MNH-D3X-HDXX-XX-INVERTER",
		"Non-Frost":                       "নন-ফ্রস্ট",
		"430 Ltr.":                        "৪৩০ লিটার",
		"370 Ltr.":                        "৩৭০ লিটার",
		"Five Star (*****)":               "ফাইভ স্টার (*****)",
		"360 kWh/year":                    "৩৬০ কিলোওয়াট-ঘন্টা/বছর",
		"T":                               "T",
		"220-240 V/ 50 Hz":                "২২০-২৪০ ভি / ৫০ হার্জ",
		"76.5 \u00b1 2 Kg":                "৭৬.৫ ± ২ কেজি",
		"86 \u00b1 2 Kg":                  "৮৬ ± ২ কেজি",
		"BLDC Inverter":                   "BLDC ইনভার্টার",
		"Mechanical":                      "যান্ত্রিক",
		"Automatic":                       "স্বয়ংক্রিয়",
		"R600a":                           "R600a",
		"100 % Copper":                    "১০০% কপার",
		"Copper":                          "কপার",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "সাইক্লোপেন্টিন (পরিবেশবান্ধব, CFC/HCFC ফ্রি)",
		"1000VA or More":      "১০০০ভিএ বা বেশি",
		"Glass/4":             "গ্লাস/৪",
		"GPPS/7":              "GPPS/৭",
		"Yes/1":               "হ্যাঁ/১",
		"Yes/2":               "হ্যাঁ/২",
		"705 x 690 x 1845 mm": "৭০৫ x ৬৯০ x ১৮৪৫ মিমি",
		"760 x 775 x 1900 mm": "৭৬০ x ৭৭৫ x ১৯০০ মিমি",
		"48/ 48/ 24":          "৪৮/ ৪৮/ ২৪",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[Eco-friendly (১০০% CFC &HCFC Free) Green Technology]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[Eco-friendly (১০০% CFC &HCFC Free) Green Technology]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "Freezer Cabinet Less than -১৮০CRefrigerator Cabinet ০০Cto +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃Refrigerator Cabinet ০℃ to +৫℃",
		"Manual": "ম্যানুয়াল",
		"N ~ ST": "N ~ ST",
		"No": "না",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes": "হ্যাঁ",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
	
		"4": "৪",
		"Glass": "গ্লাস",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mnh-d3x-hdxx-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMnhD3xHdxxXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mnh-d3x-hdxx-xx-inverter"
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
		"Brand":                             310,
		"Model Name":                        316,
		"Door Type":                         142,
		"Capacity":                          138,
		"Refrigerator Capacity":             156,
		"Freezer Capacity":                  146,
		"Energy Efficiency Rating":          143,
		"Energy Star Rating":                144,
		"Annual Energy Consumption":         137,
		"Dimensions":                        25,
		"Weight":                            80,
		"Color":                             311,
		"Compressor Type":                   139,
		"Cooling Technology":                698,
		"Defrost Type":                      141,
		"Temperature Control":               158,
		"Shelf Material":                    699,
		"Number of Shelves":                 154,
		"Door Bins":                         700,
		"Crisper Drawers":                   701,
		"Ice Maker":                         702,
		"Water Dispenser":                   703,
		"Noise Level":                       150,
		"Voltage":                           160,
		"Frequency (Hz)":                    704,
		"App Control":                       705,
		"Voice Assistant Support":           385,
		"Warranty":                          323,
		"Compressor Warranty (Years)":       707,
		"Refrigerant":                       708,
		"Gross Volume":                      709,
		"Net Volume":                        710,
		"Special Features":                  69,
		"Capillary":                         711,
		"Climate Type (SN, N, ST, T)":       712,
		"Compressor Input Power (Watt)":     713,
		"Cooling Effect":                    714,
		"Defrosting (Automatic/ Manual)":    715,
		"Deodorizer":                        716,
		"Depth (mm)":                        717,
		"Door Basket":                       718,
		"Egg Tray or Pocket":                719,
		"Handle (Recessed/ Grip)":           720,
		"Interior Lamp":                     721,
		"Polyurethane Foam Blowing Agent":   722,
		"Reversible Door":                   723,
		"Thermostat":                        724,
		"Vegetable Crisper":                 725,
		"Vegetable Crisper Cover":           726,
		"Loading Capacity (40HQ/40Ft/20Ft)": 727,
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}

	specs := map[string]string{
		"Brand":                           "Marcel",
		"Model Name":                      "MNH-D3X-HDXX-XX-INVERTER",
		"Cooling Technology":              "Non-Frost",
		"Gross Volume":                    "430 Ltr.",
		"Net Volume":                      "370 Ltr.",
		"Energy Star Rating":              "Five Star (*****)",
		"Annual Energy Consumption":       "360 kWh/year",
		"Temperature Control":             "Mechanical",
		"Defrost Type":                    "Automatic",
		"Voltage":                         "220-240 V/ 50 Hz",
		"Weight":                          "76.5 ± 2 Kg / 86 ± 2 Kg",
		"Compressor Type":                 "BLDC Inverter",
		"Refrigerant":                     "R600a",
		"Condenser":                       "100 % Copper",
		"Capillary":                       "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "1000VA or More",
		"Shelf Material":                   "Glass",
		"Number of Shelves":                "4",
		"Door Bins":                        "GPPS/7",
		"Crisper Drawers":                  "Yes/1",
		"Egg Tray or Pocket":               "Yes/2",
		"Freezer Compartment - Rack Shelf": "Wire/3",
		"Freezer Compartment - Ice Box":    "Yes/1",
		"Dimensions":                       "705 x 690 x 1845 mm",
		"Packing Dimensions":               "760 x 775 x 1900 mm",
		"Loading Capacity":                 "48/ 48/ 24",
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
