package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC2xGdelXxInverter seeds specifications/options for product 'marcel-mfe-c2x-gdel-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeC2xGdelXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC2xGdelXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC2xGdelXxInverter() *SpecificationSeederRefrigeratorMarcelMfeC2xGdelXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeC2xGdelXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c2x-gdel-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC2xGdelXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfe-c2x-gdel-xx-inverter": "মার্সেল-mfe-c2x-gdel-xx-inverter",
		"MFE-C2X-GDEL-XX-INVERTER":        "MFE-C2X-GDEL-XX-INVERTER",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"341 Ltr":                         "৩৪১ লিটার",
		"320 Ltr.":                        "৩২০ লিটার",
		"67 ± 2 Kg":                       "৬৭ ± ২ কেজি",
		"72 ± 2 Kg":                       "৭২ ± ২ কেজি",
		"N ~ ST":                          "এন ~ এসটি",
		"220-240V~/50Hz":                  "২২০-২৪০V~/৫০Hz",
		"V 0601 - (38~109) V 0701 - (33~126) V 0702 - (33~126)": "V ০৬০১ - (৩৮~১০৯) V ০৭০১ - (৩৩~১২৬) V ০৭০২ - (৩৩~১২৬)",
		"V 0601 - BLDC V 0701 - BLDC V 0702 - BLDC":             "V ০৬০১ - BLDC V ০৭০১ - BLDC V ০৭০২ - BLDC",
		"V 0601 - R600a V 0701 - R600a V 0702 - R600a":          "V ০৬০১ - R৬০০a V ০৭০১ - R৬০০a V ০৭০২ - R৬০০a",
		"Wire":                "ওয়্যার",
		"2":                   "২",
		"PVC/4":               "পিভিসি/৪",
		"Yes/1":               "হ্যাঁ/১",
		"Copper":              "কপার",
		"RoHS Certified":      "RoHS সার্টিফায়েড",
		"Cyclopentene":        "সাইক্লোপেন্টিন",
		"594 x 708 x 1720 mm": "৫৯৪ x ৭০৮ x ১৭২০ মিমি",
		"635 x 740 x 1790 mm": "৬৩৫ x ৭৪০ x ১৭৯০ মিমি",
		"77/ 57/ 27":          "৭৭/ ৫৭/ ২৭",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Manual": "ম্যানুয়াল",
		"Mechanical": "মেকানিক্যাল",
		"No": "না",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"R600a": "R600a",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes": "হ্যাঁ",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/2": "হ্যাঁ/২",
	
		"594": "৫৯৪",
		"708": "৭০৮",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c2x-gdel-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeC2xGdelXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c2x-gdel-xx-inverter"
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
		"Model Name":          "MFE-C2X-GDEL-XX-INVERTER",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N ~ ST",
		"Compressor Input Power (Watt)": "V 0601 - (38~109)V 0701 - (33~126)V 0702 - (33~126)",
		"Compressor Type": "V 0601 - BLDCV 0701 - BLDCV 0702 - BLDC",
		"Cooling Efect": "Freezer Cabinet Less than -18 ̊CRefrigerator Cabinet 0 ̊C to +5 ̊C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "708",
		"Door Basket": "PVC/4",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "341 Ltr",
		"Gross Weight": "72 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height/mm": "1720",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "77/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "320 Ltr.",
		"Net Weight": "67 ± 2 Kg",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Rated Voltage/ Hz": "220-240V~/50Hz",
		"Recommended voltage stabilizer capacity": "V 0601 - No NeedV 0701 - No NeedV 0702 - No Need",
		"Refrigerant": "V 0601 - R600aV 0701 - R600aV 0702 - R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Shelf (Material/No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
		"Width/mm": "594",
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
