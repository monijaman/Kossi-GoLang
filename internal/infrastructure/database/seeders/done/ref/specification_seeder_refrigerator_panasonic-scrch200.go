package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorPanasonicScrch200 seeds specifications/options for product 'panasonic-scrch200'
type SpecificationSeederRefrigeratorPanasonicScrch200 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorPanasonicScrch200 creates a new seeder instance
func NewSpecificationSeederRefrigeratorPanasonicScrch200() *SpecificationSeederRefrigeratorPanasonicScrch200 {
	return &SpecificationSeederRefrigeratorPanasonicScrch200{
		BaseSeeder: BaseSeeder{name: "Specifications for panasonic-scrch200"},
	}
}

func (s *SpecificationSeederRefrigeratorPanasonicScrch200) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Panasonic":                       "প্যানাসনিক",
		"SCR-CH200":                       "এসসিআর-সিএইচ২০০",
		"Chest":                           "চেস্ট",
		"200 Liters":                      "২০০ লিটার",
		"0 Liters":                        "০ লিটার",
		"5 Star":                          "৫ তারা",
		"5":                               "৫",
		"250 kWh":                         "২৫০ কিলোওয়াট ঘণ্টা",
		"1045 x 605 x 844 mm":             "১০৪৫ x ৬০৫ x ৮৪৪ মিমি",
		"33 kg":                           "৩৩ কেজি",
		"Silver":                          "সিলভার",
		"Inverter":                        "ইনভার্টার",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"Manual":                          "ম্যানুয়াল",
		"Mechanical":                      "মেকানিক্যাল",
		"Wire Shelves":                    "ওয়্যার শেল্ফ",
		"2":                               "২",
		"3":                               "৩",
		"1":                               "১",
		"Yes":                             "হ্যাঁ",
		"No":                              "না",
		"40 dB":                           "৪০ ডেসিবেল",
		"220 ~ 240V":                      "২২০ ~ ২৪০ভোল্ট",
		"50":                              "৫০",
		"5 Years Service":                 "5 বছর সার্ভিস",
		"Chest Freezer, Energy Efficient": "চেস্ট ফ্রিজার, এনার্জি এফিশিয়েন্ট",
		"Refrigerant":                     "রেফ্রিজারেন্ট",
		"Gross Volume":                    "মোট ভলিউম",
		"Net Volume":                      "নেট ভলিউম",
		"200 Ltr.":                        "২০০ লিটার",
		"R-134a":                          "আর-১৩৪এ",
		"Special Features":                "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'panasonic-scrch200'
func (s *SpecificationSeederRefrigeratorPanasonicScrch200) Seed(db *gorm.DB) error {
	productSlug := "panasonic-scrch200"
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
		"Brand":                       "Panasonic",
		"Model Name":                  "SCR-CH200",
		"Door Type":                   "Chest",
		"Capacity":                    "200 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "200 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1045 x 605 x 844 mm",
		"Weight":                      "33 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire Shelves",
		"Number of Shelves":           "2",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years Service",
		"Compressor Warranty (Years)": "2",
		"Refrigerant":                 "R-134a",
		"Gross Volume":                "200 Ltr.",
		"Net Volume":                  "200 Ltr.",
		"Special Features":            "Chest Freezer, Energy Efficient",
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
