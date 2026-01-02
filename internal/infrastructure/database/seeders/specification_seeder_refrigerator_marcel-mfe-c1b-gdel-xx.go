package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC1bGdelXx seeds specifications/options for product 'marcel-mfe-c1b-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfeC1bGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC1bGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC1bGdelXx() *SpecificationSeederRefrigeratorMarcelMfeC1bGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC1bGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c1b-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-c1b-gdel-xx": "মার্সেল-mfe-c1b-gdel-xx",
		"MFE-C1B-GDEL-XX":        "MFE-C1B-GDEL-XX",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c1b-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c1b-gdel-xx"
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
		"Model Name":          "MFE-C1B-GDEL-XX",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "312 Ltr.",
		"Net Volume":          "290 Ltr.",
		"Weight":              "60.14 ± 2 Kg",
		"Refrigerant":         "V0101- R600a V0201- R600a V0301- R600a V0302- R600a V0401- R600a V 0501 -R600a",
		"Temperature Control": "V 0201 - Mechanical; V 0301 - Mechanical; V 0302 - Mechanical; V 0501 - Mechanical; V 0401 - Electrical",
		"Defrost Type":        "Manual",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "PVC/3",
		"Crisper Drawers":     "Yes (Plastic)",
		"Voltage":             "220-240V~ and 50Hz",
		"Dimensions":          "594 x 708 x 1620 mm",
		"Special Features":    "Net Weight: 60.14 ± 2 Kg; Gross Weight: 67.16 ± 2 Kg; Climate Type: N ~ ST; Compressor Input Power (Watt): V 0101-145.7 V 0201-145.7 V 0301-117 V 0302-117 V 0401 - 43.5~143.7 V 0501-123; Compressor Types: V 0101 - RSIR, V 0201 - RSIR, V 0301 - RSIR, V 0302 - RSIR, V 0501 - RSIR, V 0401 - BLDC; Recommended voltage stabilizer capacity: V 0201/0202/0203/0301: Wide Voltage Design (145V-253V) (suggested stabilizer 2100VA if out of range); V 0501: Wide Voltage Design (75V-264V) (suggested stabilizer 2100VA if out of range); Handle: Recressed/ Grip/ Built-in; Lock: Yes; Capillary: Copper; Thermostat: RoHS Certified; Polyurethane foam blowing agent: CycloPentene; Interior Lamp: Yes; Vegetable Crisper Cover: Yes (ABS/ PS); Egg Tray or Pocket: Yes; Can Storage Dispenser: No; Deodorizer: No; Freezer Drawer: No; Freezer Door Basket: No; Packaging Dimensions: 625 x 745 x 1630 mm; Loading Capacity- 40HQ/40Ft/20Ft: 79/54/27",
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
