package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorHitachiRSx800Gpbo seeds specifications/options for product 'hitachi-r-sx800gpbo'
type SpecificationSeederRefrigeratorHitachiRSx800Gpbo struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorHitachiRSx800Gpbo creates a new seeder instance
func NewSpecificationSeederRefrigeratorHitachiRSx800Gpbo() *SpecificationSeederRefrigeratorHitachiRSx800Gpbo {
	return &SpecificationSeederRefrigeratorHitachiRSx800Gpbo{
		BaseSeeder: BaseSeeder{name: "Specifications for hitachi-r-sx800gpbo"},
	}
}

func (s *SpecificationSeederRefrigeratorHitachiRSx800Gpbo) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Hitachi":             "হিটাচি",
		"R-SX800GPBO":         "আর-এসএক্স৮০০জিপিবিও",
		"Side by Side":        "সাইড বাই সাইড",
		"800 Liters":          "৮০০ লিটার",
		"480 Liters":          "৪৮০ লিটার",
		"320 Liters":          "৩২০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"480 kWh":             "৪৮০ কিলোওয়াট ঘণ্টা",
		"970 x 750 x 1820 mm": "৯৭০ x ৭৫০ x ১৮২০ মিমি",
		"125 kg":              "১২৫ কেজি",
		"Stainless Steel":     "স্টেইনলেস স্টিল",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"6":                   "৬",
		"10":                  "১০",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"Premium Inverter Compressor, Touch Control Panel, Multi Airflow System, Ice & Water Dispenser": "প্রিমিয়াম ইনভার্টার কম্প্রেসার, টাচ কন্ট্রোল প্যানেল, মাল্টি এয়ারফ্লো সিস্টেম, আইস এন্ড ওয়াটার ডিসপেন্সার",
		"Refrigerant":                   "রেফ্রিজারেন্ট",
		"Gross Volume":                  "মোট ভলিউম",
		"Net Volume":                    "নেট ভলিউম",
		"125 ± 2 Kg":                    "১২৫ ± ২ কেজি",
		"BLDC Inverter":                 "বিএলডিসি ইনভার্টার",
		"R-600a":                        "আর-৬০০এ",
		"800 Ltr.":                      "৮০০ লিটার",
		"480 Ltr.":                      "৪৮০ লিটার",
		"320 Ltr.":                      "৩২০ লিটার",
		"Advanced Side-by-Side Cooling": "অ্যাডভান্সড সাইড-বাই-সাইড কুলিং",
		"220 ~ 240V":                    "২২০ ~ ২৪০ভোল্ট",
		"Special Features":              "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'hitachi-r-sx800gpbo'
func (s *SpecificationSeederRefrigeratorHitachiRSx800Gpbo) Seed(db *gorm.DB) error {
	productSlug := "hitachi-r-sx800gpbo"
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
		"Model Name":                  "R-SX800GPBO",
		"Door Type":                   "Side by Side",
		"Capacity":                    "800 Liters",
		"Refrigerator Capacity":       "480 Liters",
		"Freezer Capacity":            "320 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "480 kWh",
		"Dimensions":                  "970 x 750 x 1820 mm",
		"Weight":                      "125 ± 2 Kg",
		"Color":                       "Stainless Steel",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "6",
		"Door Bins":                   "10",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "Yes",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "800 Ltr.",
		"Net Volume":                  "800 Ltr.",
		"Special Features":            "Premium Inverter Compressor, Touch Control Panel, Multi Airflow System, Ice & Water Dispenser",
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
