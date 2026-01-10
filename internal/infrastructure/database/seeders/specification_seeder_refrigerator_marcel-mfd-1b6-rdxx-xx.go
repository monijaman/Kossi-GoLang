package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfd1b6RdxxXx seeds specifications/options for product 'marcel-mfd-1b6-rdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfd1b6RdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfd1b6RdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfd1b6RdxxXx() *SpecificationSeederRefrigeratorMarcelMfd1b6RdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfd1b6RdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfd-1b6-rdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfd1b6RdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":         "মার্সেল",
				"marcel-mfd-1b6-rdxx-xx":         "মার্সেল-mfd-1b6-rdxx-xx",
		"MFD-1B6-RDXX-XX":   "MFD-1B6-RDXX-XX",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfd-1b6-rdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfd1b6RdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfd-1b6-rdxx-xx"
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
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 0201- 96V 0301- 63.5V 0401- 93.25V 0501- 60.4V 0601- 56V 0701- 60.4",
		"Compressor Type": "V 0201- RSIRV 0301- RSCRV 0401- RSCRV 0501- RSCRV 0601- RSCRV 0701-RSCR",
		"Cooling Efect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "535",
		"Door Basket": "PS/2",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "132 Ltr.",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height/mm": "1320",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "184/ 88/ 44",
		"Lock": "Yes",
		"Net Volume": "129 Ltr.",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Recommended voltage stabilizer capacity": "600VA",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Shelf (Material/No.)": "Wire/1",
		"Temperature Control (Electronic/    Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
		"Weight/Kg - Net/Packing:": "38 / 42.5 ±2 KG",
		"Width/mm": "512",
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
