package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter seeds specifications/options for product 'mfc-c4h-gjxb-fd-inverter'
type SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter() *SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfc-c4h-gjxb-fd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1790":               "1790",
		"2":                  "2",
		"220-240V~ and 50Hz": "220-240V~ and 50Hz",
		"333 Ltr.":           "333 Ltr.",
		"348 Ltr":            "348 Ltr",
		"5":                  "5",
		"66/ 73 ± 2":         "66/ 73 ± 2",
		"710":                "710",
		"Glass/ 2":           "Glass/ 2",
		"Manual":             "Manual",
		"No":                 "No",
		"R600a":              "R600a",
		"V 0501- BLDC V 0502- BLDC V 0503- BLDC V 0504- BLDC V 0505- BLDC":                                  "V 0501- BLDC V 0502- BLDC V 0503- BLDC V 0504- BLDC V 0505- BLDC",
		"V 0501- Mechanical V 0502- Electronics V 0503- Mechanical V 0504- Electronics V 0505- Electronics": "V 0501- Mechanical V 0502- Electronics V 0503- Mechanical V 0504- Electronics V 0505- Electronics",
		"V 0501/V0502/V0503/V0504/V0505: Wide Voltage Design (80V to 300V) N.B.: Do not use voltage stabilizer unless voltage is out of the range(80V-300V). Then suggested voltage stabilizer capacity is 2100VA.": "V 0501/V0502/V0503/V0504/V0505: Wide Voltage Design (80V to 300V) N.B.: Do not use voltage stabilizer unless voltage is out of the range(80V-300V). Then suggested voltage stabilizer capacity is 2100VA.",
		"Wire/2":              "Wire/2",
		"Yes(Glass/ Plastic)": "Yes(Glass/ Plastic)",
		"Yes/1":               "Yes/1",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfc-c4h-gjxb-fd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter) Seed(db *gorm.DB) error {
	productSlug := "mfc-c4h-gjxb-fd-inverter"
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
		"Compressor Input Power (Watt)": "V 0501-33.78~126.46V 0502-33.78~126.46V 0503-33.78~126.46V 0504-33.78~126.46V 0505-33.78~126.46",
		"Compressor Type": "V 0501- BLDCV 0502- BLDCV 0503- BLDCV 0504- BLDCV 0505- BLDC",
		"Cooling Effect": "Freezer Cabinet -15℃ to -22℃
Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "650",
		"Door Basket": "4",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes/2",
		"Gross Volume": "348 Ltr",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height/mm": "1740",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "67/ 48/ 24",
		"Lock": "Yes",
		"Net Volume": "333 Ltr.",
		"Operating voltage": "V 0501/V0502/V0503/V0504/V0505: Wide Voltage Design (80V to 300V)N.B.: Do not use voltage stabilizer unless voltage is out of the range(80V-300V). Then suggested voltage stabilizer capacity is 2100VA.",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Shelf (Material/No.)": "Glass/ 2",
		"Star rating (BDS 1850:2012)": "5",
		"Temperature Control (Electronic/  Mechanical)": "V 0501- MechanicalV 0502- ElectronicsV 0503- MechanicalV 0504- ElectronicsV 0505- Electronics",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes(Glass/ Plastic)",
		"Weight/Kg - Net/Packing": "66/ 73 ± 2",
		"Width/mm": "650",
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
