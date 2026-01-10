package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB9eGdxxXx seeds specifications/options for product 'marcel-mfe-b9e-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfeB9eGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB9eGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB9eGdxxXx() *SpecificationSeederRefrigeratorMarcelMfeB9eGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeB9eGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-b9e-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB9eGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-b9e-gdxx-xx": "মার্সেল-mfe-b9e-gdxx-xx",
		"MFE-B9E-GDXX-XX":        "MFE-B9E-GDXX-XX",
		"Direct Cool":            "ডাইরেক্ট কুল",
		"316 Ltr.":               "৩১৬ লিটার",
		"295 Ltr.":               "২৯৫ লিটার",
		"62 ± 2 Kg":              "৬২ ± ২ কেজি",
		"220-240V~ and 50Hz":     "২২০-২৪০ভি~ এবং ৫০হার্টজ",
		"V 0101 - RSCR; V 0201 - RSCR; V 0202 - RSCR; V 0301 - RSCR; V 0501 - RSCR": "V 0101 - RSCR; V 0201 - RSCR; V 0202 - RSCR; V 0301 - RSCR; V 0501 - RSCR",
		"V0102 - R600a; V0201 - R600a; V0202 - R600a; V0301 - R134a; V0501 - R600a": "V0102 - R600a; V0201 - R600a; V0202 - R600a; V0301 - R134a; V0501 - R600a",
		"Mechanical":          "যান্ত্রিক",
		"Manual":              "ম্যানুয়াল",
		"Wire/2":              "ওয়্যার/2",
		"PVC/4":               "PVC/4",
		"Yes (Plastic)":       "হ্যাঁ (প্লাস্টিক)",
		"594 x 708 x 1620 mm": "৫৯৪ x ৭০৮ x ১৬২০ মিমি",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-b9e-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeB9eGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-b9e-gdxx-xx"
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
		"Model Name":          "MFE-B9E-GDXX-XX",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "316 Ltr.",
		"Net Volume":          "295 Ltr.",
		"Weight":              "62 ± 2 Kg",
		"Voltage":             "220-240V~ and 50Hz",
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
