package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC6eI5GedeFdInverter seeds specifications/options for product 'mfc-c6e-i5-gede-fd-inverter'
type SpecificationSeederRefrigeratorMarcelMfcC6eI5GedeFdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC6eI5GedeFdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC6eI5GedeFdInverter() *SpecificationSeederRefrigeratorMarcelMfcC6eI5GedeFdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfcC6eI5GedeFdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfc-c6e-i5-gede-fd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC6eI5GedeFdInverter) getBanglaTranslations() map[string]string {
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
		"V 0401- Mechanical V 0501- Mechanical V 0601- Mechanical V 0701- Mechanical V 0702- Mechanical V 0703-Mechanical V 0801-Mechanical V 0802-Electronic V 0901-Mechanical V 0902-Electronic V 1001-Electronic":                                                                                                     "V 0401- Mechanical V 0501- Mechanical V 0601- Mechanical V 0701- Mechanical V 0702- Mechanical V 0703-Mechanical V 0801-Mechanical V 0802-Electronic V 0901-Mechanical V 0902-Electronic V 1001-Electronic",
		"V 0401- RSCR V 0501-BLDC V 0601- RSCR V 0701- RSCR V 0702- RSCR V 0703-RSCR V 0801-BLDC V 0802-BLDC V 0901-BLDC V 0902-BLDC V 1001-BLDC":                                                                                                                                                                        "V 0401- RSCR V 0501-BLDC V 0601- RSCR V 0701- RSCR V 0702- RSCR V 0703-RSCR V 0801-BLDC V 0802-BLDC V 0901-BLDC V 0902-BLDC V 1001-BLDC",
		"V 0401/V 0601/V 0701/V 0702/V 0703: Need Voltage stabilizer capacity is 2100VA V 0501: Wide Voltage Design (105V-276V) V 0801:Wide Voltage Design (75V-264V) V 0901/V 0902/V 1001:Wide Voltage Design (80V-300V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.": "V 0401/V 0601/V 0701/V 0702/V 0703: Need Voltage stabilizer capacity is 2100VA V 0501: Wide Voltage Design (105V-276V) V 0801:Wide Voltage Design (75V-264V) V 0901/V 0902/V 1001:Wide Voltage Design (80V-300V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Wire/2": "Wire/2",
		"Yes":    "Yes",
		"Yes/1":  "Yes/1",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfc-c6e-i5-gede-fd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfcC6eI5GedeFdInverter) Seed(db *gorm.DB) error {
	productSlug := "mfc-c6e-i5-gede-fd-inverter"
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
		"Type":                                  "Direct Cool",
		"Gross Volume":                          "380 Ltr.",
		"Net Volume":                            "365 Ltr.",
		"Climatic Type (SN, N, ST, T)":          "N~ST",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Compressor Input Power (Watt)":         "V0401- 130V 0501- 50.3~166.7V 0601- 130V 0701- 130V 0702- 130V 0703-130V 0801-33.78~126.46V 0802-33.78~126.46V 0901-33.78~126.46V 0902-33.78~126.47V 1001-33.78~126.48",
		"Compressor Type":                       "V0401- RSCRV 0501-BLDCV 0601- RSCRV 0701- RSCRV 0702- RSCRV 0703-RSCRV 0801-BLDCV 0802-BLDCV 0901-BLDCV 0902-BLDCV 1001-BLDC",
		"Cooling Effect":                        "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/  Mechanical)": "V0401- MechanicalV 0501- MechanicalV 0601- MechanicalV 0701- MechanicalV 0702- MechanicalV 0703-MechanicalV 0801-MechanicalV 0802-ElectronicV 0901-MechanicalV 0902-ElectronicV 1001-Electronic",
		"Defrosting (Automatic/ Manual)":                "Manual",
		"Reversible Door":                               "No",
		"Handle (Recessed/ Grip)":                       "Recessed/ Grip",
		"Lock":                                          "Yes",
		"Capillary":                                     "Copper",
		"Polyurethane foam blowing agent":               "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Operating voltage":                             "V0401/V 0601/V 0701/V 0702/V 0703: Need Voltage  stabilizer capacity is 2100VAV 0501: Wide Voltage Design (105V-276V)V 0801:Wide Voltage Design (75V-264V)V 0901/V 0902/V 1001:Wide Voltage Design (80V-300V)N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Shelf (Material/ No.)":                         "Wire/2",
		"Door Basket":                                   "2",
		"Interior Lamp":                                 "No",
		"Vegetable Crisper":                             "Yes/1",
		"Vegetable Crisper Cover":                       "Yes",
		"Egg Tray":                                      "Yes/2",
		"Can Storage Dispenser":                         "No",
		"Drawer":                                        "No",
		"Width/mm":                                      "710",
		"Depth/mm":                                      "710",
		"Height/mm":                                     "1910",
		"Weight/Kg - Net/Packing":                       "67/ 76 ± 2",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":            "66/ 48/ 24",
		"Star rating (BDS 1850:2012)":                   "5",
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
