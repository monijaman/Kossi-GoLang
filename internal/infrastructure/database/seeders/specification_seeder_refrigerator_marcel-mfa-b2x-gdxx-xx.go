package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx seeds specifications/options for product 'marcel-mfa-b2x-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx() *SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-b2x-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-b2x-gdxx-xx": "মার্সেল-mfa-b2x-gdxx-xx",
		"MFA-B2X-GDXX-XX":        "MFA-B2X-GDXX-XX",
		"205 Ltr.":               "205 লিটার",
		"220 Ltr.":               "220 লিটার",
		"545 x 635 x 1580 mm":    "545 x 635 x 1580 মিমি",
		"48.5 ± 2 Kg":            "48.5 ± 2 কেজি",
		"RSCR":                   "RSCR",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical":             "মেকানিক্যাল",
		"220 ~ 240/50":           "220 ~ 240/50",
		"50":                     "50",
		"5 Star":                 "5 তারা",
		"12":                     "12",
		"52 ± 2 Kg":              "52 ± 2 কেজি",
		"Wire/3":                 "তার/3",
		"GPPS/4":                 "GPPS/4",
		"Steel":                  "ইস্পাত",
		"Cyclopentene":           "সাইক্লোপেন্টেন",
		"108.6 W":                "108.6 W",
		"580 x 645 x 1640 mm":    "580 x 645 x 1640 মিমি",
		"100/ 100/ 50":           "100/ 100/ 50",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-b2x-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfaB2xGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-b2x-gdxx-xx"
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
        "Model Name":          "MFA-B2X-GDXX-XX",
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
