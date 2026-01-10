package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfa2a3GdelSc seeds specifications/options for product 'marcel-mfa-2a3-gdel-sc'
type SpecificationSeederRefrigeratorMarcelMfa2a3GdelSc struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfa2a3GdelSc creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfa2a3GdelSc() *SpecificationSeederRefrigeratorMarcelMfa2a3GdelSc {
	return &SpecificationSeederRefrigeratorMarcelMfa2a3GdelSc{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-2a3-gdel-sc"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfa2a3GdelSc) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-2a3-gdel-sc": "মার্সেল-mfa-2a3-gdel-sc",
		"MFA-2A3-GDEL-SC":        "MFA-2A3-GDEL-SC",
		"Direct Cool":            "ডাইরেক্ট কুল",
		"176 Ltr.":               "১৭৬ লিটার",
		"213 Ltr.":               "২১৩ লিটার",
		"220-240V/ 50Hz":         "২২০-২৪০ভি/ ৫০হার্টজ",
		"50Hz":                   "৫০হার্টজ",
		"RSCR":                   "আরএসসিআর",
		"Electronic":             "ইলেকট্রনিক",
		"Manual":                 "ম্যানুয়াল",
		"R600a":                  "আর৬০০এ",
		"Wire":                   "ওয়্যার",
		"2":                      "২",
		"3":                      "৩",
		"1":                      "১",
		"Yes":                    "হ্যাঁ",
		"557 x 645 x 1510 mm": "৫৫৭ x ৬৪৫ x ১৫১০ মিমি",
		"45.5 ± 2 Kg":            "৪৫.৫ ± ২ কেজি",
		"5 Star":                 "৫ স্টার",
		"Replacement Guarantee: 1 Year, Main Parts (Compressor): 12 Years, Door: 3 Years, Spare Parts: 4 Years, After Sales Service: 5 Years": "রিপ্লেসমেন্ট গ্যারান্টি: ১ বছর, মূল অংশ (কম্প্রেসার): ১২ বছর, দরজা: ৩ বছর, স্পেয়ার পার্টস: ৪ বছর, আফটার সেলস সার্ভিস: ৫ বছর",
		"12": "১২",
		"Lock: Yes, Interior Lamp: Yes, Handle: Recessed, Capillary: Copper, Polyurethane foam blowing agent CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology], Refrigerator Shelf: Wire/2, Refrigerator Door Basket: GPPS/3, Vegetable Crisper: Yes/1, Egg Tray: Yes/1-2, Freezer Shelf: Wire/2, Loading quantity: 102/102/50 (40HQ/40Ft/20Ft), Climatic Type: N~ST, Cooling Effect: Freezer Cabinet Less than -18°C, Refrigerator Cabinet 0°C to +5°C, Energy Rating: 5 Star, Temperature Control: Electronic, Recommended voltage stabilizer capacity: No need if within 145V-260V; 1000VA if outside range.": "লক: হ্যাঁ, ইন্টেরিয়র ল্যাম্প: হ্যাঁ, হ্যান্ডেল: রিসেসড, ক্যাপিলারি: কপার, পলিউরেথেন ফোম ব্লোয়িং এজেন্ট সাইক্লোপেন্টেন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি], রেফ্রিজারেটর শেল্ফ: ওয়্যার/২, রেফ্রিজারেটর দরজা বাস্কেট: জিপিপিএস/৩, ভেজিটেবল ক্রিসপার: হ্যাঁ/১, এগ ট্রে: হ্যাঁ/১-২, ফ্রিজার শেল্ফ: ওয়্যার/২, লোডিং কোয়ান্টিটি: ১০২/১০২/৫০ (৪০এইচকিউ/৪০এফটি/২০এফটি), ক্লাইমেটিক টাইপ: এন~এসটি, কুলিং ইফেক্ট: ফ্রিজার ক্যাবিনেট লেস থ্যান -১৮°C, রেফ্রিজারেটর ক্যাবিনেট ০°C টু +৫°C, এনার্জি রেটিং: ৫ স্টার, টেম্পারেচার কন্ট্রোল: ইলেকট্রনিক, রেকমেন্ডেড ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: ১৪৫ভি-২৬০ভি এর মধ্যে থাকলে প্রয়োজন নেই; বাইরে থাকলে ১০০০ভিএ।",
		"50 ± 2 Kg": "৫০ ± ২ কেজি",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"220 ~ 240/ 50":          "২২০ ~ ২৪০/ ৫০",
		"5 Star (BDS 1850:2012)": "৫ Star (BDS ১৮৫০:২০১২)",
		"1530":                   "১৫৩০",
		"V 1402- 102":            "V ১৪০২- ১০২",
		"Yes/1":                  "Yes/১",
		"Copper":                 "কপার",
		"645":                    "৬৪৫",
		"No":                     "না",
		"N~ST":                   "N~ST",
		"V 1402-R600a":           "V ১৪০২-R৬০০a",
		"No Need to use voltage stabilizer. If out of voltage range(145V-260V), then suggested voltage stabilizer capacity is 1000VA.": "No Need to use voltage stabilizer. If out of voltage range(১৪৫V-২৬০V), then suggested voltage stabilizer capacity is ১০০০VA.",
		"Yes/1-2": "Yes/১-২",
		"CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentane [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Wire/2":       "Wire/২",
		"580":          "৫৮০",
		"102/ 102 /50": "১০২/ ১০২ /৫০",
		"Recessed":     "Recessed",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Mechanical": "মেকানিক্যাল",
		"N ~ ST": "N ~ ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/2": "হ্যাঁ/২",
	
		"GPPS/3": "জিপিপিএস/৩",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-2a3-gdel-sc'
func (s *SpecificationSeederRefrigeratorMarcelMfa2a3GdelSc) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-2a3-gdel-sc"
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
		"Model Name":          "MFA-2A3-GDEL-SC",
		"Can Storage Dispenser":          "No",
		"Capillary":                      "Copper",
		"Climate Type (SN, N, ST, T)":   "N~ST",
		"Compressor Input Power (Watt)":  "V 1402- 102",
		"Compressor Type":                "RSCR",
		"Cooling Effect":                 "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer":                     "No",
		"Depth (mm)":                       "645",
		"Door Basket":                    "No",
		"Door Basket":                    "GPPS/3",
		"Drawer":                         "No",
		"Egg Tray or Pocket":             "Yes/1-2",
		"Energy Star Rating":                  "5 Star (BDS 1850:2012)",
		"Gross Volume": "213 Ltr.",
		"Gross Weight":                       "50 ± 2 Kg",
		"Handle (Recessed/ Grip)":            "Recessed",
		"Height (mm)":                          "1510",
		"Interior Lamp":                      "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "102/ 102 /50",
		"Lock":                               "Yes",
		"Net Volume":                         "176 Ltr.",
		"Net Weight":                         "45.5 ± 2 Kg",
		"Polyurethane Foam Blowing Agent":    "CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage":                  "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "No Need to use voltage stabilizer.If out of voltage range(145V-260V), then suggested              voltage stabilizer capacity is 1000VA.",
		"Refrigerant":           "V 1402-R600a",
		"Reversible Door":       "No",
		"Shelf Material":  "Wire/2",
		"Temperature Control": "Electronic",
		"Type":                    "Direct Cool",
		"Vegetable Crisper":       "Yes/1",
		"Vegetable Crisper Cover": "Yes/1",
		"Width":                "557",
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
