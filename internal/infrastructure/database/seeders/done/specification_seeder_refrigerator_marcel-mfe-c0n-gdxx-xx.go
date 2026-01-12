package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx seeds specifications/options for product 'marcel-mfe-c0n-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx() *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c0n-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-c0n-gdxx-xx": "মার্সেল-mfe-c0n-gdxx-xx",
		"MFE-C0N-GDXX-XX":        "MFE-C0N-GDXX-XX",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c0n-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c0n-gdxx-xx"
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
		"Model Name":          "MFE-C0N-GDXX-XX",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "309 Ltr.",
		"Net Volume":          "270 Ltr.",
		"Weight":              "64 Kg",
		"Voltage":             "220 ~ 240/ 50Hz",
		"Compressor Type":     "V 0101- RSCR V 0201- RSCR",
		"Refrigerant":         "V 0101- R600a V 0201- R600a",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "Yes/3",
		"Crisper Drawers":     "Yes/1",
		"Dimensions":          "594 x 708 x 1646 mm",
		"Special Features":    "Packing Weight: 69.5 ± 2 Kg; Compressor Input Power (W): V 0101-117 V 0201-109; Cooling Effect: Freezer Cabinet Less than -18 ̊C; Refrigerator Cabinet 0 ̊C to +5 ̊C; Handle: Recressed; Lock: Yes; Capillary: Copper; Thermostat: RoHS Certified; Polyurethane foam blowing agent: CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Refrigerator features: Door Basket Yes/3, Vegetable Box Yes/1, Vegetable Box Cover Yes, Egg Case Yes/2, Ice Case Yes/1, Ice Box Yes/1; Freezer features: Drawer Yes/3, Ice Tray Yes/1, Ice Remover spoon Yes/1; Packaging Dimensions: 625 x 745 x 1676 mm; Loading Capacity- 40HQ/40Ft/20Ft: 78/57/27",
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
