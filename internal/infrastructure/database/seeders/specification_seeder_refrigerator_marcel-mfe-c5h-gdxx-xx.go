package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx seeds specifications/options for product 'mfe-c5h-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx() *SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-c5h-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1830": "1830",
		"220~240/ 50/135": "220~240/ 50/135",
		"345 Ltr.": "345 Ltr.",
		"358 Ltr.": "358 Ltr.",
		"625": "625",
		"745": "745",
		"76.5 ± 2 Kg": "76.5 ± 2 Kg",
		"84 ± 2 Kg": "84 ± 2 Kg",
		"Manual": "Manual",
		"No": "No",
		"Rack Evaporator": "Rack Evaporator",
		"V 0102- R600a V 0301- R600a": "V 0102- R600a V 0301- R600a",
		"V 0102- RSCR V 0301- RSCR": "V 0102- RSCR V 0301- RSCR",
		"Yes": "Yes",
		"Yes/1": "Yes/1",
		"Yes/2": "Yes/2",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-c5h-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "mfe-c5h-gdxx-xx"
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
		"Compressor Type": "V 0102- RSCR V 0301- RSCR",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "745",
		"Door Basket": "No",
		"Gross Volume": "358 Ltr.",
		"Gross Weight": "84 ± 2 Kg",
		"Height/mm": "1830",
		"Ice Case": "Yes/2",
		"Ice Tray": "Yes/1",
		"Net Volume": "345 Ltr.",
		"Net Weight": "76.5 ± 2 Kg",
		"Rated Voltage/ Hz/ watt": "220~240/ 50/135",
		"Refrigerant": "V 0102- R600a V 0301- R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Rack Evaporator",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
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
