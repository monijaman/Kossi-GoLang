package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB9eCrxxXxInverter seeds specifications/options for product 'marcel-mfe-b9e-crxx-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeB9eCrxxXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB9eCrxxXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB9eCrxxXxInverter() *SpecificationSeederRefrigeratorMarcelMfeB9eCrxxXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeB9eCrxxXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-b9e-crxx-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB9eCrxxXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfe-b9e-crxx-xx-inverter": "মার্সেল-mfe-b9e-crxx-xx-inverter",
		"MFE-B9E-CRXX-XX-INVERTER":        "MFE-B9E-CRXX-XX-INVERTER",
		// Add more translations as needed
		"Mechanical": "যান্ত্রিক",
		"635":        "৬৩৫",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"316 Ltr.":      "৩১৬ লিটার",
		"Yes":           "হ্যাঁ",
		"78/ 57/ 27":    "৭৮/ ৫৭/ ২৭",
		"740":           "৭৪০",
		"Direct Cool":   "ডাইরেক্ট কুল",
		"V0401 - R600a": "V০৪০১ - R৬০০a",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Copper":                    "কপার",
		"Manual":                    "Manual",
		"RoHS Certified":            "RoHS Certified",
		"57 ± 2 Kg":                 "৫৭ ± ২ কেজি",
		"V 0401 - BLDC":             "V ০৪০১ - BLDC",
		"No":                        "না",
		"295 Ltr.":                  "২৯৫ লিটার",
		"NO":                        "NO",
		"1690":                      "১৬৯০",
		"Recressed/ Grip/ Built-in": "Recressed/ Grip/ Built-in",
		"N ~ ST":                    "N ~ ST",
		"Yes (Plastic)":             "Yes (Plastic)",
		"220-240V ~ and 50Hz":       "২২০-২৪০V ~ and ৫০Hz",
		"V 0401:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.": "V ০৪০১:Wide Voltage Design (৭৫V-২৬৪V) N.B.: If out of voltage range(৭৫V-২৬৪V) then suggested voltage stabilizer capacity is ২১০০VA.",
		"Wire/2":          "Wire/২",
		"V 0401 - 33~126": "V ০৪০১ - ৩৩~১২৬",
		"65 ± 2 Kg":       "৬৫ ± ২ কেজি",
		"Yes (ABS/ PS)":   "Yes (ABS/ PS)",
	
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
	
		"594": "৫৯৪",
		"PVC/4": "PVC/৪",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-b9e-crxx-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeB9eCrxxXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-b9e-crxx-xx-inverter"
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
		"Model Name":          "MFE-B9E-CRXX-XX-INVERTER",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 0401 - 33~126",
		"Compressor Type": "V 0401 - BLDC",
		"Cooling Efect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "NO",
		"Depth (mm)": "711",
		"Door Basket": "PVC/4",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "316 Ltr.",
		"Gross Weight": "65 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recressed/ Grip/ Built-in",
		"Height (mm)": "1626",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "78/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "295 Ltr.",
		"Net Weight": "57 ± 2 Kg",
		"Polyurethane Foam Blowing Agent": "CycloPentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~and 50Hz",
		"Recommended voltage stabilizer capacity": "V 0401:Wide Voltage Design (75V-264V)N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Refrigerant": "V0401 - R600a",
		"Reversible Door": "No",
		"Shelf Material": "Wire/2",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
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
