package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMniF1nGdneDd seeds specifications/options for product 'marcel-mni-f1n-gdne-dd'
type SpecificationSeederRefrigeratorMarcelMniF1nGdneDd struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMniF1nGdneDd creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMniF1nGdneDd() *SpecificationSeederRefrigeratorMarcelMniF1nGdneDd {
	return &SpecificationSeederRefrigeratorMarcelMniF1nGdneDd{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mni-f1n-gdne-dd"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMniF1nGdneDd) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mni-f1n-gdne-dd": "মার্সেল-mni-f1n-gdne-dd",
		"MNI-F1N-GDNE-DD":        "MNI-F1N-GDNE-DD",
		"No Frost":               "নো ফ্রস্ট",
		"Direct Cool":            "ডাইরেক্ট কুল",
		"R600a":                  "R600a",
		"Mechanical":             "যান্ত্রিক",
		"Electronic":             "ইলেকট্রনিক",
		"220 ~ 240":              "২২০ ~ ২৪০",
		"50":                     "৫০",
		"Automatic":              "স্বয়ংক্রিয়",
		"Double Door":            "ডাবল দরজা",
		"Rotary":                 "রোটারি",
		"BLDC Inverter":          "বিএলডিসি ইনভার্টার",
		"No Frost Technology, LED Lighting, Anti-fungal Door Gasket, Dynamic Flow, Power Cooler": "নো ফ্রস্ট প্রযুক্তি, LED লাইটিং, অ্যান্টি-ফাঙ্গাল ডোর গ্যাসকেট, ডায়নামিক ফ্লো, পাওয়ার কুলার",
		"591 Ltr.":            "৫৯১ লিটার",
		"548 Ltr.":            "৫৪৮ লিটার",
		"115 ± 2 Kg":          "১১৫ ± ২ কেজি",
		"125 ± 2 Kg":          "১২৫ ± ২ কেজি",
		"865 x 725 x 1780 mm": "৮৬৫ x ৭২৫ x ১৭৮০ মিমি",
		"900 x 775 x 1885 mm": "৯০০ x ৭৭৫ x ১৮৮৫ মিমি",
		"39/39/18":            "৩৯/৩৯/১৮",
		"5":                   "৫",
		"1":                   "১",
		"Glass":               "গ্লাস",
		"GPPS/4":              "GPPS/4",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"Freezer < -18℃; Refrigerator 0℃ to +5℃": "ফ্রিজার: < -18℃; ফ্রিজ: 0℃ থেকে +5℃",
		"775 mm": "৭৭৫ মিমি",
		"GLASS/5": "GLASS/৫",
		"Steel": "স্টীল",
		"220~240 V/ 50 Hz": "২২০~২৪০ V/ ৫০ Hz",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"1885 mm": "১৮৮৫ মিমি",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Yes/1": "Yes/১",
		"900 mm": "৯০০ মিমি",
		"Copper": "কপার",
		"No-Frost": "নো ফ্রস্ট",
		"591": "৫৯১",
		"39/39/18 (Vertical Loading)": "৩৯/৩৯/১৮ (Vertical Loading)",
		"Built In": "বিল্ট ইন",
		"45.4 - 197": "৪৫.৪ - ১৯৭",
		"Glass/5": "Glass/৫",
		"Yes/ 1": "Yes/ ১",
		"N~T": "N~T",
		"548": "৫৪৮",

	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Manual": "ম্যানুয়াল",
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
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/2": "হ্যাঁ/২",
	
		"725 mm": "৭২৫ মিমি",
		"865 mm": "৮৬৫ মিমি",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mni-f1n-gdne-dd'
func (s *SpecificationSeederRefrigeratorMarcelMniF1nGdneDd) Seed(db *gorm.DB) error {
	productSlug := "marcel-mni-f1n-gdne-dd"
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
		"Bottle Pocket": "GPPS/4",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N~T",
		"Compressor": "BLDC Inverter",
		"Compressor Input Power (watt)": "45.4 - 197",
		"Condenser": "Steel",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Depth/mm": "725 mm",
		"Drawer": "Yes/1",
		"Egg Tray or Pocket": "Yes/ 1",
		"Gross Volume": "591",
		"Gross Weight": "125 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Built In",
		"Height/mm": "1780 mm",
		"Ice Box": "Yes/1",
		"Ice Case": "Yes/1",
		"Interior LED Lamp": "Yes",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/40Ft/20Ft": "39/39/18 (Vertical Loading)",
		"Lock": "No",
		"Net Volume": "548",
		"Net Weight": "115 ± 2 Kg",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz": "220~240 V/ 50 Hz",
		"Shelf (Material/ No.)": "GLASS/5",
		"Shelf (Material/No.)": "Glass/5",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Type": "No-Frost",
		"Vegetable Crisper": "Yes/1",
		"Water Dispenser": "Yes",
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
