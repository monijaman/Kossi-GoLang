package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcfA4eGeleDl seeds specifications/options for product 'marcel-mcf-a4e-gele-dl'
type SpecificationSeederRefrigeratorMarcelMcfA4eGeleDl struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcfA4eGeleDl creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcfA4eGeleDl() *SpecificationSeederRefrigeratorMarcelMcfA4eGeleDl {
	return &SpecificationSeederRefrigeratorMarcelMcfA4eGeleDl{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcf-a4e-gele-dl"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcfA4eGeleDl) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                         "মার্সেল",
		"marcel-mcf-a4e-gele-dl":         "মার্সেল-mcf-a4e-gele-dl",
		"MCF-A4E-GELE-DL":                "MCF-A4E-GELE-DL",
		"Single Door":                    "সিঙ্গেল ডোর",
		"145 Liters":                     "১৪৫ লিটার",
		"0 Liters":                       "০ লিটার",
		"744 x 585 x 853 mm (W x D x H)": "৭৪৪ x ৫৮৫ x ৮৫৩ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"37 kg":                          "৩৭ কেজি",
		"Silver":                         "সিলভার",
		"RSCR":                           "RSCR",
		"Direct Cool":                    "ডাইরেক্ট কুল",
		"Manual":                         "ম্যানুয়াল",
		"Electronic":                     "ইলেকট্রনিক",
		"Wire":                           "ওয়্যার",
		"1":                              "১",
		"No":                             "না",
		"Yes":                            "হ্যাঁ",
		"40 dB":                          "৪০ ডেসিবেল",
		"220-240V":                       "২২০-২৪০ ভোল্ট",
		"50":                             "৫০",
		"2 Years":                        "২ বছর",
		"5":                              "৫",
		"R600a":                          "R600a",
		"Single Door Freezer, Manual Defrost, Electronic Control, Lock, Interior Lamp, Reversible Door": "সিঙ্গেল ডোর ফ্রিজার, ম্যানুয়াল ডিফ্রস্ট, ইলেকট্রনিক কন্ট্রোল, লক, ইন্টেরিয়র ল্যাম্প, রিভার্সিবল ডোর",
		// Add more translations as needed
		"145 Ltr.": "১৪৫ লিটার",
		"Foaming Door": "Foaming Door",
		"RSCR for All Version": "RSCR for All Version",
		"42 ± 2 (Kg)": "৪২ ± ২ ( কেজি)",
		"37 ± 2 (Kg)": "৩৭ ± ২ ( কেজি)",
		"Steel": "স্টীল",
		"V 06.03-64 V 07.02-104": "V ০৬.০৩-৬৪ V ০৭.০২-১০৪",
		"V 06.03-N/A V 07.02-(600VA)": "V ০৬.০৩-N/A V ০৭.০২-(৬০০VA)",
		"795": "৭৯৫",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"Copper": "কপার",
		"RoHS Certified": "RoHS Certified",
		"V 06.03-R600a V 07.02-R600a": "V ০৬.০৩-R৬০০a V ০৭.০২-R৬০০a",
		"Painted Steel (PCM)": "Painted Steel (PCM)",
		"N/A": "N/A",
		"Wire/1": "Wire/১",
		"CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentane [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"890": "৮৯০",
		"50/106/159": "৫০/১০৬/১৫৯",
		"N ~ T": "N ~ T",
		"620": "৬২০",
		"Freezer Cabinet Less than -18℃": "Freezer Cabinet Less than -১৮℃",
		"Embossed Aluminium Sheet": "Embossed Aluminium Sheet",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcf-a4e-gele-dl'
func (s *SpecificationSeederRefrigeratorMarcelMcfA4eGeleDl) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcf-a4e-gele-dl"
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
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Dimensions":                  25,
		"Weight":                      80,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
		"Special Features":            69,
	}

	
    		
    specs := map[string]string{
		"Type": "Direct Cool",
		"Gross Volume": "145 Ltr.",
		"Net Volume": "145 Ltr.",
		"Gross Weight": "42 ± 2 (Kg)",
		"Net Weight": "37 ± 2 (Kg)",
		"Climatic Type (SN, N, ST, T)": "N ~ T",
		"Rated Operating Voltage and Frequency": "220 ~ 240/ 50",
		"Compressor Input Power (Watt)": "V 06.03-64 V 07.02-104",
		"Compressor Type": "RSCR for All Version",
		"Cooling Efect": "Freezer Cabinet Less than -18℃",
		"Energy Rating": "N/A",
		"Temperature Control (Electronic/ Mechanical):": "Electronic",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Reversible Door": "Foaming Door",
		"Handle (Recessed/ Grip)": "Yes",
		"Lock": "Yes",
		"Refrigerant": "V 06.03-R600a V 07.02-R600a",
		"Condenser": "Steel",
		"Thermostat": "RoHS Certified",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity (600VA)": "V 06.03-N/A V 07.02-(600VA)",
		"Exterior Material": "Painted Steel (PCM)",
		"Interior Material": "Embossed Aluminium Sheet",
		"Ice/ cold water Dispenser": "No",
		"Shelf (Material/No)": "Wire/1",
		"Drawer": "No",
		"Basket": "Wire/1",
		"Interior Lamp": "Yes",
		"Width/mm": "795",
		"Depth/mm": "620",
		"Height/mm": "890",
		"Loading quantity (20ft/40ft.40HQ)": "50/106/159",
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
