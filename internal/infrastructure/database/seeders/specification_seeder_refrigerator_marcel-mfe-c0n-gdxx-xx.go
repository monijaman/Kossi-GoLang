package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx seeds specifications/options for product 'marcel-mfe-c0n-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx() *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c0n-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-c0n-gdxx-xx": "মার্সেল-mfe-c0n-gdxx-xx",
		"MFE-C0N-GDXX-XX":        "MFE-C0N-GDXX-XX",
		// Add more translations as needed
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"Copper": "কপার",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Manual": "ম্যানুয়াল",
		"Mechanical": "মেকানিক্যাল",
		"N ~ ST": "N ~ ST",
		"No": "না",
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
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes": "হ্যাঁ",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"309 Ltr.": "৩০৯ লিটার",
		"594": "৫৯৪",
		"708": "৭০৮",
		"78/ 57/ 27": "৭৮/ ৫৭/ ২৭",
		"Yes/3": "হ্যাঁ/3",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c0n-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC0nGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c0n-gdxx-xx"
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
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Class": "N ~ ST",
		"Compressor Input Power (Watt)": "V 0101- 117V 0201- 109",
		"Compressor Type": "V 0101- RSCRV 0201- RSCR",
		"Cooling Efect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00C to +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "708",
		"Door Basket": "Yes/3",
		"Drawer": "Yes/3",
		"Egg Case": "Yes/2",
		"Gross Volume": "309 Ltr.",
		"Handle (Recessed/ Grip)": "Recressed",
		"Height/mm": "1646",
		"Ice Box": "Yes/1",
		"Ice Case": "Yes/1",
		"Ice Remover spoon": "Yes/1",
		"Ice Tray": "Yes/1",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "78/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "270 Ltr.",
		"Polyurethane foam blowing agent": "CycloPentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Refrigerant": "V 0101- R600aV 0201- R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Box": "Yes/1",
		"Vegetable Box Cover": "Yes",
		"Weight/Kg - Net/Packing:": "64 / 69.5 ± 2 Kg",
		"Width/mm": "594",
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
