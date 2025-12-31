package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaB4dNexxXx seeds specifications/options for product 'marcel-mfa-b4d-nexx-xx'
type SpecificationSeederRefrigeratorMarcelMfaB4dNexxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaB4dNexxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaB4dNexxXx() *SpecificationSeederRefrigeratorMarcelMfaB4dNexxXx {
	return &SpecificationSeederRefrigeratorMarcelMfaB4dNexxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-b4d-nexx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaB4dNexxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-b4d-nexx-xx": "মার্সেল-mfa-b4d-nexx-xx",
		"MFA-B4D-NEXX-XX":        "MFA-B4D-NEXX-XX",

		// Common values used in this seeder
		"Direct Cool":                    "ডিরেক্ট কুল",
		"244 Ltr.":                       "২৪৪ লিটার",
		"220 Ltr.":                       "২২০ লিটার",
		"N~ST":                           "N~ST",
		"220-240V~":                      "২২০-২৪০V~",
		"50Hz":                           "৫০Hz",
		"V 0701-108.6; V 0801-99.4":      "V 0701-108.6; V 0801-99.4",
		"RSCR":                           "RSCR",
		"Mechanical":                     "মেকানিক্যাল",
		"Manual":                         "ম্যানুয়াল",
		"Recessed":                       "রিসেসড",
		"Yes":                            "হ্যাঁ",
		"No":                             "না",
		"Copper":                         "তামা",
		"Cyclopentene":                   "সাইক্লোপেন্টেন",
		"R600a":                          "R600a",
		"51/57 (Net/Packing) Kg (±2 Kg)": "৫১/৫৭ (নেট/প্যাকিং) কেজি (±২ কেজি)",
		"545 x 640 x 1760 mm":            "৫৪৫ x ৬৪০ x ১৭৬০ মিমি",
		"V 0701-R600a; V 0801-R600a":     "V 0701-R600a; V 0801-R600a",
		"5 star (BDS 1850:2012)":         "৫ স্টার (BDS 1850:2012)",

		// Feature clusters
		"Wire/3":                "ওয়্যার/৩",
		"Wire/2":                "ওয়্যার/২",
		"Egg Case":              "ডিম ধারণ বাক্স",
		"Vegetable Box":         "শাকসবজি বক্স",
		"Interior Lamp":         "ইন্টারিয়র ল্যাম্প",
		"Can Storage Dispenser": "ক্যান স্টোরেজ ডিসপেনসার",
		"Deodorizer":            "ডিওডোরাইজার",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft 103/ 75/ 36": "লোডিং ক্যাপাসিটি- 40HQ/40Ft/20Ft 103/75/36",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-b4d-nexx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfaB4dNexxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-b4d-nexx-xx"
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
        "Model Name":          "MFA-B4D-NEXX-XX",
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
