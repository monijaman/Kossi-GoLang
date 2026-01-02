package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB9eGdenXxInverter seeds specifications/options for product 'marcel-mfe-b9e-gden-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeB9eGdenXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB9eGdenXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB9eGdenXxInverter() *SpecificationSeederRefrigeratorMarcelMfeB9eGdenXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeB9eGdenXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-b9e-gden-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB9eGdenXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfe-b9e-gden-xx-inverter": "মার্সেল-mfe-b9e-gden-xx-inverter",
		"MFE-B9E-GDEN-XX-INVERTER":        "MFE-B9E-GDEN-XX-INVERTER",
		// Add more translations as needed
		"Mechanical": "যান্ত্রিক",
		"635":        "৬৩৫",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"316 Ltr.":       "৩১৬ লিটার",
		"Yes":            "হ্যাঁ",
		"78/ 57/ 27":     "৭৮/ ৫৭/ ২৭",
		"740":            "৭৪০",
		"Direct Cool":    "ডাইরেক্ট কুল",
		"220-240V~/50Hz": "২২০-২৪০V~/৫০Hz",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Copper":                    "কপার",
		"70 ± 2 Kg":                 "৭০ ± ২ কেজি",
		"Manual":                    "Manual",
		"RoHS Certified":            "RoHS Certified",
		"No":                        "না",
		"295 Ltr.":                  "২৯৫ লিটার",
		"1690":                      "১৬৯০",
		"Recressed/ Grip/ Built-in": "Recressed/ Grip/ Built-in",
		"N ~ ST":                    "N ~ ST",
		"Yes (Plastic)":             "Yes (Plastic)",
		"V 0501:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.": "V ০৫০১:Wide Voltage Design (৭৫V-২৬৪V) N.B.: If out of voltage range(৭৫V-২৬৪V) then suggested voltage stabilizer capacity is ২১০০VA.",
		"Wire/2":        "Wire/২",
		"Yes (ABS/ PS)": "Yes (ABS/ PS)",
		"62 ± 2 Kg":     "৬২ ± ২ কেজি",
		"V 0101 - RSCR; V 0201 - RSCR; V 0202 - RSCR; V 0301 - RSCR; V 0501 - RSCR": "V ০১০১ - RSCR; V ০২০১ - RSCR; V ০২০২ - RSCR; V ০৩০১ - RSCR; V ০৫০১ - RSCR",
		"V0102 - R600a; V0201 - R600a; V0202 - R600a; V0301 - R134a; V0501 - R600a": "V০১০২ - R৬০০a; V০২০১ - R৬০০a; V০২০২ - R৬০০a; V০৩০১ - R১৩৪a; V০৫০১ - R৬০০a",
		"PVC/4":               "PVC/৪",
		"594 x 708 x 1620 mm": "৫৯৪ x ৭০৮ x ১৬২০ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-b9e-gden-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeB9eGdenXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-b9e-gden-xx-inverter"
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
		"Brand":               "Marcel",
		"Model Name":          "MFE-B9E-GDEN-XX-INVERTER",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "316 Ltr.",
		"Net Volume":          "295 Ltr.",
		"Weight":              "62 ± 2 Kg",
		"Voltage":             "220-240V~/50Hz",
		"Compressor Type":     "V 0101 - RSCR; V 0201 - RSCR; V 0202 - RSCR; V 0301 - RSCR; V 0501 - RSCR",
		"Refrigerant":         "V0102 - R600a; V0201 - R600a; V0202 - R600a; V0301 - R134a; V0501 - R600a",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "PVC/4",
		"Crisper Drawers":     "Yes (Plastic)",
		"Dimensions":          "594 x 708 x 1620 mm",
		"Special Features":    "Gross Weight: 70 ± 2 Kg; Climate Type: N ~ ST; Compressor Input Power (Watt): V 0101 - 118; V 0201 - 118; V 0202 - 118; V 0301 - 134; V 0501 - 118; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Reversible Door: No; Handle (Recessed/ Grip): Recressed/ Grip/ Built-in; Lock: Yes; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Recommended voltage stabilizer capacity: V 0501:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.; Interior Lamp: Yes; Vegetable Crisper Cover: Yes (ABS/ PS); Egg Tray or Pocket: Yes; Can Storage Dispenser: No; Deodorizer: No; Freezer Shelf (Material/ No.): Wire/2; Freezer Drawer: No; Freezer Door Basket: No; Freezer Interior Lamp: No; Packing Dimensions: 635 x 740 x 1690 mm; Loading Capacity: 78/ 57/ 27",
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
