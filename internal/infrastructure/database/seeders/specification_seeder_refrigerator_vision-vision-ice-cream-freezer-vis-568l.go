package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVisionIceCreamFreezerVIS568L seeds specifications/options for product 'vision-vision-ice-cream-freezer-vis-568l'
type SpecificationSeederRefrigeratorVisionVisionIceCreamFreezerVIS568L struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVisionIceCreamFreezerVIS568L creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVisionIceCreamFreezerVIS568L() *SpecificationSeederRefrigeratorVisionVisionIceCreamFreezerVIS568L {
	return &SpecificationSeederRefrigeratorVisionVisionIceCreamFreezerVIS568L{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vision-ice-cream-freezer-vis-568l"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVisionIceCreamFreezerVIS568L) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision":                            "ভিশন",
		"Vision Ice Cream Freezer VIS 568L": "ভিশন আইস ক্রিম ফ্রিজার ভিআইএস ৫৬৮এল",
		"Single Door":                       "একক দরজা",
		"568 Liters":                        "৫৬৮ লিটার",
		"0 Liters":                          "০ লিটার",
		"5 Star":                            "৫ তারা",
		"5":                                 "৫",
		"295 kWh":                           "২৯৫ কিলোওয়াট ঘণ্টা",
		"1590 x 680 x 785 mm":               "১৫৯০ x ৬৮০ x ৭৮৫ মিমি",
		"65 kg":                             "৬৫ কেজি",
		"Silver":                            "সিলভার",
		"Direct Cool":                       "সরাসরি কুলিং",
		"Frost Free":                        "ফ্রস্ট মুক্ত",
		"Electronic":                        "ইলেকট্রনিক",
		"Toughened Glass":                   "শক্ত গ্লাস",
		"3":                                 "৩",
		"4":                                 "৪",
		"2":                                 "২",
		"Yes":                               "হ্যাঁ",
		"No":                                "না",
		"40 dB":                             "৪০ ডেসিবেল",
		"220 ~ 240V":                        "২২০ ~ ২৪০ ভোল্ট",
		"50":                                "৫০",
		"5 Years":                           "৫ বছর",
		"10":                                "১০",
		"R-600a":                            "আর-৬০০এ",
		"High efficient cooling system, Low noise energy saving compressor, Use compressor Cooling Fan, with lock and key, Adjustable thermostat": "উচ্চ কার্যকর কুলিং সিস্টেম, কম শব্দ এনার্জি সেভিং কম্প্রেসর, কম্প্রেসর কুলিং ফ্যান ব্যবহার করুন, লক এবং কী সহ, অ্যাডজাস্টেবল থার্মোস্ট্যাট",
		"220V":             "২২০ ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
		"Refrigerant":      "শীতলক",
		"Gross Volume":     "সামগ্রিক ভলিউম",
		"Net Volume":       "নিট ভলিউম",
		"BLDC Inverter":    "বিএলডিসি ইনভার্টার",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vision-ice-cream-freezer-vis-568l'
func (s *SpecificationSeederRefrigeratorVisionVisionIceCreamFreezerVIS568L) Seed(db *gorm.DB) error {
	productSlug := "vision-vision-ice-cream-freezer-vis-568l"
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
		"Model Name":                  "Vision Ice Cream Freezer VIS 568L",
		"Door Type":                   "Single Door",
		"Capacity":                    "568 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "568 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "1590 x 680 x 785 mm",
		"Weight":                      "65 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Normal",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Wire Basket",
		"Number of Shelves":           "3",
		"Door Bins":                   "0",
		"Crisper Drawers":             "0",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "568 Liters",
		"Net Volume":                  "568 Liters",
		"Special Features":            "High efficient cooling system, Low noise energy saving compressor, Use compressor Cooling Fan, with lock and key, Adjustable thermostat",
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
