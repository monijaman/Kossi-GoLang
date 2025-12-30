package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVISIONRefrigeratorVis50LtrGray seeds specifications/options for product 'vision-vision-refrigerator-vis-50-ltr-gray'
type SpecificationSeederRefrigeratorVisionVISIONRefrigeratorVis50LtrGray struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVISIONRefrigeratorVis50LtrGray creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVISIONRefrigeratorVis50LtrGray() *SpecificationSeederRefrigeratorVisionVISIONRefrigeratorVis50LtrGray {
	return &SpecificationSeederRefrigeratorVisionVISIONRefrigeratorVis50LtrGray{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vision-refrigerator-vis-50-ltr-gray"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVISIONRefrigeratorVis50LtrGray) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision":                               "ভিশন",
		"VISION Refrigerator Vis- 50 Ltr Gray": "ভিশন রেফ্রিজারেটর ভিআইএস- ৫০ লিটার গ্রে",
		"Single Door":                          "সিঙ্গেল ডোর",
		"50 Liters":                            "৫০ লিটার",
		"35 Liters":                            "৩৫ লিটার",
		"15 Liters":                            "১৫ লিটার",
		"5 Star":                               "৫ তারা",
		"5":                                    "৫",
		"100 kWh":                              "১০০ কিলোওয়াট ঘণ্টা",
		"454 x 450 x 515 mm":                   "৪৫৪ x ৪৫০ x ৫১৫ মিমি",
		"14.56 kg":                             "১৪.৫৬ কেজি",
		"Gray":                                 "গ্রে",
		"Direct Cool":                          "সরাসরি কুলিং",
		"Frost Free":                           "ফ্রস্ট ফ্রি",
		"Mechanical":                           "মেকানিকাল",
		"Wire":                                 "ওয়্যার",
		"1":                                    "১",
		"2":                                    "২",
		"Yes":                                  "হ্যাঁ",
		"No":                                   "না",
		"40 dB":                                "৪০ ডেসিবেল",
		"220V":                                 "২২০ভোল্ট",
		"50":                                   "৫০",
		"5 Years":                              "৫ বছর",
		"10":                                   "১০",
		"R600a":                                "আর-৬০০এ",
		"Reversible door hinge, Adjustable leg, Eco-friendly Technology": "রিভার্সিবল ডোর হিঞ্জ, অ্যাডজাস্টেবল লেগ, ইকো-ফ্রেন্ডলি টেকনোলজি",
		"50 Ltr.": "৫০ লিটার",
		"Manual":  "ম্যানুয়াল",
		"0":       "০",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vision-refrigerator-vis-50-ltr-gray'
func (s *SpecificationSeederRefrigeratorVisionVISIONRefrigeratorVis50LtrGray) Seed(db *gorm.DB) error {
	productSlug := "vision-vision-refrigerator-vis-50-ltr-gray"
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
		"Model Name":                  "VISION Refrigerator Vis- 50 Ltr Gray",
		"Door Type":                   "Single Door",
		"Capacity":                    "50 Liters",
		"Refrigerator Capacity":       "35 Liters",
		"Freezer Capacity":            "15 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "100 kWh",
		"Dimensions":                  "454 x 450 x 515 mm",
		"Weight":                      "14.56 kg",
		"Color":                       "Gray",
		"Compressor Type":             "Normal",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire",
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
		"Special Features":            "Reversible door hinge, Adjustable leg, Eco-friendly Technology",
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
