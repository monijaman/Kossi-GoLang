package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2cGdshXx seeds specifications/options for product 'marcel-mfb-b2c-gdsh-xx'
type SpecificationSeederRefrigeratorMarcelMfbB2cGdshXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2cGdshXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2cGdshXx() *SpecificationSeederRefrigeratorMarcelMfbB2cGdshXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB2cGdshXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2c-gdsh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2cGdshXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b2c-gdsh-xx": "মার্সেল-mfb-b2c-gdsh-xx",
		"MFB-B2C-GDSH-XX":        "MFB-B2C-GDSH-XX",

		"Direct Cool":    "ডিরেক্ট কুল",
		"223 Ltr.":       "২২৩ লিটার",
		"219 Ltr.":       "২১৯ লিটার",
		"51.5 ± 2 Kg":    "৫১.৫ ± ২ কেজি",
		"56 ± 2 Kg":      "৫৬ ± ২ কেজি",
		"N~ST":           "N~ST",
		"220 ~ 240/ 50":  "২২০ ~ ২৪০/ ৫০",
		"V 01.01-R600a":  "V 01.01-R600a",
		"V 01.02-R600a":  "V 01.02-R600a",
		"RSCR":           "RSCR",
		"Mechanical":     "ম্যানুয়াল/মেকানিক্যাল",
		"Manual":         "ম্যানুয়াল",
		"Recessed/ Grip": "রিসেসড/গ্রিপ",
		"Copper":         "তামা",
		"Cyclopentene":   "সাইক্লোপেন্টেন",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"RoHS Certified":      "RoHS Certified",
		"555 x 630 x 1550 mm": "৫৫৫ x ৬৩০ x ১৫৫০ মিমি",
		"580 x 645 x 1620 mm": "৫৮০ x ৬৪৫ x ১৬২০ মিমি",
		"72/ 72/ 36":          "৭২/ ৭২/ ৩৬",
		"Wire/2":              "ওয়্যার/২",
		"GPPS/3":              "GPPS/৩",
		"Yes":                 "হ্যাঁ",

		// Special features long translation
		`Gross Weight: 56 ± 2 Kg; Packing Dimensions: 580 x 645 x 1620 mm; Loading Capacity (40HQ/ 40Ft/ 20Ft): 72/ 72/ 36; Climatic Type: N~ST; Compressor Input Power: V 01.01-R600a; V 01.02-R600a; Compressor Type: RSCR; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene; Recommended voltage stabilizer capacity: V 01.01, V 01.02-Low Voltage(140~260V). For V 01.01, V 01.02 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.`: `গ্রস ওয়েট: 56 ± 2 কেজি; প্যাকেজিং ডাইমেনশন: 580 x 645 x 1620 মিমি; লোডিং ক্ষমতা (40HQ/ 40Ft/ 20Ft): 72/ 72/ 36; ক্লাইম্যাটিক টাইপ: N~ST; কম্প্রেসার ইনপুট পাওয়ার: V 01.01-R600a; V 01.02-R600a; কম্প্রেসার টাইপ: RSCR; কুলিং ইফেক্ট: ফ্রিজার কেবিনেট -18℃ এর নিচে; রেফ্রিজারেটর কেবিনেট 0℃ থেকে +5℃; থার্মোস্ট্যাট: RoHS সার্টিফায়েড; ক্যাপিলারি: তামা; পলিইউরেথেন ফোম ব্লোয়িং এজেন্ট: সাইক্লোপেন্টেন; সুপারিশকৃত ভোল্টেজ স্ট্যাবিলাইজার ক্ষমতা: V 01.01, V 01.02-লো ভোল্টেজ(140~260V). V 01.01, V 01.02 এর ক্ষেত্রে - ওয়াইড ভোল্টেজ রেঞ্জ (140Vac - 260Vac). ভোল্টেজ স্ট্যাবিলাইজার প্রয়োজন নেই। উল্লিখিত রেঞ্জের বাইরে ভোল্টেজ থাকলে 1000VA সুপারিশ করা হয়।`,
	
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
		"R600a": "R600a",
		"RSIR": "RSIR",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"555 mm": "৫৫৫ মিমি",
		"630 mm": "৬৩০ মিমি",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2c-gdsh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2cGdshXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2c-gdsh-xx"
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
		"Model Name":          "MFB-B2C-GDSH-XX",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 01.01-R600aV 01.02-R600a",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "630 mm",
		"Door Basket": "GPPS/3",
		"Drawer": "No",
		"Egg Tray": "Yes/1",
		"Gross Volume": "223 Ltr.",
		"Gross Weight": "56± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height (mm)": "1550 mm",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "72/ 72/ 36",
		"Lock": "Yes",
		"Net Volume": "219 Ltr.",
		"Net Weight": "51.5± 2 Kg",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 01.01, V 01.02-Low Voltage(140~260V)For V 01.01, V 01.02 - Wide Voltage Range(140Vac - 260Vac). Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended",
		"Refrigerant": "R600a",
		"Refrigerator Compartment Light": "Yes",
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
