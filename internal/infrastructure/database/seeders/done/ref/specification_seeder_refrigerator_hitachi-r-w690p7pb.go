package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorHitachiRW690P7Pb seeds specifications/options for product 'hitachi-r-w690p7pb'
type SpecificationSeederRefrigeratorHitachiRW690P7Pb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorHitachiRW690P7Pb creates a new seeder instance
func NewSpecificationSeederRefrigeratorHitachiRW690P7Pb() *SpecificationSeederRefrigeratorHitachiRW690P7Pb {
	return &SpecificationSeederRefrigeratorHitachiRW690P7Pb{
		BaseSeeder: BaseSeeder{name: "Specifications for hitachi-r-w690p7pb"},
	}
}

func (s *SpecificationSeederRefrigeratorHitachiRW690P7Pb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Hitachi":             "হিটাচি",
		"R-W690P7PB":          "আর-ডব্লিউ৬৯০পি৭পিবি",
		"Side by Side":        "সাইড বাই সাইড",
		"690 Liters":          "৬৯০ লিটার",
		"410 Liters":          "৪১০ লিটার",
		"280 Liters":          "২৮০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"420 kWh":             "৪২০ কিলোওয়াট ঘণ্টা",
		"910 x 700 x 1780 mm": "৯১০ x ৭০০ x ১৭৮০ মিমি",
		"110 kg":              "১১০ কেজি",
		"Stainless Steel":     "স্টেইনলেস স্টিল",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"8":                   "৮",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Inverter Compressor, Ice & Water Dispenser, LED Lighting, Smart Diagnosis": "ইনভার্টার কম্প্রেসার, আইস এন্ড ওয়াটার ডিসপেন্সার, এলইডি লাইটিং, স্মার্ট ডায়াগনোসিস",
		"Refrigerant":                "রেফ্রিজারেন্ট",
		"Gross Volume":               "মোট ভলিউম",
		"Net Volume":                 "নেট ভলিউম",
		"110 ± 2 Kg":                 "১১০ ± ২ কেজি",
		"BLDC Inverter":              "বিএলডিসি ইনভার্টার",
		"R-600a":                     "আর-৬০০এ",
		"690 Ltr.":                   "৬৯০ লিটার",
		"410 Ltr.":                   "৪১০ লিটার",
		"280 Ltr.":                   "২৮০ লিটার",
		"Premium Cooling Technology": "প্রিমিয়াম কুলিং টেকনোলজি",
		"220 ~ 240V":                 "২২০ ~ ২৪০ভোল্ট",
		"Special Features":           "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'hitachi-r-w690p7pb'
func (s *SpecificationSeederRefrigeratorHitachiRW690P7Pb) Seed(db *gorm.DB) error {
	productSlug := "hitachi-r-w690p7pb"
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
		"Model Name":                  "R-W690P7PB",
		"Door Type":                   "Side by Side",
		"Capacity":                    "690 Liters",
		"Refrigerator Capacity":       "410 Liters",
		"Freezer Capacity":            "280 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "420 kWh",
		"Dimensions":                  "910 x 700 x 1780 mm",
		"Weight":                      "110 ± 2 Kg",
		"Color":                       "Stainless Steel",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "5",
		"Door Bins":                   "8",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "Yes",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "690 Ltr.",
		"Net Volume":                  "690 Ltr.",
		"Special Features":            "Inverter Compressor, Ice & Water Dispenser, LED Lighting, Smart Diagnosis",
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
