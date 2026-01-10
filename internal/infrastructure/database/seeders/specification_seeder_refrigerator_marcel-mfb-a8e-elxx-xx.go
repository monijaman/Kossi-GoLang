package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbA8eElxxXx seeds specifications/options for product 'marcel-mfb-a8e-elxx-xx'
type SpecificationSeederRefrigeratorMarcelMfbA8eElxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbA8eElxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbA8eElxxXx() *SpecificationSeederRefrigeratorMarcelMfbA8eElxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfbA8eElxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-a8e-elxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbA8eElxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-a8e-elxx-xx": "মার্সেল-mfb-a8e-elxx-xx",
		"MFB-A8E-ELXX-XX":        "MFB-A8E-ELXX-XX",

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
		"555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"105/ 105/ 52":               "১০৫/ ১০৫/ ৫২",
		"207 Ltr.":                   "২০৭ লিটার",
		"193 Ltr.":                   "১৯৩ লিটার",
		"53 ± 2 Kg":                  "৫৩ ± ২ কেজি",
		"N/A":                        "N/A",
		"Wire":                       "ওয়্যার",
		"2":                          "২",
		"GPPS/4":                     "জিপিপিএস/৪",
		"555 x 630 x 1485 mm": "৫৫৫ x ৬৩০ x ১৪৮৫ মিমি",
		"Lock: Yes, Interior Lamp: Yes, Handle: Recessed/Grip, Capillary: Copper, Polyurethane foam blowing agent Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology], Refrigerator Compartment Shelves: Wire/2, Refrigerator Door Baskets: GPPS/4, Vegetable Crisper: Yes, Egg Tray or Pocket: Yes, Freezer Compartment Shelves: Wire/2, Drawer: HIPS/4, Loading quantity: 102/102/50 (40HQ/40Ft/20Ft), Climatic Type: N~ST, Cooling Effect: Freezer Cabinet Less than -18°C, Refrigerator Cabinet 0°C to +5°C, Energy Rating: N/A, Recommended voltage stabilizer capacity: V 05.01,V 06.01, V 06.02,V 07.01-1000VA V 07.02-Low Voltage(155~260V) For V 07.02 - Wide Voltage Range (155Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.": "লক: হ্যাঁ, ইন্টেরিয়র ল্যাম্প: হ্যাঁ, হ্যান্ডেল: রিসেসড/গ্রিপ, ক্যাপিলারি: কপার, পলিউরেথেন ফোম ব্লোয়িং এজেন্ট সাইক্লোপেন্টেন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি], রেফ্রিজারেটর কম্পার্টমেন্ট শেল্ফস: ওয়্যার/২, রেফ্রিজারেটর দরজা বাস্কেটস: জিপিপিএস/৪, ভেজিটেবল ক্রিসপার: হ্যাঁ, এগ ট্রে অর পকেট: হ্যাঁ, ফ্রিজার কম্পার্টমেন্ট শেল্ফস: ওয়্যার/২, ড্রয়ার: এইচআইপিএস/৪, লোডিং কোয়ান্টিটি: ১০২/১০২/৫০ (৪০এইচকিউ/৪০এফটি/২০এফটি), ক্লাইমেটিক টাইপ: এন~এসটি, কুলিং ইফেক্ট: ফ্রিজার ক্যাবিনেট লেস থ্যান -১৮°C, রেফ্রিজারেটর ক্যাবিনেট ০°C টু +৫°C, এনার্জি রেটিং: N/A, রেকমেন্ডেড ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: V ০৫.০১,V ০৬.০১, V ০৬.০২,V ০৭.০১-১০০০VA V ০৭.০২-Low Voltage(১৫৫~২৬০V) For V ০৭.০২ - Wide Voltage Range (১৫৫Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is recommended.",
	
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
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"102/ 102/ 50": "১০২/ ১০২/ ৫০",
		"1485 mm": "১৪৮৫ মিমি",
		"555 mm": "৫৫৫ মিমি",
		"630 mm": "৬৩০ মিমি",
		"HIPS/4": "HIPS/৪",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-a8e-elxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbA8eElxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-a8e-elxx-xx"
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
		"Model Name":          "MFB-A8E-ELXX-XX",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 05.01-111V 06.01-111V 06.02-111V 07.01-104V 07.02-104",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "630 mm",
		"Door Basket": "GPPS/4",
		"Drawer": "HIPS/4",
		"Egg Tray or Pocket": "Yes",
		"Energy Rating": "N/A",
		"Gross Volume": "207 Ltr.",
		"Gross Weight": "58± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height (mm)": "1485 mm",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "102/ 102/ 50",
		"Lock": "Yes",
		"Net Volume": "193 Ltr.",
		"Net Weight": "53± 2 Kg",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 05.01,V 06.01, V 06.02,V 07.01-1000VAV 07.02-Low Voltage(155~260V)For V 07.02 - Wide Voltage Range (155Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range,1000VA is recommended",
		"Refrigerant": "R600a",
		"Shelf (Material/No.)": "Wire/2",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
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
