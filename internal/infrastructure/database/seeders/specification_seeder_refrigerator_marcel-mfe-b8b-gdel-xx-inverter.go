package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB8bGdelXxInverter seeds specifications/options for product 'mfe-b8b-gdel-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeB8bGdelXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB8bGdelXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB8bGdelXxInverter() *SpecificationSeederRefrigeratorMarcelMfeB8bGdelXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeB8bGdelXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-b8b-gdel-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdelXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1530": "1530",
		"220-240V ~ and 50Hz": "220-240V ~ and 50Hz",
		"265 Ltr.": "265 Ltr.",
		"282 Ltr.": "282 Ltr.",
		"59 ± 2 Kg": "59 ± 2 Kg",
		"625": "625",
		"745": "745",
		"Manual": "Manual",
		"Mechanical": "Mechanical",
		"No": "No",
		"R600a": "R600a",
		"V 0501- BLDC": "V 0501- BLDC",
		"Wire/2": "Wire/2",
		"Yes (ABS/ PS)": "Yes (ABS/ PS)",
		"Yes (Plastic)": "Yes (Plastic)",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-b8b-gdel-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdelXxInverter) Seed(db *gorm.DB) error {
	productSlug := "mfe-b8b-gdel-xx-inverter"
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
		"Compressor Type": "V 0501- BLDC",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "745",
		"Door Basket": "No",
		"Gross Volume": "282 Ltr.",
		"Height/mm": "1530",
		"Net Volume": "265 Ltr.",
		"Net Weight": "59 ± 2 Kg",
		"Rated Operating Voltage and Frequency": "220-240V ~ and 50Hz",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Vegetable Crisper": "Yes (Plastic)",
		"Vegetable Crisper Cover": "Yes (ABS/ PS)",
		"Width/mm": "625",
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
