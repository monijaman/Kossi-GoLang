package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter seeds specifications/options for product 'marcel-mfe-c5h-gden-dd-inverter'
type SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter() *SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c5h-gden-dd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":         "মার্সেল",
				"marcel-mfe-c5h-gden-dd-inverter":         "মার্সেল-mfe-c5h-gden-dd-inverter",
		"MFE-C5H-GDEN-DD-INVERTER":   "MFE-C5H-GDEN-DD-INVERTER",
		// Add more translations as needed
		"1890": "১৮৯০",
		"Rack Evaporator": "Rack Evaporator",
		"635": "৬৩৫",
		"Yes": "হ্যাঁ",
		"740": "৭৪০",
		"345 Ltr.": "৩৪৫ লিটার",
		"V 0102- 130 V 0301- 123 V 0401- 126.46": "V ০১০২- ১৩০ V ০৩০১- ১২৩ V ০৪০১- ১২৬.৪৬",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Electronic": "ইলেকট্রনিক",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"V 0102- Copper V 0301- Copper V 0401- Copper": "V ০১০২- Copper V ০৩০১- Copper V ০৪০১- Copper",
		"Yes/2": "Yes/২",
		"220-240V~/50Hz": "২২০-২৪০V~/৫০Hz",
		"Yes/1": "Yes/১",
		"Manual": "Manual",
		"RoHS Certified": "RoHS Certified",
		"No": "না",
		"76/ 57/ 27": "৭৬/ ৫৭/ ২৭",
		"No Need To use Voltage stabilizer. NB: If out of voltage range(140V-260V), Then suggested 2100VA.": "No Need To use Voltage stabilizer. NB: If out of voltage range(১৪০V-২৬০V), Then suggested ২১০০VA.",
		"84 ± Kg": "৮৪ ± কেজি",
		"Yes/4": "Yes/৪",
		"N ~ ST": "N ~ ST",
		"76.5 ± Kg": "৭৬.৫ ± কেজি",
		"V 0102- RSIR V 0301- RSIR V 0401- BLDC Inverter": "V ০১০২- RSIR V ০৩০১- RSIR V ০৪০১- BLDC Inverter",
		"Freezer Cabinet Less than -18 ̊C Refrigerator Cabinet 0 ̊C to +5 ̊C": "Freezer Cabinet Less than -১৮ ̊C Refrigerator Cabinet ০ ̊C to +৫ ̊C",
		"V 0102- R600a V 0301- R600a V 0401- R600a": "V ০১০২- R৬০০a V ০৩০১- R৬০০a V ০৪০১- R৬০০a",
		"358 Ltr.": "৩৫৮ লিটার",
		"Recessed": "Recessed",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c5h-gden-dd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c5h-gden-dd-inverter"
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
		"Type": "Direct Cool",
		"Gross Volume": "358 Ltr.",
		"Net Volume": "345 Ltr.",
		"Net Weight": "76.5 ± Kg",
		"Gross Weight": "84 ± Kg",
		"Climate Type (SN, N, ST, T)": "N ~ ST",
		"Rated Voltage/ Hz": "220-240V~/50Hz",
		"Compressor Input Power (Watt)": "V 0102- 130 V 0301- 123 V 0401- 126.46",
		"Compressor Type": "V 0102- RSIR V 0301- RSIR V 0401- BLDC Inverter",
		"Cooling Efect": "Freezer Cabinet Less than -18 ̊C Refrigerator Cabinet 0 ̊C to +5 ̊C",
		"Temperature Control (Electronic/ Mechanical)": "Electronic",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Reversible Door": "No",
		"Handle (Recessed/ Grip):": "Recessed",
		"Lock": "Yes",
		"Refrigerant": "V 0102- R600a V 0301- R600a V 0401- R600a",
		"Thermostat": "RoHS Certified",
		"Capillary": "V 0102- Copper V 0301- Copper V 0401- Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "No Need To use Voltage stabilizer. NB: If out of voltage range(140V-260V), Then suggested 2100VA.",
		"Shelf (Material/ No.)": "Rack Evaporator",
		"Door Basket": "No",
		"Interior Lamp": "No",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
		"Egg Tray or Pocket": "Yes",
		"Can Storage Dispenser": "No",
		"Deodorizer": "No",
		"Drawer": "Yes/4",
		"Ice Tray": "Yes/1",
		"Ice Case": "Yes/2",
		"Ice Remover spoon": "Yes/1",
		"Width/mm": "635",
		"Depth/mm": "740",
		"Height/mm": "1890",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "76/ 57/ 27",
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
