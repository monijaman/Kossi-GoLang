package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMniE6cGdelDd seeds specifications/options for product 'marcel-mni-e6c-gdel-dd'
type SpecificationSeederRefrigeratorMarcelMniE6cGdelDd struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMniE6cGdelDd creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMniE6cGdelDd() *SpecificationSeederRefrigeratorMarcelMniE6cGdelDd {
	return &SpecificationSeederRefrigeratorMarcelMniE6cGdelDd{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mni-e6c-gdel-dd"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMniE6cGdelDd) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mni-e6c-gdel-dd": "মার্সেল-mni-e6c-gdel-dd",
		"MNI-E6C-GDEL-DD":        "MNI-E6C-GDEL-DD",

		// specs values -> Bangla
		"563 Ltr.":            "৫৬৩ লিটার",
		"501 Ltr.":            "৫০১ লিটার",
		"BLDC Inverter":       "বিএলডিসি ইনভার্টার",
		"220~240 V/ 50 Hz":    "২২০~২৪০ V/ ৫০ Hz",
		"Mechanical":          "যান্ত্রিক",
		"R600a":               "R600a",
		"865 x 715 x 1725 mm": "৮৬৫ x ৭১৫ x ১৭২৫ মিমি",
		"890 x 775 x 1770 mm": "৮৯০ x ৭৭৫ x ১৭৭০ মিমি",
		"101 / 115.5 ± 2":     "১০১ / ১১৫.৫ ± ২",
		"39/ 39/ 18":          "৩৯/ ৩৯/ ১৮",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"Copper": "কপার",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Manual": "ম্যানুয়াল",
		"N ~ ST": "N ~ ST",
		"No": "না",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
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
	
		"100 % Copper": "১০০% কপার",
		"865 mm": "৮৬৫ মিমি",
		"Automatic": "স্বয়ংক্রিয়",
		"Built In": "বিল্ট ইন",
		"GPPS/5": "GPPS/5",
		"Glass/5": "Glass/৫",
		"Non-Frost": "নন-ফ্রস্ট",
		"ST": "এসটি",
		"Yes/ 1": "Yes/ ১",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mni-e6c-gdel-dd'
func (s *SpecificationSeederRefrigeratorMarcelMniE6cGdelDd) Seed(db *gorm.DB) error {
	productSlug := "marcel-mni-e6c-gdel-dd"
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
		"Bottle Pocket": "GPPS/5",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "ST",
		"Compressor": "BLDC Inverter",
		"Condenser": "100 % Copper",
		"Cooling Effect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Depth/mm": "715 mm",
		"Drawer": "Yes/1",
		"Egg Tray or Pocket": "Yes/ 1",
		"Gross Volume": "563 Ltr.",
		"Handle (Recessed/ Grip)": "Built In",
		"Height/mm": "1725 mm",
		"Ice Box": "Yes/1",
		"Ice Case": "Yes/1",
		"Interior LED Lamp": "Yes",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "39/ 39/ 18",
		"Lock": "No",
		"Net Volume": "501 Ltr.",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz/ watt": "220~240 V/ 50 Hz",
		"Refrigerant": "R600a",
		"Shelf (Material/ No.)": "GPPS/5",
		"Shelf (Material/No.)": "Glass/5",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Non-Frost",
		"Vegetable Crisper": "Yes/1",
		"Water Dispenser": "No",
		"Weight/Kg - Net/Packing": "101/115.5 ± 2",
		"Width/mm": "865 mm",
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
