package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC5hElexXx seeds specifications/options for product 'marcel-mfe-c5h-elex-xx'
type SpecificationSeederRefrigeratorMarcelMfeC5hElexXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC5hElexXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC5hElexXx() *SpecificationSeederRefrigeratorMarcelMfeC5hElexXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC5hElexXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c5h-elex-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC5hElexXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-c5h-elex-xx": "মার্সেল-mfe-c5h-elex-xx",
		"MFE-C5H-ELEX-XX":        "MFE-C5H-ELEX-XX",
		"Direct Cool":            "ডাইরেক্ট কুল",
		"358 Ltr.":               "৩৫৮ লিটার",
		"345 Ltr.":               "৩৪৫ লিটার",
		"68.5 ± 2 Kg":            "৬৮.৫ ± ২ কেজি",
		"76 ± 2 Kg":              "৭৬ ± ২ কেজি",
		"N ~ ST":                 "এন ~ এসটি",
		"220~240/ 50/135":        "২২০~২৪০/ ৫০/১৩৫",
		"RSCR":                   "আরএসসিআর",
		"Mechanical":             "যান্ত্রিক",
		"Manual":                 "ম্যানুয়াল",
		"Recessed":               "রিসেসড",
		"Yes":                    "হ্যাঁ",
		"No":                     "না",
		"R600a":                  "R600a",
		"Copper":                 "কপার",
		"Cyclopentene":           "সাইক্লোপেন্টিন",
		"2000VA":                 "২০০০ভিএ",
		"594 x 732 x 1820 mm":    "৫৯৪ x ৭৩২ x ১৮২০ মিমি",
		"625 x 745 x 1830 mm":    "৬২৫ x ৭৪৫ x ১৮৩০ মিমি",
		"76/ 57/ 27":             "৭৬/ ৫৭/ ২৭",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c5h-elex-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC5hElexXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c5h-elex-xx"
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
		"Brand":                           "Marcel",
		"Model Name":                      "MFE-C5H-ELEX-XX",
		"Cooling Technology":              "Direct Cool",
		"Gross Volume":                    "358 Ltr.",
		"Net Volume":                      "345 Ltr.",
		"Net Weight":                      "68.5 ± 2 Kg",
		"Gross Weight":                    "76 ± 2 Kg",
		"Climate Type":                    "N ~ ST",
		"Voltage":                         "220~240/ 50/135",
		"Compressor Type":                 "RSCR",
		"Temperature Control":             "Mechanical",
		"Defrost Type":                    "Manual",
		"Reversible Door":                 "No",
		"Handle":                          "Recessed",
		"Lock":                            "Yes",
		"Refrigerant":                     "R600a",
		"Thermostat":                      "RoHS Certified",
		"Capillary":                       "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene",
		"Recommended voltage stabilizer capacity": "2000VA",
		"Number of Shelves":                       "2",
		"Shelf Material":                          "Wire",
		"Door Bins":                               "Yes/3",
		"Interior Lamp":                           "Yes",
		"Crisper Drawers":                         "Yes/1",
		"Vegetable Crisper":                       "Yes/1",
		"Vegetable Crisper Cover":                 "Yes",
		"Egg Tray or Pocket":                      "Yes",
		"Can Storage Dispenser":                   "No",
		"Freezer Shelf":                           "Rack Evaporator",
		"Freezer Drawer":                          "Yes/4",
		"Ice Tray":                                "Yes/1",
		"Ice Case":                                "Yes/2",
		"Ice Remover spoon":                       "Yes/1",
		"Freezer Door Baskets":                    "No",
		"Freezer Interior Lamp":                   "No",
		"Cooling Effect":                          "Freezer Cabinet Less than -18°C; Refrigerator Cabinet 0°C to +50°C",
		"Dimensions":                              "594 x 732 x 1820 mm",
		"Packing Dimensions":                      "625 x 745 x 1830 mm",
		"Loading Capacity":                        "76/ 57/ 27",
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
