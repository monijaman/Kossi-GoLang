package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB9eGdenXx seeds specifications/options for product 'marcel-mfe-b9e-gden-xx'
type SpecificationSeederRefrigeratorMarcelMfeB9eGdenXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB9eGdenXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB9eGdenXx() *SpecificationSeederRefrigeratorMarcelMfeB9eGdenXx {
	return &SpecificationSeederRefrigeratorMarcelMfeB9eGdenXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-b9e-gden-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB9eGdenXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-b9e-gden-xx": "মার্সেল-mfe-b9e-gden-xx",
		"MFE-B9E-GDEN-XX":        "MFE-B9E-GDEN-XX",
		// Add more translations as needed
		"Mechanical": "যান্ত্রিক",
		"635":        "৬৩৫",
		"V 0101 - RSCR V 0201 - RSCR V 0202 - RSCR V 0301 - RSCR V 0501 - RSCR": "V ০১০১ - RSCR V ০২০১ - RSCR V ০২০২ - RSCR V ০৩০১ - RSCR V ০৫০১ - RSCR",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃":         "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes":            "হ্যাঁ",
		"316 Ltr.":       "৩১৬ লিটার",
		"78/ 57/ 27":     "৭৮/ ৫৭/ ২৭",
		"740":            "৭৪০",
		"Direct Cool":    "ডাইরেক্ট কুল",
		"220-240V~/50Hz": "২২০-২৪০V~/৫০Hz",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"V 0101 - 118 V 0201 - 118 V 0202 - 118 V 0301 - 134 V 0501 - 118":    "V ০১০১ - ১১৮ V ০২০১ - ১১৮ V ০২০২ - ১১৮ V ০৩০১ - ১৩৪ V ০৫০১ - ১১৮",
		"70 ± 2 Kg":      "৭০ ± ২ কেজি",
		"Manual":         "Manual",
		"Copper":         "কপার",
		"RoHS Certified": "RoHS Certified",
		"No":             "না",
		"295 Ltr.":       "২৯৫ লিটার",
		"V0102 - R600a V0201 - R600a V0202 - R600a V0301 - R134a V0501 - R600a": "V০১০২ - R৬০০a V০২০১ - R৬০০a V০২০২ - R৬০০a V০৩০১ - R১৩৪a V০৫০১ - R৬০০a",
		"1690": "১৬৯০",
		"V 0501:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.": "V ০৫০১:Wide Voltage Design (৭৫V-২৬৪V) N.B.: If out of voltage range(৭৫V-২৬৪V) then suggested voltage stabilizer capacity is ২১০০VA.",
		"Recressed/ Grip/ Built-in": "Recressed/ Grip/ Built-in",
		"N ~ ST":                    "N ~ ST",
		"Yes (Plastic)":             "Yes (Plastic)",
		"Wire/2":                    "Wire/২",
		"Wire":                      "ওয়্যার",
		"2":                         "২",
		"Yes (ABS/ PS)":             "Yes (ABS/ PS)",
		"62 ± 2 Kg":                 "৬২ ± ২ কেজি",
		"V 0101 - RSCR; V 0201 - RSCR; V 0202 - RSCR; V 0301 - RSCR; V 0501 - RSCR": "V ০১০১ - RSCR; V ০২০১ - RSCR; V ০২০২ - RSCR; V ০৩০১ - RSCR; V ০৫০১ - RSCR",
		"V0102 - R600a; V0201 - R600a; V0202 - R600a; V0301 - R134a; V0501 - R600a": "V০১০২ - R৬০০a; V০২০১ - R৬০০a; V০২০২ - R৬০০a; V০৩০১ - R১৩৪a; V০৫০১ - R৬০০a",
		"PVC/4":               "PVC/৪",
		"594 x 708 x 1620 mm": "৫৯৪ x ৭০৮ x ১৬২০ মিমি",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"R600a": "R600a",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"1620": "১৬২০",
		"594": "৫৯৪",
		"708": "৭০৮",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-b9e-gden-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeB9eGdenXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-b9e-gden-xx"
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
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N ~ ST",
		"Compressor Input Power (Watt)": "V 0101 - 118V 0201 - 118V 0202 - 118V 0301 - 134V 0501 - 118",
		"Compressor Type": "V 0101 - RSCRV 0201 - RSCRV 0202 - RSCRV 0301 - RSCRV 0501 - RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "708",
		"Door Basket": "PVC/4",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "316 Ltr.",
		"Gross Weight": "70 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recressed/ Grip/ Built-in",
		"Height/mm": "1620",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "78/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "295 Ltr.",
		"Net Weight": "62 ± 2 Kg",
		"Polyurethane foam blowing agent": "CycloPentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~/50Hz",
		"Recommended voltage stabilizer capacity": "V 0501:Wide Voltage Design (75V-264V)N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Refrigerant": "V0102 - R600aV0201 - R600aV0202 - R600aV0301 - R134aV0501 - R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes (Plastic)",
		"Vegetable Crisper Cover": "Yes (ABS/ PS)",
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
