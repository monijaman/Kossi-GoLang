package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE180LDigital seeds specifications/options for product 'vision-vision-glass-door-refrigerator-re-180l-digital'
type SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE180LDigital struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE180LDigital creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE180LDigital() *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE180LDigital {
	return &SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE180LDigital{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vision-glass-door-refrigerator-re-180l-digital"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE180LDigital) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VISION Glass Door Refrigerator RE-180L Dahlia Red Top Mount": "VISION Glass Door Refrigerator RE-180L Dahlia Red Top Mount",
		"Single Door":         "Single Door",
		"180 Liters":          "180 লিটার",
		"0 Liters":            "০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"295 kWh":             "২৯৫ কিলোওয়াট ঘণ্টা",
		"572 x 662 x 1355 mm": "৫৭২ x ৬৬২ x ১৩৫৫ মিমি",
		"55.96 kg":            "৫৫.৯৬ কেজি",
		"Dahlia Red":          "ডাহলিয়া রেড",
		"LG":                  "এলজি",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"3":                   "৩",
		"2":                   "২",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"160 ~ 260V":          "১৬০ ~ ২৬০ভোল্ট",
		"50":                  "৫০",
		"4 Years":             "৪ বছর",
		"10":                  "১০",
		"R-600a":              "আর-৬০০এ",
		"Low noise compressor, 100% copper condenser": "লো নয়েজ কম্প্রেসর, ১০০% কপার কন্ডেনসার",
		"220V":             "২২০ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"202 Ltr.":         "২০২ লিটার",
		"BLDC Inverter":    "বিএলডিসি ইনভার্টার",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vision-glass-door-refrigerator-re-180l-digital'
func (s *SpecificationSeederRefrigeratorVisionVISIONGlassDoorRefrigeratorRE180LDigital) Seed(db *gorm.DB) error {
	productSlug := "vision-vision-glass-door-refrigerator-re-180l-digital"
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
		"Model Name":                  "VISION Glass Door Refrigerator RE-180L Dahlia Red Top Mount",
		"Door Type":                   "Single Door",
		"Capacity":                    "180 Liters",
		"Refrigerator Capacity":       "180 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "572 x 662 x 1355 mm",
		"Weight":                      "55.96 kg",
		"Color":                       "Dahlia Red",
		"Compressor Type":             "LG",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
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
		"Warranty":                    "4 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "202 Ltr.",
		"Net Volume":                  "180 Ltr.",
		"Special Features":            "Low noise compressor, 100% copper condenser",
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
