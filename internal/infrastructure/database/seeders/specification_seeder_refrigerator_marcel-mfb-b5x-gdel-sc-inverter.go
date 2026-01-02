package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5xGdelScInverter seeds specifications/options for product 'marcel-mfb-b5x-gdel-sc-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB5xGdelScInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5xGdelScInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5xGdelScInverter() *SpecificationSeederRefrigeratorMarcelMfbB5xGdelScInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB5xGdelScInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b5x-gdel-sc-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdelScInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":         "মার্সেল",
				"marcel-mfb-b5x-gdel-sc-inverter":         "মার্সেল-mfb-b5x-gdel-sc-inverter",
		"MFB-B5X-GDEL-SC-INVERTER":   "MFB-B5X-GDEL-SC-INVERTER",
		// Add more translations as needed
		"54.5 ± 2 Kg": "৫৪.৫ ± ২ কেজি",
		"1725": "১৭২৫",
		"V 01.01-97.4 V 02.01-97.4 V 03.01-97.4 V 03.02-97.4": "V ০১.০১-৯৭.৪ V ০২.০১-৯৭.৪ V ০৩.০১-৯৭.৪ V ০৩.০২-৯৭.৪",
		"97/ 72/ 36": "৯৭/ ৭২/ ৩৬",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes": "হ্যাঁ",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Electronic": "ইলেকট্রনিক",
		"50.5/55.5 ± 2": "৫০.৫/৫৫.৫ ± ২",
		"RSCR": "RSCR",
		"R600a": "R৬০০a",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"V 01.01, V 02.01-Low Voltage(140~260V) For V01.01, V 02.01-Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 03.01, V 03.02-Low Voltage(150~260V) For V 03.01, V 03.02 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.": "V ০১.০১, V ০২.০১-Low Voltage(১৪০~২৬০V) For V০১.০১, V ০২.০১-Wide Voltage Range (১৪০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended. V ০৩.০১, V ০৩.০২-Low Voltage(১৫০~২৬০V) For V ০৩.০১, V ০৩.০২ - Wide Voltage Range (১৫০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended.",
		"Copper": "কপার",
		"Manual": "Manual",
		"RoHS Certified": "RoHS Certified",
		"645": "৬৪৫",
		"No": "না",
		"N~ST": "N~ST",
		"250 Ltr.": "২৫০ লিটার",
		"244 Ltr.": "২৪৪ লিটার",
		"59.5 ± 2 Kg": "৫৯.৫ ± ২ কেজি",
		"Wire/2": "Wire/২",
		"580": "৫৮০",
		"Recessed/Grip": "Recessed/Grip",
		"GPPS/3": "GPPS/৩",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b5x-gdel-sc-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdelScInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b5x-gdel-sc-inverter"
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
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Dimensions":                  25,
		"Weight":                      80,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
		"Special Features":            69,
	}

	
    		
    specs := map[string]string{
		"Type": "Direct Cool",
		"Gross Volume": "250 Ltr.",
		"Net Volume": "244 Ltr.",
		"Net Weight": "54.5 ± 2 Kg",
		"Gross Weight": "59.5 ± 2 Kg",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Compressor Input Power (Watt)": "V 01.01-97.4 V 02.01-97.4 V 03.01-97.4 V 03.02-97.4",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/ Mechanical)": "Electronic",
		"Defrosting (Automatic/ Manual):": "Manual",
		"Handle (Recessed/ Grip)": "Recessed/Grip",
		"Lock": "Yes",
		"Refrigerant": "R600a",
		"Thermostat": "RoHS Certified",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "V 01.01, V 02.01-Low Voltage(140~260V) For V01.01, V 02.01-Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 03.01, V 03.02-Low Voltage(150~260V) For V 03.01, V 03.02 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.",
		"Shelf (Material/No.)": "Wire/2",
		"Door Basket": "GPPS/3",
		"Interior Lamp": "Yes",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Egg Tray": "Yes",
		"Shelf (Material/ No.)": "Wire/2",
		"Drawer": "No",
		"Width/mm": "580",
		"Depth/mm": "645",
		"Height/mm": "1725",
		"Weight/Kg - Net/Packing": "50.5/55.5 ± 2",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "97/ 72/ 36",
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
