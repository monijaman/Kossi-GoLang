package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC1gGdxxXxInverter seeds specifications/options for product 'marcel-mfc-c1g-gdxx-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfcC1gGdxxXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC1gGdxxXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC1gGdxxXxInverter() *SpecificationSeederRefrigeratorMarcelMfcC1gGdxxXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfcC1gGdxxXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfc-c1g-gdxx-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC1gGdxxXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfc-c1g-gdxx-xx-inverter": "মার্সেল-mfc-c1g-gdxx-xx-inverter",
		"MFC-C1G-GDXX-XX-INVERTER":        "MFC-C1G-GDXX-XX-INVERTER",

		"Direct Cool":         "ডিরেক্ট কুল",
		"177 Ltr.":            "১৭৭ লিটার",
		"175 Ltr.":            "১৭৫ লিটার",
		"50 ± 2 Kg":           "৫০ ± ২ কেজি",
		"R600a":               "R600a",
		"Mechanical":          "ম্যানুয়াল/মেকানিক্যাল",
		"220 ~ 240":           "২২0 ~ ২৪0",
		"555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"61/ 69 ± 2": "৬১/ ৬৯ ± ২",
		"317 Ltr": "৩১৭ Ltr",
		"337 Ltr": "৩৩৭ Ltr",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes": "হ্যাঁ",
		"Yes (Glass/ plastic)": "Yes (Glass/ plastic)",
		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",
		"710": "৭১০",
		"V02.01- 118 V03.01- 118 V0302-118 V0401-33.78~126.46": "V০২.০১- ১১৮ V০৩.০১- ১১৮ V০৩০২-১১৮ V০৪০১-৩৩.৭৮~১২৬.৪৬",
		"Yes/2": "Yes/২",
		"CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Wire/3": "Wire/৩",
		"Copper": "কপার",
		"Manual": "Manual",
		"No": "না",
		"N~ST": "N~ST",
		"2": "২",
		"V 0301- R600a": "V ০৩০১- R৬০০a",
		"V02.01- RSCR V03.01- RSCR V0302-RSCR V0401-BLDC": "V০২.০১- RSCR V০৩.০১- RSCR V০৩০২-RSCR V০৪০১-BLDC",
		"69/ 48/ 24": "৬৯/ ৪৮/ ২৪",
		"1660": "১৬৬০",
		"5": "৫",
		"V 0301/V 0302: Wide Voltage Design (150V-260V) N.B.: If out of voltage range(150V-260V), then suggested voltage stabilizer capacity is 2100VA. V 0401:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.": "V ০৩০১/V ০৩০২: Wide Voltage Design (১৫০V-২৬০V) N.B.: If out of voltage range(১৫০V-২৬০V), then suggested voltage stabilizer capacity is ২১০০VA. V ০৪০১:Wide Voltage Design (৭৫V-২৬৪V) N.B.: If out of voltage range(৭৫V-২৬৪V) then suggested voltage stabilizer capacity is ২১০০VA.",
		"Wire/2": "Wire/২",
		"Recessed/Grip": "Recessed/Grip",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfc-c1g-gdxx-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfcC1gGdxxXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfc-c1g-gdxx-xx-inverter"
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
		"Gross Volume": "337 Ltr",
		"Net Volume": "317 Ltr",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Compressor Input Power (Watt)": "V02.01- 118 V03.01- 118 V0302-118 V0401-33.78~126.46",
		"Compressor Type": "V02.01- RSCR V03.01- RSCR V0302-RSCR V0401-BLDC",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Reversible Door": "No",
		"Handle (Recessed/Grip)": "Recessed/Grip",
		"Lock": "Yes",
		"Refrigerant": "V 0301- R600a",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "CycloPentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Operating voltage": "V 0301/V 0302: Wide Voltage Design (150V-260V) N.B.: If out of voltage range(150V-260V), then suggested voltage stabilizer capacity is 2100VA. V 0401:Wide Voltage Design (75V-264V) N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Shelf (Material/ No.)": "Wire/3",
		"Door Basket": "2",
		"Interior Lamp": "No",
		"Vegetable Crisper Cover": "Yes (Glass/ plastic)",
		"Egg Tray or Pocket": "Yes/2",
		"Can Storage Dispenser": "No",
		"Deodorizer": "No",
		"Shelf (Material/No.)": "Wire/2",
		"Drawer": "No",
		"Width/mm": "710",
		"Depth/mm": "710",
		"Height/mm": "1660",
		"Weight/Kg - Net/Packing": "61/ 69 ± 2",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "69/ 48/ 24",
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
