package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE280Liter seeds specifications/options for product 'vision-vision-glass-door-refrigerator-re-280-liter'
type SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE280Liter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE280Liter creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE280Liter() *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE280Liter {
	return &SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE280Liter{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vision-glass-door-refrigerator-re-280-liter"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE280Liter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VISION Glass Door Refrigerator RE-280 Liter Mirror Jaba Flower Top Mount": "VISION Glass Door Refrigerator RE-280 Liter Mirror Jaba Flower Top Mount",
		"Single Door":         "সিঙ্গেল ডোর",
		"280 Liters":          "২৮০ লিটার",
		"0 Liters":            "০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"295 kWh":             "২৯৫ কিলোওয়াট ঘণ্টা",
		"576 x 550 x 1574 mm": "৫৭৬ x ৫৫০ x ১৫৭৪ মিমি",
		"330 kg":              "330 কেজি",
		"Mirror Jaba Flower":  "মিরর জাবা ফ্লাওয়ার",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"3":                   "৩",
		"2":                   "২",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"160 ~ 260V":          "১৬০ ~ ২৬০ভোল্ট",
		"50":                  "৫০",
		"5 Years":             "৫ বছর",
		"10":                  "১০",
		"R-600a":              "আর-৬০০এ",
		"Anti-bacterial gasket, Low noise compressor, Eco-friendly Technology": "অ্যান্টি-ব্যাকটেরিয়াল গ্যাসকেট, লো নয়েজ কম্প্রেসর, ইকো-ফ্রেন্ডলি টেকনোলজি",
		"220V":             "২২০ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"BLDC Inverter":    "বিএলডিসি ইনভার্টার",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vision-glass-door-refrigerator-re-280-liter'
func (s *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE280Liter) Seed(db *gorm.DB) error {
	productSlug := "vision-vision-glass-door-refrigerator-re-280-liter"
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
		"Model Name":                  "VISION Glass Door Refrigerator RE-280 Liter Mirror Jaba Flower Top Mount",
		"Door Type":                   "Single Door",
		"Capacity":                    "280 Liters",
		"Refrigerator Capacity":       "280 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "576 x 550 x 1574 mm",
		"Weight":                      "330 kg",
		"Color":                       "Mirror Jaba Flower",
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
		"Gross Volume":                "280 Ltr.",
		"Net Volume":                  "280 Ltr.",
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
