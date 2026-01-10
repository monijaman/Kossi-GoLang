package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx seeds specifications/options for product 'mfc-c4h-gdne-xx'
type SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC4hGdneXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC4hGdneXx() *SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx {
	return &SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfc-c4h-gdne-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1790": "1790",
		"220-240V~ and 50Hz": "220-240V~ and 50Hz",
		"333 Ltr.": "333 Ltr.",
		"348 Ltr": "348 Ltr",
		"5": "5",
		"66/ 73 ± 2": "66/ 73 ± 2",
		"710": "710",
		"Glass/2": "Glass/2",
		"Manual": "Manual",
		"No": "No",
		"V 0201- R600a V 0301- R600a V 0302- R600a": "V 0201- R600a V 0301- R600a V 0302- R600a",
		"V 0201- RSCR V 0301- RSIR V 0302- RSIR": "V 0201- RSCR V 0301- RSIR V 0302- RSIR",
		"V 0201/0301/0302: Wide Voltage Design (140V-260V) N.B.: If out of voltage range(140V-260V), then suggested voltage stabilizer capacity is 2100VA.": "V 0201/0301/0302: Wide Voltage Design (140V-260V) N.B.: If out of voltage range(140V-260V), then suggested voltage stabilizer capacity is 2100VA.",
		"Wire/2": "Wire/2",
		"Yes(Glass/ Plastic)": "Yes(Glass/ Plastic)",
		"Yes/1": "Yes/1",
		// Add more translations as needed
	
		"Copper": "কপার",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Mechanical": "মেকানিক্যাল",
		"N ~ ST": "N ~ ST",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"R600a": "R600a",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes": "হ্যাঁ",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/2": "হ্যাঁ/২",
	
		"645": "৬৪৫",}
}

// Seed inserts specification records for the product identified by slug 'mfc-c4h-gdne-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx) Seed(db *gorm.DB) error {
	productSlug := "mfc-c4h-gdne-xx"
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
		"Model Name":          "MFC-C4H-GDNE-XX",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 0201- 145.7V 0301- 115V 0302- 115",
		"Compressor Type": "V 0201- RSCRV 0301- RSIRV 0302- RSIR",
		"Cooling Efect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth (mm)": "645",
		"Door Basket": "PS/4",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "348 Ltr",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height (mm)": "1740",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "67/ 48/ 24",
		"Lock": "Yes",
		"Net Volume": "333 Ltr.",
		"Operating voltage": "V 0201/0301/0302: Wide Voltage Design (140V-260V)N.B.: If out of voltage range(140V-260V), then suggested voltage stabilizer capacity is 2100VA.",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Refrigerant": "V 0201- R600aV 0301- R600aV 0302- R600a",
		"Reversible Door": "No",
		"Shelf Material": "Wire/2",
		"Shelf (Material/No.)": "Glass/2",
		"Star rating (BDS 1850:2012)": "5",
		"Temperature Control": "Mechanical",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes(Glass/ Plastic)",
		"Weight/Kg - Net/Packing": "66/ 73 ± 2",
		"Width": "650",
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
