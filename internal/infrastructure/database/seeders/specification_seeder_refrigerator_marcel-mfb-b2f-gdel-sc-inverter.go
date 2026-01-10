package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2fGdelScInverter seeds specifications/options for product 'marcel-mfb-b2f-gdel-sc-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB2fGdelScInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2fGdelScInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2fGdelScInverter() *SpecificationSeederRefrigeratorMarcelMfbB2fGdelScInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB2fGdelScInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2f-gdel-sc-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdelScInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfb-b2f-gdel-sc-inverter": "মার্সেল-mfb-b2f-gdel-sc-inverter",
		"MFB-B2F-GDEL-SC-INVERTER":        "MFB-B2F-GDEL-SC-INVERTER",

		"Direct Cool":         "ডিরেক্ট কুল",
		"177 Ltr.":            "১৭৭ লিটার",
		"175 Ltr.":            "১৭৫ লিটার",
		"50 ± 2 Kg":           "৫০ ± ২ কেজি",
		"R600a":               "R600a",
		"Mechanical":          "ম্যানুয়াল/মেকানিক্যাল",
		"Manual":              "ম্যানুয়াল",
		"Recessed/ Grip":      "রিসেসড/গ্রিপ",
		"Copper":              "কপার",
		"Cyclopentene":        "সাইক্লোপেন্টেন",
		"220 ~ 240":           "২২0 ~ ২৪0",
		"Hermetic":            "হার্মেটিক",
		"555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"Compressor: V 01.01-RSCR; V 01.02-RSCR; Compressor Input Power (Watt): V 01.01-88; V 01.02-88; Voltage/Hz: 220 ~ 240/ 50; Stabilizer: 5 Ampere": "কমপ্রেসার: V 01.01-RSCR; V 01.02-RSCR; কমপ্রেসার ইনপুট পাওয়ার (ওয়াট): V 01.01-88; V 01.02-88; ভোল্টেজ/হার্জ: 220 ~ 240/50; স্ট্যাবিলাইজার: ৫ অ্যাম্পিয়ার",
		"66 ± 2 Kg": "৬৬ ± ২ কেজি",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes":        "হ্যাঁ",
		"645 mm":     "৬৪৫ মিমি",
		"Electronic": "ইলেকট্রনিক",
		"RSCR":       "RSCR",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"220 ~ 240/ 50":  "২২০ ~ ২৪০/ ৫০",
		"HIPS/5":         "HIPS/৫",
		"Wire/3":         "Wire/৩",
		"RoHS Certified": "RoHS Certified",
		"1765 mm":        "১৭৬৫ মিমি",
		"N~ST":           "N~ST",
		"V 06.01-124 V 07.01-109 V 07.02-109 V 07.03-109 V 08.01-127": "V ০৬.০১-১২৪ V ০৭.০১-১০৯ V ০৭.০২-১০৯ V ০৭.০৩-১০৯ V ০৮.০১-১২৭",
		"N/A":    "N/A",
		"580 mm": "৫৮০ মিমি",
		"V 06.01-Low Voltage(130~260V) For V 06.01- Wide Voltage Range (130Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 07.01,V 07.02,V 07.03-Low Voltage(140~260V) For V07.01, V 07.02,V 07.03 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended V 08.01-Low Voltage(150~260V) For V 08.01 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.": "V ০৬.০১-Low Voltage(১৩০~২৬০V) For V ০৬.০১- Wide Voltage Range (১৩০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended. V ০৭.০১,V ০৭.০২,V ০৭.০৩-Low Voltage(১৪০~২৬০V) For V০৭.০১, V ০৭.০২,V ০৭.০৩ - Wide Voltage Range (১৪০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended V ০৮.০১-Low Voltage(১৫০~২৬০V) For V ০৮.০১ - Wide Voltage Range (১৫০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended.",
		"238 Ltr.":      "২৩৮ লিটার",
		"Wire/2":        "Wire/২",
		"Recessed/Grip": "Recessed/Grip",
		"252 Ltr.":      "২৫২ লিটার",
		"GPPS/4":        "GPPS/৪",
		"59 ± 2 Kg":     "৫৯ ± ২ কেজি",
		"98/ 72/ 36":    "৯৮/ ৭২/ ৩৬",

		// Special features long translation for consolidated unmapped specs
		"Gross Weight: 66 ± 2 Kg; Climatic Type: N~ST; Compressor Input Power (Watt): V 06.01-124; V 07.01-109; V 07.02-109; V 07.03-109; V 08.01-127; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Handle: Recessed/Grip; Lock: Yes; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene (Eco-friendly, 100% CFC & HCFC Free); Recommended voltage stabilizer capacity: V 06.01 - Low Voltage(130~260V) (Wide Range 130Vac - 260Vac) — stabilizer usually not required; V 07.01/V 07.02/V 07.03 - Low Voltage(140~260V) (Wide Range 140Vac - 260Vac) — stabilizer usually not required; V 08.01 - Low Voltage(150~260V) (Wide Range 150Vac - 260Vac) — stabilizer usually not required; if voltages exceed these ranges, 1000VA recommended; Refrigerator Compartment: Interior Lamp: Yes; Vegetable Crisper Cover: Yes; Egg Tray: Yes; Freezer Compartment: Shelf (Wire/3); Drawer: HIPS/5; Net Dimensions: 555 x 630 x 1720 mm; Loading Capacity (40HQ/40Ft/20Ft): 98/ 72/ 36": "কম্প্রেসার ইনপুট পাওয়ার: ৬৬ ± ২ কেজি; ক্লাইম্যাটিক টাইপ: N~ST; কম্প্রেসার ইনপুট পাওয়ার (ওয়াট): V ০৬.০১-১২৪; V ০৭.০১-১০৯; V ০৭.০২-১০৯; V ০৭.০৩-১০৯; V ০৮.০১-১২৭; কুলিং ইফেক্ট: ফ্রিজার অংশ −১৮℃ এর নিচে; রেফ্রিজারেটর অংশ ০℃ থেকে +৫℃; হ্যান্ডেল: রিসেসড/গ্রিপ; লক: হ্যাঁ; থার্মোস্ট্যাট: RoHS সার্টিফায়েড; ক্যাপিলারি: তামা; পলিউরেথেন ফোম ব্লোয়িং এজেন্ট: সাইক্লোপেন্টেন (ইকো-ফ্রেন্ডলি, ১০০% CFC/HCFC মুক্ত); সুপারিশকৃত স্ট্যাবিলাইজার: V ০৬.০১ - লো ভোল্টেজ (১৩০~২৬০V) (ওয়াইড রেঞ্জ ১৩০Vac - ২৬০Vac) — সাধারণত স্ট্যাবিলাইজার প্রয়োজন হয় না; V ০৭.০১/V ০৭.০২/V ০৭.০৩ - লো ভোল্টেজ (১৪০~২৬০V) (ওয়াইড রেঞ্জ ১৪০Vac - ২৬০Vac) — সাধারণত স্ট্যাবিলাইজার প্রয়োজন হয় না; V ০৮.০১ - লো ভোল্টেজ (১৫০~২৬০V) (ওয়াইড রেঞ্জ ১৫০Vac - ২৬০Vac) — সাধারণত স্ট্যাবিলাইজার প্রয়োজন হয় না; যদি ভোল্টেজ এই রেঞ্জের বাইরে যায়, তবে ১০০০VA স্ট্যাবিলাইজার ব্যবহার করার পরামর্শ দেওয়া হয়; রেফ্রিজারেটর অংশ: অভ্যন্তরীণ বাতি: হ্যাঁ; সবজি ক্রিস্পার ঢাকনা: হ্যাঁ; ডিম ধারক: হ্যাঁ; ফ্রিজার অংশ: তাক (তারের) (৩টি); ড্রয়ার: HIPS (৫টি); নেট মাত্রা: ৫৫৫ x ৬৩০ x ১৭২০ মিমি; লোডিং ক্যাপাসিটি (40HQ/40Ft/20Ft): ৯৮/৭২/৩৬",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST": "N ~ ST",
		"No": "না",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSIR": "RSIR",
		"Wire/1": "Wire/১",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"555 mm": "৫৫৫ মিমি",
		"630 mm": "৬৩০ মিমি",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2f-gdel-sc-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdelScInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2f-gdel-sc-inverter"
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
		"Model Name":          "MFB-B2F-GDEL-SC-INVERTER",
		"Capillary":                          "Copper",
		"Climate Type (SN, N, ST, T)":       "N~ST",
		"Compressor Input Power (Watt)":      "V 06.01-124V 07.01-109V 07.02-109V 07.03-109V 08.01-127",
		"Compressor Type":                    "RSCR",
		"Cooling Effect":                     "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual):":    "Manual",
		"Depth (mm)":                           "630 mm",
		"Door Basket":                        "GPPS/4",
		"Drawer":                             "HIPS/5",
		"Egg Tray":                           "Yes",
		"Energy Rating":                      "N/A",
		"Gross Volume":                       "252 Ltr.",
		"Gross Weight":                       "66 ± 2 Kg",
		"Handle (Recessed/ Grip)":            "Recessed/Grip",
		"Height (mm)":                          "1720 mm",
		"Interior Lamp":                      "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "98/ 72/ 36",
		"Lock":                               "Yes",
		"Net Volume":                         "238 Ltr.",
		"Net Weight":                         "59 ± 2 Kg",
		"Polyurethane Foam Blowing Agent":    "Cyclopentene [Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage":                  "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 06.01-Low Voltage(130~260V)For  V 06.01- Wide Voltage Range (130Vac - 260Vac). Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended.  V 07.01,V 07.02,V 07.03-Low Voltage(140~260V)For V07.01, V 07.02,V 07.03 - Wide Voltage Range (140Vac - 260Vac).Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommendedV 08.01-Low Voltage(150~260V)For V 08.01 - Wide Voltage Range (150Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended.",
		"Refrigerant":           "R600a",
		"Shelf Material": "Wire/3",
		"Shelf (Material/No.)":  "Wire/2",
		"Temperature Control": "Electronic",
		"Thermostat":              "RoHS Certified",
		"Type":                    "Direct Cool",
		"Vegetable Crisper":       "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Width":                "555 mm",
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
