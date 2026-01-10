package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2cGdehScInverter seeds specifications/options for product 'marcel-mfb-b2c-gdeh-sc-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB2cGdehScInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2cGdehScInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2cGdehScInverter() *SpecificationSeederRefrigeratorMarcelMfbB2cGdehScInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB2cGdehScInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2c-gdeh-sc-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2cGdehScInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfb-b2c-gdeh-sc-inverter": "মার্সেল-mfb-b2c-gdeh-sc-inverter",
		"MFB-B2C-GDEH-SC-INVERTER":        "MFB-B2C-GDEH-SC-INVERTER",

		// Common values used in this seeder
		"Direct Cool":                "ডিরেক্ট কুল",
		"177 Ltr.":                   "১৭৭ লিটার",
		"175 Ltr.":                   "১৭৫ লিটার",
		"N~ST":                       "N~ST",
		"220 ~ 240/ 50":              "২২০ ~ ২৪০/ ৫০",
		"V 01.01-88; V 01.02-88":     "V 01.01-88; V 01.02-88",
		"V 01.01-RSCR; V 01.02-RSCR": "V 01.01-RSCR; V 01.02-RSCR",
		"RSCR":                       "RSCR",
		"Mechanical":                 "মেকানিক্যাল",
		"Manual":                     "ম্যানুয়াল",
		"Recessed/ Grip":             "রিসেসড/গ্রিপ",
		"Yes":                        "হ্যাঁ",
		"No":                         "না",
		"Copper":                     "তামা",
		"Cyclopentene":               "সাইক্লোপেন্টেন",
		"R600a":                      "R600a",
		"50 ± 2 Kg":                  "৫০ ± ২ কেজি",
		"54 ± 2 Kg":                  "৫৪ ± ২ কেজি",
		"555 x 630 x 1410 mm":        "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm":        "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"105/ 105/ 52":               "১০৫/ ১০৫/ ৫২",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃":       "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"51.5 ± 2 Kg":               "৫১.৫ ± ২ কেজি",
		"RoHS Certified":            "RoHS Certified",
		"645":                       "৬৪৫",
		"56 ± 2 Kg":                 "৫৬ ± ২ কেজি",
		"1620":                      "১৬২০",
		"Yes/ 1":                    "Yes/ ১",
		"V 01.01-97.4 V 01.02-97.4": "V ০১.০১-৯৭.৪ V ০১.০২-৯৭.৪",
		"72/ 72/ 36":                "৭২/ ৭২/ ৩৬",
		"Wire/2":                    "Wire/২",
		"580":                       "৫৮০",
		"V 01.01, V 01.02-Low Voltage(140~260V) For V 01.01, V 01.02 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended": "V ০১.০১, V ০১.০২-Low Voltage(১৪০~২৬০V) For V ০১.০১, V ০১.০২ - Wide Voltage Range (১৪০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended",
		"GPPS/3":              "GPPS/৩",
		"223 Ltr.":            "২২৩ লিটার",
		"220~240V/50Hz":       "২২০~২৪০V/৫০Hz",
		"219 Ltr.":            "২১৯ লিটার",
		"555 x 630 x 1550 mm": "৫৫৫ x ৬৩০ x ১৫৫০ মিমি",
		"580 x 645 x 1620 mm": "৫৮০ x ৬৪৫ x ১৬২০ মিমি",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST": "N ~ ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSIR": "RSIR",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"555": "৫৫৫",
		"630": "৬৩০",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2c-gdeh-sc-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2cGdehScInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2c-gdeh-sc-inverter"
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
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 01.01-97.4V 01.02-97.4",
		"Compressor Type": "RSCR",
		"Cooling Efect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "630",
		"Door Basket": "GPPS/3",
		"Drawer": "No",
		"Egg Tray": "Yes/ 1",
		"Freezer Compartment Light": "Yes",
		"Gross Volume": "223 Ltr.",
		"Gross Weight": "56 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height/mm": "1550",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "72/ 72/ 36",
		"Lock": "Yes",
		"Net Volume": "219 Ltr.",
		"Net Weight": "51.5 ± 2 Kg",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rack Shelf (Material/ No.)": "Wire/2",
		"Rated Operating Voltage and Frequency": "220~240V/50Hz",
		"Recommended voltage stabilizer capacity": "V 01.01, V 01.02-Low Voltage(140~260V)For V 01.01, V 01.02 - Wide Voltage Range (140Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended",
		"Refrigerant": "R600a",
		"Refrigerator Compartment Light": "Yes",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Width/mm": "555",
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
