package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2fGdxxXx seeds specifications/options for product 'marcel-mfb-b2f-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfbB2fGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2fGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2fGdxxXx() *SpecificationSeederRefrigeratorMarcelMfbB2fGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB2fGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2f-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b2f-gdxx-xx": "মার্সেল-mfb-b2f-gdxx-xx",
		"MFB-B2F-GDXX-XX":        "MFB-B2F-GDXX-XX",

		"Direct Cool":                    "ডিরেক্ট কুল",
		"252 Ltr.":                       "২৫২ লিটার",
		"238 Ltr.":                       "২৩৮ লিটার",
		"59 ± 2 Kg":                      "৫৯ ± ২ কেজি",
		"66 ± 2 Kg":                      "৬৬ ± ২ কেজি",
		"RSCR":                           "আরএসসিআর",
		"No":                             "না",
		"Freezer Cabinet Less than -18℃": "ফ্রিজার কেবিনেট -১৮℃ এর কম",
		"Refrigerator Cabinet 0℃ to +5℃": "রেফ্রিজেরেটর কেবিনেট ০℃ থেকে +৫℃",
		"Mechanical":                     "ম্যানুয়াল/মেকানিক্যাল",
		"Manual":                         "ম্যানুয়াল",
		"Recessed/ Grip":                 "রিসেসড/গ্রিপ",
		"Yes":                            "হ্যাঁ",
		"V 06.01-R600a; V 07.01-R600a; V 07.02-R600a; V 07.03-R600a; V 08.01-R600a": "ভি ০৬.০১-আর৬০০এ; ভি ০৭.০১-আর৬০০এ; ভি ০৭.০২-আর৬০০এ; ভি ০৭.০৩-আর৬০০এ; ভি ০৮.০১-আর৬০০এ",
		"RoHS Certified": "রোএইচএস সনদপ্রাপ্ত",
		"Capillary":      "ক্যাপিলারি",
		"Copper":         "কপার",
		"Cyclopentene":   "সাইক্লোপেন্টেন",
		"Eco-friendly (100% CFC & HCFC Free) Green Technology": "পরিবেশবান্ধব (১০০% সিএফসি ও এইচসিএফসি মুক্ত) সবুজ প্রযুক্তি",
		"Wire":                "তারা",
		"2":                   "২",
		"GPPS/4":              "জিপিপিএস/৪",
		"555 x 630 x 1720 mm": "৫৫৫ x ৬৩০ x ১৭২০ মিমি",
		"580 x 645 x 1765 mm": "৫৮০ x ৬৪৫ x ১৭৬৫ মিমি",
		"98/ 72/ 36":          "৯৮/ ৭২/ ৩৬",
		"HIPS/5":              "এইচআইপিএস/৫",
		"Net Weight: 59 ± 2 Kg; Gross Weight: 66 ± 2 Kg; Performance: Climatic Type: N~ST; Compressor Input Power (Watt): V 06.01-124; V 07.01-109; V 07.02-109; V 07.03-109; V 08.01-127; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Refrigerant: V 06.01-R600a; V 07.01-R600a; V 07.02-R600a; V 07.03-R600a; V 08.01-R600a; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene; Eco-friendly (100% CFC & HCFC Free) Green Technology; Recommended voltage stabilizer: V 06.01-Low Voltage(130~260V); V 07.01,V 07.02,V 07.03-Low Voltage(140~260V); V 08.01-Low Voltage(150~260V); Refrigerator Shelf: Wire/2; Door Basket: GPPS/4; Interior Lamp: Yes; Vegetable Crisper: Yes; Vegetable Crisper Cover: Yes; Egg Tray: Yes; Freezer Shelf: Wire/3; Freezer Drawer: HIPS/5; Packing Dimensions: 580 x 645 x 1765 mm; Loading Capacity- 40HQ/ 40Ft/ 20Ft: 98/ 72/ 36": "নেট ওজন: ৫৯ ± ২ কেজি; গ্রস ওজন: ৬৬ ± ২ কেজি; পারফরম্যান্স: ক্লাইম্যাটিক টাইপ: N~ST; কম্প্রেসার ইনপুট পাওয়ার (ওয়াট): ভি ০৬.০১-১২৪; ভি ০৭.০১-১০৯; ভি ০৭.০২-১০৯; ভি ০৭.০৩-১০৯; ভি ০৮.০১-১২৭; কুলিং ইফেক্ট: ফ্রিজার কেবিনেট -১৮℃ এর কম; রেফ্রিজেরেটর কেবিনেট ০℃ থেকে +৫℃; রেফ্রিজারেন্ট: ভি ০৬.০১-আর৬০০এ; ভি ০৭.০১-আর৬০০এ; ভি ০৭.০২-আর৬০০এ; ভি ০৭.০৩-আর৬০০এ; ভি ০৮.০১-আর৬০০এ; থার্মোস্ট্যাট: রোএইচএস সনদপ্রাপ্ত; ক্যাপিলারি: কপার; পলিউরেথেন ফোম ব্লোয়িং এজেন্ট: সাইক্লোপেন্টেন; পরিবেশবান্ধব (১০০% সিএফসি ও এইচসিএফসি মুক্ত) সবুজ প্রযুক্তি; সুপারিশকৃত ভোল্টেজ স্ট্যাবিলাইজার: ভি ০৬.০১-নিম্ন ভোল্টেজ (১৩০~২৬০ভি); ভি ০৭.০১, ভি ০৭.০২, ভি ০৭.০৩ - নিম্ন ভোল্টেজ (১৪০~২৬০ভি); ভি ০৮.০১ - নিম্ন ভোল্টেজ (১৫০~২৬০ভি); রেফ্রিজারেটর শেলফ: তারে/২; ডোর বাস্কেট: জিপিপিএস/৪; ইন্টারিয়র ল্যাম্প: হ্যাঁ; ভেজিটেবল ক্রিস্পার: হ্যাঁ; ভেজিটেবল ক্রিস্পার কভাৰ: হ্যাঁ; ডিম ট্রে: হ্যাঁ; ফ্রিজার শেলফ: তারে/৩; ফ্রিজার ড্রয়ার: এইচআইপিএস/৫; প্যাকিং ডাইমেনশনস: ৫৮০ x ৬৪৫ x ১৭৬৫ মিমি; লোডিং ক্যাপাসিটি- ৪০এইচকিউ/ ৪০এফটি/ ২০এফটি: ৯৮/ ৭২/ ৩৬",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST": "N ~ ST",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"R600a": "R600a",
		"RSIR": "RSIR",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"555 mm": "৫৫৫ মিমি",
		"630 mm": "৬৩০ মিমি",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2f-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2f-gdxx-xx"
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
		"Model Name":          "MFB-B2F-GDXX-XX",
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
		"Net Weight": "59± 2 Kg",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 06.01-Low Voltage(130~260V)For V 06.01- Wide Voltage Range (130Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is   recommended.V 07.01,V 07.02,V 07.03-Low Voltage(140~260V)For V07.01, V 07.02,V 07.03 - Wide Voltage Range (140Vac -   260Vac).Voltage stabilizer is not required. In case of voltages beyond   this range,1000VA is recommended V 08.01-Low Voltage(150~260V)For V 08.01 - Wide Voltage Range (150Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is   recommended.",
		"Refrigerant": "V 06.01-R600aV 07.01-R600aV 07.02-R600aV 07.03-R600aV 08.01-R600a",
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
