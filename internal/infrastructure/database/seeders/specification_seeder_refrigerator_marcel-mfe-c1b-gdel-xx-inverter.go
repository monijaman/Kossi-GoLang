package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC1bGdelXxInverter seeds specifications/options for product 'marcel-mfe-c1b-gdel-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeC1bGdelXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC1bGdelXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC1bGdelXxInverter() *SpecificationSeederRefrigeratorMarcelMfeC1bGdelXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeC1bGdelXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c1b-gdel-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdelXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfe-c1b-gdel-xx-inverter": "মার্সেল-mfe-c1b-gdel-xx-inverter",
		"MFE-C1B-GDEL-XX-INVERTER":        "MFE-C1B-GDEL-XX-INVERTER",
		// Add more translations as needed
		"V 0101-145.7 V 0201-145.7 V 0301-117 V 0302-117 V 0501-123": "V ০১০১-১৪৫.৭ V ০২০১-১৪৫.৭ V ০৩০১-১১৭ V ০৩০২-১১৭ V ০৫০১-১২৩",
		"67.16 ± 2 Kg":  "৬৭.১৬ ± ২ কেজি",
		"60.14 ± 2 Kg":  "৬০.১৪ ± ২ কেজি",
		"635":           "৬৩৫",
		"Yes (ABS/ PS)": "Yes (ABS/ PS)",
		"V 0201/0202/0203/0301: Wide Voltage Design (145V-253V) N.B.: If out of voltage range (145V-253V) then suggested voltage stabilizer capacity is 2100VA. V 0501:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.": "V ০২০১/০২০২/০২০৩/০৩০১: Wide Voltage Design (১৪৫V-২৫৩V) N.B.: If out of voltage range (১৪৫V-২৫৩V) then suggested voltage stabilizer capacity is ২১০০VA. V ০৫০১:Wide Voltage Design (৭৫V-২৬৪V) N.B.: If out of voltage range(৭৫V-২৬৪V) then suggested voltage stabilizer capacity is ২১০০VA.",
		"Yes":         "হ্যাঁ",
		"740":         "৭৪০",
		"312 Ltr.":    "৩১২ লিটার",
		"Direct Cool": "ডাইরেক্ট কুল",
		"79/ 54/ 27":  "৭৯/ ৫৪/ ২৭",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Copper": "কপার",
		"V 0101- RSIR V 0201- RSIR V 0301- RSIR V 0302- RSIR V 0501- RSIR": "V ০১০১- RSIR V ০২০১- RSIR V ০৩০১- RSIR V ০৩০২- RSIR V ০৫০১- RSIR",
		"Manual":         "Manual",
		"RoHS Certified": "RoHS Certified",
		"No":             "না",
		"V0101- R600a V0201- R600a V0301- R600a V0302- R600a V 0501 -R600a": "V০১০১- R৬০০a V০২০১- R৬০০a V০৩০১- R৬০০a V০৩০২- R৬০০a V ০৫০১ -R৬০০a",
		"1690":                      "১৬৯০",
		"Recressed/ Grip/ Built-in": "Recressed/ Grip/ Built-in",
		"N ~ ST":                    "N ~ ST",
		"Yes (Plastic)":             "Yes (Plastic)",
		"V 0201 - Mechanical V 0301 - Mechanical V 0302 - Mechanical V 0501 - Mechanical": "V ০২০১ - Mechanical V ০৩০১ - Mechanical V ০৩০২ - Mechanical V ০৫০১ - Mechanical",
		"220-240V ~ and 50Hz": "২২০-২৪০V ~ and ৫০Hz",
		"Wire/2":              "Wire/২",
		"Freezer Cabinet Less than -18 ̊C Refrigerator Cabinet 0 ̊C to +5 ̊C": "Freezer Cabinet Less than -১৮ ̊C Refrigerator Cabinet ০ ̊C to +৫ ̊C",
		"290 Ltr.": "২৯০ লিটার",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
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
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"1620": "১৬২০",
		"594": "৫৯৪",
		"708": "৭০৮",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c1b-gdel-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdelXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c1b-gdel-xx-inverter"
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
		"Model Name":          "MFE-C1B-GDEL-XX-INVERTER",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 0101-145.7V 0201-145.7V 0301-117V 0302-117V 0501-123",
		"Compressor Type": "V 0101- RSIRV 0201- RSIRV 0301- RSIRV 0302- RSIRV 0501- RSIR",
		"Cooling Efect": "Freezer Cabinet Less than -18 ̊CRefrigerator Cabinet 0 ̊C to +5 ̊C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "708",
		"Door Basket": "PVC/3",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "312 Ltr.",
		"Gross Weight": "67.16 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recressed/ Grip/ Built-in",
		"Height/mm": "1620",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "79/ 54/ 27",
		"Lock": "Yes",
		"Net Volume": "290 Ltr.",
		"Net Weight": "60.14 ± 2 Kg",
		"Polyurethane foam blowing agent": "CycloPentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~and 50Hz",
		"Recommended voltage stabilizer capacity": "V 0201/0202/0203/0301: Wide Voltage Design (145V-253V)N.B.: If out of voltage range (145V-253V) then suggested voltage stabilizer capacity is 2100VA.V 0501:Wide Voltage Design (75V-264V)N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Refrigerant": "V0101- R600aV0201- R600aV0301- R600aV0302- R600aV 0501 -R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "V 0201 - MechanicalV 0301 - MechanicalV 0302 - MechanicalV 0501 - Mechanical",
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
