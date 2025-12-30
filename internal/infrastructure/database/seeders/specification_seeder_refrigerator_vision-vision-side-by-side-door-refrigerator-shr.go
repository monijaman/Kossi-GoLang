package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVISIONSideBySideDoorRefrigeratorSHR seeds specifications/options for product 'vision-vision-side-by-side-door-refrigerator-shr'
type SpecificationSeederRefrigeratorVisionVISIONSideBySideDoorRefrigeratorSHR struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVISIONSideBySideDoorRefrigeratorSHR creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVISIONSideBySideDoorRefrigeratorSHR() *SpecificationSeederRefrigeratorVisionVISIONSideBySideDoorRefrigeratorSHR {
	return &SpecificationSeederRefrigeratorVisionVISIONSideBySideDoorRefrigeratorSHR{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vision-side-by-side-door-refrigerator-shr"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVISIONSideBySideDoorRefrigeratorSHR) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"Vision GD Refrigerator Side By Side Inverter SHR-566 EN-2": "ভিশন জিডি রেফ্রিজারেটর সাইড বাই সাইড ইনভার্টার এসএইচআর-৫৬৬ ইএন-২",
		"Side by Side":        "সাইড বাই সাইড",
		"566 Liters":          "৫৬৬ লিটার",
		"400 Liters":          "৪০০ লিটার",
		"166 Liters":          "১৬৬ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"400 kWh":             "৪০০ কিলোওয়াট ঘণ্টা",
		"700 x 650 x 1800 mm": "৭০০ x ৬৫০ x ১৮০০ মিমি",
		"100 kg":              "১০০ কেজি",
		"Silver":              "সিলভার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"4":                   "৪",
		"3":                   "৩",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"160 ~ 260V":          "১৬০ ~ ২৬০ভোল্ট",
		"50":                  "৫০",
		"5 Years":             "৫ বছর",
		"10":                  "১০",
		"R600a":               "আর-৬০০এ",
		"Inverter Compressor, Eco-friendly Technology": "ইনভার্টার কম্প্রেসর, ইকো-ফ্রেন্ডলি টেকনোলজি",
		"566 Ltr.":      "৫৬৬ লিটার",
		"BLDC Inverter": "বিএলডিসি ইনভার্টার",
		"Automatic":     "অটোমেটিক",
		"2":             "২",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vision-side-by-side-door-refrigerator-shr'
func (s *SpecificationSeederRefrigeratorVisionVISIONSideBySideDoorRefrigeratorSHR) Seed(db *gorm.DB) error {
	productSlug := "vision-vision-side-by-side-door-refrigerator-shr"
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
		"Model Name":                  "Vision GD Refrigerator Side By Side Inverter SHR-566 EN-2",
		"Door Type":                   "Side by Side",
		"Capacity":                    "566 Liters",
		"Refrigerator Capacity":       "400 Liters",
		"Freezer Capacity":            "166 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "400 kWh",
		"Dimensions":                  "700 x 650 x 1800 mm",
		"Weight":                      "100 kg",
		"Color":                       "Silver",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "4",
		"Crisper Drawers":             "3",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "42 dB",
		"Voltage":                     "160 ~ 260V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "566 Ltr.",
		"Net Volume":                  "566 Ltr.",
		"Special Features":            "Inverter Compressor, Eco-friendly Technology",
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
