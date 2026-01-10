package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5xGdxxXx seeds specifications/options for product 'marcel-mfb-b5x-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfbB5xGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5xGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5xGdxxXx() *SpecificationSeederRefrigeratorMarcelMfbB5xGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB5xGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b5x-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"220 ~ 240/ 50": "220 ~ 240/ 50",
		"250 Ltr.":      "250 Ltr.",
		"274 Ltr.":      "274 Ltr.",
		"54.5 ± 2 Kg":   "54.5 ± 2 Kg",
		"Manual":        "Manual",
		"R600a":         "R600a",
		"RSCR":          "RSCR",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b5x-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b5x-gdxx-xx"
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
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 01.01-97.4V 02.01-97.4V 03.01-97.4V 03.02-97.4",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "630 mm",
		"Door Basket": "GPPS/3",
		"Drawer": "No",
		"Egg Tray": "Yes",
		"Gross Volume": "250 Ltr.",
		"Gross Weight": "59.5± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height/mm": "1675 mm",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "97/ 72/ 36",
		"Lock": "Yes",
		"Net Volume": "274 Ltr.",
		"Net Weight": "54.5± 2 Kg",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 01.01, V 02.01-Low Voltage(140~260V)For V01.01, V 02.01-Wide Voltage Range (140Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is   recommended.V 03.01, V 03.02-Low Voltage(150~260V)For V 03.01, V 03.02 - Wide Voltage Range (150Vac -   260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is   recommended.",
		"Refrigerant": "R600a",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Width/mm": "555 mm",
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
