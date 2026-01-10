package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB8bGdenDdInverter seeds specifications/options for product 'mfe-b8b-gden-dd-inverter'
type SpecificationSeederRefrigeratorMarcelMfeB8bGdenDdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB8bGdenDdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB8bGdenDdInverter() *SpecificationSeederRefrigeratorMarcelMfeB8bGdenDdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeB8bGdenDdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-b8b-gden-dd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdenDdInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1590": "1590",
		"220-240V ~ and 50Hz": "220-240V ~ and 50Hz",
		"265 Ltr.": "265 Ltr.",
		"282 Ltr.": "282 Ltr.",
		"59 ± 2 Kg": "59 ± 2 Kg",
		"635": "635",
		"66 ± 2 Kg": "66 ± 2 Kg",
		"740": "740",
		"Manual": "Manual",
		"No": "No",
		"V 0702- Electronic V 0902- Electronic": "V 0702- Electronic V 0902- Electronic",
		"V 0702- R600a V 0902- R600a": "V 0702- R600a V 0902- R600a",
		"V 0702- RSCR V 0902- BLDC": "V 0702- RSCR V 0902- BLDC",
		"Wire/2": "Wire/2",
		"Yes (ABS/ PS)": "Yes (ABS/ PS)",
		"Yes (Plastic)": "Yes (Plastic)",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-b8b-gden-dd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdenDdInverter) Seed(db *gorm.DB) error {
	productSlug := "mfe-b8b-gden-dd-inverter"
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
		"Brand":                             310,
		"Model Name":                        316,
		"Door Type":                         142,
		"Capacity":                          138,
		"Refrigerator Capacity":             156,
		"Freezer Capacity":                  146,
		"Energy Efficiency Rating":          143,
		"Energy Star Rating":                144,
		"Annual Energy Consumption":         137,
		"Dimensions":                        25,
		"Weight":                            80,
		"Color":                             311,
		"Compressor Type":                   139,
		"Cooling Technology":                698,
		"Defrost Type":                      141,
		"Temperature Control":               158,
		"Shelf Material":                    699,
		"Number of Shelves":                 154,
		"Door Bins":                         700,
		"Crisper Drawers":                   701,
		"Ice Maker":                         702,
		"Water Dispenser":                   703,
		"Noise Level":                       150,
		"Voltage":                           160,
		"Frequency (Hz)":                    704,
		"App Control":                       705,
		"Voice Assistant Support":           385,
		"Warranty":                          323,
		"Compressor Warranty (Years)":       707,
		"Refrigerant":                       708,
		"Gross Volume":                      709,
		"Net Volume":                        710,
		"Special Features":                  69,
		"Capillary":                         711,
		"Climate Type (SN, N, ST, T)":       712,
		"Compressor Input Power (Watt)":     713,
		"Cooling Effect":                    714,
		"Defrosting (Automatic/ Manual)":    715,
		"Deodorizer":                        716,
		"Depth (mm)":                        717,
		"Door Basket":                       718,
		"Egg Tray or Pocket":                719,
		"Handle (Recessed/ Grip)":           720,
		"Interior Lamp":                     721,
		"Polyurethane Foam Blowing Agent":   722,
		"Reversible Door":                   723,
		"Thermostat":                        724,
		"Vegetable Crisper":                 725,
		"Vegetable Crisper Cover":           726,
		"Loading Capacity (40HQ/40Ft/20Ft)": 727,
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}
	specs := map[string]string{
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 0702- 109V 0902- 96",
		"Compressor Type": "V 0702- RSCRV 0902- BLDC",
		"Condenser": "Steel",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "711",
		"Door Basket": "PVC/3",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "282 Ltr.",
		"Gross Weight": "66 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recressed/ Grip/ Built-in",
		"Height/mm": "1520",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "80/ 80/ 38",
		"Lock": "Yes",
		"Net Volume": "265 Ltr.",
		"Net Weight": "59 ± 2 Kg",
		"Polyurethane foam blowing agent": "CycloPentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V ~ and 50Hz",
		"Recommended voltage stabilizer capacity": "V0902: Wide Voltage Design (80V-300V)N.B: No need to volatage stabilizerV0702: Wide Voltage Design (155V-260V)N.B.: If out of voltage range(155V-260V) then suggested voltage stabilizer capacity is 2100VA.",
		"Refrigerant": "V 0702- R600aV 0902- R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "V 0702- ElectronicV 0902- Electronic",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes (Plastic)",
		"Vegetable Crisper Cover": "Yes (ABS/ PS)",
		"Width/mm": "594",
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
