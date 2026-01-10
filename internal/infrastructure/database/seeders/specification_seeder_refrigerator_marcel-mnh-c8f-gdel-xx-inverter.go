package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMnhC8fGdelXxInverter seeds specifications/options for product 'marcel-mnh-c8f-gdel-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMnhC8fGdelXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMnhC8fGdelXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMnhC8fGdelXxInverter() *SpecificationSeederRefrigeratorMarcelMnhC8fGdelXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMnhC8fGdelXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mnh-c8f-gdel-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMnhC8fGdelXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":         "মার্সেল",
				"marcel-mnh-c8f-gdel-xx-inverter":         "মার্সেল-mnh-c8f-gdel-xx-inverter",
		"MNH-C8F-GDEL-XX-INVERTER":   "MNH-C8F-GDEL-XX-INVERTER",
		// Add more translations as needed
		"Glass/3": "Glass/৩",
		"Glass/1": "Glass/১",
		"Yes/ 2": "Yes/ ২",
		"220-240/ 50 Hz": "২২০-২৪০/ ৫০ Hz",
		"Steel": "স্টীল",
		"Yes": "হ্যাঁ",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"74/82 ± 2": "৭৪/৮২ ± ২",
		"1775 mm": "১৭৭৫ মিমি",
		"Electronic": "ইলেকট্রনিক",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Yes/2": "Yes/২",
		"Yes/1": "Yes/১",
		"BLDC Inverter": "বিএলডিসি ইনভার্টার",
		"Copper": "কপার",
		"RoHS Certified": "RoHS Certified",
		"750 mm": "৭৫০ মিমি",
		"740 mm": "৭৪০ মিমি",
		"No-Frost": "নো ফ্রস্ট",
		"GPPS/2": "GPPS/২",
		"328 Ltr.": "৩২৮ লিটার",
		"T": "T",
		"Automatic": "স্বয়ংক্রিয়",
		"386 Ltr.": "৩৮৬ লিটার",
		"66/ 48/ 24": "৬৬/ ৪৮/ ২৪",

	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
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
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
	
		"GPPS/7": "GPPS/৭",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mnh-c8f-gdel-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMnhC8fGdelXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mnh-c8f-gdel-xx-inverter"
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
		"Model Name":          "MNH-C8F-GDEL-XX-INVERTER",
		"Bottle Pocket": "GPPS/7",
		"Capillary": "Copper",
		"Chilled Room": "Yes/1",
		"Chilled Room Door": "Yes/1",
		"Climatic Type (SN, N, ST, T)": "T",
		"Compressor": "BLDC Inverter",
		"Condenser": "Steel",
		"Cooling Efect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Depth/mm": "680 mm",
		"Egg Tray or Pocket": "Yes/ 2",
		"Gross Volume": "386 Ltr.",
		"Height/mm": "1735 mm",
		"Ice Box": "Yes/1",
		"Interior LED Lamp": "Yes",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "66/ 48/ 24",
		"Net Volume": "328 Ltr.",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz/ watt": "220-240/ 50 Hz",
		"Shelf (Material/ No.)": "Glass/1",
		"Shelf (Material/No.)": "Glass/3",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Thermostat": "RoHS Certified",
		"Twist Ice tray": "Yes/2",
		"Type": "No-Frost",
		"Vegetable Crisper": "Yes/1",
		"Weight(Kg):  Net/Packing": "74/82  ± 2",
		"Width/mm": "705 mm",
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
