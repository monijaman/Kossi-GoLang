package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC5hGdenXx seeds specifications/options for product 'marcel-mfe-c5h-gden-xx'
type SpecificationSeederRefrigeratorMarcelMfeC5hGdenXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC5hGdenXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC5hGdenXx() *SpecificationSeederRefrigeratorMarcelMfeC5hGdenXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC5hGdenXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c5h-gden-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdenXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                     "মার্সেল",
		"marcel-mfe-c5h-gden-xx":     "মার্সেল-mfe-c5h-gden-xx",
		"MFE-C5H-GDEN-XX":            "MFE-C5H-GDEN-XX",
		"Direct Cool":                "ডাইরেক্ট কুল",
		"358 Ltr.":                   "৩৫৮ লিটার",
		"345 Ltr.":                   "৩৪৫ লিটার",
		"68.5 \u00b1 2 Kg":           "৬৮.৫ ± ২ কেজি",
		"76 \u00b1 2 Kg":             "৭৬ ± ২ কেজি",
		"N ~ ST":                     "N ~ ST",
		"220~240/ 50/135":            "২২০~২৪০/৫০/১৩৫",
		"V 0102- 130\nV 0301- 123":   "V 0102- 130\nV 0301- 123",
		"V 0102- RSCR\nV 0301- RSCR": "V 0102- RSCR\nV 0301- RSCR",
		"Mechanical":                 "যান্ত্রিক",
		"Manual":                     "ম্যানুয়াল",
		"No":                         "না",
		"Recessed":                   "রিসেসড",
		"Yes":                        "হ্যাঁ",
		"R600a":                      "R600a",
		"Copper":                     "কপার",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "সাইক্লোপেন্টিন (পরিবেশবান্ধব, CFC/HCFC ফ্রি)",
		"2000VA":              "২০০০ভিএ",
		"Wire/2":              "ওয়্যার/2",
		"Yes/3":               "হ্যাঁ/3",
		"Yes/4":               "হ্যাঁ/4",
		"594 x 711 x 1820 mm": "৫৯৪ x ৭১১ x ১৮২০ মিমি",
		"625 x 745 x 1830 mm": "৬২৫ x ৭৪৫ x ১৮৩০ মিমি",
		"76/ 57/ 27":          "৭৬/ ৫৭/ ২৭",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
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
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c5h-gden-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdenXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c5h-gden-xx"
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
		"Model Name":          "MFE-C5H-GDEN-XX",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "N ~ ST",
		"Compressor": "RSCR",
		"Cooling Efect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "711",
		"Door Basket": "Yes/3",
		"Drawer": "Yes/4",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "358 Ltr.",
		"Gross Weight": "76± 2 Kg",
		"Handle (Recessed/ Grip):": "Recessed",
		"Height (mm)": "1825",
		"Ice Case": "Yes/2",
		"Ice Remover spoon": "Yes/1",
		"Ice Tray": "Yes/1",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "76/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "345 Ltr.",
		"Net Weight": "68.5± 2 Kg",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Voltage": "220~240/ 50/135",
		"Recommended voltage stabilizer capacity": "2000VA",
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
