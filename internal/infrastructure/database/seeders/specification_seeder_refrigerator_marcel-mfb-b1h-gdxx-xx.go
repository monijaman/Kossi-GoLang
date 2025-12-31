package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx seeds specifications/options for product 'marcel-mfb-b1h-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx() *SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b1h-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b1h-gdxx-xx": "মার্সেল-mfb-b1h-gdxx-xx",
		"MFB-B1H-GDXX-XX":        "MFB-B1H-GDXX-XX",

		// Common values used in this seeder
		"Direct Cool":                "ডিরেক্ট কুল",
		"177 Ltr.":                   "১৭৭ লিটার",
		"175 Ltr.":                   "১৭৫ লিটার",
		"N~ST":                       "N~ST",
		"220 ~ 240/ 50":              "২২০ ~ ২৪০/ ৫০",
		"V 01.01-88; V 01.02-88":     "V 01.01-88; V 01.02-88",
		"V 01.01-RSCR; V 01.02-RSCR": "V 01.01-RSCR; V 01.02-RSCR",
		"RSCR":                       "RSCR",
		"Mechanical":                 "মেকানিক্যাল",
		"Manual":                     "ম্যানুয়াল",
		"Recessed/ Grip":             "রিসেসড/গ্রিপ",
		"Yes":                        "হ্যাঁ",
		"No":                         "না",
		"Copper":                     "তামা",
		"Cyclopentene":               "সাইক্লোপেন্টেন",
		"R600a":                      "R600a",
		"50 ± 2 Kg":                  "৫০ ± ২ কেজি",
		"54 ± 2 Kg":                  "৫৪ ± ২ কেজি",
		"555 x 630 x 1410 mm":        "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm":        "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"105/ 105/ 52":               "১০৫/ ১০৫/ ৫২",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b1h-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB1hGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b1h-gdxx-xx"
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
        "Model Name":          "MFB-B1H-GDXX-XX",
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
