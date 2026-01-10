package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5dGaxbWdInverter seeds specifications/options for product 'mfb-b5d-gaxb-wd-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB5dGaxbWdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5dGaxbWdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5dGaxbWdInverter() *SpecificationSeederRefrigeratorMarcelMfbB5dGaxbWdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB5dGaxbWdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfb-b5d-gaxb-wd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5dGaxbWdInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1845 mm":       "1845 mm",
		"220 ~ 240/ 50": "220 ~ 240/ 50",
		"254 Ltr. (Freezer 120L, Refrigerator: 134L)": "254 Ltr. (Freezer 120L, Refrigerator: 134L)",
		"268 Ltr. (Freezer 127L, Refrigerator: 141L)": "268 Ltr. (Freezer 127L, Refrigerator: 141L)",
		"5 star (BDS 1850:2012)":                      "5 star (BDS 1850:2012)",
		"580 mm":                                      "580 mm",
		"62 ± 2 Kg":                                   "62 ± 2 Kg",
		"645 mm":                                      "৬৪৫ মিমি",
		"67 ± 2 Kg":                                   "67 ± 2 Kg",
		"Electronic":                                  "Electronic",
		"Manual":                                      "Manual",
		"No":                                          "No",
		"R600a":                                       "R600a",
		"V.0801: BLDC":                                "V.0801: BLDC",
		"Wire/3":                                      "Wire/3",
		"Yes":                                         "Yes",
		"Yes (New Feature with 3.5L Water Reserving Capacity)": "Yes (New Feature with 3.5L Water Reserving Capacity)",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfb-b5d-gaxb-wd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5dGaxbWdInverter) Seed(db *gorm.DB) error {
	productSlug := "mfb-b5d-gaxb-wd-inverter"
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
		"Climatic Type (SN, N, ST, T)": "N~T",
		"Compressor Input Power (Watt)": "V.0801: 34 ~128",
		"Compressor Type": "V.0801: BLDC",
		"Cooling Effect": "Freezer Cabinet between -16℃ to -24℃Fresh Food Cabinet  between 0℃ to +6℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "635 mm",
		"Door Pocket": "GPPS/4",
		"Drawer": "HIPS/5",
		"Egg Tray": "Yes",
		"Energy Rating": "5 star (BDS 1850:2012)",
		"Gross Volume": "268 Ltr. (Freezer 127L, Refrigerator: 141L)",
		"Gross Weight": "67 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height/mm": "1810 mm",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "97/ 72/ 36",
		"Lock": "Yes",
		"Net Volume": "254 Ltr. (Freezer 120L, Refrigerator: 134L)",
		"Net Weight": "62 ± 2 Kg",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "For V.0801 - Wide Voltage Range at AC Input (75 V - 264 V). Voltage stabilizer is not required.",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Water Dispenser": "Yes (New Feature with 3.5L Water Reserving Capacity)",
		"Width/mm": "550 mm",
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
