package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE305Liter seeds specifications/options for product 'vision-vision-glass-door-refrigerator-re-305-liter'
type SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE305Liter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE305Liter creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE305Liter() *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE305Liter {
	return &SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE305Liter{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vision-glass-door-refrigerator-re-305-liter"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE305Liter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VISION Glass Door Refrigerator RE-305 Liter": "ভিশন গ্লাস ডোর রেফ্রিজারেটর RE-305 লিটার",
		"Single Door":         "একক দরজা",
		"305 Liters":          "৩০৫ লিটার",
		"0 Liters":            "০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"295 kWh":             "২৯৫ কিলোওয়াট ঘণ্টা",
		"550 x 570 x 1680 mm": "৫৫০ x ৫৭০ x ১৬৮০ মিমি",
		"355 kg":              "৩৫৫ কেজি",
		"SS":                  "স্টেইনলেস স্টিল",
		"Direct Cool":         "সরাসরি কুলিং",
		"Frost Free":          "ফ্রস্ট মুক্ত",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "শক্ত গ্লাস",
		"3":                   "৩",
		"4":                   "৪",
		"2":                   "২",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"160 ~ 260V":          "১৬০ ~ ২৬০ ভোল্ট",
		"50":                  "৫০",
		"5 Years":             "৫ বছর",
		"10":                  "১০",
		"R-600a":              "আর-৬০০এ",
		"Anti-bacterial gasket, Low noise compressor, Eco-friendly Technology": "ব্যাকটেরিয়া প্রতিরোধী গ্যাসকেট, কম শব্দ কম্প্রেসর, পরিবেশ বান্ধব প্রযুক্তি",
		"220V":             "২২০ ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
		"Refrigerant":      "শীতলক",
		"Gross Volume":     "সামগ্রিক ভলিউম",
		"Net Volume":       "নিট ভলিউম",
		"BLDC Inverter":    "বিএলডিসি ইনভার্টার",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vision-glass-door-refrigerator-re-305-liter'
func (s *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE305Liter) Seed(db *gorm.DB) error {
	productSlug := "vision-vision-glass-door-refrigerator-re-305-liter"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
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
		"Brand":                       "Vision",
		"Model Name":                  "VISION Glass Door Refrigerator RE-305 Liter",
		"Door Type":                   "Single Door",
		"Capacity":                    "305 Liters",
		"Refrigerator Capacity":       "305 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "550 x 570 x 1680 mm",
		"Weight":                      "355 kg",
		"Color":                       "SS",
		"Compressor Type":             "Normal",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "4",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "160 ~ 260V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "305 Ltr.",
		"Net Volume":                  "305 Ltr.",
		"Special Features":            "Anti-bacterial gasket, Low noise compressor, Eco-friendly Technology",
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
				if err := db.Create(sModel).Last(&sModel).Error; err != nil {
					return err
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
