package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfdA7xGdshXx seeds specifications/options for product 'marcel-mfd-a7x-gdsh-xx'
type SpecificationSeederRefrigeratorMarcelMfdA7xGdshXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfdA7xGdshXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfdA7xGdshXx() *SpecificationSeederRefrigeratorMarcelMfdA7xGdshXx {
	return &SpecificationSeederRefrigeratorMarcelMfdA7xGdshXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfd-a7x-gdsh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfdA7xGdshXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":         "মার্সেল",
				"marcel-mfd-a7x-gdsh-xx":         "মার্সেল-mfd-a7x-gdsh-xx",
		"MFD-A7X-GDSH-XX":   "MFD-A7X-GDSH-XX",
		// Add more translations as needed
		"HIPS/3": "HIPS/৩",
		"Mechanical": "যান্ত্রিক",
		"170 Ltr.": "১৭০ লিটার",
		"1535 mm": "১৫৩৫ মিমি",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes": "হ্যাঁ",
		"Recessed/ Grip": "Recessed/ Grip",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"R600a": "R৬০০a",
		"555 mm": "৫৫৫ মিমি",
		"Yes/1": "Yes/১",
		"Copper": "কপার",
		"Manual": "Manual",
		"RoHS Certified": "RoHS Certified",
		"N~ST": "N~ST",
		"184 Ltr.": "১৮৪ লিটার",
		"600VA": "৬০০VA",
		"630 mm": "৬৩০ মিমি",
		"V 0101 - RSCR": "V ০১০১ - RSCR",
		"V 0101 - 104": "V ০১০১ - ১০৪",
		"PS/3": "PS/৩",
		"48.4 ± 2 Kg": "৪৮.৪ ± ২ কেজি",
		"220 ~ 240/ 50 Hz": "২২০ ~ ২৪০/ ৫০ Hz",
		"Wire/2": "Wire/২",
		"108/ 108/ 52": "১০৮/ ১০৮/ ৫২",
		"53.4 ± 2 Kg": "৫৩.৪ ± ২ কেজি",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfd-a7x-gdsh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfdA7xGdshXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfd-a7x-gdsh-xx"
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
		"Gross Volume": "184 Ltr.",
		"Net Volume": "170 Ltr.",
		"Net Weight": "48.4 ± 2 Kg",
		"Gross Weight": "53.4 ± 2 Kg",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Rated Voltage/ Hz": "220 ~ 240/ 50 Hz",
		"Compressor Input Power (Watt)": "V 0101 - 104",
		"Compressor Type": "V 0101 - RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Lock": "Yes",
		"Refrigerant": "R600a",
		"Thermostat": "RoHS Certified",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "600VA",
		"Shelf (Material/No.)": "Wire/2",
		"Door Basket": "PS/3",
		"Interior Lamp": "Yes",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
		"Egg Tray or Pocket": "Yes",
		"Drawer": "HIPS/3",
		"Width/mm": "555 mm",
		"Depth/mm": "630 mm",
		"Height/mm": "1535 mm",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "108/ 108/ 52",
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
