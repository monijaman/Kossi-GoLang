package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5xGdelXx seeds specifications/options for product 'marcel-mfb-b5x-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfbB5xGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5xGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5xGdelXx() *SpecificationSeederRefrigeratorMarcelMfbB5xGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB5xGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b5x-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1520":               "১৫২০",
		"220 ~ 240/ 50/ 109": "২২০ ~ ২৪০/ ৫০/ ১০৯",
		"265 Ltr.":           "২৬৫ লিটার",
		"282 Ltr.":           "২৮২ লিটার",
		"54 ± 2 Kg":          "৫৪ ± ২ কেজি",
		"54/61 ± 2":          "৫৪/৬১ ± ২",
		"594":                "৫৯৪",
		"61 ± 2 Kg":          "৬১ ± ২ কেজি",
		"708":                "৭০৮",
		"80/ 80/ 38":         "৮০/ ৮০/ ৩৮",
		"Copper":             "কপার",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Direct Cool, Eco Friendly, Energy Saving, Large Vegetable Crisper, Stabilizer Free Operation, Anti Bacterial Gasket, Door Lock, LED Light, Adjustable Shelves, Egg Tray, Ice Tray": "ডাইরেক্ট কুল, ইকো ফ্রেন্ডলি, এনার্জি সেভিং, লার্জ ভেজিটেবল ক্রিস্পার, স্ট্যাবিলাইজার ফ্রি অপারেশন, অ্যান্টি ব্যাকটেরিয়াল গ্যাসকেট, ডোর লক, LED লাইট, অ্যাডজাস্টেবল শেলভস, এগ ট্রে, আইস ট্রে",
		"Freezer Cabinet Less than -180C Refrigerator Cabinet 00C to +50C": "Freezer Cabinet Less than -১৮০C Refrigerator Cabinet ০০C to +৫০C",
		"Marcel":                    "মার্সেল",
		"Mechanical":                "মেকানিক্যাল",
		"MFB-B5X-GDEL-XX":           "MFB-B5X-GDEL-XX",
		"Manual":                    "ম্যানুয়াল",
		"N ~ ST":                    "N ~ ST",
		"No":                        "না",
		"PVC/3":                     "PVC/3",
		"Recressed/ Grip/ Built-in": "রিসেসড/ গ্রিপ/ বিল্ট-ইন",
		"RoHS Certified":            "RoHS Certified",
		"R600a":                     "R600a",
		"RSCR":                      "RSCR",
		"Wire/2":                    "ওয়্যার/2",
		"Yes":                       "হ্যাঁ",
		"Yes (ABS/ PS)":             "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)":             "হ্যাঁ (প্লাস্টিক)",
		// Add more translations as needed
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"1675 mm": "১৬৭৫ mm",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"274 Ltr.": "২৭৪ লিটার",
		"54/59 ± 2": "৫৪/৫৯ ± ২",
		"555 mm": "৫৫৫ মিমি",
		"630 mm": "৬৩০ মিমি",
		"97/ 72/ 36": "৯৭/ ৭২/ ৩৬",
		"GPPS/3": "জিপিপিএস/৩",}
}

// Seed inserts specification records for the product identified by slug 'mfb-b5x-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b5x-gdel-xx"
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
		"Brand":               "Marcel",
		"Model Name":          "MFB-B5X-GDEL-XX",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 01.01-97.4V 02.01-97.4V 03.01-97.4V 03.02-97.4",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "630 mm",
		"Door Basket": "GPPS/3",
		"Drawer": "No",
		"Egg Tray": "Yes",
		"Gross Volume": "250 Ltr.",
		"Gross Weight": "59± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height (mm)": "1675 mm",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "97/ 72/ 36",
		"Lock": "Yes",
		"Net Volume": "274 Ltr.",
		"Weight": "54± 2 Kg",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 01.01, V 02.01-Low Voltage(140~260V)For V01.01, V 02.01-Wide Voltage Range (140Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is   recommended.V 03.01, V 03.02-Low Voltage(150~260V)For V 03.01, V 03.02 - Wide Voltage Range (150Vac -   260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is   recommended.",
		"Refrigerant": "R600a",
		"Shelf Material": "Wire/2",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Weight/Kg - Net/Packing": "54/59 ± 2",
		"Width": "555 mm",
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
