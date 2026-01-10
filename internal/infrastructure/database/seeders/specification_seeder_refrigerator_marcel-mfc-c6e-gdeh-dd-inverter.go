package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC6eGdehDdInverter seeds specifications/options for product 'mfc-c6e-gdeh-dd-inverter'
type SpecificationSeederRefrigeratorMarcelMfcC6eGdehDdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC6eGdehDdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC6eGdehDdInverter() *SpecificationSeederRefrigeratorMarcelMfcC6eGdehDdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfcC6eGdehDdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfc-c6e-gdeh-dd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC6eGdehDdInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1910":               "1910",
		"2":                  "2",
		"220-240V~ and 50Hz": "220-240V~ and 50Hz",
		"365 Ltr.":           "365 Ltr.",
		"380 Ltr.":           "380 Ltr.",
		"5":                  "5",
		"67/ 76 ± 2":         "67/ 76 ± 2",
		"710":                "710",
		"Manual":             "Manual",
		"No":                 "No",
		"V 0401- Mechanical V 0501- Mechanical V 0601- Mechanical V 0701- Mechanical V 0702- Mechanical V 0703-Mechanical V 0801-Mechanical V 0802-Electronic":                                                                                                       "V 0401- Mechanical V 0501- Mechanical V 0601- Mechanical V 0701- Mechanical V 0702- Mechanical V 0703-Mechanical V 0801-Mechanical V 0802-Electronic",
		"V 0401- RSCR V 0501-BLDC V 0601- RSCR V 0701- RSCR V 0702- RSCR V 0703-RSCR V 0801-BLDC V 0802-BLDC":                                                                                                                                                        "V 0401- RSCR V 0501-BLDC V 0601- RSCR V 0701- RSCR V 0702- RSCR V 0703-RSCR V 0801-BLDC V 0802-BLDC",
		"V 0401/V 0601/V 0701/V 0702/V 0703: Need Voltage stabilizer capacity is 2100VA V 0501: Wide Voltage Design (105V-276V) V 0801:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.": "V 0401/V 0601/V 0701/V 0702/V 0703: Need Voltage stabilizer capacity is 2100VA V 0501: Wide Voltage Design (105V-276V) V 0801:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Wire/2": "Wire/2",
		"Yes":    "Yes",
		"Yes/1":  "Yes/1",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfc-c6e-gdeh-dd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfcC6eGdehDdInverter) Seed(db *gorm.DB) error {
	productSlug := "mfc-c6e-gdeh-dd-inverter"
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
		"Can Storage Dispenser":                 "No",
		"Capillary":                             "Copper",
		"Climatic Type (SN, N, ST, T)":          "N~ST",
		"Compressor Input Power (Watt)":         "V0401- 130V 0501- 50.3~166.7V 0601- 130V 0701- 130V 0702- 130V 0703-130V 0801-33.78~126.46V 0802-33.78~126.46",
		"Compressor Type":                       "V0401- RSCRV 0501-BLDCV 0601- RSCRV 0701- RSCRV 0702- RSCRV 0703-RSCRV 0801-BLDCV 0802-BLDC",
		"Cooling Effect":                        "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)":        "Manual",
		"Depth/mm":                              "650",
		"Door Basket":                           "4",
		"Drawer":                                "No",
		"Egg Tray":                              "Yes/2",
		"Gross Volume":                          "380 Ltr.",
		"Handle (Recessed/ Grip)":               "Recessed/ Grip",
		"Height/mm":                             "1860",
		"Interior Lamp":                         "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":    "66/ 48/ 24",
		"Lock":                                  "Yes",
		"Net Volume":                            "365 Ltr.",
		"Operating voltage":                     "V0401/V 0601/V 0701/V 0702/V 0703: Need Voltage  stabilizer capacity is 2100VAV 0501: Wide Voltage Design (105V-276V)V 0801:Wide Voltage Design (75V-264V)N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Polyurethane foam blowing agent":       "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Reversible Door":                       "No",
		"Shelf (Material/ No.)":                 "Wire/3",
		"Star rating (BDS 1850:2012)":           "5",
		"Temperature Control (Electronic/  Mechanical)": "V0401- MechanicalV 0501- MechanicalV 0601- MechanicalV 0701- MechanicalV 0702- MechanicalV 0703-MechanicalV 0801-MechanicalV 0802-Electronic",
		"Type":                    "Direct Cool",
		"Vegetable Crisper":       "Yes/1",
		"Vegetable Crisper Cover": "Yes",
		"Weight/Kg - Net/Packing": "67/ 76 ± 2",
		"Width/mm":                "650",
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
