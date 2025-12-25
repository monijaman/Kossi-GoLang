package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWhirlpoolChestFreezerCf420t seeds specifications/options for product 'whirlpool-chest-freezer-cf-420t'
type SpecificationSeederRefrigeratorWhirlpoolChestFreezerCf420t struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWhirlpoolChestFreezerCf420t creates a new seeder instance
func NewSpecificationSeederRefrigeratorWhirlpoolChestFreezerCf420t() *SpecificationSeederRefrigeratorWhirlpoolChestFreezerCf420t {
	return &SpecificationSeederRefrigeratorWhirlpoolChestFreezerCf420t{
		BaseSeeder: BaseSeeder{name: "Specifications for whirlpool-chest-freezer-cf-420t"},
	}
}

func (s *SpecificationSeederRefrigeratorWhirlpoolChestFreezerCf420t) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Whirlpool":                "হুইরলপুল",
		"Chest Freezer CF 420T":    "চেস্ট ফ্রিজার সিএফ ৪২০টি",
		"Chest Door":               "চেস্ট দরজা",
		"420 Liters":               "৪২০ লিটার",
		"0 Liters":                 "০ লিটার",
		"5 Star":                   "৫ তারা",
		"5":                        "৫",
		"300 kWh":                  "৩০০ কিলোওয়াট ঘণ্টা",
		"1200 x 600 x 850 mm":      "১২০০ x ৬০০ x ৮৫০ মিমি",
		"65 kg":                    "৬৫ কেজি",
		"Silver":                   "সিলভার",
		"Inverter":                 "ইনভার্টার",
		"Frost Free":               "ফ্রস্ট ফ্রি",
		"Mechanical":               "মেকানিক্যাল",
		"Wire":                     "ওয়াইর",
		"1":                        "১",
		"0":                        "০",
		"No":                       "না",
		"45 dB":                    "৪৫ ডেসিবেল",
		"220V":                     "২২০ভোল্ট",
		"50":                       "৫০",
		"3 Years":                  "৩ বছর",
		"10":                       "১০",
		"R-600a":                   "আর-৬০০এ",
		"Manual":                   "ম্যানুয়াল",
		"Refrigerant":              "রেফ্রিজারেন্ট",
		"Gross Volume":             "মোট ভলিউম",
		"Net Volume":               "নেট ভলিউম",
		"BLDC Inverter":            "বিএলডিসি ইনভার্টার",
		"Chest Freezer Technology": "চেস্ট ফ্রিজার প্রযুক্তি",
		"220 ~ 240V":               "২২০ ~ ২৪০ভোল্ট",
		"Special Features":         "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'whirlpool-chest-freezer-cf-420t'
func (s *SpecificationSeederRefrigeratorWhirlpoolChestFreezerCf420t) Seed(db *gorm.DB) error {
	productSlug := "whirlpool-chest-freezer-cf-420t"
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
		"Brand":                       "Whirlpool",
		"Model Name":                  "Chest Freezer CF 420T",
		"Door Type":                   "Chest Door",
		"Capacity":                    "420 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "420 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "300 kWh",
		"Dimensions":                  "1200 x 600 x 850 mm",
		"Weight":                      "65 kg",
		"Color":                       "Silver",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire",
		"Number of Shelves":           "1",
		"Door Bins":                   "0",
		"Crisper Drawers":             "0",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "45 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "420 Ltr.",
		"Net Volume":                  "420 Ltr.",
		"Special Features":            "Chest Freezer Technology",
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
