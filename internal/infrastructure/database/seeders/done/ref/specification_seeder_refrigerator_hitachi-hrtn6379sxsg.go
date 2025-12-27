package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorHitachiHrtn6379Sxsg seeds specifications/options for product 'hitachi-hrtn6379sxsg'
type SpecificationSeederRefrigeratorHitachiHrtn6379Sxsg struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorHitachiHrtn6379Sxsg creates a new seeder instance
func NewSpecificationSeederRefrigeratorHitachiHrtn6379Sxsg() *SpecificationSeederRefrigeratorHitachiHrtn6379Sxsg {
	return &SpecificationSeederRefrigeratorHitachiHrtn6379Sxsg{
		BaseSeeder: BaseSeeder{name: "Specifications for hitachi-hrtn6379sxsg"},
	}
}

func (s *SpecificationSeederRefrigeratorHitachiHrtn6379Sxsg) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Hitachi":             "হিটাচি",
		"HRTN6379SXSG":        "এইচআরটিএন৬৩৭৯এসএক্সএসজি",
		"Top Mount":           "টপ মাউন্ট",
		"379 Liters":          "৩৭৯ লিটার",
		"237 Liters":          "২৩৭ লিটার",
		"142 Liters":          "১৪২ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"280 kWh":             "২৮০ কিলোওয়াট ঘণ্টা",
		"600 x 650 x 1700 mm": "৬০০ x ৬৫০ x ১৭০০ মিমি",
		"75 kg":               "৭৫ কেজি",
		"Silver":              "সিলভার",
		"Inverter":            "ইনভার্টার",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Mechanical":          "মেকানিক্যাল",
		"Wire Shelves":        "ওয়্যার শেল্ফ",
		"3":                   "৩",
		"4":                   "৪",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Inverter Compressor, Stabilizer Free Operation, Large Vegetable Crisper": "ইনভার্টার কম্প্রেসার, স্ট্যাবিলাইজার ফ্রি অপারেশন, লার্জ ভেজিটেবল ক্রিস্পার",
		"Refrigerant":                 "রেফ্রিজারেন্ট",
		"Gross Volume":                "মোট ভলিউম",
		"Net Volume":                  "নেট ভলিউম",
		"75 ± 2 Kg":                   "৭৫ ± ২ কেজি",
		"BLDC Inverter":               "বিএলডিসি ইনভার্টার",
		"R-600a":                      "আর-৬০০এ",
		"379 Ltr.":                    "৩৭৯ লিটার",
		"237 Ltr.":                    "২৩৭ লিটার",
		"142 Ltr.":                    "১৪২ লিটার",
		"Energy Efficient Technology": "এনার্জি এফিশিয়েন্ট টেকনোলজি",
		"220 ~ 240V":                  "২২০ ~ ২৪০ভোল্ট",
		"Special Features":            "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'hitachi-hrtn6379sxsg'
func (s *SpecificationSeederRefrigeratorHitachiHrtn6379Sxsg) Seed(db *gorm.DB) error {
	productSlug := "hitachi-hrtn6379sxsg"
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
		"Model Name":                  "HRTN6379SXSG",
		"Door Type":                   "Top Mount",
		"Capacity":                    "379 Liters",
		"Refrigerator Capacity":       "237 Liters",
		"Freezer Capacity":            "142 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "280 kWh",
		"Dimensions":                  "600 x 650 x 1700 mm",
		"Weight":                      "75 ± 2 Kg",
		"Color":                       "Silver",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire Shelves",
		"Number of Shelves":           "3",
		"Door Bins":                   "4",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "379 Ltr.",
		"Net Volume":                  "379 Ltr.",
		"Special Features":            "Inverter Compressor, Stabilizer Free Operation, Large Vegetable Crisper",
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
