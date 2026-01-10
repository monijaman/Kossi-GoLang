package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXx seeds specifications/options for product 'marcel-mfe-c1b-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC1bGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC1bGdxxXx() *SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c1b-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfe-c1b-gdxx-xx": "মার্সেল-mfe-c1b-gdxx-xx",
		"MFE-C1B-GDXX-XX":        "MFE-C1B-GDXX-XX",
		"Direct Cool":            "ডাইরেক্ট কুল",
		"312 Ltr":                "৩১২ লিটার",
		"290 Ltr.":               "২৯০ লিটার",
		"59.22 ± 2 Kg":           "৫৯.২২ ± ২ কেজি",
		"66.06 ± 2 Kg":           "৬৬.০৬ ± ২ কেজি",
		"N ~ ST":                 "এন ~ এসটি",
		"220-240V~/50Hz":         "২২০-২৪০V~/৫০Hz",
		"V 0101 - RSIR V 0201 - RSIR V 0301 - RSIR V 0302 - RSIR V 0501 - RSIR":      "V ০১০১ - RSIR V ০২০১ - RSIR V ০৩০১ - RSIR V ০৩০২ - RSIR V ০৫০১ - RSIR",
		"V 0101 - R600a V 0201 - R600a V 0301 - R600a V 0302 - R600a V 0501 - R600a": "V ০১০১ - R৬০০a V ০২০১ - R৬০০a V ০৩০১ - R৬০০a V ০৩০২ - R৬০০a V ০৫০১ - R৬০০a",
		"Mechanical":          "ম্যানুয়াল",
		"Manual":              "ম্যানুয়াল",
		"Wire":                "ওয়্যার",
		"2":                   "২",
		"PVC/3":               "পিভিসি/৩",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"Copper":              "কপার",
		"RoHS Certified":      "RoHS সার্টিফায়েড",
		"Cyclopentene":        "সাইক্লোপেন্টিন",
		"594 x 708 x 1620 mm": "৫৯৪ x ৭০৮ x ১৬২০ মিমি",
		"625 x 745 x 1630 mm": "৬২৫ x ৭৪৫ x ১৬৩০ মিমি",
		"79/ 54/ 27":          "৭৯/ ৫৪/ ২৭",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"R600a": "R600a",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"1620": "১৬২০",
		"594": "৫৯৪",
		"708": "৭০৮",
		"Recressed/ Grip/ Built-in": "রিসেসড/ গ্রিপ/ বিল্ট-ইন",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c1b-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC1bGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c1b-gdxx-xx"
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
		"Climate Type (SN, N, ST, T)": "N ~ ST",
		"Compressor Input Power (Watt)": "V 0101 - 145.7V 0201 - 145.7V 0301 - 117V 0302 - 117V 0501 - 123",
		"Compressor Type": "V 0101 - RSIRV 0201 - RSIRV 0301 - RSIRV 0302 - RSIRV 0501 - RSIR",
		"Cooling Efect": "Freezer Cabinet Less than -18 ̊CRefrigerator Cabinet 0 ̊C to +5 ̊C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "708",
		"Door Basket": "PVC/3",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "312 Ltr.",
		"Gross Weight": "66.06 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recressed/ Grip/ Built-in",
		"Height/mm": "1620",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "79/ 54/ 27",
		"Lock": "Yes",
		"Net Volume": "290 Ltr.",
		"Net Weight": "59.22 ± 2 Kg",
		"Polyurethane foam blowing agent": "CycloPentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Recommended voltage stabilizer capacity": "V 0201/0202/0203/0301: Wide Voltage Design (145V-253V)N.B.: If out of voltage range (145V-253V) then suggested voltage stabilizer capacity is 2100VA.V 0501:Wide Voltage Design (75V-264V)N.B.: If out of voltage range(75V-264V) then suggested voltage stabilizer capacity is 2100VA.",
		"Refrigerant": "V 0101 - R600aV 0201 - R600aV 0301 - R600aV 0302 - R600aV 0501 - R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes (Plastic)",
		"Vegetable Crisper Cover": "Yes (ABS/ PS)",
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
