package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC0nGdelXx seeds specifications/options for product 'marcel-mfe-c0n-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfeC0nGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC0nGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC0nGdelXx() *SpecificationSeederRefrigeratorMarcelMfeC0nGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC0nGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c0n-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC0nGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                       "মার্সেল",
		"marcel-mfe-c0n-gdel-xx":       "মার্সেল-mfe-c0n-gdel-xx",
		"MFE-C0N-GDEL-XX":              "MFE-C0N-GDEL-XX",
		"Direct Cool":                  "ডাইরেক্ট কুল",
		"309 Ltr.":                     "৩০৯ লিটার",
		"270 Ltr.":                     "২৭০ লিটার",
		"64 Kg":                        "৬৪ কেজি",
		"220 ~ 240/ 50":                "২২০ ~ ২৪০/ ৫০",
		"V 0101- RSCR; V 0201- RSCR":   "V 0101- RSCR; V 0201- RSCR",
		"V 0101- R600a; V 0201- R600a": "V 0101- R600a; V 0201- R600a",
		"Mechanical":                   "যান্ত্রিক",
		"Manual":                       "ম্যানুয়াল",
		"Wire/2":                       "ওয়্যার/2",
		"Yes/3":                        "হ্যাঁ/3",
		"Yes/1":                        "হ্যাঁ/1",
		"594 x 708 x 1646 mm":          "৫৯৪ x ৭০৮ x ১৬৪৬ মিমি",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c0n-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC0nGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c0n-gdel-xx"
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
		"Model Name":          "MFE-C0N-GDEL-XX",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "309 Ltr.",
		"Net Volume":          "270 Ltr.",
		"Weight":              "64 Kg",
		"Voltage":             "220 ~ 240/ 50",
		"Compressor Type":     "V 0101- RSCR; V 0201- RSCR",
		"Refrigerant":         "V 0101- R600a; V 0201- R600a",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "Yes/3",
		"Crisper Drawers":     "Yes/1",
		"Dimensions":          "594 x 708 x 1646 mm",
		"Special Features":    "Climate Class: N ~ ST; Compressor Input Power (Watt): V 0101- 117; V 0201- 109; Cooling Effect: Freezer Cabinet Less than -180C; Refrigerator Cabinet 00C to +50C; Reversible Door: No; Handle (Recessed/ Grip): Recressed; Lock: Yes; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Interior Lamp: Yes; Vegetable Box Cover: Yes; Egg Case: Yes/2; Ice Case: Yes/1; Ice Box: Yes/1; Can Storage Dispenser: No; Freezer Shelf (Material/ No.): No; Freezer Door Basket: No; Freezer Drawer: Yes/3; Freezer Ice Tray: Yes/1; Freezer Ice Remover spoon: Yes/1; Freezer Interior Lamp: No; Packing Dimensions: 625 x 745 x 1676 mm; Gross Weight: 69.5 ± 2 Kg; Loading Capacity: 78/ 57/ 27",
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
