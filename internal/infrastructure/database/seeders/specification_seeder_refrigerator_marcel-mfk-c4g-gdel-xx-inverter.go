package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfkC4gGdelXxInverter seeds specifications/options for product 'marcel-mfk-c4g-gdel-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfkC4gGdelXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfkC4gGdelXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfkC4gGdelXxInverter() *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfkC4gGdelXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfk-c4g-gdel-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":         "মার্সেল",
				"marcel-mfk-c4g-gdel-xx-inverter":         "মার্সেল-mfk-c4g-gdel-xx-inverter",
		"MFK-C4G-GDEL-XX-INVERTER":   "MFK-C4G-GDEL-XX-INVERTER",
		// Add more translations as needed
		"Mechanical": "যান্ত্রিক",
		"715": "৭১৫",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes": "হ্যাঁ",
		"347 Ltr": "৩৪৭ Ltr",
		"1645": "১৬৪৫",
		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",
		"345 Ltr.": "৩৪৫ লিটার",
		"Recessed/ Grip": "Recessed/ Grip",
		"710": "৭১০",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Yes/2": "Yes/২",
		"Copper": "কপার",
		"V 01.01- 147": "V ০১.০১- ১৪৭",
		"Manual": "Manual",
		"No": "না",
		"GPPS/2": "GPPS/২",
		"N~ST": "N~ST",
		"V01.01 R600a": "V০১.০১ R৬০০a",
		"V.01.01 Wide Voltage Range (145V ~ 260V) NB: If out of voltage range then suggested voltage stabilizer capacity is 2100VA.": "V.০১.০১ Wide Voltage Range (১৪৫V ~ ২৬০V) NB: If out of voltage range then suggested voltage stabilizer capacity is ২১০০VA.",
		"Direct Evaporating Cooling System (DECS)": "Direct Evaporating Cooling System (DECS)",
		"Wire/2": "Wire/২",
		"Wire/ 3": "Wire/ ৩",
		"V.01.01- RSCR": "V.০১.০১- RSCR",
		"67.5 / 74 ± 2": "৬৭.৫ / ৭৪ ± ২",
		"69/69/34": "৬৯/৬৯/৩৪",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfk-c4g-gdel-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfk-c4g-gdel-xx-inverter"
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
		"Type": "Direct Evaporating Cooling System (DECS)",
		"Gross Volume": "347 Ltr",
		"Net Volume": "345 Ltr.",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Compressor Input Power (Watt)": "V 01.01- 147",
		"Compressor Type": "V.01.01- RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Lock": "Yes",
		"Refrigerant": "V01.01 R600a",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Operating voltage": "V.01.01 Wide Voltage Range (145V ~ 260V) NB: If out of voltage range then suggested voltage stabilizer capacity is 2100VA.",
		"Shelf (Material/ No.)": "Wire/2",
		"Door Basket": "GPPS/2",
		"Interior Lamp": "Yes",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Egg Tray": "Yes/2",
		"Rack Shelf (Material/No)": "Wire/ 3",
		"Drawer": "No",
		"Width/mm": "715",
		"Depth/mm": "710",
		"Height/mm": "1645",
		"Weight/Kg - Net/Packing": "67.5 / 74 ± 2",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "69/69/34",
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
