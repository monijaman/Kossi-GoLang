package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfdA6cGdelScInverter seeds specifications/options for product 'marcel-mfd-a6c-gdel-sc-inverter'
type SpecificationSeederRefrigeratorMarcelMfdA6cGdelScInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfdA6cGdelScInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfdA6cGdelScInverter() *SpecificationSeederRefrigeratorMarcelMfdA6cGdelScInverter {
	return &SpecificationSeederRefrigeratorMarcelMfdA6cGdelScInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfd-a6c-gdel-sc-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfdA6cGdelScInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfd-a6c-gdel-sc-inverter": "মার্সেল-mfd-a6c-gdel-sc-inverter",
		"MFD-A6C-GDEL-SC-INVERTER":        "MFD-A6C-GDEL-SC-INVERTER",
		// Add more translations as needed
		"176 Ltr.": "১৭৬ লিটার",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes":                "হ্যাঁ",
		"163 Ltr.":           "১৬৩ লিটার",
		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",
		"Recessed/ Grip":     "Recessed/ Grip",
		"Low Voltage (80~300V) Voltage Stabilizer is not required": "Low Voltage (৮০~৩০০V) Voltage Stabilizer is not required",
		"555":         "৫৫৫",
		"109/109 /49": "১০৯/১০৯ /৪৯",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Electronic":  "ইলেকট্রনিক",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"51.5 ± 2 Kg":    "৫১.৫ ± ২ কেজি",
		"Yes/1":          "Yes/১",
		"Copper":         "কপার",
		"Manual":         "Manual",
		"RoHS Certified": "RoHS Certified",
		"No":             "না",
		"R600A":          "R৬০০A",
		"595":            "৫৯৫",
		"Wire/2":         "Wire/২",
		"V 0802- 22~84":  "V ০৮০২- ২২~৮৪",
		"N ~ T":          "N ~ T",
		"46 ± 2 Kg":      "৪৬ ± ২ কেজি",
		"V 0802- BLDC":   "V ০৮০২- BLDC",
		"1630":           "১৬৩০",

		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Mechanical":    "মেকানিক্যাল",
		"N ~ ST":        "N ~ ST",
		"N~ST":          "N~ST",
		"PVC/1":         "PVC/১",
		"PVC/2":         "PVC/২",
		"PVC/3":         "PVC/৩",
		"R600a":         "R600a",
		"RSCR":          "RSCR",
		"RSIR":          "RSIR",
		"Wire/1":        "Wire/১",
		"Wire/3":        "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/2":         "হ্যাঁ/২",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfd-a6c-gdel-sc-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfdA6cGdelScInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfd-a6c-gdel-sc-inverter"
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
		"Type":                                  "Direct Cool",
		"Gross Volume":                          "176 Ltr.",
		"Net Volume":                            "163 Ltr.",
		"Net Weight":                            "46 ± 2 Kg",
		"Gross Weight":                          "51.5 ± 2 Kg",
		"Climate Type (SN, N, ST, T)":           "N ~ T",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Compressor Input Power (Watt)":         "V 0802- 22~84",
		"Compressor Type":                       "V 0802- BLDC",
		"Cooling Efect":                         "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Defrosting (Automatic/ Manual)":                "Manual",
		"Reversible Door":                               "No",
		"Handle (Recessed/ Grip)":                       "Recessed/ Grip",
		"Lock":                                          "Yes",
		"Refrigerant":                                   "R600A",
		"Capillary":                                     "Copper",
		"Thermostat":                                    "RoHS Certified",
		"Polyurethane foam blowing agent":               "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Operating Voltage":                             "Low Voltage (80~300V) Voltage Stabilizer is not required",
		"Shelf (Material/No.)":                          "Wire/2",
		"Door Basket":                                   "No",
		"Interior Lamp":                                 "No",
		"Vegetable Crisper":                             "Yes/1",
		"Vegetable Crisper Cover":                       "Yes",
		"Egg Tray or Pocket":                            "Yes",
		"Can Storage Dispenser":                         "No",
		"Deodorizer":                                    "No",
		"Shelf (Material/ No.)":                         "Wire/2",
		"Drawer":                                        "No",
		"Width/mm":                                      "555",
		"Depth/mm":                                      "595",
		"Height/mm":                                     "1630",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":            "109/109 /49",
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
