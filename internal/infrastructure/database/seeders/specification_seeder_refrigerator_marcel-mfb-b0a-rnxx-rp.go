package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB0aRnxxRp seeds specifications/options for product 'marcel-mfb-b0a-rnxx-rp'
type SpecificationSeederRefrigeratorMarcelMfbB0aRnxxRp struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB0aRnxxRp creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB0aRnxxRp() *SpecificationSeederRefrigeratorMarcelMfbB0aRnxxRp {
	return &SpecificationSeederRefrigeratorMarcelMfbB0aRnxxRp{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b0a-rnxx-rp"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB0aRnxxRp) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b0a-rnxx-rp": "মার্সেল-mfb-b0a-rnxx-rp",
		"MFB-B0A-RNXX-RP":        "MFB-B0A-RNXX-RP",

		// Common values used in this seeder
		"Direct Cool":         "ডিরেক্ট কুল",
		"223 Ltr.":            "২২৩ লিটার",
		"209 Ltr.":            "২০৯ লিটার",
		"N~ST":                "N~ST",
		"220 ~ 240/ 50":       "২২০ ~ ২৪০/ ৫০",
		"V 04.01-111":         "V 04.01-111",
		"V 05.01-111":         "V 05.01-111",
		"V 05.02-111":         "V 05.02-111",
		"V 05.03-111":         "V 05.03-111",
		"V 06.01-108.6":       "V 06.01-108.6",
		"RSCR":                "RSCR",
		"Mechanical":          "মেকানিক্যাল",
		"Manual":              "ম্যানুয়াল",
		"Recessed/ Grip":      "রিসেসড/গ্রিপ",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"Copper":              "তামা",
		"Cyclopentene":        "সাইক্লোপেন্টেন",
		"R600a":               "R600a",
		"56 ± 2 Kg":           "৫৬ ± ২ কেজি",
		"60 ± 2 Kg":           "৬০ ± ২ কেজি",
		"555 x 630 x 1560 mm": "৫৫৫ x ৬৩০ x ১৫৬০ মিমি",
		"580 x 645 x 1610 mm": "৫৮০ x ৬৪৫ x ১৬১০ মিমি",
		"101/ 101/ 49":        "১০১/ ১০১/ ৪৯",
		"Wire/3":              "ওয়্যার/৩",
		"GPPS/4":              "GPPS/৪",
		"Wire/2":              "ওয়্যার/২",
		"HIPS/5":              "HIPS/৫",

		// Special features long translation (Bangla)
		`Gross Weight: 60 ± 2 Kg; Packing Dimensions: 580 x 645 x 1610 mm; Loading Capacity (40HQ/ 40Ft/ 20Ft): 101/ 101/ 49; Climatic Type: N~ST; Compressor Input Power: V 04.01-111; V 05.01-111; V 05.02-111; V 05.03-111; V 06.01-108.6; Cooling Effect: Freezer Cabinet Less than -18C; Refrigerator Cabinet 0C to +5C; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene; Recommended voltage stabilizer capacity: V 04.01,V 05.01,V 05.02,V 05.03-1000VA V 06.01-Low Voltage(130~260V) For V 06.01 - Wide Voltage Range (130Vac -260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.`: `গ্রস ওয়েট: 60 ± 2 কেজি; প্যাকেজিং ডাইমেনশন: 580 x 645 x 1610 মিমি; লোডিং ক্ষমতা (40HQ/ 40Ft/ 20Ft): 101/ 101/ 49; ক্লাইম্যাটিক টাইপ: N~ST; কম্প্রেসার ইনপুট পাওয়ার: V 04.01-111; V 05.01-111; V 05.02-111; V 05.03-111; V 06.01-108.6; কুলিং ইফেক্ট: ফ্রিজার কেবিনেট -18C এর নিচে; রেফ্রিজারেটর কেবিনেট 0C থেকে +5C; ক্যাপিলারি: তামা; পলিইউরেথেন ফোম ব্লোয়িং এজেন্ট: সাইক্লোপেন্টেন; সুপারিশকৃত ভোল্টেজ স্ট্যাবিলাইজার ক্ষমতা: V 04.01,V 05.01,V 05.02,V 05.03-1000VA V 06.01-লো ভোল্টেজ(130~260V) V 06.01 এর ক্ষেত্রে - ওয়াইড ভোল্টেজ রেঞ্জ (130Vac -260Vac). ভোল্টেজ স্ট্যাবিলাইজার প্রয়োজন নেই। উল্লিখিত রেঞ্জের বাইরে ভোল্টেজ থাকলে 1000VA সুপারিশ করা হয়।`,
	
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
	
		"555 mm": "৫৫৫ মিমি",
		"630 mm": "৬৩০ মিমি",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b0a-rnxx-rp'
func (s *SpecificationSeederRefrigeratorMarcelMfbB0aRnxxRp) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b0a-rnxx-rp"
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
		"Model Name":          "MFB-B0A-RNXX-RP",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 04.01-111V 05.01-111V 05.02-111V 05.03-111V 06.01-108.6",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "630 mm",
		"Door Basket": "GPPS/4",
		"Drawer": "HIPS/5",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "223 Ltr.",
		"Gross Weight": "60± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height (mm)": "1560 mm",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "101/ 101/ 49",
		"Lock": "Yes",
		"Net Volume": "209 Ltr.",
		"Net Weight": "56± 2 Kg",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 04.01, V 05.01, V 05.02, V 05.03-1000VA, V 06.01-Low Voltage (130~260V)For V 06.01 - Wide Voltage Range (130Vac -260Vac). Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf Material": "Wire/3",
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
