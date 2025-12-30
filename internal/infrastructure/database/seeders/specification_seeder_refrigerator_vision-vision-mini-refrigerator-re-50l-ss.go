package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVISIONMiniRefrigeratorRE50LSS seeds specifications/options for product 'vision-vision-mini-refrigerator-re-50l-ss'
type SpecificationSeederRefrigeratorVisionVISIONMiniRefrigeratorRE50LSS struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVISIONMiniRefrigeratorRE50LSS creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVISIONMiniRefrigeratorRE50LSS() *SpecificationSeederRefrigeratorVisionVISIONMiniRefrigeratorRE50LSS {
	return &SpecificationSeederRefrigeratorVisionVISIONMiniRefrigeratorRE50LSS{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vision-mini-refrigerator-re-50l-ss"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVISIONMiniRefrigeratorRE50LSS) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision":                             "ভিশন",
		"VISION Mini Refrigerator RE-50L SS": "ভিশন মিনি রেফ্রিজারেটর আরই-৫০এল এসএস",
		"Glass Door":                         "গ্লাস ডোর",
		"50 Liters":                          "৫০ লিটার",
		"0 Liters":                           "০ লিটার",
		"5 Star":                             "৫ তারা",
		"5":                                  "৫",
		"100 kWh":                            "১০০ কিলোওয়াট ঘণ্টা",
		"470 x 496 x 445 mm":                 "৪৭০ x ৪৯৬ x ৪৪৫ মিমি",
		"18.61 kg":                           "১৮.৬১ কেজি",
		"SS":                                 "এসএস",
		"Frost Free":                         "ফ্রস্ট ফ্রি",
		"Electronic":                         "ইলেকট্রনিক",
		"Toughened Glass":                    "টাফেন্ড গ্লাস",
		"1":                                  "১",
		"2":                                  "২",
		"Yes":                                "হ্যাঁ",
		"No":                                 "না",
		"40 dB":                              "৪০ ডেসিবেল",
		"220V":                               "২২০ভোল্ট",
		"50":                                 "৫০",
		"5 Years":                            "৫ বছর",
		"10":                                 "১০",
		"R600a":                              "আর-৬০০এ",
		"3 layer anti-bacterial gasket, 100% copper condenser, Eco-friendly Technology": "৩ লেয়ার অ্যান্টি-ব্যাকটেরিয়াল গাস্কেট, ১০০% কপার কন্ডেন্সার, ইকো-ফ্রেন্ডলি টেকনোলজি",
		"50 Ltr.":   "৫০ লিটার",
		"Normal":    "নরমাল",
		"Automatic": "অটোমেটিক",
		"0":         "০",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vision-mini-refrigerator-re-50l-ss'
func (s *SpecificationSeederRefrigeratorVisionVISIONMiniRefrigeratorRE50LSS) Seed(db *gorm.DB) error {
	productSlug := "vision-vision-mini-refrigerator-re-50l-ss"
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
		"Model Name":                  "VISION Mini Refrigerator RE-50L SS",
		"Door Type":                   "Glass Door",
		"Capacity":                    "50 Liters",
		"Refrigerator Capacity":       "50 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "100 kWh",
		"Dimensions":                  "470 x 496 x 445 mm",
		"Weight":                      "18.61 kg",
		"Color":                       "SS",
		"Compressor Type":             "Normal",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "1",
		"Door Bins":                   "2",
		"Crisper Drawers":             "0",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "50 Ltr.",
		"Net Volume":                  "50 Ltr.",
		"Special Features":            "3 layer anti-bacterial gasket, 100% copper condenser, Eco-friendly Technology",
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
