package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC2xGdelXx seeds specifications/options for product 'marcel-mfe-c2x-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfeC2xGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC2xGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC2xGdelXx() *SpecificationSeederRefrigeratorMarcelMfeC2xGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC2xGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c2x-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC2xGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                             "মার্সেল",
		"marcel-mfe-c2x-gdel-xx":             "মার্সেল-mfe-c2x-gdel-xx",
		"MFE-C2X-GDEL-XX":                    "MFE-C2X-GDEL-XX",
		"Direct Cool":                        "ডাইরেক্ট কুল",
		"341 Ltr":                            "৩৪১ লিটার",
		"341 Ltr.":                           "৩৪১ লিটার",
		"320 Ltr.":                           "৩২০ লিটার",
		"61 ± 2 Kg":                          "৬১ ± ২ কেজি",
		"68 ± 2 Kg":                          "৬৮ ± ২ কেজি",
		"N ~ ST":                             "এন ~ এসটি",
		"220-240V~/50Hz":                     "২২০-২৪০V~/৫০Hz",
		"220~240/ 50/130":                    "২২০~২৪০/ ৫০/১৩০",
		"RSCR":                               "আরএসসি আর",
		"Mechanical":                         "ম্যানুয়াল",
		"Manual":                             "ম্যানুয়াল",
		"Recessed/ Grip":                     "রিসেসড/ গ্রিপ",
		"Yes":                                "হ্যাঁ",
		"No":                                 "না",
		"Wire/2":                             "ওয়্যার/২",
		"PVC/4":                              "পিভিসি/৪",
		"Copper":                             "কপার",
		"RoHS Certified":                     "RoHS সার্টিফায়েড",
		"Cyclopentene":                       "সাইক্লোপেন্টিন",
		"2000VA or More":                     "২০০০VA বা বেশি",
		"Freezer Cabinet Less than -18 ̊C":   "ফ্রিজার কেবিনেট: -১৮°C-এর নিচে",
		"Refrigerator Cabinet 0 ̊C to +5 ̊C": "রেফ্রিজারেটর: ০°C থেকে +৫°C",
		"585 x 711 x 1726 mm": "৫৮৫ x ৭১১ x ১৭২৬ মিমি",
		"625 x 745 x 1730 mm": "৬২৫ x ৭৪৫ x ১৭৩০ মিমি",
		"77/ 57/ 27":                         "৭৭/ ৫৭/ ২৭",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"R600a": "R600a",
		"RSIR": "RSIR",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c2x-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC2xGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c2x-gdel-xx"
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
		"Model Name":          "MFE-C2X-GDEL-XX",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N ~ ST",
		"Compressor": "RSCR",
		"Cooling Efect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "711",
		"Door Basket": "PVC/4",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "341 Ltr.",
		"Gross Weight": "68± 2Kg",
		"Handle (Recessed/ Grip):": "Recessed/ Grip",
		"Height (mm)": "1726",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "77/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "320 Ltr.",
		"Net Weight": "61± 2Kg",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage": "220~240/ 50/130",
		"Recommended voltage stabilizer capacity": "2000VA or More",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf Material": "Wire/2",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
		"Width": "585",
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
