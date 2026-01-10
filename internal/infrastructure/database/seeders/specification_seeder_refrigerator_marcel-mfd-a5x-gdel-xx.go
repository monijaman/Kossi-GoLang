package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfdA5xGdelXx seeds specifications/options for product 'marcel-mfd-a5x-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfdA5xGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfdA5xGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfdA5xGdelXx() *SpecificationSeederRefrigeratorMarcelMfdA5xGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfdA5xGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfd-a5x-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfdA5xGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfd-a5x-gdel-xx": "মার্সেল-mfd-a5x-gdel-xx",
		"MFD-A5X-GDEL-XX":        "MFD-A5X-GDEL-XX",
		// Add more translations as needed
		"47 ± 2Kg":         "৪৭ ± ২ কেজি",
		"Mechanical":       "যান্ত্রিক",
		"Tempered GLass/2": "Tempered GLass/২",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes":            "হ্যাঁ",
		"220~240/ 50":    "২২০~২৪০/ ৫০",
		"Recessed/ Grip": "Recessed/ Grip",
		"555":            "৫৫৫",
		"Direct Cool":    "ডাইরেক্ট কুল",
		"112/112 /52":    "১১২/১১২ /৫২",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"R600a":          "R৬০০a",
		"Yes/1":          "Yes/১",
		"V 0101- 88":     "V ০১০১- ৮৮",
		"Copper":         "কপার",
		"Manual":         "Manual",
		"RoHS Certified": "RoHS Certified",
		"V 0101- RSCR":   "V ০১০১- RSCR",
		"No":             "না",
		"T":              "T",
		"595":            "৫৯৫",
		"43 ± 2Kg":       "৪৩ ± ২ কেজি",
		"Low Voltage (140~280V ) Voltage Stabilizer is not required": "Low Voltage (১৪০~২৮০V ) Voltage Stabilizer is not required",
		"150 Ltr.": "১৫০ লিটার",
		"Wire/2":   "Wire/২",
		"1490":     "১৪৯০",

		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST":        "N ~ ST",
		"N~ST":          "N~ST",
		"PVC/1":         "PVC/১",
		"PVC/2":         "PVC/২",
		"PVC/3":         "PVC/৩",
		"RSCR":          "RSCR",
		"RSIR":          "RSIR",
		"Wire/1":        "Wire/১",
		"Wire/3":        "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/2":         "হ্যাঁ/২",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfd-a5x-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfdA5xGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfd-a5x-gdel-xx"
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
		"Type":                              456,
		"Lock":                              361,
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}
	specs := map[string]string{
		"Brand":               "Marcel",
		"Model Name":          "MFD-A5X-GDEL-XX",
		"Can Storage Dispenser":                   "No",
		"Capillary":                               "Copper",
		"Climatic Type (SN, N, ST, T)":            "T",
		"Compressor Input Power (Watt)":           "V 0101- 88",
		"Compressor Type":                         "V 0101- RSCR",
		"Cooling Effect":                          "Freezer Cabinet Less than -18℃     Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)":          "Manual",
		"Deodorizer":                              "No",
		"Depth (mm)":                                "556",
		"Door Basket":                             "PS/2",
		"Drawer":                                  "No",
		"Egg Tray or Pocket":                      "Yes",
		"Gross Weight":                            "47 ± 2Kg",
		"Handle (Recessed/ Grip)":                 "Recessed/ Grip",
		"Height (mm)":                               "1440",
		"Interior Lamp":                           "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":      "112/112 /52",
		"Lock":                                    "Yes",
		"Net Weight":                              "43 ± 2Kg",
		"Polyurethane foam blowing agent":         "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Rated Voltage/Hz":                        "220~240/ 50",
		"Recommended voltage stabilizer capacity": "Low Voltage (140~280V ) Voltage Stabilizer is not required",
		"Refrigerant":                             "R600a",
		"Reversible Door":                         "No",
		"Shelf (Material/ No.)":                   "Tempered GLass/2",
		"Shelf (Material/No.)":                    "Wire/2",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Thermostat":              "RoHS Certified",
		"Total Volume":            "150 Ltr.",
		"Type":                    "Direct Cool",
		"Vegetable Crisper":       "Yes/1",
		"Vegetable Crisper Cover": "Yes",
		"Width":                "512",
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
