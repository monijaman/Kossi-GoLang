package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2fElxxXx seeds specifications/options for product 'marcel-mfb-b2f-elxx-xx'
type SpecificationSeederRefrigeratorMarcelMfbB2fElxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2fElxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2fElxxXx() *SpecificationSeederRefrigeratorMarcelMfbB2fElxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB2fElxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2f-elxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2fElxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b2f-elxx-xx": "মার্সেল-mfb-b2f-elxx-xx",
		"MFB-B2F-ELXX-XX":        "MFB-B2F-ELXX-XX",

		"Direct Cool":                    "সরাসরি কুলিং",
		"252 Ltr.":                       "২৫২ লিটার",
		"238 Ltr.":                       "২৩৮ লিটার",
		"59 ± 2 Kg":                      "৫৯ ± ২ কেজি",
		"66 ± 2 Kg":                      "৬৬ ± ২ কেজি",
		"N~ST":                           "N~ST",
		"220 ~ 240/ 50":                  "220 ~ 240 ভোল্ট, 50 হার্টজ",
		"V 06.01-124":                    "V 06.01-124",
		"V 07.01-109":                    "V 07.01-109",
		"V 07.02-109":                    "V 07.02-109",
		"V 07.03-109":                    "V 07.03-109",
		"V 08.01-127":                    "V 08.01-127",
		"RSCR":                           "RSCR",
		"Energy Rating":                  "Energy Rating",
		"Freezer Cabinet Less than -18℃": "ফ্রিজার অংশ -18℃ এর নিচে",
		"Refrigerator Cabinet 0℃ to +5℃": "রেফ্রিজারেটর অংশ 0℃ থেকে +5℃",
		"R600a":                          "R600a",
		"Mechanical":                     "ম্যানুয়াল",
		"Manual":                         "ম্যানুয়াল",
		"Recessed/ Grip":                 "রিসেসড (গ্রিপ)",
		"Copper":                         "তামা",
		"Cyclopentene":                   "সাইক্লোপেন্টেন",
		"555 x 630 x 1720 mm": "৫৫৫ x ৬৩০ x ১৭২০ মিমি",
		"580 x 645 x 1765 mm": "৫৮০ x ৬৪৫ x ১৭৬৫ মিমি",
		"Wire/2":                         "তারের তাক (২টি)",
		"Wire/3":                         "তারের তাক (৩টি)",
		"GPPS/4":                         "GPPS (৪টি)",
		"HIPS/5":                         "HIPS (৫টি)",
		"Yes":                            "হ্যাঁ",
		"Egg Tray":                       "ডিম ধারক",
		"Interior Lamp":                  "অভ্যন্তরীণ বাতি",
		"72/ 36":                         "৭২/৩৬",
		"98/ 72/ 36":                     "৯৮/৭২/৩৬",

		// Special features long translation
		`Gross Weight: 66 ± 2 Kg; Packing Dimensions: 580 x 645 x 1765 mm; Loading Capacity (40HQ/ 40Ft/ 20Ft): 98/ 72/ 36; Climatic Type (SN, N, ST, T): N~ST; Compressor Input Power (Watt): V 06.01-124; V 07.01-109; V 07.02-109; V 07.03-109; V 08.01-127; Compressor Type: RSCR; Energy Rating: No; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene; Recommended voltage stabilizer capacity: V 06.01-Low Voltage(130~260V) For V 06.01 - Wide Voltage Range (130Vac - 260Vac). Voltage stabilizer is not required. V 07.01,V 07.02,V 07.03-Low Voltage(140~260V) For V07.01, V 07.02,V 07.03 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. V 08.01-Low Voltage(150~260V) For V 08.01 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended; Refrigerator Compartment: Shelf (Material/ No.): Wire/2; Door Basket: GPPS/4; Interior Lamp: Yes; Vegetable Crisper: Yes; Vegetable Crisper Cover: Yes; Egg Tray: Yes; Freezer Compartment: Shelf (Material/ No.): Wire/3; Drawer: HIPS/5; Loading Capacity- 40HQ/ 40Ft/ 20Ft: 98/ 72/ 36`: `গ্রস ওজন: ৬৬ ± ২ কেজি; প্যাকেজিং মাত্রা: ৫৮০ x ৬৪৫ x ১৭৬৫ মিমি; লোডিং ক্যাপাসিটি (40HQ/40Ft/20Ft): ৯৮/৭২/৩৬; আবহাওয়া টাইপ: N~ST; কম্প্রেসার ইনপুট পাওয়ার (ওয়াট): V 06.01-124; V 07.01-109; V 07.02-109; V 07.03-109; V 08.01-127; কম্প্রেসার টাইপ: RSCR; এনার্জি রেটিং: নেই; কুলিং পারফরম্যান্স: ফ্রিজার অংশ -18℃ এর নিচে; রেফ্রিজারেটর অংশ 0℃ থেকে +5℃; থার্মোস্ট্যাট: RoHS সার্টিফায়েড; ক্যাপিলারি: তামা; পলিউরেথেন ফোম ব্লোয়িং এজেন্ট: সাইক্লোপেন্টেন (ইকো-ফ্রেন্ডলি, CFC/HCFC মুক্ত); সুপারিশকৃত ভোল্টেজ স্ট্যাবিলাইজার: V 06.01 (লো ভোল্টেজ 130~260V) — V 06.01 এর ক্ষেত্রে উইড ভোল্টেজ রেঞ্জ 130Vac–260Vac; V 07.01, V 07.02, V 07.03 (লো ভোল্টেজ 140~260V) — এইগুলোর ক্ষেত্রে উইড ভোল্টেজ রেঞ্জ 140Vac–260Vac; V 08.01 (লো ভোল্টেজ 150~260V) — উইড ভোল্টেজ রেঞ্জ 150Vac–260Vac; সাধারণত স্ট্যাবিলাইজার প্রয়োজন নেই; এই সীমার বাইরে ভোল্টেজ হলে 1000VA স্ট্যাবিলাইজার ব্যবহার করার পরামর্শ দেওয়া হয়; রেফ্রিজারেটর অংশ: তাক (তারের) (২টি); দরজার বাস্কেট: GPPS (৪টি); অভ্যন্তরীণ বাতি: হ্যাঁ; সবজি ক্রিস্পার: হ্যাঁ; সবজি ক্রিস্পার ঢাকনা: হ্যাঁ; ডিম ধারক: হ্যাঁ; ফ্রিজার অংশ: তাক (তারের) (৩টি); ড্রয়ার: HIPS (৫টি);`,
	
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
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"555 mm": "৫৫৫ মিমি",
		"630 mm": "৬৩০ মিমি",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2f-elxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2fElxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2f-elxx-xx"
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
		"Model Name":          "MFB-B2F-ELXX-XX",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 06.01-124V 07.01-109V 07.02-109V 07.03-109V 08.01-127",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "630 mm",
		"Door Basket": "GPPS/4",
		"Drawer": "HIPS/5",
		"Egg Tray": "Yes",
		"Energy Rating": "No",
		"Gross Volume": "252 Ltr.",
		"Gross Weight": "66± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height (mm)": "1720 mm",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "98/ 72/ 36",
		"Lock": "Yes",
		"Net Volume": "238 Ltr.",
		"Weight": "59± 2 Kg",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 06.01-Low Voltage(130~260V)For V 06.01- Wide Voltage Range (130Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is   recommended.V 07.01,V 07.02,V 07.03-Low Voltage(140~260V)For V07.01, V 07.02,V 07.03 - Wide Voltage Range (140Vac -   260Vac).Voltage stabilizer is not required. In case of voltages beyond   this range,1000VA is recommended V 08.01-Low Voltage(150~260V)For V 08.01 - Wide Voltage Range (150Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is   recommended.",
		"Refrigerant": "R600a",
		"Shelf Material": "Wire/2",
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
