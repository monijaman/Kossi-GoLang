package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx seeds specifications/options for product 'marcel-mfa-b2x-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx() *SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-b2x-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-b2x-gdxx-xx": "মার্সেল-mfa-b2x-gdxx-xx",
		"MFA-B2X-GDXX-XX":        "MFA-B2X-GDXX-XX",
		"205 Ltr.":               "205 লিটার",
		"220 Ltr.":               "220 লিটার",
		"545 x 635 x 1580 mm": "৫৪৫ x ৬৩৫ x ১৫৮০ মিমি",
		"48.5 ± 2 Kg":            "48.5 ± 2 কেজি",
		"RSCR":                   "RSCR",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical":             "মেকানিক্যাল",
		"220 ~ 240/50":           "220 ~ 240/50",
		"50":                     "50",
		"5 Star":                 "5 তারা",
		"12":                     "12",
		"52 ± 2 Kg":              "52 ± 2 কেজি",
		"Wire/3":                 "তার/3",
		"GPPS/4":                 "GPPS/4",
		"Steel":                  "ইস্পাত",
		"Cyclopentene":           "সাইক্লোপেন্টেন",
		"108.6 W":                "108.6 W",
		"580 x 645 x 1640 mm": "৫৮০ x ৬৪৫ x ১৬৪০ মিমি",
		"100/ 100/ 50":           "100/ 100/ 50",
		// Add more translations as needed
		"": "",
		"V05.01: No Need to use voltage stabilizer. If out of voltage range(140V-260V), then suggested voltage stabilizer capacity is 1000VA.": "V০৫.০১: No Need to use voltage stabilizer. If out of voltage range(১৪০V-২৬০V), then suggested voltage stabilizer capacity is ১০০০VA.",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃":                                                                        "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes":            "হ্যাঁ",
		"Recessed/ Grip": "Recessed/ Grip",
		"V05.01-R600a":   "V০৫.০১-R৬০০a",
		"Direct Cool":    "ডাইরেক্ট কুল",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"5 Star (BDS 1850:2012)": "৫ Star (BDS ১৮৫০:২০১২)",
		"Yes/1":                  "Yes/১",
		"RoHS Certified":         "RoHS Certified",
		"645":                    "৬৪৫",
		"No":                     "না",
		"N~ST":                   "N~ST",
		"V05.01-108.6":           "V০৫.০১-১০৮.৬",
		"1640":                   "১৬৪০",
		"Wire/2":                 "Wire/২",
		"580":                    "৫৮০",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"Copper": "কপার",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST": "N ~ ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"R600a": "R600a",
		"RSIR": "RSIR",
		"Wire/1": "Wire/১",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/2": "হ্যাঁ/২",
	
		"635": "635",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-b2x-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-b2x-gdxx-xx"
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
		"Model Name":          "MFA-B2X-GDXX-XX",
		"Can Storage Dispenser":          "No",
		"Capillary":                      "Steel",
		"Climatic Type (SN, N, ST, T)":   "N~ST",
		"Compressor Input Power (Watt)":  "V05.01-108.6",
		"Compressor Type":                "RSCR",
		"Cooling Effect":                 "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer":                     "No",
		"Depth/mm":                       "635",
		"Door Basket":                    "GPPS/4",
		"Drawer":                         "No",
		"Egg Tray or Pocket":             "Yes/1",
		"Energy rating":                  "5 Star (BDS 1850:2012)",
		"Gross Volume (Outer Dimension, Manufacturer declared)": "220 Ltr.",
		"Gross Weight":                            "52 ± 2 Kg",
		"Handle (Recessed/ Grip)":                 "Recessed/ Grip",
		"Height/mm":                               "1580",
		"Interior Lamp":                           "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":      "100/ 100/ 50",
		"Lock":                                    "Yes",
		"Net Volume":                              "205 Ltr.",
		"Net Weight":                              "48.5 ± 2 Kg",
		"Polyurethane foam blowing agent":         "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency":   "220 ~ 240/50",
		"Recommended voltage stabilizer capacity": "V05.01: No Need to use voltage stabilizer.If out of voltage range(140V-260V),then suggested voltage             stabilizer capacity is 1000VA.", "Refrigerant": "V05.01-R600a",
		"Reversible Door":                               "No",
		"Shelf (Material/ No.)":                         "Wire/3",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat":                                    "RoHS Certified",
		"Type":                                          "Direct Cool",
		"Vegetable Crisper":                             "Yes/1",
		"Vegetable Crisper Cover":                       "Yes/1",
		"Width/mm":                                      "545",
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
