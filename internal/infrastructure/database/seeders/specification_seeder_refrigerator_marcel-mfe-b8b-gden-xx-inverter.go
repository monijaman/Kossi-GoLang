package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter seeds specifications/options for product 'marcel-mfe-b8b-gden-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter() *SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-b8b-gden-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":         "মার্সেল",
				"marcel-mfe-b8b-gden-xx-inverter":         "মার্সেল-mfe-b8b-gden-xx-inverter",
		"MFE-B8B-GDEN-XX-INVERTER":   "MFE-B8B-GDEN-XX-INVERTER",
		// Add more translations as needed
		"66 ± 2 Kg": "৬৬ ± ২ কেজি",
		"1590": "১৫৯০",
		"V 0501 - Mechanical": "V ০৫০১ - Mechanical",
		"635": "৬৩৫",
		"80/ 80/ 38": "৮০/ ৮০/ ৩৮",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes": "হ্যাঁ",
		"740": "৭৪০",
		"V 0501 - R600a": "V ০৫০১ - R৬০০a",
		"Direct Cool": "ডাইরেক্ট কুল",
		"265 Ltr.": "২৬৫ লিটার",
		"220-240V~/50Hz": "২২০-২৪০V~/৫০Hz",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Copper": "কপার",
		"Manual": "Manual",
		"RoHS Certified": "RoHS Certified",
		"No": "না",
		"Recressed/ Grip/ Built-in": "Recressed/ Grip/ Built-in",
		"N ~ ST": "N ~ ST",
		"V 0501 - 42~145": "V ০৫০১ - ৪২~১৪৫",
		"Yes (Plastic)": "Yes (Plastic)",
		"V 0501 - BLDC": "V ০৫০১ - BLDC",
		"V0501/V0502/V0503 : Wide Voltage Design (110V-270V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.": "V০৫০১/V০৫০২/V০৫০৩ : Wide Voltage Design (১১০V-২৭০V) N.B.: If out of voltage range(৭৫V-২৬৪V) then suggested voltage stabilizer capacity is ২১০০VA.",
		"Wire/2": "Wire/২",
		"Yes (ABS/ PS)": "Yes (ABS/ PS)",
		"59 ± 2 Kg": "৫৯ ± ২ কেজি",
		"282 Ltr.": "২৮২ লিটার",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-b8b-gden-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-b8b-gden-xx-inverter"
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
		"Gross Volume": "282 Ltr.",
		"Net Volume": "265 Ltr.",
		"Net Weight": "59 ± 2 Kg",
		"Gross Weight": "66 ± 2 Kg",
		"Climate Type (SN, N, ST, T)": "N ~ ST",
		"Rated Operating Voltage and Frequency": "220-240V~/50Hz",
		"Compressor Input Power (Watt)": "V 0501 - 42~145",
		"Compressor Type": "V 0501 - BLDC",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/ Mechanical)": "V 0501 - Mechanical",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Reversible Door": "No",
		"Handle (Recessed/ Grip)": "Recressed/ Grip/ Built-in",
		"Lock": "Yes",
		"Refrigerant": "V 0501 - R600a",
		"Thermostat": "RoHS Certified",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "V0501/V0502/V0503 : Wide Voltage Design (110V-270V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Shelf (Material/ No.)": "Wire/2",
		"Door Basket": "No",
		"Interior Lamp": "No",
		"Vegetable Crisper": "Yes (Plastic)",
		"Vegetable Crisper Cover": "Yes (ABS/ PS)",
		"Egg Tray or Pocket": "Yes",
		"Can Storage Dispenser": "No",
		"Deodorizer": "No",
		"Drawer": "No",
		"Width/mm": "635",
		"Depth/mm": "740",
		"Height/mm": "1590",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "80/ 80/ 38",
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
