package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB9eElnxXx seeds specifications/options for product 'marcel-mfe-b9e-elnx-xx'
type SpecificationSeederRefrigeratorMarcelMfeB9eElnxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB9eElnxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB9eElnxXx() *SpecificationSeederRefrigeratorMarcelMfeB9eElnxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeB9eElnxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-b9e-elnx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB9eElnxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-b9e-elnx-xx": "মার্সেল-mfe-b9e-elnx-xx",
		"MFE-B9E-ELNX-XX":        "MFE-B9E-ELNX-XX",
		// Add more translations as needed
		"Direct Cool":                    "ডাইরেক্ট কুল",
		"316 Ltr.":                       "৩১৬ লিটার",
		"295 Ltr.":                       "২৯৫ লিটার",
		"57 ± 2 Kg":                      "৫৭ ± ২ কেজি",
		"220-240V~/50Hz":                 "২২০-২৪০V~/৫০Hz",
		"V 0101 - RSCR; V 0201 - RSCR":   "V ০১০১ - RSCR; V ০২০১ - RSCR",
		"V 0101 - R600a; V 0201 - R600a": "V ০১০১ - R৬০০a; V ০২০১ - R৬০০a",
		"Mechanical":                     "যান্ত্রিক",
		"Manual":                         "Manual",
		"Wire/2":                         "Wire/২",
		"PVC/4":                          "PVC/৪",
		"Yes (Plastic)":                  "Yes (Plastic)",
		"594 x 711 x 1620 mm":            "৫৯৪ x ৭১১ x ১৬২০ মিমি",
		"70 ± 2 Kg":                      "৭০ ± ২ কেজি",
		"N ~ ST":                         "N ~ ST",
		"V 0101 - 118; V 0201 - 118":     "V ০১০১ - ১১৮; V ০২০১ - ১১৮",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"No":                        "না",
		"Yes":                       "হ্যাঁ",
		"Recressed/ Grip/ Built-in": "Recressed/ Grip/ Built-in",
		"RoHS Certified":            "RoHS Certified",
		"Copper":                    "কপার",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"2100VA":        "২১০০VA",
		"Yes (ABS/ PS)": "Yes (ABS/ PS)",
		"78/ 57/ 27":    "৭৮/ ৫৭/ ২৭",
		"625":           "৬২৫",
		"745":           "৭৪৫",
		"1630":          "১৬৩০",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-b9e-elnx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeB9eElnxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-b9e-elnx-xx"
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
		"Model Name":          "MFE-B9E-ELNX-XX",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "316 Ltr.",
		"Net Volume":          "295 Ltr.",
		"Weight":              "57 ± 2 Kg",
		"Voltage":             "220-240V~/50Hz",
		"Compressor Type":     "V 0101 - RSCR; V 0201 - RSCR",
		"Refrigerant":         "V 0101 - R600a; V 0201 - R600a",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "PVC/4",
		"Crisper Drawers":     "Yes (Plastic)",
		"Dimensions":          "594 x 711 x 1620 mm",
		"Special Features":    "Gross Weight: 70 ± 2 Kg; Climate Type: N ~ ST; Compressor Input Power (Watt): V 0101 - 118; V 0201 - 118; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Reversible Door: No; Handle (Recessed/ Grip): Recressed/ Grip/ Built-in; Lock: Yes; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Recommended voltage stabilizer capacity: 2100VA; Interior Lamp: Yes; Vegetable Crisper Cover: Yes (ABS/ PS); Egg Tray or Pocket: Yes; Can Storage Dispenser: No; Deodorizer: No; Freezer Shelf (Material/ No.): Wire/2; Freezer Drawer: No; Freezer Door Basket: No; Freezer Interior Lamp: No; Packing Dimensions: 625 x 745 x 1630 mm; Loading Capacity: 78/ 57/ 27",
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
