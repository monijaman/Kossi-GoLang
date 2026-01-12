package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaB4dRlxxXxInverter seeds specifications/options for product 'marcel-mfa-b4d-rlxx-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfaB4dRlxxXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaB4dRlxxXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaB4dRlxxXxInverter() *SpecificationSeederRefrigeratorMarcelMfaB4dRlxxXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfaB4dRlxxXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-b4d-rlxx-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaB4dRlxxXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfa-b4d-rlxx-xx-inverter": "মার্সেল-mfa-b4d-rlxx-xx-inverter",
		"MFA-B4D-RLXX-XX-INVERTER":        "MFA-B4D-RLXX-XX-INVERTER",

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

		// Feature clusters
		"Wire/3":                "ওয়্যার/৩",
		"Wire/2":                "ওয়্যার/২",
		"Egg Case":              "ডিম ধারণ বাক্স",
		"Vegetable Box":         "শাকসবজি বক্স",
		"Interior Lamp":         "ইন্টারিয়র ল্যাম্প",
		"Deodorizer":            "ডিওডোরাইজার",
		"Loading Capacity (40HQ/40Ft/20Ft)": "লোডিং ক্যাপাসিটি- 40HQ/40Ft/20Ft 103/75/36",
		// Add more translations as needed
	
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
		"Recessed/ Grip": "Recessed/ Grip",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"175 Ltr.": "175 লিটার",
		"177 Ltr.": "১৭৭ লিটার",
		"220 ~ 240": "220 ~ 240 ভোল্ট",
		"50 ± 2 Kg": "৫০ ± ২ কেজি",
		"555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-b4d-rlxx-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfaB4dRlxxXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-b4d-rlxx-xx-inverter"
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
        "Model Name":          "MFA-B4D-RLXX-XX-INVERTER",
        "Cooling Technology":  "Direct Cool",
        "Gross Volume":        "177 Ltr.",
        "Net Volume":          "175 Ltr.",
        "Weight":              "50 ± 2 Kg",
        "Refrigerant":         "R600a",
        "Temperature Control": "Mechanical",
        "Voltage":             "220 ~ 240",
        "Dimensions":          "555 x 630 x 1410 mm",
        "Packing Dimensions":  "580 x 645 x 1455 mm",
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
