package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorHitachiRBg410P6Pbx seeds specifications/options for product 'hitachi-r-bg410p6pbx'
type SpecificationSeederRefrigeratorHitachiRBg410P6Pbx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorHitachiRBg410P6Pbx creates a new seeder instance
func NewSpecificationSeederRefrigeratorHitachiRBg410P6Pbx() *SpecificationSeederRefrigeratorHitachiRBg410P6Pbx {
	return &SpecificationSeederRefrigeratorHitachiRBg410P6Pbx{
		BaseSeeder: BaseSeeder{name: "Specifications for hitachi-r-bg410p6pbx"},
	}
}

func (s *SpecificationSeederRefrigeratorHitachiRBg410P6Pbx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Hitachi":             "হিটাচি",
		"R-BG410P6PBX":        "আর-বিজি৪১০পি৬পিবিএক্স",
		"Bottom Mount":        "বটম মাউন্ট",
		"410 Liters":          "৪১০ লিটার",
		"280 Liters":          "২৮০ লিটার",
		"130 Liters":          "১৩০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"320 kWh":             "৩২০ কিলোওয়াট ঘণ্টা",
		"600 x 650 x 1700 mm": "৬০০ x ৬৫০ x ১৭০০ মিমি",
		"85 kg":               "৮৫ কেজি",
		"Stainless Steel":     "স্টেইনলেস স্টিল",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"3":                   "৩",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"37 dB":               "৩৭ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Inverter Compressor, LED Interior Lighting, Humidity Control Crisper, Antibacterial Gasket": "ইনভার্টার কম্প্রেসার, এলইডি ইন্টেরিয়র লাইটিং, হিউমিডিটি কন্ট্রোল ক্রিস্পার, অ্যান্টিব্যাকটেরিয়াল গ্যাসকেট",
		"Refrigerant":             "রেফ্রিজারেন্ট",
		"Gross Volume":            "মোট ভলিউম",
		"Net Volume":              "নেট ভলিউম",
		"85 ± 2 Kg":               "৮৫ ± ২ কেজি",
		"BLDC Inverter":           "বিএলডিসি ইনভার্টার",
		"R-600a":                  "আর-৬০০এ",
		"410 Ltr.":                "৪১০ লিটার",
		"280 Ltr.":                "২৮০ লিটার",
		"130 Ltr.":                "১৩০ লিটার",
		"Bottom Mount Efficiency": "বটম মাউন্ট এফিশিয়েন্সি",
		"220 ~ 240V":              "২২০ ~ ২৪০ভোল্ট",
		"Special Features":        "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'hitachi-r-bg410p6pbx'
func (s *SpecificationSeederRefrigeratorHitachiRBg410P6Pbx) Seed(db *gorm.DB) error {
	productSlug := "hitachi-r-bg410p6pbx"
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
		"Model Name":                  "R-BG410P6PBX",
		"Door Type":                   "Bottom Mount",
		"Capacity":                    "410 Liters",
		"Refrigerator Capacity":       "280 Liters",
		"Freezer Capacity":            "130 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "320 kWh",
		"Dimensions":                  "600 x 650 x 1700 mm",
		"Weight":                      "85 ± 2 Kg",
		"Color":                       "Stainless Steel",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "5",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "37 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "410 Ltr.",
		"Net Volume":                  "410 Ltr.",
		"Special Features":            "Inverter Compressor, LED Interior Lighting, Humidity Control Crisper, Antibacterial Gasket",
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
