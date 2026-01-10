package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC1bGdenDd seeds specifications/options for product 'marcel-mfe-c1b-gden-dd'
type SpecificationSeederRefrigeratorMarcelMfeC1bGdenDd struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC1bGdenDd creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC1bGdenDd() *SpecificationSeederRefrigeratorMarcelMfeC1bGdenDd {
	return &SpecificationSeederRefrigeratorMarcelMfeC1bGdenDd{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c1b-gden-dd"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdenDd) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-c1b-gden-dd": "মার্সেল-mfe-c1b-gden-dd",
		"MFE-C1B-GDEN-DD":        "MFE-C1B-GDEN-DD",
		// Add more translations as needed
		"V 0401 - Electrical": "V ০৪০১ - Electrical",
		"635":                 "৬৩৫",
		"Yes (ABS/ PS)":       "Yes (ABS/ PS)",
		"78/ 57/ 27":          "৭৮/ ৫৭/ ২৭",
		"Yes":                 "হ্যাঁ",
		"220-240V~ and 50Hz":  "২২০-২৪০V~ and ৫০Hz",
		"740":                 "৭৪০",
		"312 Ltr.":            "৩১২ লিটার",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"72 ± 2 Kg":           "৭২ ± ২ কেজি",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Copper":                    "কপার",
		"Manual":                    "Manual",
		"V 0401 - BLDC":             "V ০৪০১ - BLDC",
		"No":                        "না",
		"V 0401 - 43.5~143.7":       "V ০৪০১ - ৪৩.৫~১৪৩.৭",
		"1690":                      "১৬৯০",
		"Recressed/ Grip/ Built-in": "Recressed/ Grip/ Built-in",
		"N ~ ST":                    "N ~ ST",
		"Yes (Plastic)":             "Yes (Plastic)",
		"V 0401 - R600a":            "V ০৪০১ - R৬০০a",
		"V 0401:Wide Voltage Design (105V-185V) N.B.: If out of voltage range(105V-185V) then suggested voltage stabilizer capacity is 2100VA.": "V ০৪০১:Wide Voltage Design (১০৫V-১৮৫V) N.B.: If out of voltage range(১০৫V-১৮৫V) then suggested voltage stabilizer capacity is ২১০০VA.",
		"Wire/2": "Wire/২",
		"Freezer Cabinet Less than -18 ̊C Refrigerator Cabinet 0 ̊C to +5 ̊C": "Freezer Cabinet Less than -১৮ ̊C Refrigerator Cabinet ০ ̊C to +৫ ̊C",
		"65 ± 2 Kg": "৬৫ ± ২ কেজি",
		"290 Ltr.":  "২৯০ লিটার",
	
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Mechanical": "মেকানিক্যাল",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"R600a": "R600a",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"1620": "১৬২০",
		"594": "৫৯৪",
		"708": "৭০৮",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c1b-gden-dd'
func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdenDd) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c1b-gden-dd"
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
		"Model Name":          "MFE-C1B-GDEN-DD",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N ~ ST",
		"Compressor Input Power (Watt)": "V 0401 - 43.5~143.7",
		"Compressor Type": "V 0401 - BLDC",
		"Cooling Effect": "Freezer Cabinet Less than -18 ̊CRefrigerator Cabinet 0 ̊C to +5 ̊C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth (mm)": "708",
		"Door Basket": "PVC/3",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "312 Ltr.",
		"Gross Weight": "72 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recressed/ Grip/ Built-in",
		"Height (mm)": "1620",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "78/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "290 Ltr.",
		"Net Weight": "65 ± 2 Kg",
		"Polyurethane Foam Blowing Agent": "CycloPentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Recommended voltage stabilizer capacity": "V 0401:Wide Voltage Design (105V-185V)N.B.: If out of voltage range(105V-185V)  then suggested voltage stabilizer capacity is 2100VA.",
		"Refrigerant": "V 0401 - R600a",
		"Reversible Door": "No",
		"Shelf Material": "Wire/2",
		"Temperature Control": "V 0401 - Electrical",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes (Plastic)",
		"Vegetable Crisper Cover": "Yes (ABS/ PS)",
		"Width": "594",
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
