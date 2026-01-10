package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXxInverter seeds specifications/options for product 'marcel-mfe-c0n-gdxx-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC0nGdxxXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC0nGdxxXxInverter() *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c0n-gdxx-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfe-c0n-gdxx-xx-inverter": "মার্সেল-mfe-c0n-gdxx-xx-inverter",
		"MFE-C0N-GDXX-XX-INVERTER":        "MFE-C0N-GDXX-XX-INVERTER",
		// Add more translations as needed
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"Copper": "কপার",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[Eco-friendly (১০০% CFC &HCFC Free) Green Technology]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[Eco-friendly (১০০% CFC &HCFC Free) Green Technology]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "Freezer Cabinet Less than -১৮০CRefrigerator Cabinet ০০Cto +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃Refrigerator Cabinet ০℃ to +৫℃",
		"Manual": "ম্যানুয়াল",
		"Mechanical": "মেকানিক্যাল",
		"N ~ ST": "N ~ ST",
		"No": "না",
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
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes": "হ্যাঁ",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"290 Ltr.": "২৯০ লিটার",
		"312 Ltr.": "৩১২ লিটার",
		"585 x 711 x 1626 mm": "৫৮৫ x ৭১১ x ১৬২৬ মিমি",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c0n-gdxx-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c0n-gdxx-xx-inverter"
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
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}

	specs := map[string]string{
		"Brand":               "Marcel",
		"Model Name":          "MFE-C0N-GDXX-XX-INVERTER",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "312 Ltr.",
		"Net Volume":          "290 Ltr.",
		"Weight":              "59.4 Kg",
		"Voltage":             "220 ~ 240/ 50Hz",
		"Compressor Type":     "RSCR",
		"Refrigerant":         "R600a",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "PVC/3",
		"Crisper Drawers":     "Yes (Plastic)",
		"Dimensions":          "585 x 711 x 1626 mm",
		"Special Features":    "Gross Weight: 66.4 Kg; Climate Class: N ~ ST; Rated watt: 145.7; Reversible Door: No; Handle: Recressed/ Grip/ Built-in; Lock: Yes; Capillary: Copper; Polyurethane foam blowing agent: CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Interior Lamp: Yes; Vegetable Crisper Cover: Yes (ABS/ PS); Egg Tray or Pocket: Yes; Can Storage Dispenser: No; Deodorizer: No; Freezer Shelf: Wire/2; Freezer Drawer: No; Freezer Door Basket: No; Freezer Interior Lamp: No; Cooling Effect: Freezer Cabinet Less than -18 ̊C Refrigerator Cabinet 0 ̊C to +5 ̊C; Packaging Dimensions: 625 x 745 x 1630 mm; Loading Capacity- 40HQ/40Ft/20Ft: 78/57/27",
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
