package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaB4dRxxxRp seeds specifications/options for product 'marcel-mfa-b4d-rxxx-rp'
type SpecificationSeederRefrigeratorMarcelMfaB4dRxxxRp struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaB4dRxxxRp creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaB4dRxxxRp() *SpecificationSeederRefrigeratorMarcelMfaB4dRxxxRp {
	return &SpecificationSeederRefrigeratorMarcelMfaB4dRxxxRp{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-b4d-rxxx-rp"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaB4dRxxxRp) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-b4d-rxxx-rp": "মার্সেল-mfa-b4d-rxxx-rp",
		"MFA-B4D-RXXX-RP":        "MFA-B4D-RXXX-RP",

		// Common values used in this seeder
		"Direct Cool":                    "ডিরেক্ট কুল",
		"244 Ltr.":                       "২৪৪ লিটার",
		"220 Ltr.":                       "২২০ লিটার",
		"N~ST":                           "N~ST",
		"220-240V~":                      "২২০-২৪০V~",
		"50Hz":                           "৫০Hz",
		"V 0701-108.6; V 0801-99.4":      "V 0701-108.6; V 0801-99.4",
		"RSCR":                           "RSCR",
		"Mechanical":                     "মেকানিক্যাল",
		"Manual":                         "ম্যানুয়াল",
		"Recessed":                       "রিসেসড",
		"Yes":                            "হ্যাঁ",
		"No":                             "না",
		"Copper":                         "তামা",
		"Cyclopentene":                   "সাইক্লোপেন্টেন",
		"R600a":                          "R600a",
		"51/57 (Net/Packing) Kg (±2 Kg)": "৫১/৫৭ (নেট/প্যাকিং) কেজি (±২ কেজি)",
		"545 x 605 x 1760 mm": "৫৪৫ x ৬০৫ x ১৭৬০ মিমি",
		"V 0701-R600a; V 0801-R600a":     "V 0701-R600a; V 0801-R600a",
		"5 star (BDS 1850:2012)":         "৫ স্টার (BDS 1850:2012)",

		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",

		"V 0701-R600a V 0801-R600a": "V ০৭০১-R৬০০a V ০৮০১-R৬০০a",
		"Wire":                      "ওয়্যার",
		"3":                         "৩",
		"4":                         "৪",
		"Yes/1":                     "Yes/১",

		// Feature clusters
		"Wire/3":                "ওয়্যার/৩",
		"Wire/2":                "ওয়্যার/২",
		"Egg Case":              "ডিম ধারণ বাক্স",
		"Vegetable Box":         "শাকসবজি বক্স",
		"Interior Lamp":         "ইন্টারিয়র ল্যাম্প",
		"Deodorizer":            "ডিওডোরাইজার",
		"Loading Capacity (40HQ/40Ft/20Ft)": "লোডিং ক্যাপাসিটি- 40HQ/40Ft/20Ft 103/75/36",
		// Add more translations as needed
		"51/57": "৫১/৫৭",
		"1770":  "১৭৭০",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃":                                                                               "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"V 0701/V 0801 : No Need to use voltage stabilizer If out of voltage range(145V-260V), then suggested voltage stabilizer capacity is 1000VA.": "V ০৭০১/V ০৮০১ : No Need to use voltage stabilizer If out of voltage range(১৪৫V-২৬০V), then suggested voltage stabilizer capacity is ১০০০VA.",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]":                                                                         "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"RoHS Certified":           "RoHS Certified",
		"645":                      "৬৪৫",
		"V 0701-108.6 V 0801-99.4": "V ০৭০১-১০৮.৬ V ০৮০১-৯৯.৪",
		"Yes/ 1":                   "Yes/ ১",
		"103/ 75/ 36":              "১০৩/ ৭৫/ ৩৬",
		"580":                      "৫৮০",
	
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST": "N ~ ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"Wire/1": "Wire/১",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/2": "হ্যাঁ/২",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-b4d-rxxx-rp'
func (s *SpecificationSeederRefrigeratorMarcelMfaB4dRxxxRp) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-b4d-rxxx-rp"
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
		"Model Name":          "MFA-B4D-RXXX-RP",
		"Capillary":                      "Copper",
		"Climate Type (SN, N, ST, T)":   "N~ST",
		"Compressor Input Power (Watt)":  "V 0701-108.6V 0801-99.4",
		"Compressor Type":                "RSCR",
		"Cooling Effect":                 "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer":                     "No",
		"Depth (mm)":                       "605",
		"Door Basket":                    "No",
		"Drawer":                         "No",
		"Egg Case":                       "Yes/ 1",
		"Energy Rating":                  "5 star (BDS 1850:2012)",
		"Gross Volume": "244 Ltr.",
		"Handle (Recessed/ Grip)":                               "Recessed",
		"Height (mm)":                                             "1760",
		"Interior Lamp":                                         "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)":                    "103/ 75/ 36",
		"Lock":                                                  "Yes",
		"Net Volume":                                            "220 Ltr.",
		"Polyurethane Foam Blowing Agent":                       "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage":                 "220-240V~ and 50Hz",
		"Recommended voltage stabilizer capacity":               "V 0701/V 0801 : No Need to use voltage stabilizer If out of voltage range(145V-260V), then suggested voltage stabilizer capacity is 1000VA.",
		"Refrigerant":                                           "V 0701-R600aV 0801-R600a",
		"Reversible Door":                                       "No",
		"Shelf Material":                                 "Wire/2",
		"Shelf (Material/No.)":                                  "Wire/3",
		"Temperature Control":         "Mechanical",
		"Thermostat":                                            "RoHS Certified",
		"Type":                                                  "Direct Cool",
		"Vegetable Box":                                         "Yes/1",
		"Vegetable Box Cover":                                   "Yes",
		"Weight/Kg - Net/Packing (±2 KG)":                       "51/57",
		"Width":                                              "545",
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
