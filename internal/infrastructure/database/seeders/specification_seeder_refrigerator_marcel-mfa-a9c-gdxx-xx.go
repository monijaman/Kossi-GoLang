package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaA9cGdxxXx seeds specifications/options for product 'marcel-mfa-a9c-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfaA9cGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaA9cGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaA9cGdxxXx() *SpecificationSeederRefrigeratorMarcelMfaA9cGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfaA9cGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-a9c-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaA9cGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-a9c-gdxx-xx": "মার্সেল-mfa-a9c-gdxx-xx",
		"MFA-A9C-GDXX-XX":        "MFA-A9C-GDXX-XX",
		"175 Ltr.":               "175 লিটার",
		"193 Ltr.":               "193 লিটার",
		"538 x 600 x 1230 mm":    "538 x 600 x 1230 মিমি",
		"45 ± 2 Kg":              "45 ± 2 কেজি",
		"RSCR":                   "RSCR",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical":             "মেকানিক্যাল",
		"R600a":                  "R600a",
		"220-240 V/ 50 Hz":       "220-240 ভোল্ট/ 50 হার্জ",
		"50":                     "50",
		"3 glass shelves":        "3 গ্লাস তাক",
		"Glass":                  "গ্লাস",
		"5 PS":                   "5 PS",
		"1":                      "1",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-a9c-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfaA9cGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-a9c-gdxx-xx"
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
		"Model Name":          "MFA-A9C-GDXX-XX",
		"Capacity":            "175 Ltr.",
		"Gross Volume":        "193 Ltr.",
		"Net Volume":          "175 Ltr.",
		"Dimensions":          "538 x 600 x 1230 mm",
		"Weight":              "45 ± 2 Kg",
		"Compressor Type":     "RSCR",
		"Defrost Type":        "Manual",
		"Temperature Control": "Mechanical",
		"Refrigerant":         "R600a",
		"Voltage":             "220-240 V/ 50 Hz",
		"Frequency (Hz)":      "50",
		"Special Features":    "Lock, Interior Lamp, Vegetable Crisper (1), Egg Tray (1), Medicine Box (1), Recessed Handle, Eco-friendly (100% CFC & HCFC Free) Green Technology, 2.25 L Bottle Accommodation, Double Layer Freezer Door, Uniform Flow",
		"Number of Shelves":   "3 glass shelves",
		"Shelf Material":      "Glass",
		"Door Bins":           "5 PS",
		"Crisper Drawers":     "1",
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
