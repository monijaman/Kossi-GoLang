package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfa2a3NexxXx seeds specifications/options for product 'marcel-mfa-2a3-nexx-xx'
type SpecificationSeederRefrigeratorMarcelMfa2a3NexxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfa2a3NexxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfa2a3NexxXx() *SpecificationSeederRefrigeratorMarcelMfa2a3NexxXx {
	return &SpecificationSeederRefrigeratorMarcelMfa2a3NexxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-2a3-nexx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfa2a3NexxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-2a3-nexx-xx": "মার্সেল-mfa-2a3-nexx-xx",
		"MFA-2A3-NEXX-XX":        "MFA-2A3-NEXX-XX",
		"176 Ltr.":               "176 লিটার",
		"213 Ltr.":               "213 লিটার",
		"545 x 640 x 1510 mm":    "545 x 640 x 1510 মিমি",
		"Net - Width: 545 mm; Depth: 640 mm; Height: 1510 mm (545 x 640 x 1510 mm)": "নেট - প্রস্থ: 545 মিমি; গভীরতা: 640 মিমি; উচ্চতা: 1510 মিমি (545 x 640 x 1510 মিমি)",
		"45.5 ± 2 Kg":      "45.5 ± 2 কেজি",
		"RSCR":             "RSCR",
		"Manual":           "ম্যানুয়াল",
		"Mechanical":       "মেকানিক্যাল",
		"R600a":            "R600a",
		"220 ~ 240/ 50 Hz": "220 ~ 240/ 50 হার্জ",
		"50":               "50",
		"12":               "12",
		"Refrigerator: 2 wire shelves; Freezer: 2 wire shelves": "রেফ্রিজারেটর: 2 তারের তাক; ফ্রিজার: 2 তারের তাক",
		"Wire":                               "তার",
		"Refrigerator: 2 PVC; Freezer: None": "রেফ্রিজারেটর: 2 PVC; ফ্রিজার: কোনোটি নয়",
		"1":                                  "1",
		"V 1101 - 102":                       "V 1101 - 102",
		"V 1102 - 102":                       "V 1102 - 102",
		"V 1301 - 108.6":                     "V 1301 - 108.6",
		"V 1302 - 108.6":                     "V 1302 - 108.6",
		"V 1401 - 102":                       "V 1401 - 102",
		"V 1501 - 99.4":                      "V 1501 - 99.4",
		"V 1601 - 108.6":                     "V 1601 - 108.6",
		"Copper":                             "তামা",
		"CycloPentane":                       "সাইক্লোপেনটেন",
		"580 x 645 x 1530 mm":                "580 x 645 x 1530 মিমি",
		"102/ 102 /50":                       "102/ 102 /50",
		"No Need":                            "প্রয়োজন নেই",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-2a3-nexx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfa2a3NexxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-2a3-nexx-xx"
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
        "Brand":               "Marcel",
        "Model Name":          "MFA-2A3-NEXX-XX",
        "Cooling Technology":  "Direct Cool",
        "Gross Volume":        "177 Ltr.",
        "Net Volume":          "175 Ltr.",
        "Weight":              "50 ± 2 Kg",
        "Refrigerant":         "R600a",
        "Temperature Control": "Mechanical",
        "Voltage":             "220 ~ 240",
        "Dimensions":          "555 x 630 x 1410 mm",
        "Packing Dimensions":  "580 x 645 x 1455 mm",
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
