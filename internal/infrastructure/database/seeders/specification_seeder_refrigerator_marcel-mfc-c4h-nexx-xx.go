package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC4hNexxXx seeds specifications/options for product 'mfc-c4h-nexx-xx'
type SpecificationSeederRefrigeratorMarcelMfcC4hNexxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC4hNexxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC4hNexxXx() *SpecificationSeederRefrigeratorMarcelMfcC4hNexxXx {
	return &SpecificationSeederRefrigeratorMarcelMfcC4hNexxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfc-c4h-nexx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC4hNexxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"MFC-C4H-NEXX-XX":        "MFC-C4H-NEXX-XX",
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
		"R600a": "R600a",
		"V 0103: 105V-276V (Wide Voltage Design)": "V 0103: 105V-276V (Wide Voltage Design)",
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
	}
}

// Seed inserts specification records for the product identified by slug 'mfc-c4h-nexx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfcC4hNexxXx) Seed(db *gorm.DB) error {
	productSlug := "mfc-c4h-nexx-xx"
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
		"Model Name":          "MFC-C4H-NEXX-XX",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "710",
		"Door Basket": "No",
		"Gross Volume": "348 Ltr",
		"Height (mm)": "1790",
		"Net Volume": "333 Ltr.",
		"Operating voltage": "V 0103: 105V-276V (Wide Voltage Design)",
		"Voltage": "220-240V~ and 50Hz",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf Material": "Wire/2",
		"Shelf (Material/No.)": "Glass/2",
		"Star rating (BDS 1850:2012)": "5",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes(Glass/ Plastic)",
		"Weight/Kg - Net/Packing": "66/ 73 ± 2",
		"Width": "710",
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
