package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5xGdelSc seeds specifications/options for product 'marcel-mfb-b5x-gdel-sc'
type SpecificationSeederRefrigeratorMarcelMfbB5xGdelSc struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5xGdelSc creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5xGdelSc() *SpecificationSeederRefrigeratorMarcelMfbB5xGdelSc {
	return &SpecificationSeederRefrigeratorMarcelMfbB5xGdelSc{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b5x-gdel-sc"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdelSc) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1675":          "১৬৭৫",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"244 Ltr.":      "২৪৪ লিটার",
		"250 Ltr.":      "২৫০ লিটার",
		"50.5/55.5 ± 2": "৫০.৫/৫৫.৫ ± ২",
		"54.5 ± 2 Kg":   "৫৪.৫ ± ২ কেজি",
		"555":           "৫৫৫",
		"59.5 ± 2 Kg":   "৫৯.৫ ± ২ কেজি",
		"630":           "৬৩০",
		"97/ 72/ 36":    "৯৭/ ৭২/ ৩৬",
		"Copper":        "কপার",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Direct Cool, Eco Friendly, Energy Saving, Large Vegetable Crisper, Stabilizer Free Operation, Anti Bacterial Gasket, Door Lock, LED Light, Adjustable Shelves, Egg Tray, Ice Tray": "ডাইরেক্ট কুল, ইকো ফ্রেন্ডলি, এনার্জি সেভিং, লার্জ ভেজিটেবল ক্রিস্পার, স্ট্যাবিলাইজার ফ্রি অপারেশন, অ্যান্টি ব্যাকটেরিয়াল গ্যাসকেট, ডোর লক, LED লাইট, অ্যাডজাস্টেবল শেলভস, এগ ট্রে, আইস ট্রে",
		"Electronic": "ইলেকট্রনিক",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"GPPS/3":          "GPPS/3",
		"Manual":          "ম্যানুয়াল",
		"Marcel":          "মার্সেল",
		"Mechanical":      "মেকানিক্যাল",
		"MFB-B5X-GDEL-SC": "MFB-B5X-GDEL-SC",
		"N~ST":            "N~ST",
		"No":              "না",
		"Recessed/Grip":   "রিসেসড/গ্রিপ",
		"RoHS Certified":  "RoHS Certified",
		"R600a":           "R600a",
		"RSCR":            "RSCR",
		"V 01.01-97.4 V 02.01-97.4 V 03.01-97.4 V 03.02-97.4": "V ০১.০১-৯৭.৪ V ০২.০১-৯৭.৪ V ০৩.০১-৯৭.৪ V ০৩.০২-৯৭.৪",
		"V 01.01, V 02.01-Low Voltage(140~260V) For V01.01, V 02.01-Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 03.01, V 03.02-Low Voltage(150~260V) For V 03.01, V 03.02 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.": "V ০১.০১, V ০২.০১-Low Voltage(১৪০~২৬০V) For V০১.০১, V ০২.০১-Wide Voltage Range (১৪০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is recommended. V ০৩.০১, V ০৩.০২-Low Voltage(১৫০~২৬০V) For V ০৩.০১, V ০৩.০২ - Wide Voltage Range (১৫০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is recommended.",
		"Wire/2": "ওয়্যার/2",
		"Yes":    "হ্যাঁ",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfb-b5x-gdel-sc'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdelSc) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b5x-gdel-sc"
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
		"Brand":                              "Marcel",
		"Capillary":                          "Copper",
		"Climatic Type (SN, N, ST, T)":       "N~ST",
		"Compressor Input Power (Watt)":      "V 01.01-97.4 V 02.01-97.4 V 03.01-97.4 V 03.02-97.4",
		"Compressor Type":                    "RSCR",
		"Cooling Effect":                     "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual):":    "Manual",
		"Depth/mm":                           "630",
		"Door Basket":                        "GPPS/3",
		"Drawer":                             "No",
		"Egg Tray":                           "Yes",
		"Gross Volume":                       "250 Ltr.",
		"Gross Weight":                       "59.5 ± 2 Kg",
		"Handle (Recessed/ Grip)":            "Recessed/Grip",
		"Height/mm":                          "1675",
		"Interior Lamp":                      "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "97/ 72/ 36",
		"Lock":                               "Yes",
		"Model Name":                         "MFB-B5X-GDEL-SC",
		"Net Volume":                         "244 Ltr.",
		"Net Weight":                         "54.5 ± 2 Kg",
		"Polyurethane foam blowing agent":    "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Rated Voltage/ Hz":                  "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 01.01, V 02.01-Low Voltage(140~260V) For V01.01, V 02.01-Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 03.01, V 03.02-Low Voltage(150~260V) For V 03.01, V 03.02 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.",
		"Refrigerant":           "R600a",
		"Shelf (Material/ No.)": "Wire/2",
		"Shelf (Material/No.)":  "Wire/2",
		"Special Features":      "Direct Cool, Eco Friendly, Energy Saving, Large Vegetable Crisper, Stabilizer Free Operation, Anti Bacterial Gasket, Door Lock, LED Light, Adjustable Shelves, Egg Tray, Ice Tray",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Thermostat":              "RoHS Certified",
		"Type":                    "Direct Cool",
		"Vegetable Crisper":       "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Weight/Kg - Net/Packing": "50.5/55.5 ± 2",
		"Width/mm":                "555",
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
