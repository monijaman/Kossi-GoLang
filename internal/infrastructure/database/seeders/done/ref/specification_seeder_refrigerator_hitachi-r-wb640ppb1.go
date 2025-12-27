package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorHitachiRWb640Ppb1 seeds specifications/options for product 'hitachi-r-wb640ppb1'
type SpecificationSeederRefrigeratorHitachiRWb640Ppb1 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorHitachiRWb640Ppb1 creates a new seeder instance
func NewSpecificationSeederRefrigeratorHitachiRWb640Ppb1() *SpecificationSeederRefrigeratorHitachiRWb640Ppb1 {
	return &SpecificationSeederRefrigeratorHitachiRWb640Ppb1{
		BaseSeeder: BaseSeeder{name: "Specifications for hitachi-r-wb640ppb1"},
	}
}

func (s *SpecificationSeederRefrigeratorHitachiRWb640Ppb1) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Hitachi":             "হিটাচি",
		"R-WB640PPB1":         "আর-ডব্লিউবি৬৪০পিপিবি১",
		"Bottom Mount":        "বটম মাউন্ট",
		"640 Liters":          "৬৪০ লিটার",
		"430 Liters":          "৪৩০ লিটার",
		"210 Liters":          "২১০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"400 kWh":             "৪০০ কিলোওয়াট ঘণ্টা",
		"700 x 750 x 1800 mm": "৭০০ x ৭৫০ x ১৮০০ মিমি",
		"105 kg":              "১০৫ কেজি",
		"Stainless Steel":     "স্টেইনলেস স্টিল",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"4":                   "৪",
		"6":                   "৬",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"39 dB":               "৩৯ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Inverter Compressor, Humidity Control, Fresh Air Filter, LED Lighting": "ইনভার্টার কম্প্রেসার, হিউমিডিটি কন্ট্রোল, ফ্রেশ এয়ার ফিল্টার, এলইডি লাইটিং",
		"Refrigerant":          "রেফ্রিজারেন্ট",
		"Gross Volume":         "মোট ভলিউম",
		"Net Volume":           "নেট ভলিউম",
		"105 ± 2 Kg":           "১০৫ ± ২ কেজি",
		"BLDC Inverter":        "বিএলডিসি ইনভার্টার",
		"R-600a":               "আর-৬০০এ",
		"640 Ltr.":             "৬৪০ লিটার",
		"430 Ltr.":             "৪৩০ লিটার",
		"210 Ltr.":             "২১০ লিটার",
		"Bottom Mount Cooling": "বটম মাউন্ট কুলিং",
		"220 ~ 240V":           "২২০ ~ ২৪০ভোল্ট",
		"Special Features":     "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'hitachi-r-wb640ppb1'
func (s *SpecificationSeederRefrigeratorHitachiRWb640Ppb1) Seed(db *gorm.DB) error {
	productSlug := "hitachi-r-wb640ppb1"
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
		"Brand":                       "Hitachi",
		"Model Name":                  "R-WB640PPB1",
		"Door Type":                   "Bottom Mount",
		"Capacity":                    "640 Liters",
		"Refrigerator Capacity":       "430 Liters",
		"Freezer Capacity":            "210 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "400 kWh",
		"Dimensions":                  "700 x 750 x 1800 mm",
		"Weight":                      "105 ± 2 Kg",
		"Color":                       "Stainless Steel",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "6",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "39 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "640 Ltr.",
		"Net Volume":                  "640 Ltr.",
		"Special Features":            "Inverter Compressor, Humidity Control, Fresh Air Filter, LED Lighting",
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
