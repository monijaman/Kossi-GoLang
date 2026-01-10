package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcfB0eGdelGx seeds specifications/options for product 'marcel-mcf-b0e-gdel-gx'
type SpecificationSeederRefrigeratorMarcelMcfB0eGdelGx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcfB0eGdelGx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcfB0eGdelGx() *SpecificationSeederRefrigeratorMarcelMcfB0eGdelGx {
	return &SpecificationSeederRefrigeratorMarcelMcfB0eGdelGx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcf-b0e-gdel-gx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcfB0eGdelGx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                         "মার্সেল",
		"marcel-mcf-b0e-gdel-gx":         "মার্সেল-mcf-b0e-gdel-gx",
		"MCF-B0E-GDEL-GX":                "MCF-B0E-GDEL-GX",
		"Single Door":                    "সিঙ্গেল ডোর",
		"205 Liters":                     "২০৫ লিটার",
		"0 Liters":                       "০ লিটার",
		"962 x 585 x 852 mm (W x D x H)": "৯৬২ x ৫৮৫ x ৮৫২ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"47 kg":                          "৪৭ কেজি",
		"Silver":                         "সিলভার",
		"RSIR/RSCR":                      "RSIR/RSCR",
		"Direct Cool":                    "ডাইরেক্ট কুল",
		"Manual":                         "ম্যানুয়াল",
		"Mechanical":                     "মেকানিক্যাল",
		"Wire":                           "ওয়্যার",
		"1":                              "১",
		"No":                             "না",
		"Yes":                            "হ্যাঁ",
		"40 dB":                          "৪০ ডেসিবেল",
		"220-240V":                       "২২০-২৪০ ভোল্ট",
		"50":                             "৫০",
		"2 Years":                        "২ বছর",
		"5":                              "৫",
		"R134a/R600a":                    "R134a/R600a",
		"Single Door Freezer, Manual Defrost, Lock, Interior Lamp": "সিঙ্গেল ডোর ফ্রিজার, ম্যানুয়াল ডিফ্রস্ট, লক, ইন্টেরিয়র ল্যাম্প",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcf-b0e-gdel-gx'
func (s *SpecificationSeederRefrigeratorMarcelMcfB0eGdelGx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcf-b0e-gdel-gx"
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
			"Gross Volume": "205 Ltr.",
			"Net Volume": "205 Ltr.",
			"Net Weight(Kg)": "47±2",
			"Gross Weight(Kg)": "53±2",
			"Climatic Type (SN, N, ST, T)": "N ~ ST",
			"Rated Voltage/ Hz": "220-240V/ 50Hz",
			"Compressor Input Power (Watt)": "V 01.01- 134V 02.01- 98V 03.01- 98V 04.01- 98V 05.01- 98V 06.01- 118",
			"Compressor Type": "V 01.01- RSIRV 02.01- RSCRV 03.01- RSCRV 04.01- RSCRV 05.01- RSCRV 06.01- RSCRV 05.01- RSCR",
			"Wide voltage range": "V 05.01, 145V-260VV 06.01, 145V-260V",
			"Cooling Efect": "Freezer Cabinet Less than -180C",
			"Temperature Control (Electronic/  Mechanical):": "Mechanical",
			"Defrosting (Automatic/ Manual)": "Manual",
			"Reversible Door": "N/A",
			"Handle (Recessed/ Grip)": "Yes",
			"Lock": "Yes",
			"Refrigerant": "V 01.01-R134aV 02.01-R600aV 03.01-R600aV 04.01-R600aV 05.01-R600aV 06.01-R600a",
			"Condenser": "Steel",
			"Capillary": "Copper",
			"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
			"Recommended voltage stabilizer capacity": "V 01.01",
			"Exterior Material": "Painted Steel (PCM)",
			"Interior Material": "Embossed Aluminium (Al2)",
			"Ice/ cold water Dispenser": "No",
			"Shelf (Material/No)": "Wire/1",
			"Drawer": "No",
			"Basket": "Wire/1",
			"Interior Lamp": "Yes",
			"Width/mm": "1010",
			"Depth/mm": "620",
			"Height/mm": "890",
			"Loading quantity(20ft/40ft.40 HQ)": "36/76/114",
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
