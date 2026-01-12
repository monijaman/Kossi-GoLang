package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC1bGdenXxInverter seeds specifications/options for product 'marcel-mfe-c1b-gden-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeC1bGdenXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC1bGdenXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC1bGdenXxInverter() *SpecificationSeederRefrigeratorMarcelMfeC1bGdenXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeC1bGdenXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c1b-gden-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdenXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfe-c1b-gden-xx-inverter": "মার্সেল-mfe-c1b-gden-xx-inverter",
		"MFE-C1B-GDEN-XX-INVERTER":        "MFE-C1B-GDEN-XX-INVERTER",
		// basic values
		"312 Ltr.":           "৩১২ লিটার",
		"290 Ltr.":           "২৯০ লিটার",
		"59.22 ± 2 Kg":       "৫৯.২২ ± ২ কেজি",
		"66.06 ± 2 Kg":       "৬৬.০৬ ± ২ কেজি",
		"Direct Cool":        "ডাইরেক্ট কুল",
		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",
		"N ~ ST":             "N ~ ST",

		// dimensions & packaging
		"594":        "৫৯৪",
		"708":        "৭০৮",
		"1620":       "১৬২০",
		"625":        "৬২৫",
		"745":        "৭৪৫",
		"1630":       "১৬৩০",
		"79/ 54/ 27": "৭৯/ ৫৪/ ২৭",

		// components / features
		"V 0101 - RSIR V 0201 - RSIR V 0301 - RSIR V 0302 - RSIR V 0501 - RSIR":      "V ০১০১ - RSIR V ০২০১ - RSIR V ০৩০১ - RSIR V ০৩০২ - RSIR V ০৫০১ - RSIR",
		"V 0101 - R600a V 0201 - R600a V 0301 - R600a V 0302 - R600a V 0501 - R600a": "V ০১০১ - R৬০০a V ০২০১ - R৬০০a V ০৩০১ - R৬০০a V ০৩০২ - R৬০০a V ০৫০১ - R৬০০a",
		"V 0101 -145.7 V 0201 - 145.7 V 0301 - 117 V 0302 - 117 V 0501 - 123":        "V ০১০১ -১৪৫.৭ V ০২০১ - ১৪৫.৭ V ০৩০১ - ১১৭ V ০৩০২ - ১১৭ V ০৫০১ - ১২৩",

		"Wire/2": "Wire/২",
		"PVC/3":  "PVC/৩",
		"Yes":    "হ্যাঁ",
		"No":     "না",
		"Manual": "ম্যানুয়াল",

		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Freezer Cabinet Less than -18 ̊C Refrigerator Cabinet 0 ̊C to +5 ̊C": "Freezer Cabinet Less than -১৮ ̊C Refrigerator Cabinet ০ ̊C to +৫ ̊C",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c1b-gden-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdenXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c1b-gden-xx-inverter"
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
		"Gross Volume":        "312 Ltr.",
		"Net Volume":          "290 Ltr.",
		"Weight":              "59.22 ± 2 Kg",
		"Voltage":             "220-240V~ and 50Hz",
		"Compressor Type":     "V 0101 - RSIR V 0201 - RSIR V 0301 - RSIR V 0302 - RSIR V 0501 - RSIR",
		"Refrigerant":         "V 0101 - R600a V 0201 - R600a V 0301 - R600a V 0302 - R600a V 0501 - R600a",
		"Cooling Technology":  "Freezer Cabinet Less than -18 ̊C Refrigerator Cabinet 0 ̊C to +5 ̊C",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "PVC/3",
		"Crisper Drawers":     "Yes (Plastic)",
		"Dimensions":          "594 x 708 x 1620 mm",
		"Special Features":    "Net Weight: 59.22 ± 2 Kg; Gross Weight: 66.06 ± 2 Kg; Compressor Input Power (Watt): V 0101 -145.7 V 0201 -145.7 V 0301 -117 V 0302 -117 V 0501 -123; Recommended voltage stabilizer capacity: V 0201/0202/0203/0301: Wide Voltage Design (145V-253V) (suggested stabilizer 2100VA if out of range); V 0501: Wide Voltage Design (75V-264V) (suggested stabilizer 2100VA if out of range); Reversible Door: No; Handle: Recressed/ Grip/ Built-in; Lock: Yes; Capillary: Copper; Thermostat: RoHS Certified; Polyurethane foam blowing agent: CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Interior Lamp: Yes; Vegetable Crisper Cover: Yes (ABS/ PS); Egg Tray or Pocket: Yes; Can Storage Dispenser: No; Deodorizer: No; Freezer Drawer: No; Freezer Door Basket: No; Packaging Dimensions: 625 x 745 x 1630 mm; Loading Capacity- 40HQ/40Ft/20Ft: 79/54/27",
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
