package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx seeds specifications/options for product 'mfk-c4g-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfkC4gGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfkC4gGdelXx() *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfk-c4g-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1645": "1645",
		"220-240V~ and 50Hz": "220-240V~ and 50Hz",
		"345 Ltr.": "345 Ltr.",
		"347 Ltr": "347 Ltr",
		"67.5 / 74 ± 2": "67.5 / 74 ± 2",
		"710": "710",
		"715": "715",
		"GPPS/2": "GPPS/2",
		"Manual": "Manual",
		"Mechanical": "Mechanical",
		"V.01.01 Wide Voltage Range (145V ~ 260V)NB: If out of voltage range then suggested voltage stabilizer capacity is 2100VA.": "V.01.01 Wide Voltage Range (145V ~ 260V)NB: If out of voltage range then suggested voltage stabilizer capacity is 2100VA.",
		"V.01.01- RSCR": "V.01.01- RSCR",
		"V01.01 R600a": "V01.01 R600a",
		"Wire/ 3": "Wire/ 3",
		"Wire/2": "Wire/2",
		"Yes": "Yes",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfk-c4g-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx) Seed(db *gorm.DB) error {
	productSlug := "mfk-c4g-gdel-xx"
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
		"Compressor Type": "V.01.01- RSCR",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "710",
		"Door Basket": "GPPS/2",
		"Gross Volume": "347 Ltr",
		"Height/mm": "1645",
		"Net Volume": "345 Ltr.",
		"Operating voltage": "V.01.01 Wide Voltage Range (145V ~ 260V)NB: If out of voltage range then suggested voltage stabilizer capacity is 2100VA.",
		"Rack Shelf (Material/No)": "Wire/ 3",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Refrigerant": "V01.01 R600a",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Weight/Kg - Net/Packing": "67.5 / 74 ± 2",
		"Width/mm": "715",
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
