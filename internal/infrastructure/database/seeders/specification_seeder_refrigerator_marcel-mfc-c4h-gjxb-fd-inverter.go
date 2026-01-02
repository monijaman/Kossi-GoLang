package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter seeds specifications/options for product 'marcel-mfc-c4h-gjxb-fd-inverter'
type SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter() *SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfc-c4h-gjxb-fd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":         "মার্সেল",
				"marcel-mfc-c4h-gjxb-fd-inverter":         "মার্সেল-mfc-c4h-gjxb-fd-inverter",
		"MFC-C4H-GJXB-FD-INVERTER":   "MFC-C4H-GJXB-FD-INVERTER",
		// Add more translations as needed
		"348 Ltr": "৩৪৮ Ltr",
		"Freezer Cabinet -15℃ to -22℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet -১৫℃ to -২২℃ Refrigerator Cabinet ০℃ to +৫℃",
		"67/ 48/ 24": "৬৭/ ৪৮/ ২৪",
		"V 0501- Mechanical V 0502- Electronics V 0503- Mechanical V 0504- Electronics V 0505- Electronics": "V ০৫০১- Mechanical V ০৫০২- Electronics V ০৫০৩- Mechanical V ০৫০৪- Electronics V ০৫০৫- Electronics",
		"Glass/ 2": "Glass/ ২",
		"Yes": "হ্যাঁ",
		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",
		"Recessed/ Grip": "Recessed/ Grip",
		"710": "৭১০",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Yes(Glass/ Plastic)": "Yes(Glass/ Plastic)",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"R600a": "R৬০০a",
		"1790": "১৭৯০",
		"Yes/2": "Yes/২",
		"V 0501-33.78~126.46 V 0502-33.78~126.46 V 0503-33.78~126.46 V 0504-33.78~126.46 V 0505-33.78~126.46": "V ০৫০১-৩৩.৭৮~১২৬.৪৬ V ০৫০২-৩৩.৭৮~১২৬.৪৬ V ০৫০৩-৩৩.৭৮~১২৬.৪৬ V ০৫০৪-৩৩.৭৮~১২৬.৪৬ V ০৫০৫-৩৩.৭৮~১২৬.৪৬",
		"Yes/1": "Yes/১",
		"Copper": "কপার",
		"Manual": "Manual",
		"No": "না",
		"N~ST": "N~ST",
		"2": "২",
		"66/ 73 ± 2": "৬৬/ ৭৩ ± ২",
		"V 0501- BLDC V 0502- BLDC V 0503- BLDC V 0504- BLDC V 0505- BLDC": "V ০৫০১- BLDC V ০৫০২- BLDC V ০৫০৩- BLDC V ০৫০৪- BLDC V ০৫০৫- BLDC",
		"V 0501/V0502/V0503/V0504/V0505: Wide Voltage Design (80V to 300V) N.B.: Do not use voltage stabilizer unless voltage is out of the range(80V-300V). Then suggested voltage stabilizer capacity is 2100VA.": "V ০৫০১/V০৫০২/V০৫০৩/V০৫০৪/V০৫০৫: Wide Voltage Design (৮০V to ৩০০V) N.B.: Do not use voltage stabilizer unless voltage is out of the range(৮০V-৩০০V). Then suggested voltage stabilizer capacity is ২১০০VA.",
		"5": "৫",
		"333 Ltr.": "৩৩৩ লিটার",
		"Wire/2": "Wire/২",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfc-c4h-gjxb-fd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfcC4hGjxbFdInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfc-c4h-gjxb-fd-inverter"
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
		"Gross Volume": "348 Ltr",
		"Net Volume": "333 Ltr.",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Compressor Input Power (Watt)": "V 0501-33.78~126.46 V 0502-33.78~126.46 V 0503-33.78~126.46 V 0504-33.78~126.46 V 0505-33.78~126.46",
		"Compressor Type": "V 0501- BLDC V 0502- BLDC V 0503- BLDC V 0504- BLDC V 0505- BLDC",
		"Cooling Effect": "Freezer Cabinet -15℃ to -22℃ Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/ Mechanical)": "V 0501- Mechanical V 0502- Electronics V 0503- Mechanical V 0504- Electronics V 0505- Electronics",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Reversible Door": "No",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Lock": "Yes",
		"Refrigerant": "R600a",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Operating voltage": "V 0501/V0502/V0503/V0504/V0505: Wide Voltage Design (80V to 300V) N.B.: Do not use voltage stabilizer unless voltage is out of the range(80V-300V). Then suggested voltage stabilizer capacity is 2100VA.",
		"Shelf (Material/ No.)": "Wire/2",
		"Door Basket": "2",
		"Interior Lamp": "No",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes(Glass/ Plastic)",
		"Egg Tray or Pocket": "Yes/2",
		"Can Storage Dispenser": "No",
		"Deodorizer": "No",
		"Shelf (Material/No.)": "Glass/ 2",
		"Drawer": "No",
		"Width/mm": "710",
		"Depth/mm": "710",
		"Height/mm": "1790",
		"Weight/Kg - Net/Packing": "66/ 73 ± 2",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "67/ 48/ 24",
		"Star rating (BDS 1850:2012)": "5",
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
