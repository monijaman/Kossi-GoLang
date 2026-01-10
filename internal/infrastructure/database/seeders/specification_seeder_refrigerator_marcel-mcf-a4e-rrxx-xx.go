package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcfA4eRrxxXx seeds specifications/options for product 'marcel-mcf-a4e-rrxx-xx'
type SpecificationSeederRefrigeratorMarcelMcfA4eRrxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcfA4eRrxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcfA4eRrxxXx() *SpecificationSeederRefrigeratorMarcelMcfA4eRrxxXx {
	return &SpecificationSeederRefrigeratorMarcelMcfA4eRrxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcf-a4e-rrxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcfA4eRrxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                         "মার্সেল",
		"marcel-mcf-a4e-rrxx-xx":         "মার্সেল-mcf-a4e-rrxx-xx",
		"MCF-A4E-RRXX-XX":                "MCF-A4E-RRXX-XX",
		"Single Door":                    "সিঙ্গেল ডোর",
		"145 Liters":                     "১৪৫ লিটার",
		"0 Liters":                       "০ লিটার",
		"744 x 585 x 845 mm (W x D x H)": "৭৪৪ x ৫৮৫ x ৮৪৫ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"34 kg":                          "৩৪ কেজি",
		"Silver":                         "সিলভার",
		"RSCR":                           "RSCR",
		"Direct Cool":                    "ডাইরেক্ট কুল",
		"Manual":                         "ম্যানুয়াল",
		"Mechanical":                     "মেকানিক্যাল",
		"Wire":                           "ওয়্যার",
		"1":                              "১",
		"No":                             "না",
		"40 dB":                          "৪০ ডেসিবেল",
		"220-240V":                       "২২০-২৪০ ভোল্ট",
		"50":                             "৫০",
		"2 Years":                        "২ বছর",
		"5":                              "৫",
		"R600a":                          "R600a",
		"Single Door Freezer, Manual Defrost, Reversible Door": "সিঙ্গেল ডোর ফ্রিজার, ম্যানুয়াল ডিফ্রস্ট, রিভার্সিবল ডোর",
		// Add more translations as needed
		"145 Ltr.": "১৪৫ লিটার",
		"Foaming Door": "Foaming Door",
		"V 05.02-(600VA) V 06.01-N/A V 06.02-N/A V 07.01-(600VA)": "V ০৫.০২-(৬০০VA) V ০৬.০১-N/A V ০৬.০২-N/A V ০৭.০১-(৬০০VA)",
		"RSCR for All Version": "RSCR for All Version",
		"V 05.02-R600a V 06.01-R600a V 06.02-R600a V 07.01-R600a": "V ০৫.০২-R৬০০a V ০৬.০১-R৬০০a V ০৬.০২-R৬০০a V ০৭.০১-R৬০০a",
		"Steel": "স্টীল",
		"Yes": "হ্যাঁ",
		"795": "৭৯৫",
		"39 ± 2 (Kg)": "৩৯ ± ২ ( কেজি)",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"Copper": "কপার",
		"RoHS Certified": "RoHS Certified",
		"Painted Steel (PCM)": "Painted Steel (PCM)",
		"N/A": "N/A",
		"Wire/1": "Wire/১",
		"V 05.02-140 V 06.01-118 V 06.02-118 V 07.01-140": "V ০৫.০২-১৪০ V ০৬.০১-১১৮ V ০৬.০২-১১৮ V ০৭.০১-১৪০",
		"CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentane [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"890": "৮৯০",
		"50/106/159": "৫০/১০৬/১৫৯",
		"34 ± 2 (Kg)": "৩৪ ± ২ ( কেজি)",
		"N ~ T": "N ~ T",
		"620": "৬২০",
		"Freezer Cabinet Less than -18℃": "Freezer Cabinet Less than -১৮℃",
		"Embossed Aluminium Sheet": "Embossed Aluminium Sheet",

	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[Eco-friendly (১০০% CFC &HCFC Free) Green Technology]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[Eco-friendly (১০০% CFC &HCFC Free) Green Technology]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "Freezer Cabinet Less than -১৮০CRefrigerator Cabinet ০০Cto +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃Refrigerator Cabinet ০℃ to +৫℃",
		"N ~ ST": "N ~ ST",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcf-a4e-rrxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMcfA4eRrxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcf-a4e-rrxx-xx"
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
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}
	specs := map[string]string{
		"Basket": "Wire/1",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N ~ T",
		"Compressor Input Power (Watt)": "V 05.02-140V 06.01-118V 06.02-118V 07.01-140",
		"Compressor Type": "RSCR for All Version",
		"Condenser": "Steel",
		"Cooling Efect": "Freezer Cabinet Less than -18℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "585",
		"Drawer": "No",
		"Energy Rating": "N/A",
		"Exterior Material": "Painted Steel (PCM)",
		"Gross Volume": "145 Ltr.",
		"Gross Weight": "39 ± 2 (Kg)",
		"Handle (Recessed/ Grip)": "Yes",
		"Height/mm": "845",
		"Ice/ cold water Dispenser": "No",
		"Interior Lamp": "No",
		"Interior Material": "Embossed Aluminium Sheet",
		"Loading quantity (20ft/40ft.40HQ)": "50/106/159",
		"Lock": "N/A",
		"Net Volume": "145 Ltr.",
		"Net Weight": "34 ± 2 (Kg)",
		"Polyurethane foam blowing agent": "CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity (600VA)": "V 05.02-(600VA)V 06.01-N/AV 06.02-N/AV 07.01-(600VA)",
		"Refrigerant": "V 05.02-R600aV 06.01-R600aV 06.02-R600aV 07.01-R600a",
		"Reversible Door": "Foaming Door",
		"Shelf (Material/No)": "Wire/1",
		"Temperature Control (Electronic/  Mechanical):": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Width/mm": "744",
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
