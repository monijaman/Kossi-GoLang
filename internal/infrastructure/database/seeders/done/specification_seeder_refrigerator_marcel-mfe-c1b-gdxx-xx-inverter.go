package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXxInverter seeds specifications/options for product 'marcel-mfe-c1b-gdxx-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC1bGdxxXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC1bGdxxXxInverter() *SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c1b-gdxx-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfe-c1b-gdxx-xx-inverter": "মার্সেল-mfe-c1b-gdxx-xx-inverter",
		"MFE-C1B-GDXX-XX-INVERTER":        "MFE-C1B-GDXX-XX-INVERTER",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"312 Ltr":                         "৩১২ লিটার",
		"290 Ltr.":                        "২৯০ লিটার",
		"59.22 ± 2 Kg":                    "৫৯.২২ ± ২ কেজি",
		"66.06 ± 2 Kg":                    "৬৬.০৬ ± ২ কেজি",
		"N ~ ST":                          "এন ~ এসটি",
		"220-240V~/50Hz":                  "২২০-২৪০V~/৫০Hz",
		"V 0101 - RSIR V 0201 - RSIR V 0301 - RSIR V 0302 - RSIR V 0501 - RSIR":      "V ০১০১ - RSIR V ০২০১ - RSIR V ০৩০১ - RSIR V ০৩০২ - RSIR V ০৫০১ - RSIR",
		"V 0101 - R600a V 0201 - R600a V 0301 - R600a V 0302 - R600a V 0501 - R600a": "V ০১০১ - R৬০০a V ০২০১ - R৬০০a V ০৩০১ - R৬০০a V ০৩০২ - R৬০০a V ০৫০১ - R৬০০a",
		"Mechanical":          "ম্যানুয়াল",
		"Manual":              "ম্যানুয়াল",
		"Wire":                "ওয়্যার",
		"2":                   "২",
		"PVC/3":               "পিভিসি/৩",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"Copper":              "কপার",
		"RoHS Certified":      "RoHS সার্টিফায়েড",
		"Cyclopentene":        "সাইক্লোপেন্টিন",
		"594 x 708 x 1620 mm": "৫৯৪ x ৭০৮ x ১৬২০ মিমি",
		"625 x 745 x 1630 mm": "৬২৫ x ৭৪৫ x ১৬৩০ মিমি",
		"79/ 54/ 27":          "৭৯/ ৫৪/ ২৭",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c1b-gdxx-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c1b-gdxx-xx-inverter"
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
		"Model Name":          "MFE-C1B-GDXX-XX-INVERTER",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "312 Ltr",
		"Net Volume":          "290 Ltr.",
		"Weight":              "59.22 ± 2 Kg",
		"Voltage":             "220-240V~/50Hz",
		"Compressor Type":     "V 0101 - RSIR V 0201 - RSIR V 0301 - RSIR V 0302 - RSIR V 0501 - RSIR",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Refrigerant":         "V 0101 - R600a V 0201 - R600a V 0301 - R600a V 0302 - R600a V 0501 - R600a",
		"Shelf Material":      "Wire",
		"Number of Shelves":   "2",
		"Door Bins":           "PVC/3",
		"Crisper Drawers":     "Yes",
		"Dimensions":          "594 x 708 x 1620 mm",
		"Special Features":    "Gross Weight: 66.06 ± 2 Kg; Compressor Input Power: V 0101 - 145.7 V 0201 - 145.7 V 0301 - 117 V 0302 - 117 V 0501 - 123; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene; Recommended stabilizer: V 0201/0202/0203/0301: Wide Voltage Design (145V-253V) (if out of range then suggested stabilizer 2100VA); V 0501: Wide Voltage Design (75V-264V) (if out of range then suggested stabilizer 2100VA); Packing Dimensions: 625 x 745 x 1630 mm; Loading Capacity: 79/ 54/ 27; Freezer Cooling: Less than -18 ̊C; Refrigerator Cooling: 0 ̊C to +5 ̊C; Interior Lamp: Yes; Vegetable Crisper: Yes (Plastic); Vegetable Crisper Cover: Yes (ABS/ PS); Egg Tray: Yes; Can Storage Dispenser: No; Deodorizer: No",
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
