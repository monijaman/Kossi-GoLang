package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx seeds specifications/options for product 'marcel-mfk-c4g-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfkC4gGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfkC4gGdelXx() *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfk-c4g-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfk-c4g-gdel-xx": "মার্সেল-mfk-c4g-gdel-xx",
		"MFK-C4G-GDEL-XX":        "MFK-C4G-GDEL-XX",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfk-c4g-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfkC4gGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfk-c4g-gdel-xx"
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
		"Model Name":          "MFK-C4G-GDEL-XX",
		"Cooling Technology":  "Direct Evaporating Cooling System (DECS)",
		"Gross Volume":        "347 Ltr",
		"Net Volume":          "345 Ltr.",
		"Compressor Type":     "V.01.01- RSCR (Input Power: V 01.01-147 W)",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Refrigerant":         "R600a",
		"Shelf Material":      "Wire",
		"Number of Shelves":   "2",
		"Door Bins":           "5",
		"Crisper Drawers":     "Yes",
		"Dimensions":          "680 x 675 x 1630 mm",
		"Packing Dimensions":  "715 x 710 x 1645 mm",
		"Weight":              "67.5 / 74 ± 2",
		"Voltage":             "V.01.01 Wide Voltage Range (145V ~ 260V)",
		"Loading Capacity":    "69/69/34",
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
