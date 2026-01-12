package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx seeds specifications/options for product 'marcel-mfb-b1h-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx() *SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b1h-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b1h-gdxx-xx": "মার্সেল-mfb-b1h-gdxx-xx",
		"MFB-B1H-GDXX-XX":        "MFB-B1H-GDXX-XX",

		// Common values used in this seeder
		"Direct Cool":         "ডিরেক্ট কুল",
		"218 Ltr.":            "২১৮ লিটার",
		"190 Ltr.":            "১৯০ লিটার",
		"N~ST":                "N~ST",
		"220 ~ 240/ 50":       "২২০ ~ ২৪০/ ৫০",
		"V 01.01-88":          "V 01.01-88",
		"V 02.01-111":         "V 02.01-111",
		"V 02.02-111":         "V 02.02-111",
		"V 03.01-109":         "V 03.01-109",
		"RSCR":                "RSCR",
		"Mechanical":          "মেকানিক্যাল",
		"Manual":              "ম্যানুয়াল",
		"Recessed/ Grip":      "রিসেসড/গ্রিপ",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"Copper":              "তামা",
		"Cyclopentene":        "সাইক্লোপেন্টেন",
		"R600a":               "R600a",
		"57 ± 2 Kg":           "৫৭ ± ২ কেজি",
		"62 ± 2 Kg":           "৬২ ± ২ কেজি",
		"555 x 630 x 1580 mm": "৫৫৫ x ৬৩০ x ১৫৮০ মিমি",
		"580 x 645 x 1625 mm": "৫৮০ x ৬৪৫ x ১৬২৫ মিমি",
		"101/ 101/ 49":        "১০১/ ১০১/ ৪৯",
		"Wire/2":              "ওয়্যার/২",
		"Wire/3":              "ওয়্যার/৩",
		"GPPS/4":              "GPPS/৪",
		"HIPS/5":              "HIPS/৫",

		// Special features long translation (Bangla)
		`Gross Weight: 62 ± 2 Kg; Packing Dimensions: 580 x 645 x 1625 mm; Loading Capacity (40HQ/ 40Ft/ 20Ft): 101/ 101/ 49; Climatic Type: N~ST; Compressor Input Power: V 01.01-88; V 02.01-111; V 02.02-111; V 03.01-109; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene; Recommended voltage stabilizer capacity: V 01.01,V 02.01,V 02.02- 1000VA; V 03.01-Low Voltage(135~260V) For V 03.01 - Wide Voltage Range (135Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.`: `গ্রস ওয়েট: 62 ± 2 কেজি; প্যাকেজিং ডাইমেনশন: 580 x 645 x 1625 মিমি; লোডিং ক্ষমতা (40HQ/ 40Ft/ 20Ft): 101/ 101/ 49; ক্লাইম্যাটিক টাইপ: N~ST; কম্প্রেসার ইনপুট পাওয়ার: V 01.01-88; V 02.01-111; V 02.02-111; V 03.01-109; কুলিং ইফেক্ট: ফ্রিজার কেবিনেট -18℃ এর নিচে; রেফ্রিজারেটর কেবিনেট 0℃ থেকে +5℃; ক্যাপিলারি: তামা; পলিইউরেথেন ফোম ব্লোয়িং এজেন্ট: সাইক্লোপেন্টেন; সুপারিশকৃত ভোল্টেজ স্ট্যাবিলাইজার ক্ষমতা: V 01.01,V 02.01,V 02.02- 1000VA; V 03.01-লো ভোল্টেজ(135~260V) V 03.01 এর ক্ষেত্রে - ওয়াইড ভোল্টেজ রেঞ্জ (135Vac - 260Vac). ভোল্টেজ স্ট্যাবিলাইজার প্রয়োজন নেই। উল্লিখিত রেঞ্জের বাইরে ভোল্টেজ থাকলে 1000VA সুপারিশ করা হয়।`,
	
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
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b1h-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b1h-gdxx-xx"
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
		"Model Name":          "MFB-B1H-GDXX-XX",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "218 Ltr.",
		"Net Volume":          "190 Ltr.",
		"Weight":              "57 ± 2 Kg",
		"Refrigerant":         "R600a",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Voltage":             "220 ~ 240/ 50",
		"Dimensions":          "555 x 630 x 1580 mm",
		"Compressor Type":     "RSCR",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "GPPS/4",
		"Crisper Drawers":     "Yes",
		"Special Features":    `Gross Weight: 62 ± 2 Kg; Packing Dimensions: 580 x 645 x 1625 mm; Loading Capacity (40HQ/ 40Ft/ 20Ft): 101/ 101/ 49; Climatic Type: N~ST; Compressor Input Power: V 01.01-88; V 02.01-111; V 02.02-111; V 03.01-109; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene; Recommended voltage stabilizer capacity: V 01.01,V 02.01,V 02.02- 1000VA; V 03.01-Low Voltage(135~260V) For V 03.01 - Wide Voltage Range (135Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.`,
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
