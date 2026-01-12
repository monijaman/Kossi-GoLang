package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMnhD3xGdelXxInverter seeds specifications/options for product 'marcel-mnh-d3x-gdel-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMnhD3xGdelXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMnhD3xGdelXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMnhD3xGdelXxInverter() *SpecificationSeederRefrigeratorMarcelMnhD3xGdelXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMnhD3xGdelXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mnh-d3x-gdel-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMnhD3xGdelXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mnh-d3x-gdel-xx-inverter": "মার্সেল-mnh-d3x-gdel-xx-inverter",
		"MNH-D3X-GDEL-XX-INVERTER":        "MNH-D3X-GDEL-XX-INVERTER",

		// specs values -> Bangla
		"430 Ltr.":                          "৪৩০ লিটার",
		"370 Ltr.":                          "৩৭০ লিটার",
		"76.5 ± 2 Kg":                       "৭৬.৫ ± ২ কেজি",
		"86 ± 2 Kg":                         "৮৬ ± ২ কেজি",
		"BLDC Inverter":                     "বিএলডিসি ইনভার্টার",
		"No-Frost":                          "নো-ফ্রস্ট",
		"R600a":                             "R600a",
		"705 x 690 x 1845 mm": "৭০৫ x ৬৯০ x ১৮৪৫ মিমি",
		"760 x 775 x 1900 mm": "৭৬০ x ৭৭৫ x ১৯০০ মিমি",
		"48/ 48/ 24":                        "৪৮/ ৪৮/ ২৪",
		"Five Star (*****)":                 "পাঁচ তারা (*****)",
		"360 kWh/year as per BDS 1850:2012": "৩৬০ kWh/বছর (BDS 1850:2012 অনুযায়ী)",
		"T":                                 "T",
		"220-240 V/ 50 Hz":                  "২২০-২৪০ V/ ৫০ Hz",
		"100 % Copper":                      "১০০% কপার",
		"Copper":                            "কপার",
		"RoHS Certified":                    "RoHS সার্টিফাইড",
		"Cyclopentene":                      "সাইক্লোপেন্টেন",
		"Glass":                             "গ্লাস",
		"GPPS":                              "GPPS",
		"Automatic":                         "স্বয়ংক্রিয়",
		"Yes":                               "হ্যাঁ",
		"Yes/1":                             "হ্যাঁ/১",
		"Yes/2":                             "হ্যাঁ/২",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Manual": "ম্যানুয়াল",
		"Mechanical": "মেকানিক্যাল",
		"N ~ ST": "N ~ ST",
		"No": "না",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
	
		"1": "১",
		"4": "৪",
		"GPPS/7": "GPPS/৭",
		"Yes/ 2": "Yes/ ২",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mnh-d3x-gdel-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMnhD3xGdelXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mnh-d3x-gdel-xx-inverter"
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
		"Type":                              456,
		"Lock":                              361,
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}

	specs := map[string]string{
		"Brand":                           "Marcel",
		"Model Name":                      "MNH-D3X-GDEL-XX-INVERTER",
		"Cooling Technology":              "No-Frost",
		"Gross Volume":                    "430 Ltr.",
		"Net Volume":                      "370 Ltr.",
		"Energy Star Rating":              "Five Star (*****)",
		"Annual Energy Consumption":       "360 kWh/year as per BDS 1850:2012",
		"Climate Type (SN, N, ST, T)":                   "T",
		"Voltage":                         "220-240 V/ 50 Hz",
		"Weight":                          "76.5 ± 2 Kg",
		"Compressor Type":                 "BLDC Inverter",
		"Temperature Control":             "Mechanical",
		"Defrost Type":                    "Automatic",
		"Refrigerant":                     "R600a",
		"Capillary":                       "Copper",
		"Thermostat":                      "RoHS Certified",
		"Polyurethane Foam Blowing Agent": "Cyclopentene",

		// Refrigerator compartment
		"Shelf Material":     "Glass",
		"Number of Shelves":  "4",
		"Door Bins":          "GPPS/7",
		"Interior Lamp":      "Yes",
		"Crisper Drawers":    "1",
		"Egg Tray or Pocket": "Yes/ 2",

		// Freezer compartment
		"Ice Maker":                "Yes/2",
		"Ice Box":                  "Yes/1",
		"Shelf Material (Freezer)": "GPPS",
		"Interior LED Lamp":        "Yes",
		"Cooling Effect":           "Freezer Cabinet Less than -18°C; Refrigerator Cabinet 0°C to +50°C",

		"Dimensions":         "705 x 690 x 1845 mm",
		"Packing Dimensions": "760 x 775 x 1900 mm",
		"Loading Capacity (40HQ/40Ft/20Ft)":   "48/ 48/ 24",
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
